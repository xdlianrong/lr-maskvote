package cryptolib

import (
	"crypto/rand"
	"errors"
	"io"
	"math/big"
	MathRand "math/rand"
	"time"
)

//同态加密的密文，加密的信息在c1内
type AHsecret struct {
	c1 *big.Int
	c2 *big.Int
}

type PrivateKey struct {
	PublicKey
	x *big.Int
}

type PublicKey struct {
	G, H, P, Y *big.Int
}

//生成ah-elgamal的生成元g,h和大素数P,Q
func AH_generateG() (*big.Int, *big.Int, *big.Int, *big.Int) {
	r := MathRand.New(MathRand.NewSource(time.Now().UnixNano()))
	temp := big.NewInt(0)
	p := big.NewInt(1)
	g := big.NewInt(2)
	h := big.NewInt(2)
	q := generatePrime(256)
	p = temp.Add(temp.Mul(q, big.NewInt(2)), big.NewInt(1))
	for !isPrime(p) {
		q = generatePrime(256)
		p = temp.Add(temp.Mul(q, big.NewInt(2)), big.NewInt(1))
	}
	max := big.NewInt(0)
	max.Sub(p, big.NewInt(1))
	g.Rand(r, max)
	h.Rand(r, max)
	if g.Cmp(big.NewInt(3)) < 1 && h.Cmp(big.NewInt(3)) < 1 {
		g, h, p, q = AH_generateG()
	}
	return g, h, p, q
}

//生成ah-elgamal加密的公私钥对
func GenerateAHkey(g, h, p, x *big.Int) (PublicKey, PrivateKey) {
	var pubKey PublicKey
	var priKey PrivateKey

	pubKey.G, pubKey.H, pubKey.P = g, h, p
	priKey.PublicKey = pubKey
	priKey.x = x
	pubKey.Y = new(big.Int).Exp(priKey.G, priKey.x, priKey.P)
	priKey.PublicKey = pubKey
	return pubKey, priKey
}

//ah-elgamal加密
func AH_Encrypt(random io.Reader, pub PublicKey, m *big.Int) (AHsecret, error) {
	var secret AHsecret
	r, err := rand.Int(random, pub.P)
	if err != nil {
		errors.New("生成随机数失败")
	}
	hM := big.NewInt(1)
	yR := big.NewInt(1)
	//c1 := new(big.Int).Mul(hM.Exp(pub.H, m, big.NewInt(0)),yR.Exp(pub.Y, r, big.NewInt(0)))
	hM.Exp(pub.H, m, pub.P)
	//yR.Exp(pub.Y, r, big.NewInt(0))
	yR.Exp(pub.Y, r, pub.P)
	c1 := big.NewInt(0)
	c1.Mul(hM, yR)
	c1.Mod(c1, pub.P)
	c2 := new(big.Int).Exp(pub.G, r, pub.P)
	secret.c1 = c1
	secret.c2 = c2
	return secret, nil
}

//ah-elgamal同态加法
func AH_add(s1, s2 AHsecret, P *big.Int) AHsecret {
	var s AHsecret
	c1 := new(big.Int).Mul(s1.c1, s2.c1)
	c2 := new(big.Int).Mul(s1.c2, s2.c2)
	s.c1 = new(big.Int).Mod(c1, P)
	s.c2 = new(big.Int).Mod(c2, P)
	return s
}

func AH_addC1(s []AHsecret, P *big.Int) *big.Int {
	mul := big.NewInt(1)
	for _, value := range s {
		mul.Mul(mul, value.c1)
		mul.Mod(mul, P)
	}
	return mul
}

//ah-elgamal解密
func AH_Decrypt(priv PrivateKey, secret AHsecret) (result *big.Int, err error) {
	s := new(big.Int).Exp(secret.c2, priv.x, priv.P) //g^(rx)
	if s.ModInverse(s, priv.P) == nil {
		return nil, errors.New("elgamal: invalid private key")
	}
	s.Mul(s, secret.c1)
	s.Mod(s, priv.P)
	return s, nil
}

//ah-elgamal同态解密（c2已经被各方加密）
func AH_DecryptC2(c1 *big.Int, c2 []*big.Int, P *big.Int) (result *big.Int, err error) {
	mul := big.NewInt(1)
	for _, value := range c2 {
		mul.Mul(mul, value)
		mul.Mod(mul, P)
	}
	if mul.ModInverse(mul, P) == nil {
		return nil, errors.New("ah-elgamal: invalid c2")
	}
	mul.Mul(mul, c1)
	mul.Mod(mul, P)
	return mul, nil
}
