package DKG

import (
	"MaskVote/ZKP"
	"MaskVote/common"
	"math/big"
	"math/rand"
	"time"
)

type Generator struct {
	//规定生成密钥的阈值和规模
	t int
	n int

	//生成多项式(用数组代表系数)
	fx  []*big.Int
	fx1 []*big.Int

	//生成元和大质数p,q
	G, H *big.Int
	p, q *big.Int

	//对于系数的承诺Cik
	Commit []*big.Int

	//秘密发送给Pj的sij
	s  []*big.Int
	s1 []*big.Int

	//其它节点信息
	qual []int
	sji  []*big.Int
	sji1 []*big.Int

	//节点公私钥
	zi *big.Int
	yi *big.Int
	Y  *big.Int //联合公钥

	//系数的承诺Aik
	Aik []*big.Int
}

func (g *Generator) New(t, n int, G, H *big.Int, p, q *big.Int) {
	g.t = t
	g.n = n
	g.G = G
	g.H = H
	g.p = p
	g.q = q
}

//生成私钥以及多项式系数承诺
func (g *Generator) GetSecret() {
	g.generatePoly()
	g.commitFactor()
	g.generateSecret()
}

//生成多项式
func (g *Generator) generatePoly() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i <= g.t; i++ {
		g.fx = append(g.fx, new(big.Int).Rand(r, g.q))
		g.fx1 = append(g.fx1, new(big.Int).Rand(r, g.q))
		//g.fx[i] = new(big.Int).Rand(r, g.q)
		//g.fx1[i] = new(big.Int).Rand(r, g.q)
	}
}

//生成多项式系数的承诺
func (g *Generator) commitFactor() {
	for index, _ := range g.fx {
		g.Commit = append(g.Commit, ZKP.Commit(g.G, g.H, g.fx[index], g.fx1[index], g.p))
	}
}

//生成发送给其他参与方的secret
func (g *Generator) generateSecret() {
	g.s = append(g.s, g.fx[0])
	g.s1 = append(g.s1, g.fx1[0])
	for j := 1; j <= g.n; j++ {
		g.s = append(g.s, common.InnerProductInP(g.fx, common.GenerateYn(big.NewInt(int64(j)), int64(g.n)), g.q))
		g.s1 = append(g.s1, common.InnerProductInP(g.fx1, common.GenerateYn(big.NewInt(int64(j)), int64(g.n)), g.q))
	}
}

//根据DLP加密系数
func (g *Generator) encryptFactor() {
	for i := 0; i <= g.t; i++ {
		g.Aik = append(g.Aik, new(big.Int).Exp(g.G, g.fx[i], g.p))
	}
}

//生成节点的公私钥
func (g *Generator) generatePk() {
	g.zi = g.fx[0]
	g.yi = new(big.Int).Exp(g.G, g.zi, g.p)
}