package voter

import (
	"MaskVote/DKG"
	"MaskVote/cryptolib"
	"golang.org/x/crypto/openpgp/elgamal"
	"math/big"
)

type Voter struct {
	//voter的身份信息
	id     int
	priKey elgamal.PrivateKey
	pubKey elgamal.PublicKey

	//投票基本信息
	t, n         int
	qual         []int      //诚实节点列表
	g, h         *big.Int   //生成元
	p, q         *big.Int   //循环群
	pubKeyOthers []*big.Int //其它voter的公钥

	//生成分布式密钥
	generator DKG.Generator
	verifier  []DKG.Verifier
}

func (voter *Voter)New(id int) {
	voter.id = id
	voter.pubKey, voter.priKey = cryptolib.GenKey()
}

//获取投票信息
func (voter *Voter)GetVoteInfo(t,n int, g,h *big.Int, p,q *big.Int, pubKeyOthers []*big.Int) {
	voter.t = t
	voter.n = n
	voter.g = g
	voter.h = h
	voter.p = p
	voter.q = q
	voter.pubKeyOthers = pubKeyOthers
}

//根据投票信息进行初始化
func (voter *Voter)Setup() {
	for i:=0;i<voter.n;i++ {
		voter.qual = append(voter.qual, 0)
	}
	voter.generator.New(voter.t, voter.n, voter.g, voter.h, voter.p, voter.q)
	voter.generator.GetSecret()
}

//公布承诺，将秘密发送给sj
func (voter *Voter)PublishSecret() {

}

