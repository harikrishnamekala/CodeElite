package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
)

func data(command string) {

	cmd := exec.Command("bash","-c",command)
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdin: %s", err.Error())
	}
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdout: %s", err.Error())
	}
	reader := bufio.NewReader(stdout)
	go func(reader io.Reader) {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			log.Printf("Reading from subprocess: %s", scanner.Text())
			stdin.Write([]byte("some sample text\n"))
		}
	}(reader)
	if err := cmd.Start(); nil != err {
		log.Fatalf("Error starting program: %s, %s", cmd.Path, err.Error())
	}
	cmd.Wait()

}

func main(){
	data("gcc main.c -o ./main")
	data("./main")
}
