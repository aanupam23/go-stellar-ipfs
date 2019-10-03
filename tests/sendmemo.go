package main

import (
        "fmt"
        "github.com/stellar/go/txnbuild"
        si "github.com/aanupam23/go-stellar-ipfs"
)

func main() {
        kp, _ := keypair.Parse(SECRETKEY)
        client := horizonclient.DefaultTestNetClient
        ar := horizonclient.AccountRequest{AccountID: kp.Address()}
        sourceAccount, err := client.AccountDetail(ar)
        check(err)

        ss := si.NewStellarShell("localhost:5001", "public")

        // Adding file to IPFS
        thash, err := ss.AddIpfsString("Meow")
        if err != nil {
                error := ss.IsListening()
                panic(error) // ends where
        }
        
        op := Payment{
            Destination: DESITINATION,
            Amount:      "10",
            Asset:       NativeAsset{},
        }

        tx := Transaction{
            SourceAccount: &sourceAccount,
            Operations:    []Operation{&op},
            Timebounds:    NewInfiniteTimeout(),
            Memo:          MemoHash(thash),
            Network:       network.TestNetworkPassphrase,
        }

        txe, err := tx.BuildSignEncode(kp.(*keypair.Full))
        if err != nil {
                panic(error)
        }
        fmt.Println(txe)
}