package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	conf := &tls.Config{
		//InsecureSkipVerify: true,
	}

	fmt.Println("1")

	conn, err := tls.Dial("tcp", "google.com:443", conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("2")

	str := "GET / HTTP/1.1\r\nHost: google.com\r\nUser-Agent: curl/7.77.0\r\nAccept: */*\r\n\r\n"

	fmt.Println("client: connected to: ", conn.RemoteAddr())
	state := conn.ConnectionState()
	for _, cert := range state.PeerCertificates {
		fmt.Println("Issuer:", cert.Issuer)
		fmt.Println("Subject:", cert.Subject)
		fmt.Println("Version:", cert.Version)
		fmt.Println("NotAfter:", cert.NotAfter)
		fmt.Println("DNS names:", cert.DNSNames)

	}
	fmt.Println("client: handshake: ", state.HandshakeComplete)
	fmt.Println("client: mutual: ", state.NegotiatedProtocolIsMutual)

	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println(n, err)
		return
	}

	fmt.Println(string(buf[:n]))
}
