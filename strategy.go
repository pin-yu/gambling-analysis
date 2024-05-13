package ga

type Strategy interface {
	SicboBet() ([2]int64, [6]int64, [15]int64, [14]int64, [6]int64, [7]int64)
	Outcome(profit int64)
}

type Strategy1324 struct {
	base int64

	scoreBoard *ScoreBoard
	lastBet    int64
	winLast    bool

	idx       int
	multiples [4]int64
}

func NewStrategy1324(base int64, scoreBoard *ScoreBoard) *Strategy1324 {
	return &Strategy1324{
		base:       base,
		scoreBoard: scoreBoard,
		idx:        0,
		multiples:  [4]int64{1, 3, 2, 4},
	}
}

func (s *Strategy1324) SicboBet() (smallBig [2]int64, single [6]int64, twoComb [15]int64, sum [14]int64, double [6]int64, leopard [7]int64) {
	if !s.winLast {
		s.idx = 0
	} else {
		s.idx = (s.idx + 1) % 4
	}

	smallBig[0] = s.base * s.multiples[s.idx]
	// smallBig[0] = s.base
	s.lastBet = s.base
	return
}

func (s *Strategy1324) Outcome(ret int64) {
	profit := ret - s.lastBet
	s.winLast = profit >= 0
	s.scoreBoard.AddProfit(profit)
}
