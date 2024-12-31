#linux
## Create the desktop file 
 can be copied copied the existing `/usr/share/applications/nvim.desktop` 
   to 
   `~/.local/share/applications/nvim-wrap.desktop` 
   and modified it to have the endresult:
```desktop
[Desktop Entry]
Name=Neovim Wrapper
TryExec=/usr/bin/nvim-wrap %F
Exec=/usr/bin/nvim-wrap %F
Terminal=true
Type=Application
Icon=nvim
Categories=Utility;TextEditor;
StartupNotify=false
```

## Install
`desktop-file-install --dir=$HOME/.local/share/applications ~/app.desktop`

---
## links
[[Set default application]]

## references
https://wiki.archlinux.org/title/Desktop_entries