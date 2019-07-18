package stellarshell

import (
	"bytes"
	"io"
)

// AddIpfs implements method for adding data from io Reader and returns the ipfs hash
func (s *StellarShell) AddIpfs(r io.Reader) (string, error) {

	// Adding file to IPFS
	ipfshash, err := s.SShell.Add(r)
	if err != nil {
		return "", err
	}
	return ipfshash, nil
}

// AddIpfsString implements method for adding data from a string and returns the ipfs hash
func (s *StellarShell) AddIpfsString(dstring string) (string, error) {

	// Adding file to IPFS
	bufferString := bytes.NewBufferString(dstring)

	ipfshash, err := s.SShell.Add(bufferString)
	if err != nil {
		return "", err
	}
	return ipfshash, nil
}
