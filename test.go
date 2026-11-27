package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n==========================================")
		fmt.Println("      PROGRAM KONVERSI SUHU (CELSIUS)     ")
		fmt.Println("==========================================")
		fmt.Println("1. Konversi Suhu")
		fmt.Println("2. Keluar")
		fmt.Print("Pilih menu (1/2): ")

		menuInput, _ := reader.ReadString('\n')
		menu := strings.TrimSpace(menuInput)

		if menu == "2" {
			fmt.Println("Terima kasih telah menggunakan program ini!")
			break
		} else if menu != "1" {
			fmt.Println("[Error] Pilihan tidak valid. Silakan masukkan 1 atau 2.")
			continue
		}

		fmt.Print("Masukkan nilai suhu dalam Celsius (°C): ")
		tempInput, _ := reader.ReadString('\n')
		tempStr := strings.TrimSpace(tempInput)

		celsius, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			fmt.Println("[Error] Input tidak valid! Harap masukkan angka (contoh: 36.5).")
			continue
		}

		if celsius < -273.15 {
			fmt.Println("[Error] Suhu tidak valid! Suhu tidak boleh di bawah nol mutlak (-273.15 °C).")
			continue
		}

		fahrenheit := (celsius * 9.0 / 5.0) + 32.0
		kelvin := celsius + 273.15
		reamur := celsius * 4.0 / 5.0

		fmt.Println("\n--- Hasil Konversi ---")
		fmt.Printf("Celsius    : %.2f °C\n", celsius)
		fmt.Printf("Fahrenheit : %.2f °F\n", fahrenheit)
		fmt.Printf("Kelvin     : %.2f K\n", kelvin)
		fmt.Printf("Reamur     : %.2f °R\n", reamur)
	}
}
