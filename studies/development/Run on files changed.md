# Run on files changed

#linux #testing

```bash
while inotifywait -qq -e close_write ./data || true; do sleep 0.1 && echo "===========================" && go test ./data; done
```
