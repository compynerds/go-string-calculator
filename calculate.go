package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//3+2+1=6
	//4+6-7=3
	//3-0+6=9
	//12-16+5=1
	fmt.Println(calculate("3+2+1"))
	fmt.Println(calculate("4+6-7"))
	fmt.Println(calculate("3-0+6"))
	fmt.Println(calculate("12-16+5"))
}

//This was the question
func calculate(expression string) (int, error) {
	total := 0
	expVariation1 := strings.Replace(expression, "+", ",", -1)
	expVariation2 := strings.Replace(expVariation1, "-", ",-", -1)
	values := strings.Split(expVariation2, ",")
	for _, value := range values {
		if value == "" {
			continue
		}
		value := strings.TrimSpace(value)
		num, err := strconv.Atoi(value)
		if err != nil {
			return 0, err
		}
		total += num
	}
	return total, nil
}

//This is me taking it as far as I can - WIP
func calculateOrderOfOperations(expression string) (int, error) {
	expression = strings.TrimSpace(expression)
	indexOfNegativeOrder := strings.Index(expression, "-(")
	if indexOfNegativeOrder == -1 {
		leftParaCleanedExpression := strings.Replace(expression, "(", "", -1)
		cleanedExpression := strings.Replace(leftParaCleanedExpression, ")", "", -1)
		return calculate(cleanedExpression)
	}
	expression = strings.Replace(expression, "-(", "+(-", 1) // this will only do the first fix other situations later
	indexOfNegativeOrder += 3                                // 3 because we convert -( to +(-
	runization := []rune(expression)
	for i := indexOfNegativeOrder; i < len(runization); i++ {
		if runization[i] == '-' {
			runization[i] = '+'
		} else if runization[i] == '+' {
			runization[i] = '-'
		} else if runization[i] == ')' {
			break
		}
	}
	newExpression := string(runization)
	newExpression = strings.TrimSpace(newExpression)
	leftParaCleanedExpression := strings.Replace(newExpression, "(", "", -1)
	cleanedExpression := strings.Replace(leftParaCleanedExpression, ")", "", -1)
	return calculate(cleanedExpression)
}

//WIP - do this later
func doNegativeParaConversion(expression string) string {
	return ""
}
