# Neovim basics

#neovim

- Format Options
  - `h formatoptions`
  - `textwidth` is responsible for `gq` or `gw`
    wrapping/formatting
- Check set value
  - `:set tw?` or `:set` for all
- Check where it was set
  - `:verbose set tw`
- Quick mapping
  - Sometimes you want to quickly make a map and use it for the session
    `map <M-q> :echo "hi" -- alt +q` or `vmap imap`

## Mapping

- see where a map is set
- `nvim -V1` then inside nvim `:verb map`
