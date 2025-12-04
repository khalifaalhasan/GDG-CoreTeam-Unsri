package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
  SOLUSI SOAL 1: Bilangan Sakti yang Hilang
  Author: Khalifaalhasan
  Language: Go (Golang)
*/

func main() {
	// Menggunakan bufio untuk Input/Output yang sangat cepat
	// Sesuai constraint waktu maksimal 0.5 detik
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Helper function untuk membaca input string
	nextString := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	// 1. Membaca jumlah test case (t)
	if scanner.Scan() {
		tStr := scanner.Text()
		t, _ := strconv.Atoi(tStr)

		// 2. Loop sebanyak t kali
		for i := 0; i < t; i++ {
			s := nextString()
			solve(s)
		}
	}
}

func solve(s string) {
	// Cek 1: Panjang string minimal 3
	// Karena format 10^n (n minimal 2), maka string minimal "102" (3 karakter)
	if len(s) < 3 {
		fmt.Println("NO")
		return
	}

	// Cek 2: Harus diawali dengan string "10"
	if s[0:2] != "10" {
		fmt.Println("NO")
		return
	}

	// Ambil sisa string sebagai eksponen n
	nPart := s[2:]

	// Cek 3: Validasi Leading Zero
	// Angka n tidak boleh diawali '0' (misal "1002" -> n="02" -> Invalid)
	if nPart[0] == '0' {
		fmt.Println("NO")
		return
	}

	// Cek 4: Konversi ke integer dan validasi n >= 2
	n, err := strconv.Atoi(nPart)
	if err != nil {
		fmt.Println("NO")
		return
	}

	if n >= 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}