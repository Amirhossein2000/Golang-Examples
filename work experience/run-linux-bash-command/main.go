package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := "nping -c 1 --udp -p 9999 --dest-ip 127.0.0.1 -g 9999 -S 127.2.2.2"
	fmt.Println(string(runcmd(cmd, true)))
}

func runcmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			log.Fatal(err)
			panic("some error found")
		}
		return out
	}
	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}
