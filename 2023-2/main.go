package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(err error){
	if err != nil {
		log.Fatal(err)
	}
}



func checkGame(st string) int{
	game := strings.Split(st, ":")
	sets := strings.Split(game[1], ";")
	maxCount := map[string]int{
		"red" : 0,
		"green" : 0,
		"blue" : 0,
	}
	for i := 0; i < len(sets); i++ {
		set := strings.Trim(sets[i], " ")
		balls := strings.Split(set, ",")
		for _, ball := range balls {
			ball = strings.Trim(ball, " ")
			ballInfo := strings.Split(ball, " ")
			ballNum, err := strconv.Atoi(ballInfo[0])
			check(err)
			ballColor := ballInfo[1]
			maxCount[ballColor] = max(maxCount[ballColor], ballNum)
		}
	}
	power := 1
	for _, v := range maxCount {
		power *= v
	}
	return power
	
}

func main(){
	f, error := os.Open("input.txt")
	check(error)
	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		st := scanner.Text()
		num := checkGame(st)
		sum += num
	}
	fmt.Println(sum)
}
