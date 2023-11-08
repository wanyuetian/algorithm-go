package main

import "fmt"

// 通过异或方式交换 不能是同一个内存地址, 否则会变成0
func swap(arr []int, i, j int) { // i != j
	if i == j {
		return
	}
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
	//arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	a, b := 101, 356
	fmt.Println(a, b)
	fmt.Println("After swap")
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)

	arr := []int{3, 1, 100}

	i, j := 0, 0

	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]

	fmt.Println(arr[i], arr[j])

	swap(arr, 0, 2)
	fmt.Println(arr[0], arr[2])
}
