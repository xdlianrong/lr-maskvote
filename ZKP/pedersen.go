//与pedersen承诺相关的函数
//基于elgamal构建

package ZKP

import (
	"math/big"
)

//生成pedersen承诺,g,h为生成元，v,r分别为承诺值和致盲因子
func Commit(g,h *big.Int, v, r *big.Int, P *big.Int) *big.Int {
	commit := new(big.Int).Mul(new(big.Int).Exp(g, v,P),new(big.Int).Exp(h, r,P))
	commit.Mod(commit, P)
	return commit
}


//对pedersen承诺矢量进行乘方
func CommitPower(commits []*big.Int, x []*big.Int, p *big.Int) *big.Int {
	result := big.NewInt(1)
	for key,_ := range commits {
		temp := new(big.Int).Exp(commits[key], x[key], p)
		result.Mul(result, temp)
		result.Mod(result, p)
	}
	return result
}

//对单个pedersen承诺进行乘方
func CommitSinglePower(commit *big.Int, x []*big.Int, p *big.Int) *big.Int {
	result := big.NewInt(1)
	for _,value := range x {
		temp := new(big.Int).Exp(commit, value, p)
		result.Mul(result, temp)
		result.Mod(result, p)
	}
	return result
}
