package main

import (
	"fmt"

	si "github.com/aanupam23/go-stellar-ipfs"
)

func main() {

	ss := si.NewStellarShell("localhost:5001", "public")

	// Adding file to IPFS
	thash, err := ss.AddIpfsString("Meow")
	if err != nil {
		error := ss.IsListening()
		panic(error)
	}

	fmt.Println(thash)

	output, err := ss.IpfsDataString(thash)
	if err != nil {
		error := ss.IsListening()
		panic(error)
	}

	fmt.Println(output)
}
