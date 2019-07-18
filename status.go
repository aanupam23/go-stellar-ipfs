package stellarshell

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"

	shi "github.com/ipfs/go-ipfs-api"
)

// IsListening implements basic method for checking if ipfs daemon is running
func (s *StellarShell) IsListening() error {
	return errors.New("IPFS Daemon not running")
}

// TODO: Check if it is actually a valid IPFS hash
// IsValid checks the IPFS hash is a valid multiaddr hash
func (s *StellarShell) IsValid(thash string) {
	shell := shi.NewShell("localhost:5001")

	// Printing IPFS hash data
	output, err := shell.Cat(thash)
	if err != nil {
		panic(err) // ends where
	}
	p := make([]byte, 4)
	var buffer bytes.Buffer
	for {
		n, err := output.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println(string(p[:n]))
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(p[:n]))
		buffer.WriteString(string(p[:n]))
	}

	//Close the Reader
	output.Close()
	fmt.Println(buffer.String())
}

// IsReady implements method for checking if ipfs daemon is ready
func (s *StellarShell) IsReady(thash string) {
	shell := shi.NewShell("localhost:5001")

	// Printing IPFS hash data
	output, err := shell.Cat(thash)
	if err != nil {
		panic(err) // ends where
	}
	p := make([]byte, 4)
	var buffer bytes.Buffer
	for {
		n, err := output.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println(string(p[:n]))
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(p[:n]))
		buffer.WriteString(string(p[:n]))
	}

	//Close the Reader
	output.Close()
	fmt.Println(buffer.String())
}
