// go-stellar-ipfs is a library that acts as a bridge between Stellar and IPFS. It combines the low fee and fast speed stellar with ipfs
package stellarshell

import (
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
// sclient is testnet client, use sclient="public" for horizon.stellar.org, else testnet for horizon-testnet.stellar.org
func NewStellarShell(hostname string, sclient string) *StellarShell {

	var S StellarShell
	shell := shi.NewShell(hostname)

	dclient := horizon.DefaultTestNetClient
	if strings.Contains(sclient, "public") {
		dclient = horizon.DefaultPublicNetClient
	}

	S.SShell = shell
	S.DClient = dclient

	return &S
}
