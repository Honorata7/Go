package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//  Declaring path variable
	// example one absolutePath := "//Users/honorata/go/src/Go/todoList/todo.txt"
	var absolutePath string

	// ANSI codes for different colors
	grey := "\033[90m"
	yellow := "\033[33m"
	green := "\033[32m"
	red := "\033[31m"
	blue := "\033[34m"
	resetColor := "\033[0m"

	//Starting program
	fmt.Println("This is a simple todo list program. Please enter the path to your todo list file. If you don't want to see the colors, please add --nocolor flag at the end of the command.")

	//read command line arguments
	flag.Parse()
	if flag.NArg() > 0 {
		absolutePath = flag.Arg(0)
		colors := flag.Arg(1)
		fmt.Println(absolutePath)
		if colors == "--nocolor" {
			grey = resetColor
			yellow = resetColor
			green = resetColor
			red = resetColor
			blue = resetColor
		}
	}

	// Check if the file has a .txt extension
	if filepath.Ext(absolutePath) != ".txt" {
		fmt.Println("Error: Not a text file.")
		return
	}

	// Check if the user has read access to the file
	_, err := os.Stat(absolutePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Error: File does not exist.")
		} else {
			fmt.Println("Error:", err)
		}
		return
	}

	// Read the file
	file, err := os.Open(absolutePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read each line and extract the first three characters
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) >= 3 {
			firstThree := line[:3]
			if firstThree[0] == '#' {
				fmt.Println(grey, line, resetColor)
				continue
			}
			switch firstThree {
			case "[ ]":
				fmt.Println(yellow, line, resetColor)
			case "[x]":
				fmt.Println(green, line, resetColor)
			case "[+]":
				fmt.Println(blue, line, resetColor)
			case "[-]":
				fmt.Println(red, line, resetColor)
			default:
				fmt.Println("Error: Invalid status.")
			}
		} else {
			fmt.Println("Error: Invalid status.")
		}
	}
}
