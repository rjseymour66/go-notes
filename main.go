package main

import (
	"fmt"
	c "go-notes/clients"
	"io"
	"os"
	"time"
)

func main() {

	c := c.TimeoutClient(5 * time.Second)

	resp, err := c.Get("https://www.example.com")

	if err != nil {
		fmt.Fprintf(os.Stdout, "%#v\n", err)
		fmt.Fprintln(os.Stdout, "This is the error")
		os.Exit(1)
	}

	// what exactly does this do, and when to do it?
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%#v\n", err)
		fmt.Fprintln(os.Stdout, "No, THIS is the error")
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s\n", body)
}
