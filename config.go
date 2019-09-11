package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Program struct {
	Name string
	Path string
	Exe string
}

func getConfigPath() string {
	rootDir := os.Getenv("SYSTEMDRIVE") + "\\"
	fullPath := fmt.Sprintf("%swinsupervisor\\winsupervisor.conf", rootDir)
	return fullPath
}

func parseConfigFile(file []byte) ([]Program, error) {
	lines := strings.Split(string(file[:]), "\n")

	var startingIndexes []int
	for index, line := range lines {
		matched, err := regexp.MatchString(`^\[program:.*\]`, line)
		if err != nil {
			return nil, err
		}
		if matched {
			// This is the beginning of program definition
			startingIndexes = append(startingIndexes, index)
		}
	}

	var blocks [][]string
	for index, startingIndex := range startingIndexes {
		if index == len(startingIndexes) - 1 {
			block := lines[startingIndex:]
			blocks = append(blocks, block)
			continue
		}
		block := lines[startingIndex:startingIndexes[index+1]]
		blocks = append(blocks, block)
	}

	var programs []Program
	for _, block := range blocks {
		program := Program{}
		for _, line := range block {
			isHeader, err := regexp.MatchString(`^\[program:.*\]`, line)
			if err != nil {
				return nil, err
			}
			if isHeader {
				name := strings.Split(line, ":")[1]
				name = strings.Split(name, "]")[0]
				program.Name = name
			} else {
				pieces := strings.Split(line, "=")
				if strings.ToLower(pieces[0]) == "exe" {
					program.Exe = pieces[1]
				}
				if strings.ToLower(pieces[0]) == "path" {
					program.Path = pieces[1]
				}
			}
		}
		programs = append(programs, program)
	}
	return programs, nil
}

func loadPrograms() ([]Program, error) {
	fullPath := getConfigPath()
	file, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	programs, err := parseConfigFile(file)
	return programs, err
}
