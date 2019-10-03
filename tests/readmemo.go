package main

import (
        "fmt"
        si "github.com/aanupam23/go-stellar-ipfs"
)

func main() {
        ss := si.NewStellarShell("localhost:5001", "public")

        thash, err: ss.ReadMemoHash(txnid)

        output, err := ss.IpfsDataString(thash)
        if err != nil {
                error := ss.IsListening()
                panic(error) // ends where
        }

        fmt.Println(output)
}