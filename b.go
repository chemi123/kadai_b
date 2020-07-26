package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type guessSet struct {
	nums  []int
	bulls int
	cows  int
}

// 方針: 数字を候補の4つまで絞りこんでから、全24パターンで位置まで見て枝狩りをする。そして最後に候補が1つになればそれを返し、そうでないならNoneという方針でできる？
// わかってないこと
// ・候補を4つまで絞るには？
// ・4つまで絞ったあとの枝狩りはどうする？
// ・他のやり方もあるのか？
func guess(sets []guessSet) ([]int, bool) {
	for _, set := range sets {
		if set.bulls == 4 {
			return set.nums, true
		}
	}
	return []int{}, false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dataSetNum, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	fmt.Println(dataSetNum)

	for i := 0; i < dataSetNum; i++ {
		scanner.Scan()
		questionsNum, _ := strconv.Atoi(scanner.Text())
		sets := make([]guessSet, 0, questionsNum)

		for j := 0; j < questionsNum; j++ {
			scanner.Scan()
			numBullsCows := strings.Split(scanner.Text(), " ")
			nums := make([]int, 0, 4)
			for _, numStr := range numBullsCows {
				num, _ := strconv.Atoi(numStr)
				nums = append(nums, num)
			}
			bulls, _ := strconv.Atoi(numBullsCows[1])
			cows, _ := strconv.Atoi(numBullsCows[2])
			sets = append(sets, guessSet{nums: nums, bulls: bulls, cows: cows})
		}
		fmt.Println(sets)
	}
}
