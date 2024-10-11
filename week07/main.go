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
	fmt.Print("Input your score : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	score, _ := strconv.ParseInt(i, 16, 32)
	if score >= 60 {
		fmt.Println("A")
		fmt.Println("%d \n", score)
	} else {
		fmt.Println("BCDF")
		fmt.Println("%d \n", score)
	}
}
