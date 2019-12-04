package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/youryharchenko/easyssh"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Create MakeConfig instance with remote username, server address and path to private key.
	host := flag.String("host", "", "Remote ssh-host: e.g. 192.168.0.100")
	port := flag.String("port", "22", "Remote ssh-port: e.g. 22")
	user := flag.String("user", "", "Remote user name")
	pwd := flag.String("pwd", "", "Remote password")
	file := flag.String("file", "", "Source file to upload")
	out := flag.String("out", ".", "Remote dir to upload")
	flag.Parse()
	log.Println(*host, *port, *user, *file, *out)
	if *host == "" || *user == "" || *file == "" {
		flag.Usage()
		return
	}
	ssh := &easyssh.MakeConfig{
		User:     *user,
		Password: *pwd,
		Server:   *host,
		Key:      "/.ssh/id_rsa",
		Port:     *port,
	}

	// Call Scp method with file you want to upload to remote server.
	err := ssh.Scp(*file, *out)

	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println("success")

		//response, _ := ssh.Run("ls -al zipkin.rb")

		//fmt.Println(response)
	}
}
