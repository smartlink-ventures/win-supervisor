package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func startProgram(path string) error {
	cmd := exec.Command("cmd.exe", "/C", "start", path)
	err := cmd.Run()
	return err
}

func supervise() {
	rootDir := os.Getenv("SYSTEMDRIVE") + "\\"
	logPath := fmt.Sprintf("%swindsupervisor\\error.log", rootDir)
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
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
