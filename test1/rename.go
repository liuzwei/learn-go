package test1

import (
	"io/ioutil"
	"os"
	"strings"
)

func renameFile() {
	// get all files in the directory
	files, err := ioutil.ReadDir("C:\\develop\\gocode\\src\\learn-go\\algorithm\\leetcode\\editor\\cn")
	if err != nil {
		panic(err)
	}
	// print file names
	for _, file := range files {
		originalFileName := file.Name()
		// judge original file name is contains " " and replace it with "_"
		if strings.Contains(originalFileName, " ") {
			newFileName := strings.ReplaceAll(originalFileName, " ", "_")
			// rename original file with new file name
			filePath := "C:\\develop\\gocode\\src\\learn-go\\algorithm\\leetcode\\editor\\cn\\"
			os.Rename(filePath+originalFileName, filePath+newFileName)
		}

	}
}
