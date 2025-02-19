package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	host := "10.1.0.22"
	port := "22"
	user := "root"
	cmd := "ps | head"

	session, client := connectSSH(host, port, user)
	defer session.Close()
	defer client.Close()

	clientSFTP, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal(err)
	}
	defer clientSFTP.Close()

	fh, _ := clientSFTP.Open("/var/log/ftp/syslog.log")

	const size = 820
	x := make([]byte, size)
	_, err = io.ReadFull(fh, x)

	fmt.Printf("%s", x)

	// setup standard out and error
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	// run single command
	err = session.Run(cmd)
	errExit(err, "error running command "+cmd)
}

func connectSSH(host string, port string, user string) (*ssh.Session, *ssh.Client) {
	// get remote host public key and local host private key
	hostKey := getHostKey(host)

	key, err := ioutil.ReadFile("/root/.ssh/id_rsa")
	errExit(err, "unable to read private key file")

	signer, err := ssh.ParsePrivateKey(key)
	errExit(err, "unable to parse private key")

	// ssh client config
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)},

		// verify host public key
		HostKeyCallback: ssh.FixedHostKey(hostKey),

		// optional host key algo list
		HostKeyAlgorithms: []string{
			ssh.KeyAlgoRSA,
			ssh.KeyAlgoDSA,
			ssh.KeyAlgoECDSA256,
			ssh.KeyAlgoECDSA384,
			ssh.KeyAlgoECDSA521,
			ssh.KeyAlgoED25519,
		},

		// optional tcp connect timeout
		Timeout: 5 * time.Second,
	}

	// connect
	client, err := ssh.Dial("tcp", host+":"+port, config)
	errExit(err, "error creating ssh connection")

	// start session
	sess, err := client.NewSession()
	errExit(err, "error creating ssh session")

	return sess, client
}

func getHostKey(host string) ssh.PublicKey {

	// parse OpenSSH known_hosts file and return public key

	file, err := os.Open("/root/.ssh/known_hosts")
	errExit(err, "can't open known_hosts")
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
			errExit(err, "error parsing"+fields[2])
			break
		}
	}

	if hostKey == nil {
		log.Fatalf("no hostkey found for %s", host)
	}

	return hostKey
}

func errExit(err error, msg string) {
	if err != nil {
		log.Fatalf("%q: %v", msg, err)
	}
}
