# Library

## Pulling library fresh published error

```sh
❯ go get github.com/sylveryte/tree-sitter-sylmark@v0.1.3
go: downloading github.com/sylveryte/tree-sitter-sylmark v0.1.3
go: github.com/sylveryte/tree-sitter-sylmark@v0.1.3: verifying module: github.com/sylveryte/tree-sitter-sylmark@v0.1.3: reading https://sum.golang.org/lookup/github.com/sylveryte/tree-sitter-sylmark@v0.1.3: 404 Not Found
	server response: not found: github.com/sylveryte/tree-sitter-sylmark@v0.1.3: invalid version: unknown revision v0.1.3
```

- This happens because server doesn't have record of .3
  because _someone tried to get that version_ before it was
  published and **poisoned** cache, **wait 30mins** for
  fresh cache.
- Quick workaround is to use `GOPRIVATE`
