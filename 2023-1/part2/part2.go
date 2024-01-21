package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func checkDigit(st string, i int) int{
	digits := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	if st[i] >= '0' && st[i] <= '9' {
		return int(st[i] - '0')
	}
	for j:=0; j<len(digits); j++ {
		num := digits[j]
		if len(num) <= i+1 && st[i-len(num)+1:i+1] == num {
			return j+1	
		}

	}
	return -1
}

func main(){
	f, error := os.Open("input.txt")
	check(error)
	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		st := scanner.Text()
		firstNum := -1
		secondNum := -1
		for i:=0; i < len(st); i++ {
			getNum := checkDigit(st, i)
			if getNum != -1{
				if firstNum == -1 {
					firstNum = getNum
				}
				secondNum = getNum
			}

		}
		total += secondNum + firstNum * 10
	}
	fmt.Println(total)
}
