package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func main() {
	serverUri := "https://your.owncloud.install.com/owncloud/ocs/v1.php/apps/files_sharing/api/v1"
	username := "your.username"
	passwd := "your.password"

	client := &http.Client{}

	// Set the form POST body
	form := url.Values{}
	form.Add("path", "/Photos/Paris.jpg")
	form.Add("shareType", "3")
	form.Add("permissions", "1")

	// Build the core request object
	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", serverUri, "shares"),
		strings.NewReader(form.Encode()),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	req.SetBasicAuth(username, passwd)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyText))
}
