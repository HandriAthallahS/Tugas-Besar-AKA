package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// ini struct/inisiasi buat nerima data dari inputan
type RequestData struct {
	Data   []int  `json:"data"`
	Metode string `json:"metode"`
}

// ini struct/inisiasi buat kirim hasil outputnya
type ResponseData struct {
	Median        float64 `json:"median"`
	ExecutionTime float64 `json:"executionTime"`
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
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || len(req.Data) == 0 {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}

		// ini buat catat waktu mulai buat hitung waktu eksekusinya
		start := time.Now()
		var median float64

		// ini diulang 100 rb kali biar perbedaan waktu iteratif sama rekursif keliatan
		for i := 0; i < 100000; i++ {
			if req.Metode == "rekursif" {
				median = medianRecursive(req.Data)
			} else {
				median = medianIterative(req.Data)
			}
		}

		// durasi ekseksinya dari nanodetik diubah ke milidetik
		duration := float64(time.Since(start).Nanoseconds()) / 1000000.0

		// ini buat kirim hasil median sama waktu eksekusi
		json.NewEncoder(w).Encode(ResponseData{
			Median:        median,
			ExecutionTime: duration,
		})
	})

	println("Server jalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// medianIterative itu disorting dulu pake insertion sort, kalau udah terurut semua baru cari mediannya
func medianIterative(A []int) float64 {
	var pass, i, n int
	var temp int

	// di sini data asli di salin semua ke slice baru biar slice asli ga keubah, jadi lebih gampang analisisnya
	var B []int = make([]int, len(A))
	copy(B, A)

	if len(B) == 0 {
		return 0
	}

	n = len(B)

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = B[pass]

		for i > 0 && temp < B[i-1] {
			B[i] = B[i-1]
			i = i - 1
		}

		B[i] = temp
		pass++
	}

	if n%2 == 0 {
		return float64(B[(n/2)-1]+B[n/2]) / 2
	} else {
		return float64(B[n/2])
	}

}

// medianRecursive itu disorting dulu pake merge sort, kalau udah terurut semua baru cari mediannya
func medianRecursive(A []int) float64 {
	var B []int = make([]int, len(A))
	copy(B, A)

	if len(B) == 0 {
		return 0
	}

	B = mergeSort(B)

	n := len(B)
	if n%2 == 0 {
		return float64(B[n/2-1]+B[n/2]) / 2
	}
	return float64(B[n/2])
}

// merge sort itu(bagi slice jadi 2, urutin kiri sama kanan, gabungin lagi secara terurut)
func mergeSort(arr []int) []int {
	var i, j, k int = 0, 0, 0
	if len(arr) <= 1 {
		return arr
	} else {
		//di sini slicenya dibagi jadi 2 dulu
		var mid int = len(arr) / 2
		var left []int = mergeSort(arr[:mid])
		var right []int = mergeSort(arr[mid:])

		//ini inisiasi variabel buat nanti nampung hasil urutan gabungan
		var result []int = make([]int, len(left)+len(right))

		//slice yang udah dibagi 2 tadi nanti digabungin di sini secara terurut
		for i = 0; i < len(result); i++ {
			if k >= len(left) {
				result[i] = right[j]
				j++
			} else if j >= len(right) {
				result[i] = left[k]
				k++
			} else if left[k] <= right[j] {
				result[i] = left[k]
				k++
			} else {
				result[i] = right[j]
				j++
			}
		}
		return result
	}
}
