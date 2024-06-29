package main

import "fmt"

func main () {
	a := []int{54,121,4,76,2,32,87,1,46,3}
	fmt.Println(a)
	sorted := quickSort(a)
	fmt.Println(sorted)
}

func quickSort(a []int) []int{
	if len(a) < 2 {
		return a
	}
	pivot := a[0]
	var left []int
	var right []int
	for i:=1;i < len(a);i++ {
		if a[i] < pivot {
			left = append(left, a[i])
		} else {
			right = append(right,a[i])
		}
	}
	rLeft := quickSort(left)
	rRight := quickSort(right)
	var result []int
	result = append(result, rLeft...)
	result = append(result,pivot)
	result = append(result,rRight...)

	return result
}
