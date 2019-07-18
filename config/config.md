# go-stellar-ipfs Config

StellarShell requires two parameters :-
- hostname
- client

#### hostname
- hostname should be the host for IPFS(Use localhost:5001 for local client)
You can also use remote ipfs clients such as infura

- sclient is testnet client. Stellar has test network and public network. By default sclient works on testnet
Use sclient="public" for horizon.stellar.org, else testnet for horizon-testnet.stellar.org
