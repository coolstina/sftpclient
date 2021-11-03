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

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/coolstina/sftpclient"
)

func main() {
	var host = os.Getenv("DESTINATION_SSH_HOST")         // example: 100.100.100.100:22
	var username = os.Getenv("DESTINATION_SSH_USERNAME") // example: hello
	var password = os.Getenv("DESTINATION_SSH_PASSWORD") // example: world

	client, err := sftpclient.NewClient(host, username, password)
	if err != nil {
		log.Panicf("init sftp client to failed: %+v\n", err)
	}

	path, err := client.UploadFile("test/data/hello.txt", "upload/test/hello.txt")
	if err != nil {
		log.Printf("upload file to failed: %+v\n", err)
	}

	fmt.Printf("upload file successfully, path: %+v\n", path)
}
