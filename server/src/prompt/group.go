package prompt

import (
	"math/rand"
	"project/pb"
	"project/zj"
)

// Group ...
type Group struct {
	Serial   uint32
	PoolLoop []Keyword
	PoolRand []Keyword
	Negative string
}

// Shuffle ...
func (g *Group) Shuffle() {
	g.Serial = rand.Uint32()
}

// Infinity alias Loop(-1)
func (g *Group) Infinity() (ch chan *pb.Predict) {
	g.Shuffle()
	return g.Loop(-1)
}

// Loop if i < 1, infinity loop
func (g *Group) Loop(i int) (ch chan *pb.Predict) {
	zj.J(`loop`, i, `times`)
	ch = make(chan *pb.Predict)

	infinite := i < 1

	go func() {
		for {
			ch <- g.row()
			g.Serial++
			if !infinite {
				i--
				if i < 1 {
					close(ch)
					return
				}
			}
		}
	}()
	return
}

// Circle ...
func (g *Group) Circle() (ch chan *pb.Predict) {
	i := 1
	for _, v := range g.PoolLoop {
		i *= len(v)
	}
	return g.Loop(i)
}

// Minimal reproduce all samples
func (g *Group) Minimal() (ch chan *pb.Predict) {
	i := 1
	for _, v := range g.PoolLoop {
		n := len(v)
		if i < n {
			i = n
		}
	}
	return g.Loop(i)
}

func (g *Group) row() (p *pb.Predict) {
	prompt := ``
	for _, v := range g.PoolLoop {
		_, s := v.Index(g.Serial)
		prompt += s
	}
	for _, v := range g.PoolRand {
		_, s := v.Rand()
		prompt += s
	}
	p = &pb.Predict{
		Prompt:         prompt,
		NegativePrompt: g.Negative,
	}
	return
}
