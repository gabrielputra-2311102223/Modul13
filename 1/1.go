package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk melakukan insertion sort
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

// Fungsi untuk memeriksa apakah jarak antar elemen tetap
func checkEqualSpacing(arr []int) (bool, int) {
	if len(arr) < 2 {
		return true, 0
	}

	spacing := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != spacing {
			return false, 0
		}
	}
	return true, spacing
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan bilangan bulat yang diakhiri oleh bilangan negatif:")

	var data []int

	for {
		fmt.Print("Masukkan bilangan: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		strNumbers := strings.Fields(text)
		for _, strNum := range strNumbers {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Println("Input tidak valid, silakan masukkan bilangan bulat.")
				continue
			}

			if num < 0 {
				// Akhiri input jika bilangan negatif ditemukan
				goto process
			} else {
				data = append(data, num)
			}
		}
	}

process:
	// Urutkan data
	insertionSort(data)

	// Cetak data yang sudah diurutkan
	fmt.Println("Data setelah diurutkan:", data)

	// Periksa jarak antar elemen
	equalSpacing, spacing := checkEqualSpacing(data)
	if equalSpacing {
		fmt.Printf("Data berjarak %d\n", spacing)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}