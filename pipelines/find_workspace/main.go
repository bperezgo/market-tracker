package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
)

func main() {
	// log.Println("[INFO] Reviewing the tests to run in the pipeline")
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal("[ERROR] failed getting the rootDir name;", err)
	}
	scriptPath := path.Join(rootDir, "./go.work")
	script, err := os.ReadFile(scriptPath)
	if err != nil {
		log.Fatal("[ERROR] failed to read the file;", err)
	}
	workspaces := findWorkspaces(script)
	// Send results to standard output
	log.SetOutput(os.Stdout)

	for _, w := range workspaces {
		fmt.Println(w)
	}
	os.Exit(0)
}

func findWorkspaces(file []byte) (res []string) {
	r, err := regexp.Compile(`\./[\w+/?]*`)
	if err != nil {
		log.Fatal("regex expresion failed")
	}
	matches := r.FindAll(file, -1)
	for _, match := range matches {
		res = append(res, string(match))
	}
	return res
}
