package DKG

import (
	"MaskVote/cryptolib"
	"fmt"
	"testing"
)

func TestGenerator_New(T *testing.T) {
	var g Generator
	G, H, p, q := cryptolib.AH_generateG()
	t := 5
	n := 7
	g.New(t,n,G,H,p,q)
	g.generatePoly()
	fmt.Println(g.fx)
	fmt.Println(g.fx1)
}
