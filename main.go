package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// ini struct/inisiasi buat nerima data dari inputan
type RequestData struct {
	Pesan  string `json:"pesan"`
	Metode string `json:"metode"`
}

// ini struct/inisiasi buat kirim hasil outputnya
type ResponseData struct {
	Hasil         string  `json:"hasil"`
	ExecutionTime float64 `json:"executionTime"`
}

var perubahan = map[rune]rune{
	'a': 'm', 'b': 'q', 'c': 'z', 'd': 'r', 'e': 't',
	'f': 'x', 'g': 'k', 'h': 'l', 'i': 's', 'j': 'w',
	'k': 'p', 'l': 'u', 'm': 'n', 'n': 'o', 'o': 'v',
	'p': 'y', 'q': 'a', 'r': 'b', 's': 'c', 't': 'd',
	'u': 'e', 'v': 'f', 'w': 'g', 'x': 'h', 'y': 'i',
	'z': 'j',

	'0': '5', '1': '8', '2': '4', '3': '9', '4': '0',
	'5': '6', '6': '1', '7': '2', '8': '7', '9': '3',
}

func main() {
	http.HandleFunc("/proses", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			return
		}

		var req RequestData
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Pesan == "" {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}

		start := time.Now()
		var hasil string

		for i := 0; i < 100000; i++ {
			if req.Metode == "rekursif" {
				hasil = enkripsirekursif(req.Pesan)
			} else {
				hasil = enkripsiiteratif(req.Pesan)
			}
		}

		duration := float64(time.Since(start).Nanoseconds()) / 1000000.0

		json.NewEncoder(w).Encode(ResponseData{
			Hasil:         hasil,
			ExecutionTime: duration,
		})
	})

	println("Server jalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func enkripsiiteratif(teks string) string {
	var huruf []rune = []rune(teks)
	var hasil []rune = make([]rune, len(huruf))
	var i int

	for i = 0; i < len(huruf); i++ {
		var karakter rune = huruf[i]

		if karakter >= 'A' && karakter <= 'Z' {
			karakter = karakter + 32
		}

		var hasilEnkripsi rune
		var ketemu bool
		hasilEnkripsi, ketemu = perubahan[karakter]

		if ketemu {
			hasil[i] = hasilEnkripsi
		} else {
			hasil[i] = huruf[i]
		}
	}
	return string(hasil)
}

func enkripsirekursif(teks string) string {
	if len(teks) == 0 {
		return ""
	}

	var karakter rune = rune(teks[0])

	if karakter >= 'A' && karakter <= 'Z' {
		karakter = karakter + 32
	}

	var hasilEnkripsi rune
	var ketemu bool
	hasilEnkripsi, ketemu = perubahan[karakter]

	if !ketemu {
		hasilEnkripsi = karakter
	}

	return string(hasilEnkripsi) + enkripsirekursif(teks[1:])
}
