package ga

import (
	"log"
)

type ScoreBoard struct {
	winCount    int64
	lossCount   int64
	totalProfit int64
	betHistory  []int64

	streakCounter int64
	streakLoss    int64
	streakDstr    map[int64]int64 // positive key means winning streak and vice versa.
	maxDrawdown   int64
}

func NewScoreBoard() *ScoreBoard {
	return &ScoreBoard{
		betHistory: make([]int64, 0),
		streakDstr: make(map[int64]int64),
	}
}

func (s *ScoreBoard) AddProfit(profit int64) {
	if profit < 0 {
		s.lossCount++
		if s.streakCounter > 0 {
			s.streakDstr[s.streakCounter] += 1
			s.streakCounter = -1
		} else {
			s.streakCounter--
		}
		s.streakLoss += profit
		if s.streakLoss < s.maxDrawdown {
			s.maxDrawdown = s.streakLoss
		}
	} else {
		s.winCount++
		if s.streakCounter < 0 {
			s.streakDstr[s.streakCounter] += 1
			s.streakCounter = 1
			s.streakLoss = 0
		} else {
			s.streakCounter++
		}
	}

	s.totalProfit += profit

	s.betHistory = append(s.betHistory, profit)
}

func (s *ScoreBoard) StopGame() {
	if s.streakCounter != 0 {
		s.streakDstr[s.streakCounter] += 1
		s.streakCounter = 0
	}
}

func (s *ScoreBoard) PrintResult() {
	log.Printf("win_count=%d, loss_count=%d, win_rate=%f", s.winCount, s.lossCount, float64(s.winCount)/float64(s.winCount+s.lossCount))
	log.Printf("profit=%d", s.totalProfit)
	log.Printf("max drawdown=%d", s.maxDrawdown)

	// log.Println("streak distribution")
	// for streak, count := range s.streakDstr {
	// 	log.Printf("streak=%d, count=%d", streak, count)
	// }
}
