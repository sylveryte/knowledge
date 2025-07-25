# Documentations

- Go has its own format for writing comments, that are
  automagically converted into docs

## Rules

- Place the comment directly before the item being
  documented, with no blanks
- Start each line with _//_ \/\* \*\/ is descouraged
- The first word should be the symbol (type func const...)
  name. You can put A or An if you want
- Use blank line with // to make multiple paragraph

## To make it snazzier

- For preformatted content put additional space after //
- For header put #, ## doesn't work like md
- Put [ pkg ] pkg is link to another package
- To link exported symbol [ pkg.SymbolName ]
- Raw URL will be converted to link
- For a text to link to url, put [ text ] and in the end
  of paragraph put mapping as `// [text]: URL`

## Lengthy Comment

- For lengthy comments for the package, the convention is
  to put comments in a file in your package called doc.go

## Go Doc pkg_name

- `go dock package` displays the documentation for the
  specified pkg and list of the identifiers in the pkg
  - also `go doc pkg.Identifier`

## HTML preview

- use pkgsite
  - `go install golang.org/x/pkgsite/cmd/pkgsite@latest`
  - `pkgsite` => localhost:8080

---

[[Learning Go]]
