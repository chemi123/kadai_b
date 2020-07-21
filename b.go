package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dataSetNum, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	fmt.Println(dataSetNum)

	for i := 0; i < dataSetNum; i++ {
		scanner.Scan()
		questionsNum, _ := strconv.Atoi(scanner.Text())
		fmt.Println(questionsNum)
		for j := 0; j < questionsNum; j++ {
			scanner.Scan()
			guessBullsCows := strings.Split(scanner.Text(), " ")
			guess := guessBullsCows[0]
			bulls, _ := strconv.Atoi(guessBullsCows[1])
			cows, _ := strconv.Atoi(guessBullsCows[2])
			fmt.Println(guess, bulls, cows)
		}
	}
}
