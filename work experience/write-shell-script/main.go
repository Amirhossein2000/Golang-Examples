//this programm will create a shell file and you can run it by command ./filename.
package main

import "os"

func main() {
	f, err := os.Create("hello_world")
	f.WriteString("echo Hello_World")
	if err != nil {
		panic(err)
	}
	f.Chmod(0755)
	f.Close()
}
