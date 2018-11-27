package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	homedir "github.com/mitchellh/go-homedir"
)

func getUserName() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filepath.FromSlash(home)
	array := strings.Split(home, "\\")
	username := array[len(array)-1]

	return username
}

func visit(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil
	}

	matched, err := filepath.Match("*.md", fi.Name())

	if err != nil {
		panic(err)
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		fmt.Println(path)

		newContents := strings.Replace(string(read), getUserName(), "<username>", -1)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}

	}

	return nil
}

// Walk get all generated cli-reference *.md file and replace currenct username with <username>
func Walk() {
	err := filepath.Walk(".", visit)
	if err != nil {
		panic(err)
	}
}
