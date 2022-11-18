package clients

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Why using time.Duration?
func TimeoutClient(d time.Duration) *http.Client {
	client := http.Client{Timeout: d}
	return &client
}

func FetchData(c *http.Client, url string) ([]byte, error) {

	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func WriteToCwdFile(fName string, data []byte) error {
	cwd, _ := os.Getwd()
	if err := os.WriteFile(filepath.Join(cwd, fName), data, 0644); err != nil {
		return fmt.Errorf("Error writing file")
	}
	return nil
}

// type Client struct {
// 	Transport RoundTripper
// 	CheckRedirect func(req *Request, via []*Request) error
// 	Jar CookieJar
// 	Timeout time.Duration
// }

/*
	Functional options pattern

1. Create a struct that contains all the options that you want to add in the NewXXX func.
	- use pointers bc the values will need to change for a specific implementation of the options struct
2. Create an Option function type that accepts a pointer to an options struct and returns an error.
	- This function type will perform some operation on a specific implementation of the options struct
3. Create a With*function that returns a function of the Option type. This allows you to define and return an implementation of the Option function.
4. Create a constructor that accepts one or more functions of the Option type and returns a new object or an error.
	a. Set default options. For example, initialize a zero-value option struct var (var options option)
	b. for range through the Option type arguments and do the following:
		i. execute one of the option arg funcs on the zero-val option struct. Each option arg func mutates an options struct using the provided argument. For example, execute WithPort(8080) will


// struct that represents options for the client
// use pointers so you can change their values
type options struct {
	tout *time.Duration
}

// function type that updates the options struct
type Option func(options *options) error

// Configuration function that returns a closure, a function that
// references a var that is outside its body.
// Convention is to start this func using 'WithXXXX'.
// The closure adds the timeout value to the struct.
func WithTimeout(d time.Duration) Option {
	// return a function that adds to the options struct
	// the timeout value passed to the wrapper function
	// 1. validate inputs
	// 2. update config struct
	return func(options *options) error {
		options.tout = &d
		return nil
	}
}

/*
	1. Create new options struct
	2. Iterate through each of the opts args
*/

/*
func NewClient(opts ...Option) (*http.Client, error) {
	var options options
	// mutate the options struct by executing the options func
	//  on each of the opts
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}
	// now the options struct values are set using the provided
	// Option funcs

	//
	var tout time.Duration
	if options.tout == nil {
		tout = 10 * time.Second
	} else {

	}
	_ = tout

	return nil, nil
}
*/
