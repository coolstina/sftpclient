// Copyright 2021 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sftpclient

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/coolstina/sshclient"
	"github.com/pkg/errors"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Client SFTP client wrapper container.
type Client struct {
	sshClient  *ssh.Client
	sftpClient *sftp.Client
}

// NewClient initialize SFTP client instance.
func NewClient(host string, username, password string) (*Client, error) {
	sc, err := sshclient.NewClient(host, username, password)
	if err != nil {
		return nil, err
	}

	sfc, err := sftp.NewClient(sc.Instance())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to start SFTP subsystem")
	}

	client := Client{
		sshClient:  sc.Instance(),
		sftpClient: sfc,
	}

	return &client, nil
}

// Close close ssh and sftp connection.
func (cli *Client) Close() error {
	var err error

	if cli.sftpClient != nil {
		err = cli.sftpClient.Close()
	}

	if cli.sshClient != nil {
		err = cli.sshClient.Close()
	}

	return err
}

// ListFiles The remote SFTP server file list is displayed.
func (cli *Client) ListFiles(remoteDir string) ([]byte, error) {

	files, err := cli.sftpClient.ReadDir(remoteDir)
	if err != nil {
		return nil, fmt.Errorf("unable to list remote dir: %v\n", err)
	}

	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("listing [%s] ...\n\n", remoteDir))

	for _, f := range files {
		var name, modTime, size string

		name = f.Name()
		modTime = f.ModTime().Format("2006-01-02 15:04:05")
		size = fmt.Sprintf("%12d", f.Size())

		if f.IsDir() {
			name = name + "/"
			modTime = ""
			size = "PRE"
		}

		// Output each file name and size in bytes
		buffer.WriteString(fmt.Sprintf("%19s %12s %s\n", modTime, size, name))
	}

	return buffer.Bytes(), nil
}

// UploadFile Upload file to sftp server.
func (cli *Client) UploadFile(localFile, remoteFile string) (string, error) {
	srcFile, err := os.Open(localFile)
	if err != nil {
		return "", fmt.Errorf("unable to open local file: %v\n", err)
	}
	defer srcFile.Close()

	// Make remote directories recursion
	dir := sftp.Join(string(filepath.Separator), filepath.Dir(remoteFile))
	if err = cli.sftpClient.MkdirAll(dir); err != nil {
		return "", errors.Wrapf(err, "mkdir directory [%s] to failed", err)
	}

	// Note: SFTP To Go doesn't support O_RDWR mode
	dstFile, err := cli.sftpClient.OpenFile(remoteFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
	if err != nil {
		return "", fmt.Errorf("unable to open remote file: %v", err)
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return "", fmt.Errorf("unable to upload local file: %v\n", err)
	}

	return filepath.Join(dir, filepath.Base(remoteFile)), nil
}
