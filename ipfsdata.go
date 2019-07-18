package stellarshell

import (
	"bytes"
	"io"
)

// IpfsData implements method for reading data from valid IPFS hash and returns data as io Reader
func (s *StellarShell) IpfsData(thash string) (io.ReadCloser, error) {

	// Get IPFS hash data
	output, err := s.SShell.Cat(thash)
	if err != nil {
		return nil, err
	}
	//Close the Reader
	output.Close()

	return output, nil
}

// IpfsDataString implements method for reading data from valid IPFS hash and returns data in string
func (s *StellarShell) IpfsDataString(thash string) (string, error) {

	// Get IPFS hash data
	output, err := s.SShell.Cat(thash)
	if err != nil {
		return "", err
	}

	// Create a paging token whose size is 4 byte
	p := make([]byte, 4)
	var buffer bytes.Buffer
	for {
		n, err := output.Read(p)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		buffer.WriteString(string(p[:n]))
	}

	//Close the Reader
	output.Close()

	return buffer.String(), nil
}
