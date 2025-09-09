package cmd 

import (
	"fmt"
	"strconv"
)

func add(first string, second string) (result string) {
   num1, err := strconv.ParseFloat(first, 64)
   if err != nil {
        fmt.Println("Error: Invalid first value")
		return
   }
    num2, err := strconv.ParseFloat(first, 64)
   if err != nil {
        fmt.Println("Error: Invalid first value")
		return
   }
   return fmt.Sprintf("%f", num1+num2)
}

func Subtract(first string, second string) (result string) {
   num1, err := strconv.ParseFloat(first, 64)
   if err != nil {
		fmt.Println("Error: Invalid first value")
		return
   }
    num2, err := strconv.ParseFloat(first, 64)
   if err != nil {
		fmt.Println("Error: Invalid first value")
		return
   }
   return fmt.Sprintf("%f", num1-num2)
}

func Multiply(first string, second string, shouldRoundUp bool) (result string) {
    num1, err := strconv.ParseFloat(first, 64)
    if err != nil {
        fmt.Println("Error: First value is not a decimal")
        return
    }
    num2, err := strconv.ParseFloat(second, 64)
    if err != nil {
        fmt.Println("Error: Second value is not a decimal")
        return
    }
    if shouldRoundUp {
        return fmt.Sprintf("%.2f", num1*num2)
    }
    return fmt.Sprintf("%f", num1*num2)
}
