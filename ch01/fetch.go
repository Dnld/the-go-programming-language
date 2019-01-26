// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	// "io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("fetching: %s\nstatus: %s\nbody: %s\n", url, resp.Status, b)

		// Exercise 1.7: copy output instead of reading it
		// _, err = io.Copy(os.Stdout, resp.Body)
		// if err != nil {
		//  fmt.Fprintf(os.Stderr, "fetch: copying: %s: %v\n", url, err)
		//  os.Exit(1)
		// }
		// io.Copy.Close()

	}
}
