# Veracode

- Download veracode cli binary from site
- Alias `alias vc .....`

## Static Scan

- Zip and then scan

```sh
git archive -o /tmp/cassini.tar HEAD:srv/webui && vc static scan /tmp/cassini.tar --results-file /tmp/vcresult.json && rm /tmp/cassini.tar
```

- You'll get `./results.json`
