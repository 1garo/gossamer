# Gossamer `network` Package

This package emulates the [peer-to-peer networking capabilities](https://crates.parity.io/sc_network/index.html)
provided by the [Substrate](https://docs.substrate.io/) framework for blockchain development, which implies that it is
built on the extensible [`libp2p` networking stack](https://docs.libp2p.io/introduction/what-is-libp2p/). `libp2p`
provides implementations of a number of battle-tested peer-to-peer networking protocols (e.g. Kademlia for peer
discovery, and Yamux for multiplexing), and also makes it possible to implement the blockchain-specific protocols
defined by Substrate (e.g. authoring and finalising blocks, and maintaining the
[transaction pool](https://docs.substrate.io/v3/concepts/tx-pool/)).

## Intro to `libp2p`

- Identities
- Substreams

## Types of Protocols

- Request/Response
- Notifications

## Gossamer Network Protocols

### `libp2p`-Provided Protocols

- [Noise](http://cryptowiki.net/index.php?title=Noise_Protocol_Framework)
  - Purpose
  - Messages
  - Location of file/implementation
- [Yamux](https://docs.libp2p.io/concepts/stream-multiplexing/)
  - Purpose: multiplexing (?)
  - Messages
  - Location of file/implementation
- [Kademlia](https://en.wikipedia.org/wiki/Kademlia)
  - Purpose: peer discovery
  - Messages
  - Location of file/implementation
- Ping
- ID

### Blockchain-Specific Protocols

- Sync
  - Purpose: accessing blocks authored by peers
  - Messages: BlockRequestMessage, BlockResponseMessage ([message.go](message.go))
  - Location of file/implementation
- Light
- Transactions
  - Purpose: transmitting valid transactions to peers
  - Messages
  - Location of file/implementation
- Block Announces
  - Purpose: transmitting new blocks authored by this host to peers
  - Messages
  - Location of file/implementation
- GRANDPA
