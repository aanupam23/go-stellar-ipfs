package stellarshell

import (
	"io"
)

// MemoHash implements method for reading data from io Reader and returns stellar memo compatible hash in byte
func (s *StellarShell) MemoHash(r io.Reader) ([32]byte, error) {

	var thash [32]byte
	// Adding file to IPFS
	ipfshash, err := s.AddIpfs(r)
	if err != nil {
		return thash, err
	}

	// Converting multiaddr hash to Stellar memo compatible
	thash, err = Ipfs2Sipfs(ipfshash)
	if err != nil {
		// invalid hash check, decoding to base58 failed
		return thash, err
	}

	return thash, nil
}

// ReadMemoHash returns valid IPFS hash from a transaction id
func (s *StellarShell) ReadMemoHash(txnid string) (string, error) {

	// Start of reading transaction Memo Hash
	transaction, err := s.DClient.LoadTransaction(txnid)
	if err != nil {
		return "", err
	}

	memobase64 := transaction.Memo

	memo, err := Decodebase64(memobase64)
	if err != nil {
		return "", err
	}

	ipfshash := Sipfs2Ipfs(memo)

	return ipfshash, nil
}
