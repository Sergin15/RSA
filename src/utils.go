package rsa

import (
	"math/big"
	"os"
)

func MakeFile(fileName string) (*os.File) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	return f
}

func CloseFile(f *os.File) {
	if err := f.Close(); err != nil {
		panic(err)
	}
}


func AreCoprime(a, b *big.Int) bool {
	gcd := new(big.Int).GCD(nil, nil, a, b)
	return gcd.Cmp(big.NewInt(1)) == 0
}

func FileExists(filename string) (bool) {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic("Erro ao verificar a existência do arquivo")
	}
}