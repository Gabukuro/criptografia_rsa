package main

import (
	"bufio"
	"math/big"
	"os"

	utils "github.com/Gabukuro/rsa/utils"
)

func Decrypt(sourceFileName string, targetFileName string, modulus *big.Int, key *big.Int) {

	sourceFile, _ := os.Open(sourceFileName)
	sourceScan := bufio.NewScanner(sourceFile)

	fullEncodedText := ""

	for sourceScan.Scan() {
		line := sourceScan.Text()

		decodedChunk, _ := new(big.Int).SetString(line, 10)
		originalChunk := decodedChunk.Exp(decodedChunk, key, modulus)

		base64EncodedChunk := utils.NewBigInt(originalChunk).Text()

		fullEncodedText += base64EncodedChunk
	}

	decryptedText := utils.DecodeBase64(fullEncodedText)

	targetFile, targetFileWriter := utils.GetFileWriter(targetFileName)

	_, _ = targetFileWriter.WriteString(decryptedText)

	_ = targetFileWriter.Flush()
	_ = targetFile.Close()
}
