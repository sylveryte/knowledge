# Sylmark

Sylmark is yet another markdown lsp focusing on simple use
case of following

## Features

- [Project link](https://codeberg.org/sylveryte/sylmark)

## Plan

- [x] Learn
  - [x] Scratch educationlsp 20% implemented rest understood
  - [>] Story markdown lsp // skipped
  - [-] Lsp lib no need
- [x] Development
  - [x] [[Sylmark#MVP]]

## Choose tools

Going through list of already [implemented lsps](https://microsoft.github.io/language-server-protocol/implementors/servers/) with golang,
we get following

- Potentials
  - https://github.com/sourcegraph/jsonrpc2 3x 3m ⭐
    - https://github.com/mattn/efm-langserver/blob/master/main.go
    - https://github.com/sqls-server/sqls/blob/master/main.go
    - https://github.com/getgauge/gauge/tree/master
    - https://github.com/StyraInc/regal/blob/main/main.go
  - https://github.com/sourcegraph/go-lsp
  - https://github.com/grpc/grpc-go 18h
  - https://github.com/grpc/grpc-web
  - https://github.com/creachadair/jrpc2 2x 3d
  - https://github.com/go-language-server/jsonrpc2 2x 3y
  - https://github.com/jdbaldry/go-language-server-protocol 4y
  - https://github.com/yinfei8/jrpc2 4y

## Notes of internal working

- glinkstore

```js
const glinkstore = {
 [target] = {
  [uri] = {
      defs:loc[],
      refs:loc[]}
    }[]
  }[]
}
```
