package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// getFileStats : Make http GET request afor a file from the
// artifactory which matches the given repo and binary file type
func getFileStats(artifactIP string, repo string, path string, name string) int {
	url := "http://" + artifactIP + "/artifactory/api/storage/" + repo + "/" + path + "/" + name + "?stats="

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Basic YWRtaW46NDlyTVU4VmpEdA==")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "64efe897-30bb-44a7-9ecf-5a5e336405cc")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("http.DefaultClient.Do() failed with:\n'%s'\n", err)
		return -1
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var artifactStats ArtifactStats
	json.Unmarshal(body, &artifactStats)

	return artifactStats.DownloadCount
}

// getAllFiles : Make http POST request to get all files from the
// artifactory which matches the given repo and binary file type
func getAllFiles(artifactIP string, repoType string, binType string) []byte {
	url := "http://" + artifactIP + "/artifactory/api/search/aql"

	postBody := "items.find(\n{\n        \"repo\":{\"$eq\":\"" + repoType + "\"}\n}\n)"
	//payload := strings.NewReader("items.find(\n{\n        \"repo\":{\"$eq\":\"jcenter-cache\"}\n}\n)")
	payload := strings.NewReader(postBody)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("authorization", "Basic YWRtaW46NDlyTVU4VmpEdA==")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("http.DefaultClient.Do() failed with:\n'%s'\n", err)
		return nil
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body
}
