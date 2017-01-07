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
	cmd := flag.String("cmd", "", "Source file to upload")

	flag.Parse()
	log.Println(*host, *port, *user, *cmd)
	if *host == "" || *user == "" || *cmd == "" {
		flag.Usage()
		return
	}
	ssh := &easyssh.MakeConfig{
		User:     *user,
		Password: *pwd,
		Server:   *host,
		// Optional key or Password without either we try to contact your agent SOCKET
		Key:  "/.ssh/id_rsa",
		Port: *port,
	}

	// Call Run method with command you want to run on remote server.
	response, err := ssh.Run(*cmd)
	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println(response)
	}
}
