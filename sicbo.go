package ga

import (
	"math/rand"
	"time"
)

type Sicbo struct {
	r     *rand.Rand
	dices [3]int64

	smallBig [2]int64  // small, big
	single   [6]int64  // 1, 2, 3, 4, 5, 6
	twoComb  [15]int64 // 12, 13, 14, 15, 16, 23, 24, 25, 26, 34, 35, 36, 45, 46, 56
	sum      [14]int64 // 4 ~ 17
	double   [6]int64  // 11, 22, 33, 44, 55, 66
	leopard  [7]int64  // all, 111, 222, 333, 444, 555, 666
}

func NewSicbo() *Sicbo {
	src := rand.NewSource(time.Now().UnixMilli())
	r := rand.New(src)

	return &Sicbo{
		r:     r,
		dices: [3]int64{1, 2, 3},
		// odds
		smallBig: [2]int64{1, 1},
		single: [6]int64{
			1, 1, 1,
			1, 1, 1,
		},
		twoComb: [15]int64{
			5, 5, 5, 5, 5,
			5, 5, 5, 5, 5,
			5, 5, 5, 5, 5,
		},
		sum: [14]int64{
			60, 30, 17, 12, 8, 6, 6,
			6, 6, 8, 12, 17, 30, 60,
		},
		double: [6]int64{
			10, 10, 10,
			10, 10, 10,
		},
		leopard: [7]int64{
			30,
			180, 180, 180,
			180, 180, 180,
		},
	}
}

func (sb *Sicbo) Play(strategies []Strategy) {
	// play
	sb.dices[0] = int64(sb.r.Intn(6) + 1)
	sb.dices[1] = int64(sb.r.Intn(6) + 1)
	sb.dices[2] = int64(sb.r.Intn(6) + 1)

	diceSum := sb.dices[0] + sb.dices[1] + sb.dices[2]
	isLeopard := sb.dices[0] == sb.dices[1] && sb.dices[1] == sb.dices[2]

	// outcome
	for _, s := range strategies {
		smallBig, single, twoComb, sum, double, leopard := s.SicboBet()

		var profit int64 = 0

		// small big
		if !isLeopard && diceSum >= 4 && diceSum <= 10 {
			profit += smallBig[0] * (1 + sb.smallBig[0])
		} else if !isLeopard && diceSum >= 11 && diceSum <= 17 {
			profit += smallBig[1] * (1 + sb.smallBig[1])
		}

		// single
		for _, d := range sb.dices {
			profit += single[d-1] * (1 + sb.single[d-1])
		}

		// twoComb
		checkBefore := [15]bool{}
		for i := 0; i < 3; i++ {
			small := sb.dices[i]
			var big int64
			if i == 2 {
				big = sb.dices[0]
			} else {
				big = sb.dices[i+1]
			}

			if small == big {
				continue
			}
			if small > big {
				small, big = big, small
			}
			var idx int
			switch small {
			case 1:
				idx = int(big - 2)
			case 2:
				idx = 5 + int(big-3)
			case 3:
				idx = 9 + int(big-4)
			case 4:
				idx = 12 + int(big-5)
			case 5:
				idx = 14
			}
			if checkBefore[idx] {
				continue
			}
			checkBefore[idx] = true
			profit += twoComb[idx] * (1 + sb.twoComb[idx])
		}

		// sum
		if diceSum >= 4 && diceSum <= 17 {
			profit += sum[diceSum-4] * (1 + sb.sum[diceSum-4])
		}

		// leopard
		if isLeopard {
			profit += leopard[0] * (1 + sb.leopard[0])                     // all
			profit += leopard[sb.dices[0]] * (1 + sb.leopard[sb.dices[0]]) // 111, 222, 333, 444, 555, 666
		} else {
			if sb.dices[0] == sb.dices[1] {
				profit += double[sb.dices[0]-1] * (1 + sb.double[sb.dices[0]-1])
			} else {
				profit += double[sb.dices[1]-1] * (1 + sb.double[sb.dices[1]-1])
			}
		}
		s.Outcome(profit)
	}
}
