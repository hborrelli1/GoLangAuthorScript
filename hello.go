package main

import (
	"fmt"
	"strings"

	// "io/fs"
	"io/ioutil"
	"log"

	// "os"
	"path/filepath"
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
			fmt.Println(imageFile, indexFile)

			// Open index file and add in images param.
			bin, err := ioutil.ReadFile(indexFile)
			if err != nil {
				log.Fatal(err)
			}

			text := string(bin)
			fmt.Println(text)

			// Split string on ---
			parts := strings.Split(text, "---")
			fmt.Println(parts, len(parts))
			doc := []string{
				parts[0],
				"---",
				parts[1],
				fmt.Sprintf(`images:
  - %s %s`, imageFile, "\n"),
				"---",
				parts[2],
			}
			content := strings.Join(doc, "")
			fmt.Println(content)

			err = ioutil.WriteFile(indexFile, []byte(content), 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
