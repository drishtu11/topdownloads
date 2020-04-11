package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileStats(t *testing.T) {
	artifactIP := "104.154.94.138"
	repoType := "jcenter-cache"
	path := "org/apache/struts/struts2-core/2.3.14"
	name := "struts2-core-2.3.14.jar"

	downloadCount := getFileStats(artifactIP, repoType, path, name)
	fmt.Printf("%s has %v downloads\n", name, downloadCount)
	assert.Equal(t, 23, downloadCount)
}

func TestGetAllFiles(t *testing.T) {
	artifactIP := "104.154.94.138"
	repoType := "jcenter-cache"
	binType := ".jar"

	body := getAllFiles(artifactIP, repoType, binType)

	var artifacts Artifacts
	json.Unmarshal(body, &artifacts)
	for i := 0; i < 2; i++ {
		if strings.Contains(artifacts.Artifacts[i].Name, binType) {
			fmt.Println("Repo: " + artifacts.Artifacts[i].Repo)
			fmt.Println("Path: " + artifacts.Artifacts[i].Path)
			fmt.Println("Name: " + artifacts.Artifacts[i].Name)
			fmt.Println("=====")
		}
	}
}
