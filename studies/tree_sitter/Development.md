# Treesitter Development

## Generate parser from grammer

```sh
npx tree-sitter generate
```

## Run playground locally

```sh
sudo ./node_modules/tree-sitter-cli/tree-sitter build --wasm && npx tree-sit
ter playground
```

- Note `sudo` is required for docker
