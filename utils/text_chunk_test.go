package utils

import (
	"math/big"
	"testing"
)

func TestConvertBigIntToString(t *testing.T) {
	bigInt, _ := new(big.Int).SetString("3163076316033238841", 0)
	resultado := NewBigInt(bigInt).Text()
	esperado := "9Ç&å+"

	if resultado != esperado {
		t.Errorf("Esperado '%s', resultado '%s'", esperado, resultado)
	}
}

func TestConvertChunkToBigInt(t *testing.T) {
	tempStr := "Habibibibibibi"
	resultado := NewString(tempStr).BigIntValue()
	esperado, _ := new(big.Int).SetString("2137449983208248718432833294721352", 0)

	if resultado.Cmp(esperado) != 0 {
		t.Errorf("Esperado '%s', resultado '%s'", esperado, resultado)
	}
}

func TestBlockSize(t *testing.T) {
	keyFileReader, _ := GetKeyFileReader("../private.txt")
	modulus, _ := GetKeyFromFile(keyFileReader)
	resultado := BlockSize(*modulus)
	esperado := 7

	if resultado != esperado {
		t.Errorf("Esperado '%d', resultado '%d'", esperado, resultado)
	}
}
