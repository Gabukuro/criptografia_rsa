package utils

import (
	"math/big"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	keyFileReader, _ := GetKeyFileReader("../assets/test.txt")
	resultado, _ := Readln(keyFileReader)
	esperado := "teste"

	if resultado != esperado {
		t.Errorf("Esperado '%s', resultado '%s'", esperado, resultado)
	}
}

func TestGetModulusFromFile(t *testing.T) {
	test, _ := GetKeyFileReader("../private.txt")
	modulus, _ := GetKeyFromFile(test)
	esperado, _ := new(big.Int).SetString("15212890864824009557", 0)

	if modulus.Cmp(esperado) != 0 {
		t.Errorf("Esperado '%s', resultado '%s'", esperado, modulus)
	}
}

func TestGetKeyFromFile(t *testing.T) {
	test, _ := GetKeyFileReader("../private.txt")
	_, key := GetKeyFromFile(test)
	esperado, _ := new(big.Int).SetString("3163076316033238841", 0)

	if key.Cmp(esperado) != 0 {
		t.Errorf("Esperado '%s', resultado '%s'", esperado, key)
	}
}

func TestEncodeBase64(t *testing.T) {
	resultado := EncodeBase64("teste")
	esperado := "dGVzdGU="

	if resultado != esperado {
		t.Errorf("Esperado '%s', resultado '%s'", esperado, resultado)
	}
}

func TestDecodeBase64ToString(t *testing.T) {
	resultado := DecodeBase64("dGVzdGU=")
	esperado := "teste"

	if resultado != esperado {
		t.Errorf("Esperado '%s', resultado '%s'", esperado, resultado)
	}
}

func TestSpliByWidth(t *testing.T) {
	resultado := SplitByWidth("testetestetesteteste", 5)
	esperado := [...]string{"teste", "teste", "teste", "teste"}

	if len(resultado) != len(esperado) {
		t.Errorf("Esperado '%s', resultado '%s'", esperado, resultado)
	}
}

func TestWriteFile(t *testing.T) {
	targetFile, targetFileWriter := GetFileWriter("../assets/test_write.txt")
	targetFileWriter.WriteString("Habibibibibibibibibibis!")
	targetFileWriter.Flush()
	targetFile.Close()

	resultado, _ := os.ReadFile("../assets/test_write.txt")
	esperado := "Habibibibibibibibibibis!"

	if esperado != string(resultado) {
		t.Errorf("Esperado '%s', resultado '%s'", esperado, resultado)
	}
}
