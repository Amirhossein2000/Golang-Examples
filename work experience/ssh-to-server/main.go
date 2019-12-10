package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	// Every client must provide a host key check.  Here is a
	// simple-minded parse of OpenSSH's known_hosts file
	host := "172.16.4.2"
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Fatalf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}

	if hostKey == nil {
		log.Fatalf("no hostkey for %s", host)
	}

	config := ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("qwdxz!@#321"),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	client, err := ssh.Dial("tcp", host+":22", &config)

	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	sftp, err := sftp.NewClient(client)

	if err != nil {
		log.Fatal(err)
	}
	defer sftp.Close()

	srcPath := "/tmp/test/"
	dstPath := "/home/amirhossein/test/"

	fileList, err := sftp.ReadDir(srcPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range fileList {
		filename := name.Name()

		// Open the source file
		srcFile, err := sftp.Open(srcPath + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer srcFile.Close()

		// // Create the destination file
		dstFile, err := os.Create(dstPath + filename)
		if err != nil {
			log.Fatal(err)
		}
		// // Copy the file
		srcFile.WriteTo(dstFile)
		dstFile.Close()
	}

}
