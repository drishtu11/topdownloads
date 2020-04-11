package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/drishtu11/topdownloads/httprequests"
	"github.com/drishtu11/topdownloads/pqheap"
)

// A PriorityQueue implements heap.Interface and holds Items.
// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type Artifacts struct {
	Artifacts []Artifact `json:"results"`
}

type Artifact struct {
	Repo       string    `json:"repo"`
	Path       string    `json:"path"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Size       int64     `json:"size"`
	Created    time.Time `json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `json:"modified"`
	ModifiedBy string    `json:"modified_by"`
	Updated    time.Time `json:"updated"`
}

// JSONBytesEqual compares the JSON in two byte slices.
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

// extractArtifactData : converts Http POST response to Map data and match artifact stats
func extractArtifactData(body []byte, artifactIP string, repoType string, binType string) map[string]int {
	var artifacts Artifacts
	json.Unmarshal(body, &artifacts)
	isAllFilesNeeded := false

	countMap := make(map[string]int)
	for i := 0; i < len(artifacts.Artifacts); i++ {
		if binType == "all" {
			isAllFilesNeeded = true
		}
		if strings.Contains(artifacts.Artifacts[i].Name, binType) || isAllFilesNeeded {
			downloadsCount := httprequests.GetFileStats(artifactIP, artifacts.Artifacts[i].Repo, artifacts.Artifacts[i].Path, artifacts.Artifacts[i].Name)
			countMap[artifacts.Artifacts[i].Name] = downloadsCount
		}
	}

	return countMap
}

// pollGetAllFiles : Polls every 5 seconds for all files from the
// artifactory which matches the given repo and binary file type
func pollGetAllFiles(artifactIP string, repoType string, binType string, num int) {
	ticker := time.NewTicker(time.Second * 5).C
	bodyCache := []byte{}
	for {
		select {
		case <-ticker:
			body := httprequests.GetAllFiles(artifactIP, repoType, binType)
			eq, _ := JSONBytesEqual(bodyCache, body)
			if eq == false {
				countMap := extractArtifactData(body, artifactIP, repoType, binType)
				pqheap.FindTopKDownloads(countMap, num)
				bodyCache = body
			}
		}
	}
}

func main() {
	/*
		artifactIP := "104.154.94.138"
		repoType := "jcenter-cache"
		binType := ".jar"
		num := 4
	*/
	// Extract all the arguments provide via input
	artifactIP := os.Args[1]
	repoType := os.Args[2]
	binType := os.Args[3]

	num, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Printf(err.Error())
	}
	// Poll Every 5 seconds for artifactory data
	pollGetAllFiles(artifactIP, repoType, binType, num)
}
