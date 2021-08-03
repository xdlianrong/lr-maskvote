package cryptolib

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/openpgp/elgamal"
	"math/big"
	"testing"
)

func TestGenerateRandom(t *testing.T) {
	fmt.Println(generateG())
}

func TestKeyGen(t *testing.T) {
	pubKey, priKey := GenKey()
	c1, c2, err := elgamal.Encrypt(rand.Reader, &pubKey, []byte("hello world!"))
	if err != nil {
		fmt.Println(err)
	}
	msg, err := elgamal.Decrypt(&priKey, c1, c2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(msg))
}

func TestAHElgamal(t *testing.T) {
	g, h, p, _ := AH_generateG()
	pubKey0, priKey0 := GenerateAHkey(g, h, p, big.NewInt(25))
	pubKey1, priKey1 := GenerateAHkey(g, h, p, big.NewInt(26))
	secret0, err := AH_Encrypt(rand.Reader, pubKey0, big.NewInt(10))
	if err != nil {
		fmt.Println(err)
	}
	secret1, err := AH_Encrypt(rand.Reader, pubKey1, big.NewInt(11))
	if err != nil {
		fmt.Println(err)
	}

	secretC1 := AH_addC1([]AHsecret{secret0, secret1}, p)
	c20 := new(big.Int).Exp(secret0.c2, priKey0.x, p)
	c21 := new(big.Int).Exp(secret1.c2, priKey1.x, p)
	result, err := AH_DecryptC2(secretC1, []*big.Int{c20, c21}, p)
	if err != nil {
		fmt.Println(err)
	}
	trueValue := new(big.Int).Exp(h, big.NewInt(21), p)
	fmt.Println("result: " + result.String())
	fmt.Println("trueValue: " + trueValue.String())
}

func TestFindFactor(t *testing.T) {
	p := generatePrime(256)
	p.Sub(p, big.NewInt(1))
	fmt.Println(p)
	fmt.Println(getPrimeFactor(p))
}