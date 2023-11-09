package main

import "fmt"

func main() {
    fmt.Println("Hello, world.")
}

out, err := exec.Command("docker", "compose", "up").Output()
if err != nil {
    log.Fatal(err)
}