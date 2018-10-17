package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"

	"github.com/mholt/caddy/caddyhttp/fastcgi"
)

func main() {
	err := Get(os.Args[1])
	if err != nil {
		panic(err)
	}
}

func Get(addr string) error {
	u, err := url.Parse(addr)
	if err != nil {
		return err
	}
	if u.Scheme != "fcgi" {
		return fmt.Errorf("Bad scheme: %s", u.Scheme)
	}
	if u.Port() == "" {
		return fmt.Errorf("Explicit port is mandatory")
	}

	client, err := fastcgi.Dial("tcp", u.Host)
	if err != nil {
		return err
	}
	defer client.Close()
	p := strings.Trim(u.Path, "/")
	fcgiParams := map[string]string{
		"REQUEST_METHOD":  "GET",
		"SERVER_PROTOCOL": "HTTP/1.0",
		"SCRIPT_FILENAME": p,
		"SCRIPT_NAME":     "/" + p,
		"CONTENT_LENGTH":  "0",
		"SERVER_SOFTWARE": "go / fcgiping ",
		"REMOTE_ADDR":     "127.0.0.1",
	}

	fmt.Println(fcgiParams)
	resp, err := client.Get(fcgiParams, nil, 0)
	if err != nil {
		return err
	}
	fmt.Println("Status: ", resp.StatusCode, resp.Status)
	fmt.Println()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}
