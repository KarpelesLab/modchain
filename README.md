# ModChain

Modular Blockchain System.

Status: in development.

## Goals

The main goal is to allow easy prototyping of blockchain based projects by
making available various components as modules that can be easily replaced or
customized.

## Module prefixes

Each part has a specific module prefix. More prefixes will be added as
development progresses.

| letter | object | description |
|:-------|:-------|:------------|
| d | [`base.Discovery`](https://godoc.org/github.com/KarpelesLab/modchain/base#Discovery) | discovery of p2p nodes |
| n | `base.Network` | networking between nodes |
