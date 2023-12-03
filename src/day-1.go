package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// res: 281
//   input := []string {
// 	"two1nine",
// 	"eightwothree",
// 	"abcone2threexyz",
// 	"xtwone3four",
// 	"4nineeightseven2",
// 	"zoneight234",
// 	"7pqrstsixteen",
//   }

  // res: 56324
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
	value := -1
	if input[charIndex] >= '0' && input[charIndex] <= '9' {
	  value = int(input[charIndex] - '0')
	} else  {
	  value = getNumberFromString(input[charIndex:])
	  fmt.Println("value: ", value)
	}

	if value == -1 {
	  continue
	}

	if firstDigit == -1 {
	  firstDigit = value
	  lastDigit = value
	} else {
	  lastDigit = value
	}
  }
  
  return firstDigit*10 + lastDigit
}

func getNumberFromString(input string) int {
  numberStrings := map[string]int {
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
  }

  for key, value := range numberStrings {
	if strings.HasPrefix(input, key) {
	  return value
	}
  }
  return -1
}