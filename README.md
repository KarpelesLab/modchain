# ModChain

Modular Blockchain System.

Status: in development.

## Goals

The main goal is to allow easy prototyping of blockchain based projects by
making available various components as modules that can be easily replaced or
customized.

The initial development cycle will aim at a simple compatibility with Bitcoin
on the network side. From that point more modules will be added, such as
smart contract system, etc. Note that compatibility with Bitcoin will be aimed
at a purely protocol level, and on disk files will most likely not be
compatible, as Go provide different tools than what Bitcoin was made with. For
example, instead of using BerkeleyDB we will prefer a Go native solution such
as boltdb.

## Module prefixes

Each part has a specific module prefix. More prefixes will be added as
development progresses.

| letter | object | description |
|:-------|:-------|:------------|
| d | [`base.Discovery`](https://godoc.org/github.com/KarpelesLab/modchain/base#Discovery) | discovery of p2p nodes |
| n | [`base.Network`](https://godoc.org/github.com/KarpelesLab/modchain/base#Network) | networking between nodes |
| r | | RPC |
| w | | Wallet |

## Code Policy

* Avoid dependency on third party packages as much as possible.
* Avoid using CGO.
