# Obsidian open md in neovim

#obsidian #neovim #linux

## open it

Create a wrapper script in `/usr/bin/nvim-wrap`
(ensure it’s executable chmod +x nvim-wrap)

```sh
#!/bin/bash
kitty -e nvim "$1" & disown
```

![[Create a desktop file]]
