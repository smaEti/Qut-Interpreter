package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`ERROR: Expected at least 2 arguments.
Usage: qut file.qut`)
		os.Exit(1)
	}

	filename := os.Args[1]
	fileExtention := path.Ext(filename)

	if strings.Compare(fileExtention, ".qut") != 0 {
		fmt.Println("ERROR: This is not a qut file.")
		os.Exit(1)
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	stringFields := strings.Fields(strings.TrimSpace(string(file)))
	instructions := make([]int, len(stringFields))
	for i, field := range stringFields {
		instructions[i], err = tokenize(field, i)
		debugPrinter("TOKENIZING ", i, "-", field, instructions[i])
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}

	jumpTable := make(map[int]int)
	stack := []int{}

	for i, inst := range instructions {
		if inst == 7 { // QUT
			stack = append(stack, i)
		} else if inst == 0 { // qut
			if len(stack) == 0 {
				fmt.Printf("ERROR: unmatched qut at instruction %d\n", i)
				os.Exit(1)
			}
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			jumpTable[start] = i
			jumpTable[i] = start
		}
	}

	if len(stack) > 0 {
		fmt.Printf("ERROR: %d unmatched QUT(s)\n", len(stack))
		os.Exit(1)
	}

	tape := make([]int, 10)
	tapeCell := 0
	register := 0

	for i := 0; i < len(instructions); i++ {
		debugPrinter(instructions[i])
		runInstruction(tape, &tapeCell, &register, &i, instructions[i], jumpTable)
		debugPrinter("TAPE: ", tape, "CELL: ", tapeCell, "REGISTER: ", register, "INST", i)
	}
}

func runInstruction(tape []int, tapeCell *int, register *int, iterator *int, instruct int, jumpTable map[int]int) {
	switch instruct {
	case 0:
		*iterator = jumpTable[*iterator] - 1
	case 1:
		*tapeCell--
		if *tapeCell < 0 {
			fmt.Println("ERROR: Tape pointer moved before memory start.")
			os.Exit(1)
		}
	case 2:
		*tapeCell++
		if *tapeCell >= len(tape) {
			fmt.Println("ERROR: Tape pointer moved beyond memory range.")
			os.Exit(1)
		}
	case 3:
		target := tape[*tapeCell]
		if target == 3 {
			fmt.Printf("ERROR: infinite loop detected! instruction %d.\n", *iterator)
			os.Exit(1)
		}
		runInstruction(tape, tapeCell, register, iterator, target, jumpTable)
	case 4:
		if tape[*tapeCell] == 0 {
			var input string
			fmt.Scan(&input)
			if len(input) > 0 {
				tape[*tapeCell] = int(input[0])
			}
		} else {
			fmt.Print(string(rune(tape[*tapeCell])))
		}
	case 5:
		tape[*tapeCell]--
	case 6:
		tape[*tapeCell]++
	case 7:
		if tape[*tapeCell] == 0 {
			*iterator = jumpTable[*iterator]
		}
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
		fmt.Print(string(rune(tape[*tapeCell])))
	case 11:
		var input string
		fmt.Scan(&input)
		if len(input) > 0 {
			tape[*tapeCell] = int(input[0])
		}
	default:
		fmt.Printf("ERROR: undefined instruction %d at position %d.\n", instruct, *iterator)
		os.Exit(1)
	}
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
	default:
		return 0, fmt.Errorf("ERROR: instruction %s is not defined! position %d\n", instruct, i)
	}
}

func debugPrinter(content ...any) {
	if os.Getenv("DEBUG") == "true" {
		fmt.Print("[DEBUG]", content, "\n")
	}
}
