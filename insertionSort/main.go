package main

import "fmt"

func main(){
	arr := []int{234,5,65,46,2,34,654,6,45,7}
	fmt.Println(arr)
	insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(a []int) []int{
	for i := 1; i< len(a); i++ {
		for j := i-1 ; j>=0; j-- {
			if a[j]>a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
	return a
}
