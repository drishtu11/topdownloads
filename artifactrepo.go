package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

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
			downloadsCount := getFileStats(artifactIP, artifacts.Artifacts[i].Repo, artifacts.Artifacts[i].Path, artifacts.Artifacts[i].Name)
			countMap[artifacts.Artifacts[i].Name] = downloadsCount
		}
	}

	return countMap
}

// findTopKDownloads : finds the top K downloads of files from given repo
func findTopKDownloads(countMap map[string]int, num int) {
	pq := make(PriorityQueue, len(countMap))
	i := 0
	for value, priority := range countMap {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Take the items out; they arrive in decreasing priority order.
	count := 0
	fmt.Printf("----------------------------------------\n")
	fmt.Printf("Top %d Downloads\n", num)
	fmt.Printf("----------------------------------------\n")
	for pq.Len() > 0 && count < num {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Artifact : %s\nDownloads : %d\n\n", item.value, item.priority)
		count++
	}
	fmt.Printf("----------------------------------------\n")
	fmt.Println("")
}

// pollGetAllFiles : Polls every 5 seconds for all files from the
// artifactory which matches the given repo and binary file type
func pollGetAllFiles(artifactIP string, repoType string, binType string, num int) {
	ticker := time.NewTicker(time.Second * 5).C
	bodyCache := []byte{}
	for {
		select {
		case <-ticker:
			body := getAllFiles(artifactIP, repoType, binType)
			eq, _ := JSONBytesEqual(bodyCache, body)
			if eq == false {
				countMap := extractArtifactData(body, artifactIP, repoType, binType)
				findTopKDownloads(countMap, num)
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
