// go-stellar-ipfs is a library that acts as a bridge between Stellar and IPFS. It combines the low fee and fast speed stellar with ipfs
package stellarshell

import (
	"net/http"
	"strings"

	shi "github.com/ipfs/go-ipfs-api"
	"github.com/stellar/go/clients/horizon"
)

// StellarShell adds new capabilities to package shell that implements a remote API interface for running a stellar compatible ipfs daemon
type StellarShell struct {
	SShell  *shi.Shell
	DClient *horizon.Client
}

// NewStellarShell create a new Stellar shell for library methods
// hostname should be the host for IPFS(Use localhost:5001 for local client)
// sclient is passing Horizon client, use sclient="testnet" for https://horizon-testnet.stellar.org, use sclient="public" for https://horizon.stellar.org
// If you wish to use custom client, than pass your custom horizon client URL as string
func NewStellarShell(hostname string, sclient string) *StellarShell {

	var S StellarShell
	shell := shi.NewShell(hostname)

	var dclient *horizon.Client
	if strings.Contains(sclient, "public") {
		dclient = horizon.DefaultPublicNetClient
	} else if strings.Contains(sclient, "testnet") {
		dclient = horizon.DefaultTestNetClient
	} else {
		dclient = &horizon.Client{
			URL:  sclient,
			HTTP: http.DefaultClient,
		}
	}

	S.SShell = shell
	S.DClient = dclient

	return &S
}
