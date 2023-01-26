/*
Write a program that issues a Linux command to the local host.
You can try anything you'd like, but something like, ping -c 4 or nslookup google.com might be fun to try.
*/


package main

import (
    "fmt"
    "log"
    "os/exec"
)

func main() {

    commandToRun := []string{"nslookup", "google.com"}

    out, err := exec.Command(commandToRun[0], commandToRun[1]).Output()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Command being run: %s %s\n", commandToRun[0], commandToRun[1])

    fmt.Println(string(out))
}
