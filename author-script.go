package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// run script recursivly in each sub-directory of /authors.
	root := "./authors"
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	// print directory name.
	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		// Go into each sub-directory and read files
		subDir := filepath.Join(root, f.Name())
		authorFiles, err := ioutil.ReadDir(subDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, af := range authorFiles {
			hasImage := strings.HasPrefix(af.Name(), "avatar.")
			if !hasImage {
				continue
			}

			indexFile := filepath.Join(subDir, "index.md")
			imageFile := filepath.Join(subDir, af.Name())

			// Open index file and add in images param.
			if fileExists(indexFile) {
				addImageParam(indexFile, imageFile)
			} else {
				addImageParam(indexFile+".md", imageFile)
			}
		}
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func addImageParam(indexFile string, imageFile string) {
	bin, err := ioutil.ReadFile(indexFile)
	if err != nil {
		log.Fatal(err)
	}

	text := string(bin)

	// Split string on ---
	parts := strings.Split(text, "---")
	doc := []string{
		parts[0],
		"---",
		parts[1],
		fmt.Sprintf(`images:
  - url: /engineering-education/%s %s`, imageFile, "\n"),
		"---",
		parts[2],
	}
	content := strings.Join(doc, "")

	err = ioutil.WriteFile(indexFile, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
