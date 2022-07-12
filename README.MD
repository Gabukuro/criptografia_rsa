# Criptografia RSA 🔐

## 🧱 Constuído com
* [GO 1.17.7](https://go.dev/)

## ☕ Rodando aplicação

```sh
// pré definidos
$ make run-decrypt
// ou
$ make run-encrypt

// manual 
// go run *go <method> <keyFileName> <sourceFile> <targetFile>
$ go run *.go decrypt private.txt source_encoded.txt target_decoded.txt
```

## 🧪 Testando aplicação

```sh
$ make test
```
