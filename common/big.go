package common

import "math/big"

//根据底数y和指数n，生成矢量y^n
func GenerateYn(y *big.Int, n int64) []*big.Int {
	var yVector []*big.Int
	var i int64 = 1
	yVector = append(yVector, big.NewInt(1))
	for ; i < n; i++ {
		yVector = append(yVector,new(big.Int).Mul(yVector[i-1], y))
	}
	return yVector
}

//计算两个矢量的内积
func InnerProductInP(a,b []*big.Int, p *big.Int) *big.Int {
	inner := big.NewInt(0)
	for index,_ := range a {
		temp := new(big.Int).Mul(a[index], b[index])
		inner.Add(inner, new(big.Int).Mod(temp, p))
	}
	return new(big.Int).Mod(inner, p)

}

//计算两个矢量相乘
func MultiVector(a, b []*big.Int, p *big.Int) []*big.Int  {
	var vector []*big.Int
	for index,_ := range a {
		temp := new(big.Int).Mul(a[index], b[index])
		vector = append(vector, new(big.Int).Mod(temp, p))
	}
	return vector
}