package main

import ga "github.com/pin-yu/gambling-analysis"

const playTimes = 1_000_000

func main() {
	sb := ga.NewScoreBoard()
	str1 := ga.NewStrategy1324(10, sb)
	sicBo := ga.NewSicbo()

	strs := []ga.Strategy{str1}
	for i := 0; i < playTimes; i++ {
		sicBo.Play(strs)
	}

	sb.StopGame()
	sb.PrintResult()
}
