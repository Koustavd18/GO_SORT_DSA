package main

import "fmt"

func main(){
	arr := []int{12,43,6,2,435,789,76,98,1,5,56}
	merged := mergeSort(arr)	
	fmt.Println(merged)
}

func merge(a []int, b []int) []int {
	var result []int
	i:= 0
	j:= 0
	for i< len(a) && j < len(b){
		if a[i] < b[j]{
			result = append(result,a[i])
			i++
		}else{
			result = append(result, b[j])
			j++
		}
	}

	for i < len(a) {
		result = append(result,a[i])
		i++
	}

	for j < len(b) {
		result = append(result, b[j])
		j++
	}

	return result
}

func mergeSort(a []int ) []int{
	if len(a) < 2 {
		return a
	}
	mid := int(len(a)/2)
	left:= mergeSort(a[:mid])
	right :=mergeSort(a[mid:])
return merge(left,right)
}
