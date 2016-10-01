package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {

	cmdName := "docker"
	cmdArgs := []string{"logs", "-f", "7f69b6bc17a3"}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	// Scan the reader
	scan(cmdReader)

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}
}

func scan(reader io.ReadCloser) {

	scanner := bufio.NewScanner(reader)
	go func() {
		for scanner.Scan() {
			s := scanner.Text()
			Encode(&s)
		}
	}()
}
