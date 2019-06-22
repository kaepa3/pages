package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

var mdFile = regexp.MustCompile(`.+\.md$`)

func main() {
	ignoreList := []string{
		".git",
	}
	fmt.Println(dirWalk("./", ignoreList))
}

func dirWalk(dir string, ignoreList []string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			if !arrayContains(ignoreList, file.Name()) {
				paths = append(paths, dirWalk(filepath.Join(dir, file.Name()), ignoreList)...)
				continue
			}
		}
		if mdFile.MatchString(file.Name()) {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths
}

func arrayContains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
