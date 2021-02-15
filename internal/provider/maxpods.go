package provider

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Functions to get maximum number of pods supported per node type.
// https://github.com/weaveworks/eksctl/blob/main/pkg/nodebootstrap/maxpods_generate.go

const maxPodsPerNodeTypeSourceText = "https://raw.github.com/awslabs/amazon-eks-ami/master/files/eni-max-pods.txt"

var maxPodsPerNodeType = getMaxPodsPerNodeType()

func getMaxPodsPerNodeType() map[string]int {
	dict := make(map[string]int)

	resp, err := http.Get(maxPodsPerNodeTypeSourceText)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, line := range strings.Split(string(body), "\n") {
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			continue
		}
		instanceType := parts[0]
		maxPods, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err.Error())
		}
		dict[instanceType] = maxPods
	}

	return dict
}
