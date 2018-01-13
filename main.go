package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Rule struct {
	id   string
	path string
	tags []string
}

const (
	ACCESSIBILITY = "ACCESSIBILITY"
	PERFORMANCE   = "PERFORMANCE"
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

func main() {

	path := os.Args[1] + "/checklist.md"

	createFile(path)
	file, _ := os.OpenFile(path, os.O_RDWR, 0644)
	defer file.Close()

	rules := []Rule{
		{id: "lang_attribute", path: "lang_attribute.md", tags: []string{ACCESSIBILITY}},
	}
	for _, rule := range rules {

		dat, err := ioutil.ReadFile("/home/manu/Documents/workspaces/golang/src/generator/rules/" + rule.path)
		if err != nil {
			panic(err)
		}
		file.WriteString(string(dat))
	}
}
