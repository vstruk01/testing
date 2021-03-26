package sort

import "fmt"

type Sort struct{}

func New() *Sort {
	return &Sort{}
}

func (s *Sort) MergeSort(arr []int) {
	if len(arr) == 1 {
		return
	}
	s.MergeSort(arr[:len(arr)/2])
	s.MergeSort(arr[len(arr)/2:])
	tmpArr := make([]int, 0, len(arr))
	i, j := 0, len(arr)/2
	for i < len(arr)/2 && j < len(arr) {
		if arr[i] < arr[j] {
			tmpArr = append(tmpArr, arr[i])
			i++
		} else {
			tmpArr = append(tmpArr, arr[j])
			j++
		}
	}
	fmt.Println(tmpArr)
	for i < len(arr)/2 {
		tmpArr = append(tmpArr, arr[i])
		i++
	}
	fmt.Println(tmpArr)
	for j < len(arr) {
		tmpArr = append(tmpArr, arr[j])
		j++
	}
	copy(arr, tmpArr)
}

func (s *Sort) BubbleSortString(array []string) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func (s *Sort) BubbleSortInt(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}
