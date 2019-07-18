package stellarshell

import (
	"io"
)

// DataHash implements method for reading data from io Reader and returns stellar manage data compatible hash in byte
func (s *StellarShell) DataHash(r io.Reader) ([]byte, error) {

	var thash []byte
	// Adding file to IPFS
	ipfshash, err := s.AddIpfs(r)
	if err != nil {
		return thash, err
	}

	// Converting multiaddr hash to Stellar memo compatible
	thash, err = Ipfs2SipfsD(ipfshash)
	if err != nil {
		// invalid hash check, decoding to base58 failed
		return thash, err
	}

	return thash, nil
}

// ReadDataHash returns valid IPFS hash from a Datakey id of managedata
func (s *StellarShell) ReadDataHash(publickey string, DataKeyid string) (string, error) {

	account, err := s.DClient.LoadAccount(publickey)
	if err != nil {
		return "", err
	}

	mdatabase64 := account.Data[DataKeyid]

	mdata, err := Decodebase64(mdatabase64)
	if err != nil {
		return "", err
	}
	return string(mdata), nil
}
