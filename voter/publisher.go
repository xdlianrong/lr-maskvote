package voter

import "math/big"

type Publisher struct {
	//生成元及循环群
	g, h *big.Int
	p, q *big.Int
}


