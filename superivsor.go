package main

import (
	"log"
	"os/exec"
)

func startProgram(path string) error {
	cmd := exec.Command("cmd.exe", "/C", "start", "/b", path)
	err := cmd.Run()
	return err
}

func supervise() {
	programs, err := loadPrograms()
	if err != nil {
		log.Fatalf("error fetching programs: %v\n", err)
	}
	processes, err := processes()
	if err != nil {
		log.Fatalf("error fetching processes: %v\n", err)
	}

	for _, program := range programs {
		process := findProcessByName(processes, program.Name)
		if process == nil {
			// No process, need to restart
			err := startProgram(program.Path)
			if err != nil {
				log.Fatalf("error starting process: %v\n", err)
			}
		}
	}
}
