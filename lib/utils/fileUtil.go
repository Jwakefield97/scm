package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

//ScanForFiles scans a directory for all files and returns them in a map {"files":["./myfile.txt",""]}
//dir is directory to scan
func ScanForFiles(dir string) map[string]interface{} {
	returnObj := map[string]interface{}{}
	files := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if len(path) < 4 {
				files = append(files, path)
			} else if path[:3] != ".scm" { //file not in .scm directory
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("error occured while walking directory")
	}
	returnObj["files"] = files
	return returnObj
}
