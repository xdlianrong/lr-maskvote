package cryptolib

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

//生成随机素数
func generatePrime(bits int) *big.Int {
	prime,err := rand.Prime(rand.Reader,bits)
	if err!= nil {
		fmt.Println(err)
		return nil
	}
	if !isPrime(prime) {
		return generatePrime(bits)
	}
	return prime
}

//验证数字是否为素数
func isPrime(num *big.Int) bool {
	return num.ProbablyPrime(10)
}

//查找最大的质因子
func getPrimeFactor (num *big.Int) *big.Int {
	max := big.NewInt(0)
	sqrt := new(big.Int).Sqrt(num)
	for i := big.NewInt(2);i.Cmp(sqrt)<=0;i.Add(i,big.NewInt(1)) {
		for new(big.Int).Mod(num, i).Cmp(big.NewInt(0)) == 0 {
			max = i
			num = new(big.Int).Div(num, i)
		}
		if num.Cmp(big.NewInt(1))==0 {
			break
		}
	}
	if num.Cmp(big.NewInt(1))!=0 {
		max =num
	}
	return max
}