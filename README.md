![sftpclient](assets/banner/sftpclient.jpg)

## Installation

```shell script
go get -u github.com/coolstina/sftpclient
```

## Example

```go
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
```

## Notice

The function is not perfect, will continue to update...
