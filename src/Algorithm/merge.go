package main

import "fmt"

func mergeSort(arr []int, l , r int)  {
	if l >= r {
		return
	}
	m := (l + r) / 2
	mergeSort(arr, l, m)
	mergeSort(arr, m + 1, r)
	if arr[m] > arr[m + 1] {
		merge(arr, l, m, r)
	}
}

func merge(arr []int, l, m, r int)  {
	aux := make([]int, r - l + 1, r + 1)
	//aux := []int{}

	for i := l; i <= r; i++ {
		aux[i - l] = arr[i]
		// append 会自动扩容切片
		//aux = append(aux, arr[i])
	}

	i := l
	j := m + 1
	for k := l; k <= r; k++ {
		if i > m {
			arr[k] = aux[j - l]
			j++
		} else if j > r {
			arr[k] = aux[i - l]
			i++
		} else if aux[j - l] > aux[i - l] {
			arr[k] = aux[i - l]
			i++
		} else {
			arr[k] = aux[j - l]
			j++
		}
	}
}

func main() {
	arr := []int{11, 13, 6, 9, 4, 8, 16}
	start := 0
	end := len(arr) - 1
	mergeSort(arr, start, end)
	fmt.Println(arr)
}
