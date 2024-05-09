package main 

import "fmt"

func main(){
	arr := []int{23,56,12,0,87,3,123,2,65}

	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(a []int) []int {
	r := len(a)
	if r <= 1 {
		return a
	}

	for i := r - 1; i>=0; i-- {
		noSwap := true
		for j := 0 ; j<i; j++ {
			if a[j] > a[j+1]{
				temp := a[j]
				a[j]=a[j+1]
				a[j+1]=temp
				noSwap = false
			}
		}
			if noSwap == true{
				break
			}
		}

		return a
	}
