package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Rule struct {
	id   string
	path string
	tags []string
}

const (
	ACCESSIBILITY = "ACCESSIBILITY"
	PERFORMANCE   = "PERFORMANCE"
	QUALITY       = "QUALITY"
	WEB           = "WEB"
	SECURITY      = "SECURITY"
	ELASTICSEARCH = "ELASTICSEARCH"
)

func createFile(path string) {
	// detect if file exists
	_, err := os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		defer file.Close()

		if err != nil {
			panic("Error when creatng the file")
		}

	}

	fmt.Println("==> done creating file", path)
}

func Filter(vs []Rule, f func(Rule) bool) []Rule {
	vsf := make([]Rule, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Contain(ruleTag string, tags []string) bool {
	for _, tag := range tags {
		if ruleTag == tag {
			return true
		}
	}
	return false
}

func Contains(ruleTags []string, tags []string) bool {
	for _, tag := range ruleTags {
		if Contain(tag, tags) {
			return true
		}
	}
	return false
}

func main() {

	var pathDir string
	var tags []string

	for i, option := range os.Args[0:] {
		switch option {
		case "-r":
			pathDir = os.Args[i+1]
			break
		case "-t":
			tags = strings.Split(os.Args[i+1], ",")
			break
		default:
			break
		}
	}

	if pathDir == "" {
		panic("You should define the path of the generated audit report file")
	}
	if len(tags) == 0 {
		panic("You should define a comma separated list of tags")
	}

	rules := []Rule{
		{id: "lang_attribute", path: "lang_attribute.md", tags: []string{ACCESSIBILITY, WEB}},
		{id: "prettier", path: "prettier.md", tags: []string{QUALITY, WEB}},
		{id: "retire", path: "retire.md", tags: []string{QUALITY, WEB, SECURITY}},
		{id: "one_h1", path: "one_h1.md", tags: []string{QUALITY, WEB, ACCESSIBILITY}},
		{id: "rel_noopener", path: "rel_noopener.md", tags: []string{QUALITY, WEB, SECURITY}},
		{id: "search_input", path: "search_input.md", tags: []string{QUALITY, WEB, ACCESSIBILITY}},
		{id: "hierarchical_title", path: "hierarchical_title.md", tags: []string{QUALITY, WEB, ACCESSIBILITY}},

		{id: "cluster_name", path: "elasticsearch/cluster_name.md", tags: []string{ELASTICSEARCH}},
		{id: "data_folders", path: "elasticsearch/data_folders.md", tags: []string{ELASTICSEARCH}},
		{id: "delete_all", path: "elasticsearch/delete_all.md", tags: []string{ELASTICSEARCH}},
		{id: "disable_swapping", path: "elasticsearch/disable_swapping.md", tags: []string{ELASTICSEARCH}},
		{id: "es_heap_size", path: "elasticsearch/es_heap_size.md", tags: []string{ELASTICSEARCH}},
		{id: "improve_bulk", path: "elasticsearch/improve_bulk.md", tags: []string{ELASTICSEARCH}},
		{id: "node_name", path: "elasticsearch/node_name.md", tags: []string{ELASTICSEARCH}},
		{id: "shard_size", path: "elasticsearch/shard_size.md", tags: []string{ELASTICSEARCH}},
		{id: "split_brain", path: "elasticsearch/split_brain.md", tags: []string{ELASTICSEARCH}},
	}

	filteredRules := Filter(rules, func(rule Rule) bool {
		return Contains(rule.tags, tags)
	})
	path := pathDir + "/checklist.md"
	createFile(path)
	file, _ := os.OpenFile(path, os.O_RDWR, 0644)
	defer file.Close()

	for _, rule := range filteredRules {

		dat, err := ioutil.ReadFile("/Users/emmanueldemey/Documents/workspaces/oss/audit-generator/rules/" + rule.path)
		if err != nil {
			panic(err)
		}
		file.WriteString(string(dat))
	}
}
