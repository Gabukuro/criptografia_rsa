package main

import (
	"math/big"
	"os"

	utils "github.com/Gabukuro/rsa/utils"
)

func Encrypt(sourceFileName string, targetFileName string, modulus *big.Int, key *big.Int) {

	sourceFile, _ := os.ReadFile(sourceFileName)
	sourceText := string(sourceFile)

	chunkSize := utils.BlockSize(*modulus)
	codedText := utils.EncodeBase64(sourceText)

	targetFile, targetFileWriter := utils.GetFileWriter(targetFileName)

	for _, chunk := range utils.SplitByWidth(codedText, chunkSize) {
		decodedChunk := utils.NewString(chunk).BigIntValue()
		encodedChunk := decodedChunk.Exp(decodedChunk, key, modulus)

		_, _ = targetFileWriter.WriteString(encodedChunk.Text(10) + "\n")
	}

	_ = targetFileWriter.Flush()
	_ = targetFile.Close()
}
