package main

import (
	"fmt"
	c "go-notes/clients/http"
	"os"
	"time"
)

func main() {

	client := c.TimeoutClient(5 * time.Second)

	resp, err := c.FetchData(client, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%#v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s\n", resp)
	c.WriteToCwdFile("test", resp)
}
