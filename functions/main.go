package main

import (
	"errors"
	"fmt"
)

// ex:1
func concat(s1, s2 string) string {
	return s1 + ", " + s2
}
func greeting(name, greeting string) {
	fmt.Println(concat(name, greeting))
}

// ex:2
func getMonthlyPrice(tier string) (float64, error) {
    basicPrice := 100.00
    premiumPrice := 150.00
    enterprisePrice := 500.00
    if tier == "basic" {
        return basicPrice, nil
    } else if tier == "premium" {
        return premiumPrice, nil
    } else if tier == "enterprise" {
        return enterprisePrice, nil
    } else {
		return 0, errors.New("invalid tier")
	}
}
func main2() {
    price, err := getMonthlyPrice("enterprise")
	if err == nil {
		fmt.Printf("Price: %.2f\n", price)
	} else {
		fmt.Println("Error:", err)
	}
}

// ex 3 ignoring return values
func calculateTotalSalary(salary, bonus, increment int) (int,int,int,int) {
	totalSalary := salary + bonus+ increment
	return totalSalary, salary,bonus,increment
}
func main3() {
	salaryOfPerson, baseSalary, _,_  := calculateTotalSalary(2200000,300000,150000)
	fmt.Println(salaryOfPerson, baseSalary)
}

// ex 3 naked return 
func calculateAge(yearOfBirth, currentYear int) (age int) {
	age = currentYear - yearOfBirth
	return 
	// return 56 : this will result into explicit return i.e even though we'd set age int, this fuction will return 56 rather than age
}

func main4() {
	age := calculateAge(2001,2024)
	fmt.Println((age))
}

func discountBasedOnAge(age uint) uint {
	if age > 0 && age < 18  {
		return 20
	}
	if age > 18 && age <35 {
		return 28
	}
	if age > 35 && age < 45 {
		return 12
	}

	return 15
}
func main() {
	discountPercent := discountBasedOnAge(38)
	fmt.Println((discountPercent))
}