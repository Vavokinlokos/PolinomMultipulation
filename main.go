package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	first := "3x^4 + 3x^5 + x^2 + 1"
	second := "5x^3 + 3"
	fmt.Println("\nFirst polynom is " + first)
	fmt.Println("\nSecond polynom is " + second)
	firstMaxPow, err := GetMaxPow(first)
	if err != nil {
		fmt.Println(err)
		return
	}
	secondMaxPow, err := GetMaxPow(second)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fmt.Sprintf("\nMAX POW OF FIRST POLYNOM IS %d", firstMaxPow))
	fmt.Println(fmt.Sprintf("MAX POW OF SECOND POLYNOM IS %d", secondMaxPow))

	firstCoeffs, err := GetCoeffs(first, firstMaxPow)
	if err != nil {
		fmt.Println(err)
		return
	}

	secondCoeffs, err := GetCoeffs(second, secondMaxPow)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("\nFIRST POLYNOM'S COEFFS ARE ")
	fmt.Print(firstCoeffs)
	fmt.Print("\nSECOND POLYNOM'S COEFFS ARE ")
	fmt.Print(secondCoeffs)

	fmt.Println("\n\nResult of multiplication (coeffs) ")
	fmt.Print(multiply(firstCoeffs, secondCoeffs))
	fmt.Println("\n\nResult of multiplication (polynom) ")
	printPolynom(multiply(firstCoeffs, secondCoeffs))
}

func multiply(first, second []float64) []float64 {
	length := len(first) + len(second) - 1
	result := make([]float64, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, 0)
	}
	for i, vf := range first {
		for j, vs := range second {
			result[i+j] += vf * vs
		}
	}
	return result
}

func printPolynom(polynom []float64) {
	reversePolynom := reverseSlice(polynom)
	for i, v := range reversePolynom {
		if i == 0 {
			continue
		}
		if v == 0 {
			continue
		}
		fmt.Print(v)
		if i == len(polynom)-1 {
			fmt.Print("x ")
			continue
		} else {
			if i != len(polynom) {
				fmt.Print("x^", len(polynom)-i, " ")
			}
			fmt.Print(" + ")
		}
	}
}

func reverseSlice(s []float64) []float64 {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func GetCoeffs(input string, maxPow int64) ([]float64, error) {
	splittedInput := strings.Split(input, " + ")
	result := make([]float64, 0, maxPow+1)
	for i := 0; i < int(maxPow+1); i++ {
		result = append(result, 0)
	}
	for _, v := range splittedInput {
		if strings.Contains(v, "x") {
			pow, _ := strconv.ParseInt(v[len(v)-1:], 10, 64)
			coefs := strings.Split(v, "x")
			if coefs[0] == "" {
				result[int(pow)] = 1
				continue
			}
			coef, err := strconv.ParseFloat(coefs[0], 64)
			if err != nil {
				return nil, errors.New("invalid polynom's coeff " + coefs[0])
			}
			result[int(pow)] = coef
		} else {
			coef, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, errors.New("invalid polynom's coeff " + v)
			}
			result[0] = coef
		}
	}
	return result, nil
}

func GetMaxPow(input string) (int64, error) {
	splittedInput := strings.Split(input, " + ")
	pows := make([]int64, 0, len(splittedInput))
	for _, v := range splittedInput {
		pow, err := strconv.ParseInt(v[len(v)-1:], 10, 64)
		if err != nil {
			return 0, errors.New("invalid pow in polynom")
		}
		pows = append(pows, pow)
	}
	return findMax(pows), nil
}

func findMax(a []int64) (max int64) {
	max = a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}
