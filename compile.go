package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"html/template"
	
	"github.com/eknkc/amber"
	"github.com/fatih/color"
	"github.com/russross/blackfriday"
)

func fileExists(filename string) bool {

	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}

} // fileExists

func toHtml(filename string) string {

	var ret string

	ext := filepath.Ext(filename)

	switch ext {
	case AMBER:
		str := strings.TrimSuffix(filename, AMBER)
		ret = fmt.Sprintf("%s.html", str)

	case MARKDOWN:
		str := strings.TrimSuffix(filename, MARKDOWN)
		ret = fmt.Sprintf("%s.html", str)
		
	case MD:
		str := strings.TrimSuffix(filename, MD)
		ret = fmt.Sprintf("%s.html", str)
		
	}

	return ret

} // toHtml

func writeHtml(filename string, t *template.Template) {

	fh, err := os.Create(filename)
	
	defer fh.Close()

	if err != nil {
		color.Red("Error: %s", err)
	} else {

		err := t.Execute(fh, nil)

		if err != nil {
			color.Red("Error: %s", err)
		} else {
			tagFile(filename)
		}

	}

} // writeHtml

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
				} else {

					htmlFilename := toHtml(f)

					if fileExists(htmlFilename) {

						if *cmdForce {
							writeHtml(htmlFilename, t)
						} else {
							color.Yellow("File %s already exists, file skipped....", f)
						}

					} else {
						writeHtml(htmlFilename, t)
					}

				}

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
