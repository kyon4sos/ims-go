package main

import "log"

func main() {
	test := []int{1, 3, 5, 2, 9, 8}
	quickSort(test, 0, len(test)-1)
	log.Println(test)
}

func quickSort(arr []int, low int, height int) {
	if low<height {
		pivot:=getPivot(arr,low,height)
		quickSort(arr,low,pivot-1)
		quickSort(arr,pivot+1,height)
	}
}

func getPivot(arr []int, low int, height int) int {
	temp:=arr[low]
	for low<height {
		for low<height && arr[height]>=temp {
			height--
		}
		arr[low] = arr[height]
		for low<height && arr[low]<=temp {
			low++
		}
		arr[height] =arr[low]
	}
	arr[low]=temp
	return low
}


