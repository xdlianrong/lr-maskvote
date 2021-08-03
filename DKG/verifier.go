package DKG

import (
	"MaskVote/ZKP"
	"MaskVote/common"
	"math/big"
)

//verifier是对某个节点进行验证的，因此对于一个player而言，应该有n-1个verifier
type Verifier struct {
	//阈值t
	t int

	//player对应的编号
	j int

	//收到的s
	s  *big.Int
	s1 *big.Int

	//收到的关于sij的承诺
	commit []*big.Int

	//收到的关于系数的承诺
	Aij *big.Int

	//生成元和大质数
	g, h *big.Int
	p, q *big.Int
}

func (v *Verifier) New(g, h *big.Int, p, q *big.Int) {
	v.g = g
	v.h = h
	v.p = p
	v.q = q
}

func (v *Verifier) getSecret(s, s1 *big.Int, commit []*big.Int) {
	v.s = s
	v.s1 = s1
	v.commit = commit
}

func (v *Verifier) verifySecret() bool {
	return ZKP.Commit(v.g, v.h, v.s, v.s1, v.p).Cmp(ZKP.CommitPower(v.commit, common.GenerateYn(big.NewInt(int64(v.j)), int64(v.t)+1), v.p)) == 0
}

func (v *Verifier) getAik(Aij *big.Int) {
	v.Aij = Aij
}

func (v *Verifier) verifyFactor() bool {
	return new(big.Int).Exp(v.g, v.s, v.p).Cmp(ZKP.CommitSinglePower(v.Aij, common.GenerateYn(big.NewInt(int64(v.j)),int64(v.t)+1),v.p)) == 0
}
