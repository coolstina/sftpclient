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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestSFTPClientSuite(t *testing.T) {
	suite.Run(t, new(SFTPClientSuite))
}

type SFTPClientSuite struct {
	suite.Suite
	client *Client
}

func (suite *SFTPClientSuite) BeforeTest(suiteName, testName string) {
	var err error
	var host = os.Getenv("DESTINATION_SSH_HOST")         // example: 100.100.100.100:22
	var username = os.Getenv("DESTINATION_SSH_USERNAME") // example: hello
	var password = os.Getenv("DESTINATION_SSH_PASSWORD") // example: world

	suite.client, err = NewClient(host, username, password)
	assert.NoError(suite.T(), err)
}

func (suite *SFTPClientSuite) Test_Client_ListFiles() {
	data, err := suite.client.ListFiles("/upload")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), data)
}

func (suite *SFTPClientSuite) Test_Client_UploadFile() {
	actual, err := suite.client.UploadFile(
		"test/data/hello.txt",
		"upload/helloworld/hello.txt",
	)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), actual)
	assert.Equal(suite.T(), `/upload/helloworld/hello.txt`, actual)
}
