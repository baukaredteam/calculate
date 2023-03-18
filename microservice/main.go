package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romanMap = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

var convToRoman = []int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

const (
	LOWCOUNT  = "Строка не является математической операцией."
	HIGHCOUNT = "Формат математической операции не удовлетворяет заданию."
	SCALE     = "Нельзя испольщовать одновременно разные системы счисления"
	NEGATIVE  = "В римской системе нет отрицательных чисел."
	ZERO      = "В римской системе нет числа 0."
	RANGE     = "Калькулятор принимает на вход числа от 1 до 10"
)

func operandsError(mathTask string) {
	re := regexp.MustCompile("[+\\-*/]")
	operands := re.Split(mathTask, -1)
	if len(operands) < 2 {
		panic(LOWCOUNT)
	}
	if len(operands) > 2 {
		panic(HIGHCOUNT)
	}
}

func detectOperation(mathTask string) string {
	if strings.Contains(mathTask, "+") {
		return "+"
	} else if strings.Contains(mathTask, "-") {
		return "-"
	} else if strings.Contains(mathTask, "*") {
		return "*"
	} else {
		return "/"
	}
}

func parseRomansToInt(result int) {
	var romanElem string
	if result == 0 {
		panic(ZERO)
	} else if result < 0 {
		panic(NEGATIVE)
	} else {
		for _, elem := range convToRoman {
			for i := elem; i <= result; {
				for idx, value := range romanMap {
					if value == elem {
						romanElem += idx
						result -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanElem)
}

func calculate(a int, b int, oper string) int {
	switch oper {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

func parse(task string) {
	var romans []int
	var stringsFound int
	operandsError(task)
	mathTask := strings.Split(task, " ")
	for i, elem := range mathTask {
		if i == 1 {
			continue
		}
		_, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
		}
	}
	switch stringsFound {
	case 0:
		num1, _ := strconv.Atoi(mathTask[0])
		num2, _ := strconv.Atoi(mathTask[2])
		errCheck := num1 < 0 && num1 > 11 && num2 < 0 && num2 > 11
		if !errCheck {
			fmt.Println(calculate(num1, num2, detectOperation(mathTask[1])))
		} else {
			panic(RANGE)
		}
	case 1:
		panic(SCALE)
	case 2:
		for i, elem := range mathTask {
			if i == 1 {
				continue
			}
			if val, ok := romanMap[elem]; ok && val > 0 && val < 11 {
				romans = append(romans, val)
			} else {
				panic(RANGE)
			}

		}
		parseRomansToInt(calculate(romans[0], romans[1], detectOperation(mathTask[1])))
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите пример:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		parse(text)
	}
}