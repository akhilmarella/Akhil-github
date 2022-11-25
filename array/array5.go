package main

import "fmt"

func main() {
	a := [5][2]int{{0, 2}, {2, 5}, {3, 6}, {5, 2}, {4, 6}}
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("a[%d][%d]=%d\n", i, j, a[i][j])

		}
	}

	arr := [3][3]int{{1, 2, 3}, {2, 4, 5}, {3, 6, 9}}
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			fmt.Printf("arr[%d][%d]=%d\n",k,l ,arr[k][l])
		}
	}
	fmt.Println(len(arr))
	fmt.Println(arr[0])
}
