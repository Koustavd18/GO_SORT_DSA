package main

import "fmt"

func main(){
	arr := []int{123,34,56,76,6,21,1,5,3,87,2}
	fmt.Println(arr)
	selectionSort(arr)
	fmt.Println(arr)
}

func selectionSort(a []int) []int {
	if len(a) <= 1{
		return a
	}

	for i := 0; i< len(a); i++{
		lowest := i
		for j := i+1; j<len(a);j++ {
			if a[j]<a[lowest] {
				lowest = j
			}
		}
		if i != lowest {
			temp := a[lowest]
			a[lowest] = a[i]
			a[i] = temp
		}
	}
	return a
}
