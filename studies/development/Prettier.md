# Prettier

#javascript

## Setup lint-staged prettier husky

- `npx prettier  '**/*.{ts,js,css,html,json}' --write`
- package.json
  -- scripts add `"prepare": "husky install"`
  -- add lint-staged to package.

```json
"lint-staged": {
  "*": "prettier --ignore-unknown --write"
},
```

- add packages

```json
"prettier": "^3.1.0",
"lint-staged": "^15.2.0",
"husky": "^8.0.3",
```

- use npx husky add .husky/pre-commit "npx lint-staged"
  -- .husky/pre-commit

```bash
#!/usr/bin/env sh
. "$(dirname -- "$0")/_/husky.sh"

npx lint-staged
```
