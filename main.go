package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[2]
	flag := os.Args[1]
	fmt.Println(fileName, flag)

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	data := bufio.NewScanner(file)

	if flag == "-l" {
		printLines(data)
	} else if flag == "-w" {
		printWords(data)
	} else if flag == "-c" {
		printBytes(data)
	} else if flag == "-m" {
		printCharacters(data)
	}

}

func printLines(data *bufio.Scanner) {
	// data.Split(bufio.ScanLines)

	var lines int

	for data.Scan() {
		lines++
	}

	fmt.Println(lines)
}

func printWords(data *bufio.Scanner) {
	data.Split(bufio.ScanWords)

	var words int

	// cnt := 0

	for data.Scan() {
		words++
		// line := data.Text()
		// cnt++
		// for i := 0; i < len(line); i++ {
		// 	if line[i] == ' ' {
		// 		words++
		// 	}
		// }

		// if len(line) != 0 {
		// 	words++
		// }

		// fmt.Println(words)
		// fmt.Println(line)

		// if cnt == 2 {
		// 	break
		// }
	}

	fmt.Println(words)
}

func printBytes(data *bufio.Scanner) {
	data.Split(bufio.ScanBytes)

	var bytes int

	for data.Scan() {
		bytes++
	}

	fmt.Println(bytes)
}

func printCharacters(data *bufio.Scanner) {
	// data.Split(bufio.ScanRunes)

	var characters int

	for data.Scan() {
		line := data.Text()
		for i := 0; i < len(line); i++ {
			characters++
		}

	}

	fmt.Println(characters)
}
