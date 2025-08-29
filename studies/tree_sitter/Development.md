# Treesitter Development

## Generate parser from grammer

```sh
npx tree-sitter generate
```

## Run playground locally

```sh
sudo ./node_modules/tree-sitter-cli/tree-sitter build --wasm && npx tree-sitter playground
```

- Note `sudo` is required for docker
