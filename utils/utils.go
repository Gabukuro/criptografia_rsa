package utils

import (
	"bufio"
	"encoding/base64"
	"log"
	"math/big"
	"os"
)

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func GetKeyFileReader(keyFileName string) (*bufio.Reader, error) {
	keyFile, err := os.Open(keyFileName)
	keyFileReader := bufio.NewReader(keyFile)
	return keyFileReader, err
}

func GetKeyFromFile(keyFileReader *bufio.Reader) (*big.Int, *big.Int) {
	modulusStr, _ := Readln(keyFileReader)
	modulus, _ := new(big.Int).SetString(modulusStr, 10)
	keyStr, _ := Readln(keyFileReader)
	key, _ := new(big.Int).SetString(keyStr, 10)
	return modulus, key
}

func SplitByWidth(str string, size int) []string {
	strLength := len(str)
	var split []string
	var stop int
	for i := 0; i < strLength; i += size {
		stop = i + size
		if stop > strLength {
			stop = strLength
		}
		split = append(split, str[i:stop])
	}
	return split
}

func GetFileWriter(targetFileName string) (*os.File, *bufio.Writer) {
	targetFile, _ := os.OpenFile(targetFileName, os.O_CREATE|os.O_WRONLY, 0644)
	targetFileWriter := bufio.NewWriter(targetFile)
	return targetFile, targetFileWriter
}

func EncodeBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func DecodeBase64(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatalf("Unable to decode: %v", err)
	}

	return string(data)
}
