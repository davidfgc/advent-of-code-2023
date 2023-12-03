package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// res: 142
//   input := []string {
// 	"1abc2",
// 	"pqr3stu8vwx",
// 	"a1b2c3d4e5f",
// 	"treb7uchet",
//   }

  // res: 55108
  data, err := os.ReadFile("/Users/davidg/davidg/practice/adventofcode/2023/src/day-1-input.txt")
  if err != nil {
	fmt.Println("Error reading file")
	fmt.Println(err)
	return
  }
  input := strings.Split(string(data), "\n")

  total := 0
  for line := 0; line < len(input); line++ {
	calibrationValue := getCalibrationValue(input[line])
	fmt.Println(calibrationValue)
	total += calibrationValue
  }
  fmt.Println("total: ", total);
}

func getCalibrationValue(input string) int {
  firstDigit := -1
  lastDigit := -1
  for charIndex := 0; charIndex < len(input); charIndex++ {
	if input[charIndex] >= '0' && input[charIndex] <= '9' {
	  value := int(input[charIndex] - '0')
	  if firstDigit == -1 {
		firstDigit = value
		lastDigit = value
	  } else {
		lastDigit = value
	  }
	}
  }
  
  return firstDigit*10 + lastDigit
}