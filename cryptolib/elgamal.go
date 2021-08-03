package cryptolib

import (
	"golang.org/x/crypto/openpgp/elgamal"
	"math/big"
	MathRand "math/rand"
	"time"
)

//获取生成元G和大质数P
func generateG() (*big.Int,*big.Int) {
	r := MathRand.New(MathRand.NewSource(time.Now().UnixNano()))
	temp := big.NewInt(0)
	p := big.NewInt(1)
	g := big.NewInt(2)
	q := generatePrime(512)
	p = temp.Add(temp.Mul(q,big.NewInt(2)),big.NewInt(1))
	for !isPrime(p) {
		q = generatePrime(512)
		p = temp.Add(temp.Mul(q,big.NewInt(2)),big.NewInt(1))
	}
	max := big.NewInt(0)
	max.Sub(p,big.NewInt(1))
	g.Rand(r,max)
	if g.Cmp(big.NewInt(3))<1 {
		generateG()
	}
	return g,p
}

//生成公钥和私钥
func GenKey() (elgamal.PublicKey, elgamal.PrivateKey){
	var publicKey elgamal.PublicKey
	var privateKey elgamal.PrivateKey
	publicKey.G, publicKey.P = generateG()
	privateKey.PublicKey=publicKey
	privateKey.X = big.NewInt(25)
	publicKey.Y = new(big.Int).Exp(privateKey.G,privateKey.X,privateKey.P)
	privateKey.PublicKey = publicKey
	return publicKey,privateKey
}
