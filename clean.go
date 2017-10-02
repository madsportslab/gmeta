package main;

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
)

func removeFiles() {

	buf, err := ioutil.ReadFile(CLEAN_FILE)

	if err != nil {
		color.Red("[Error] removeFiles(): %s", err)
	}

	tagged := map[string]int{}

	json.Unmarshal(buf, &tagged)

	for k, _ := range tagged {
		color.Yellow("Deleting file %s...", k)
		os.Remove(k)
	}

} // removeFiles

func cleanFiles() {

	if fileExists(CLEAN_FILE) {

		scanner := bufio.NewScanner(os.Stdin)

		color.Yellow("Are you sure you want to delete all .html files? [Y/n] ")

		for scanner.Scan() {

			text := strings.ToUpper(scanner.Text())

			if text == "Y" {

				removeFiles()
				color.Yellow("Deleting .static.json...")				
				os.Remove(CLEAN_FILE)
				break

			} else if text == "N" {
				break
			}

		}

	}

	} // cleanFiles

func tagFile(filename string) {

	buf, _ := ioutil.ReadFile(CLEAN_FILE)
		
	tagged := make(map[string]int)

	if len(buf) > 0 {

		err := json.Unmarshal(buf, &tagged)
			
		if err != nil {
			color.Red("[Error] tagFile(): %s", err)
			return
		}

		_, ok := tagged[filename]

		if ok {
			color.Blue("exists bitch")
		} else {
			addTag(tagged, filename)
		}

	} else {
		addTag(tagged, filename)
	}

} // tagFile

func addTag(tagged map[string]int, filename string) {

	tagged[filename] = 1
	
	j, err := json.Marshal(tagged)

	if err != nil {
		color.Red("[Error] tagFile(): %s", err)
		return
	} else {
		
		err := ioutil.WriteFile(CLEAN_FILE, j, 0755)

		if err != nil {
			color.Red("[Error] addTag(): %s", err)
		}

	}
	
} // addTag
