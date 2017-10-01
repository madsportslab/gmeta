package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	
	"github.com/eknkc/amber"
	"github.com/fatih/color"
	"github.com/russross/blackfriday"
)

func getFiles() []string {

	dir := *cmdDir

	if len(dir) == 0 {
		dir = PWD
	}

	files, err := filepath.Glob("./*[.amber|.md]")

	if err != nil {
		color.Red("Error: %s", err)
		return nil
	}

	return files

} // getFiles

func compile() {

  filenames := getFiles()

	compiler := amber.New()

	for _, f := range(filenames) {

		color.Cyan(f)

		if strings.Contains(f, AMBER) {
			
			err := compiler.ParseFile(f)

			if err != nil {
				color.Red("Error: %s", err)
			} else {

				t, err := compiler.Compile()

				if err != nil {
					color.Red("error: %s", err)
				}

				t.Execute(os.Stdout, nil)

			}
	
		} else if strings.Contains(f, MD) || strings.Contains(f, MARKDOWN) {

			// TODO: incorporate layout.amber with body

			file, err := os.Open(f)

			if err != nil {
				color.Red("Error: %s", err)
			} else {

				buf, err := ioutil.ReadAll(file)
				
				if err != nil {
					color.Red("Error: %s", err)
				}
	
				out := blackfriday.MarkdownCommon(buf)
	
				color.Green(string(out))
				
			}
			
		}
		
	}

} // compile
