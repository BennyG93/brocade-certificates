package brocade

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"crypto/tls"
)

func Delcert(url, cert, username, password, api_version string) string {

	// Building the URL using the username, pass, LB url
	urlBuild := fmt.Sprintf("https://%s:%s@%s:9070/api/tm/%s/config/active/ssl/server_keys/%s/", username, password, url, api_version, cert)

	// Bypassing secure connection
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: tr}

	// Building the request using the URL build above
	request, err := http.NewRequest("DELETE", urlBuild, nil)

	// Returning the request and catching end errors
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// Putting whole response in string
	bodyText, err := ioutil.ReadAll(response.Body)

	// Returning the response status and body
	fmt.Println("response Status", response.Status)
	s := string(bodyText)
	fmt.Println(s)
	return s
}