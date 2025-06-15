# Sylmark

Sylmark is yet another markdown lsp focusing on simple use
case of following

## Features

### MVP

- [ ] Linking
  - [ ] Wiki style `[[2025-04-April]]]`
  - [ ] Headers within files `[[2025-04-April#Goals for this month?]]`
- [ ] File link embedding `![](./../stickers/feb.png)`
  - [ ] Image Rendering linking
- [ ] Autocompletion
  - [ ] Wiki files
  - [ ] Wiki headers

## Plan

- [ ] Learn
  - [x] Scratch educationlsp 20% implemented rest understood
  - [ ] Story markdown lsp
  - [-] Lsp lib no need
- [ ] Development
  - [ ] [[Sylmark#MVP]]

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
