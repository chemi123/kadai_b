package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GuessSet struct {
	NumStr string
	Bulls  int
	Cows   int
}

// 方針: 数字を候補の4つまで絞りこんでから、全24パターンで位置まで見て枝狩りをする。そして最後に候補が1つになればそれを返し、そうでないならNoneという方針でできる？
// わかってないこと
// ・候補を4つまで絞るには？
// ・4つまで絞ったあとの枝狩りはどうする？
// ・他のやり方もあるのか？
func guess(guessSets []GuessSet) (string, bool) {
	for _, guessSet := range guessSets {
		if guessSet.Bulls == 4 {
			return guessSet.NumStr, true
		}
	}
	return "None", false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dataSetNum, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	fmt.Println(dataSetNum)

	for i := 0; i < dataSetNum; i++ {
		scanner.Scan()
		questionsNum, _ := strconv.Atoi(scanner.Text())
		guessSets := make([]GuessSet, 0, questionsNum)

		for j := 0; j < questionsNum; j++ {
			scanner.Scan()
			numBullsCows := strings.Split(scanner.Text(), " ")
			bulls, _ := strconv.Atoi(numBullsCows[1])
			cows, _ := strconv.Atoi(numBullsCows[2])
			guessSets = append(guessSets, GuessSet{NumStr: numBullsCows[0], Bulls: bulls, Cows: cows})
		}
		fmt.Println(guessSets)
	}
}
