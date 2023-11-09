package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Hello, world.")

	out, err := exec.Command("docker", "compose", "up").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output: %s\n", out)
}