package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("정수입력 : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}

	var isPrime bool = true
	// bug fix
	if n <= 1 { // A prime number is a natural number grater than 1 that has only
		isPrime = false
	} else {
		j := 2
		for j < n {
			if n%j == 0 {
				isPrime = false // 더하기 연산자
			}
			j++
		}
	}
	if isPrime { // == 비교 연산자
		fmt.Printf("%d is prime number", n)
	} else {
		fmt.Printf("%d is NOT Prime number", n)
	}
}
