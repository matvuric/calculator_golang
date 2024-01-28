package main

import (
	"bufio"
	"fmt"
	"github.com/brandenc40/romannumeral"
	"os"
	"slices"
	"strconv"
	"strings"
)

var operators = []string{
	"+",
	"-",
	"*",
	"/",
}

func parseString(expression string) (a, b int, operator string, isRoman bool) {
	expression = strings.TrimSpace(expression)
	ops := strings.Split(expression, " ")

	if len(ops) > 3 {
		panic("В выражении не может быть больше 2 операндов и 1 оператора")
	} else if len(ops) < 3 {
		panic("Выражение не является математической операцией: " + expression)
	}

	operator = ops[1]
	if !slices.Contains(operators, operator) {
		panic("Такой оператор не поддерживается: " + operator)
	}

	a, aError := strconv.Atoi(ops[0])
	if aError != nil {
		a, aError = romannumeral.StringToInt(ops[0])
		if aError != nil {
			panic("Первый операнд некорректный: " + ops[0])
		}
		isRoman = true
	}
	if a < 1 || a > 10 {
		panic("Операнды должны быть больше 0 и меньше 11: " + ops[0])
	}

	b, bError := strconv.Atoi(ops[2])
	if bError != nil {
		b, bError = romannumeral.StringToInt(ops[2])
		if bError != nil {
			panic("Второй операнд некорректный: " + ops[2])
		}
		if !isRoman {
			panic("Операнды должны быть в одной системе: либо в арабской, либо в римской")
		}
	} else if isRoman {
		panic("Операнды должны быть в одной системе: либо в арабской, либо в римской")
	}
	if b < 1 || b > 10 {
		panic("Второй операнд некорректный: " + ops[2])
	}
	return
}

func calculate(a, b int, operator string, isRoman bool) string {
	var result int
	switch operator {
	case operators[0]:
		result = a + b
	case operators[1]:
		result = a - b
	case operators[2]:
		result = a * b
	case operators[3]:
		result = a / b
	}
	if isRoman {
		romanResult, err := romannumeral.IntToString(result)
		if err != nil {
			panic("Результат в римской системе не может быть меньше или равен 0")
		}
		return romanResult
	}
	return strconv.Itoa(result)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ввод:")
	expression, _ := reader.ReadString('\n')
	fmt.Printf("Вывод:\n%s", calculate(parseString(expression)))
}
