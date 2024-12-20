package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const nMax = 7919

type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

type DaftarBuku struct {
	Pustaka  []Buku
	nPustaka int
}

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		fields := strings.Split(input, ",")
		if len(fields) != 7 {
			fmt.Println("Input tidak valid, pastikan memasukkan 7 atribut yang dipisahkan dengan koma.")
			i--
			continue
		}
		eksemplar, _ := strconv.Atoi(strings.TrimSpace(fields[4]))
		tahun, _ := strconv.Atoi(strings.TrimSpace(fields[5]))
		rating, _ := strconv.Atoi(strings.TrimSpace(fields[6]))
		buku := Buku{
			id:        strings.TrimSpace(fields[0]),
			judul:     strings.TrimSpace(fields[1]),
			penulis:   strings.TrimSpace(fields[2]),
			penerbit:  strings.TrimSpace(fields[3]),
			eksemplar: eksemplar,
			tahun:     tahun,
			rating:    rating,
		}
		pustaka.Pustaka = append(pustaka.Pustaka, buku)
		pustaka.nPustaka++
	}
}

func CetakTerfavorit(pustaka DaftarBuku) {
	if pustaka.nPustaka == 0 {
		fmt.Println("Tidak ada buku dalam pustaka.")
		return
	}
	terfavorit := pustaka.Pustaka[0]
	for _, buku := range pustaka.Pustaka {
		if buku.rating > terfavorit.rating {
			terfavorit = buku
		}
	}
	fmt.Printf("Buku terfavorit: %s, %s, %s, %s, %d, %d, %d\n", terfavorit.id, terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.eksemplar, terfavorit.tahun, terfavorit.rating)
}

func UrutBuku(pustaka *DaftarBuku) {
	sort.Slice(pustaka.Pustaka, func(i, j int) bool {
		return pustaka.Pustaka[i].rating > pustaka.Pustaka[j].rating
	})
}

func Cetak5Terbaru(pustaka DaftarBuku) {
	if pustaka.nPustaka == 0 {
		fmt.Println("Tidak ada buku dalam pustaka.")
		return
	}
	fmt.Println("5 Buku dengan rating tertinggi:")
	for i := 0; i < 5 && i < pustaka.nPustaka; i++ {
		buku := pustaka.Pustaka[i]
		fmt.Printf("%s, %s, %s, %s, %d, %d, %d\n", buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
	}
}

func CariBuku(pustaka DaftarBuku, r int) {
	left, right := 0, pustaka.nPustaka-1
	for left <= right {
		mid := (left + right) / 2
		if pustaka.Pustaka[mid].rating == r {
			buku := pustaka.Pustaka[mid]
			fmt.Printf("Buku dengan rating %d: %s, %s, %s, %s, %d, %d, %d\n", r, buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
			return
		} else if pustaka.Pustaka[mid].rating < r {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan jumlah buku: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, _ := strconv.Atoi(input)

	DaftarkanBuku(&pustaka, n)
	CetakTerfavorit(pustaka)
	UrutBuku(&pustaka)
	Cetak5Terbaru(pustaka)

	fmt.Print("Masukkan rating yang ingin dicari: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	r, _ := strconv.Atoi(input)

	CariBuku(pustaka, r)
}
