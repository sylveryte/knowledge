# Download a file binary from arraybuffer

```typescriptreact

  downloadBinaryData(data: ArrayBuffer, dataType: string, filename: string) {
    let binaryData = [];
    binaryData.push(data);
    let downloadLink = document.createElement('a');
    downloadLink.href = window.URL.createObjectURL(new Blob(binaryData, { type: dataType }));
    if (filename) downloadLink.setAttribute('download', filename);
    document.body.appendChild(downloadLink);
    downloadLink.click();
    downloadLink?.parentNode?.removeChild(downloadLink);
  }
```
