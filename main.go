package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`ERROR: Expected atleast 2 arguments.
			for example: qut file.qut
			`)
		os.Exit(1)
	}
	filename := os.Args[1]
	fileExtention := path.Ext(filename)

	if strings.Compare(fileExtention, ".qut") != 0 {
		fmt.Println(`ERROR: this is not a qut file.`)
		os.Exit(1)
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	var fileStrings string
	stringFields := strings.Fields(strings.Trim(string(file), fileStrings))
	instructions := make([]int, len(stringFields))
	for i, field := range stringFields {
		instructions[i], err = tokenize(field, i)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}
	tape := make([]int, 250)
	tapeCell := 0
	register := 0
	for i, instruct := range instructions {
		runInstruction(tape, &tapeCell, &register, &i, instruct)
	}
}

func runInstruction(tape []int, tapeCell *int, register *int, iterator *int, instruct int) {
	switch instruct {
	case 0:
	case 1:
		*tapeCell++
	case 2:
		*tapeCell--
	case 3:
		if tape[*tapeCell] == 3 {
			fmt.Printf("ERROR: infinite loop detected! instruction %d . \n", *iterator)
			os.Exit(1)
		}
		runInstruction(tape, tapeCell, register, iterator, tape[*tapeCell])
	case 4:
		if tape[*tapeCell] == 0 {
			var input string
			fmt.Scan(&input)
			char := int(input[0])
			tape[*tapeCell] = char
		} else {
			fmt.Println(string(rune(tape[*tapeCell])))
		}
	case 5:
		tape[*tapeCell]--
	case 6:
		tape[*tapeCell]++
	case 7:
	case 8:
		tape[*tapeCell] = 0
	case 9:
		if *register == 0 {
			*register = tape[*tapeCell]
		} else {
			tape[*tapeCell] = *register
			*register = 0
		}
	case 10:
		fmt.Println(string(rune(tape[*tapeCell])))
	case 11:
		var input string
		fmt.Scan(&input)
		char := int(input[0])
		tape[*tapeCell] = char
	}
	fmt.Printf("ERROR: instruction %d is not defined! instruction %d .\n", instruct, *iterator)
	os.Exit(1)
}

func tokenize(instruct string, i int) (int, error) {
	switch instruct {
	case "qut":
		return 0, nil
	case "qUt":
		return 1, nil
	case "quT":
		return 2, nil
	case "qUT":
		return 3, nil
	case "Qut":
		return 4, nil
	case "QUt":
		return 5, nil
	case "QuT":
		return 6, nil
	case "QUT":
		return 7, nil
	case "UUU":
		return 8, nil
	case "QQQ":
		return 9, nil
	case "TUQ":
		return 10, nil
	case "Tuq":
		return 11, nil
	}
	return 0, fmt.Errorf("ERROR: instruction %s is not defined! instruction %d .\n", instruct, i)
}
