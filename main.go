package main

import (
	"os"

	utils "github.com/Gabukuro/rsa/utils"
)

func main() {
	method := os.Args[1]
	keyFileName := os.Args[2]
	sourceFileName := os.Args[3]
	targetFileName := os.Args[4]

	keyFileReader, _ := utils.GetKeyFileReader(keyFileName)
	modulus, key := utils.GetKeyFromFile(keyFileReader)

	if method == "encrypt" {
		Encrypt(sourceFileName, targetFileName, modulus, key)
	} else {
		Decrypt(sourceFileName, targetFileName, modulus, key)
	}
}
