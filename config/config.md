# go-stellar-ipfs Config

StellarShell requires two parameters :-
- hostname
- client

#### hostname
- hostname should be the host for IPFS(Use localhost:5001 for local client)
You can also use remote ipfs clients such as infura

- sclient is testnet client. Stellar has test network, public network and you can also run your own instance. By default sclient works on testnet
Use sclient="testnet" for https://horizon-testnet.stellar.org, use sclient="public" for https://horizon.stellar.org
If you wish to use custom client, than pass your custom horizon client URL as string
