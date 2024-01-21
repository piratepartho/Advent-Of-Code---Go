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

func main(){
	f, error := os.Open("inputP1.txt")
	check(error)
	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		st := scanner.Text()
		firstNum := -1
		secondNum := -1
		for i:=0; i < len(st); i++ {
			if st[i] >= '0' && st[i] <= '9' {
				if firstNum == -1 {
					firstNum = int(st[i] - '0') 
				}
				secondNum = int(st[i] - '0')
			}

		}
		total += secondNum + firstNum * 10
	}
	fmt.Println(total)
}
