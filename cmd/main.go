package main

import (
	"errors"
	"fmt"
	"sonarqube/calculator"
)

func main() {
	result := calculator.Add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	result = calculator.Multiply(4, 7)
	fmt.Printf("4 * 7 = %d\n", result)

	processValueBad(5)
	processValueGood(5)
	fmt.Println(gradeStudentGood(55))
}

// NONCOMPLIANT: Missing final else clause
func processValueBad(x int) {
	if x == 0 {
		fmt.Println("Zero")
	} else if x == 1 {
		fmt.Println("One")
	} // SonarQube will flag this as S126 violation
}

// COMPLIANT: Has final else clause
func processValueGood(x int) error {
	if x == 0 {
		fmt.Println("Zero")
		return nil
	} else if x == 1 {
		fmt.Println("One")
		return nil
	} else {
		return errors.New("unsupported value")
	}
}

// COMPLIANT: All branches return (exception to the rule)
func processValueException(x int) error {
	if x == 0 {
		return nil
	} else if x == 1 {
		return errors.New("error")
	}
	// Implicit else - this is allowed because all branches return
	return errors.New("default case")
}

// NONCOMPLIANT: Multiple else-if without final else
func gradeStudentBad(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		if score >= 80 && score < 82 {
			return "B1"
		} else if score >= 82 && score < 84 {
			return "B2"
		} else {
			switch 20 {
			case 21:
				return "HI"
			}
		}

	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} // SonarQube will flag this - what if score < 60?

	return "Unknown"
}

// COMPLIANT: Has final else clause
func gradeStudentGood(score int) string {
	switch score {
	case 90:
		fmt.Println("Good")
	}

	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F" // Explicitly handle all other cases
	}
}
