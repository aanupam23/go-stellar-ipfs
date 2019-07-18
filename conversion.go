package stellarshell

import (
	"encoding/base64"

	"github.com/mr-tron/base58"
)

// Ipfs2Sipfs converts IPFS multiaddr to stellar compatible memo hash type
func Ipfs2Sipfs(ipfshash string) ([32]byte, error) {
	var thash [32]byte

	//Processing IPFS hash to compatible memo format
	dThash, err := base58.Decode(ipfshash)
	if err != nil {
		// invalid hash check, decoding to base58 failed
		return thash, err
	}

	copy(thash[:], dThash[2:])
	return thash, nil
}

// Ipfs2SipfsD converts IPFS multiaddr to stellar compatible manage data byte
func Ipfs2SipfsD(ipfshash string) ([]byte, error) {

	//Processing IPFS hash to compatible memo format
	dThash, err := base58.Decode(ipfshash)
	if err != nil {
		// invalid hash check, decoding to base58 failed
		return dThash, err
	}

	return dThash, nil
}

// Decodebase64 decodes base64 to byte
func Decodebase64(s string) ([]byte, error) {
	decodedbyte, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return decodedbyte, nil
}

// Sipfs2Ipfs Decodes Stellar memo hash to IPFS multiaddr hash
func Sipfs2Ipfs(memo []byte) string {

	var uppbyte = []byte{0x12, 0x20}
	resultbyte := append(uppbyte, memo...)
	ipfshash := base58.Encode(resultbyte)
	return ipfshash
}
