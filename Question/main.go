package main

import "fmt"

func main() {
	var answer [5][5]string

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			for j := 0; j < 5; j++ {
				if i == 0 && j == 0 {
					answer[i][j] = "North"
				} else if i%2 == 0 && j > 0 {
					answer[i][j] = "East"
				} else {
					answer[i][j] = "South"
				}
				PrintArray(answer)
			}
		} else {
			for j := 4; j >= 0; j-- {
				if i%2 == 1 && j == 4 {
					answer[i][j] = "South"
				} else {
					answer[i][j] = "West"
				}
				PrintArray(answer)
			}
		}
		fmt.Println()
	}
	fmt.Println()
	PrintArray(answer)
}

func PrintArray(answer [5][5]string) {
	for i := 0; i < len(answer); i++ {
		for j := 0; j < len(answer[i]); j++ {
			fmt.Printf("%s\t", answer[i][j])
		}
		fmt.Println()
	}
}
