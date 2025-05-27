package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Tim struct {
	NamaTim  string
	Menang   int
	Kalah    int
	Poin     int
}

type JadwalTanding struct {
	TimSatu      string
	TimDua       string
	WaktuTanding string
}

func tambahTim(daftarTim []Tim, timBaru Tim) []Tim {
	return append(daftarTim, timBaru)
}

func updateInfoTim(daftarTim []Tim, nama string, menangBaru, kalahBaru, poinBaru int) ([]Tim, bool) {
	for i := range daftarTim {
		if daftarTim[i].NamaTim == nama {
			daftarTim[i] = Tim{nama, menangBaru, kalahBaru, poinBaru}
			return daftarTim, true
		}
	}
	return daftarTim, false
}

func hapusTim(daftarTim []Tim, nama string) ([]Tim, bool) {
	for i := range daftarTim {
		if daftarTim[i].NamaTim == nama {
			return append(daftarTim[:i], daftarTim[i+1:]...), true
		}
	}
	return daftarTim, false
}

func tampilKlasemen(daftarTim []Tim) {
	SelectionSort(daftarTim)
	fmt.Println("\n--- Klasemen Sementara ---")
	for i, tim := range daftarTim {
		fmt.Printf("%d. %s - Menang: %d, Kalah: %d, Poin: %d\n", i+1, tim.NamaTim, tim.Menang, tim.Kalah, tim.Poin)
	}
}

func SequentialSearch(daftarTim []Tim, namaYangDicari string) *Tim {
	namaYangDicari = strings.ToLower(namaYangDicari)
	for i := range daftarTim {
		if strings.ToLower(daftarTim[i].NamaTim) == namaYangDicari {
			return &daftarTim[i]
		}
	}
	return nil
}

func BinarySearch(daftarTim []Tim, namaYangDicari string) *Tim {
	InsertionSortNama(daftarTim)

	low, high := 0, len(daftarTim)-1
	namaYangDicari = strings.ToLower(namaYangDicari)
	for low <= high {
		mid := (low + high) / 2
		namaTengah := strings.ToLower(daftarTim[mid].NamaTim)
		if namaTengah == namaYangDicari {
			return &daftarTim[mid]
		} else if namaYangDicari < namaTengah {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return nil
}

func SelectionSort(daftarTim []Tim) {
	for i := range daftarTim {
		idxPoinTerbanyak := i
		for j := i + 1; j < len(daftarTim); j++ {
			if daftarTim[j].Poin > daftarTim[idxPoinTerbanyak].Poin {
				idxPoinTerbanyak = j
			}
		}
		daftarTim[i], daftarTim[idxPoinTerbanyak] = daftarTim[idxPoinTerbanyak], daftarTim[i]
	}
}

func InsertionSortNama(daftarTim []Tim) {
	for i := 1; i < len(daftarTim); i++ {
		kunci := daftarTim[i]
		j := i - 1
		for j >= 0 && strings.ToLower(daftarTim[j].NamaTim) > strings.ToLower(kunci.NamaTim) {
			daftarTim[j+1] = daftarTim[j]
			j--
		}
		daftarTim[j+1] = kunci
	}
}

func tambahJadwal(daftarJadwal []JadwalTanding, jadwalBaru JadwalTanding) []JadwalTanding {
	return append(daftarJadwal, jadwalBaru)
}

func tampilJadwal(daftarJadwal []JadwalTanding) {
	if len(daftarJadwal) == 0 {
		fmt.Println("Belum ada jadwal pertandingan nih.")
		return
	}
	fmt.Println("\n--- Jadwal Pertandingan ---")
	for i, j := range daftarJadwal {
		fmt.Printf("%d. %s lawan %s - %s\n", i+1, j.TimSatu, j.TimDua, j.WaktuTanding)
	}
}

func tampilMenu() {
	fmt.Println("\n--- Menu Utama ---")
	fmt.Println("1. Lihat Klasemen")
	fmt.Println("2. Cari Tim")
	fmt.Println("3. Tambah Tim Baru")
	fmt.Println("4. Perbarui Info Tim")
	fmt.Println("5. Hapus Tim")
	fmt.Println("6. Tambah Jadwal Pertandingan")
	fmt.Println("7. Lihat Jadwal Pertandingan")
	fmt.Println("8. Keluar")
	fmt.Print("Pilih opsi (angka aja ya): ")
}

func main() {
	semuaTim := []Tim{
		{"ONIC Esports", 9, 3, 27},
		{"EVOS Legends", 7, 5, 21},
		{"Alter Ego", 6, 6, 18},
		{"RRQ Hoshi", 10, 2, 30},
		{"Bigetron Alpha", 5, 7, 15},
	}
	var semuaJadwal []JadwalTanding
	var pilihan int
	inputBaca := bufio.NewScanner(os.Stdin)

	for {
		tampilMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilKlasemen(semuaTim)

		case 2:
			var caraCari int
			fmt.Println("\nMau cari pakai cara apa?")
			fmt.Println("1. Sequential Search")
			fmt.Println("2. Binary Search")
			fmt.Print("Pilih caranya (1 atau 2): ")
			fmt.Scan(&caraCari)

			fmt.Print("Ketik nama tim yang dicari: ")
			inputBaca.Scan()
			namaTim := inputBaca.Text()

			var timKetemu *Tim
			if caraCari == 1 {
				timKetemu = SequentialSearch(semuaTim, namaTim)
			} else if caraCari == 2 {
				timKetemu = BinarySearch(semuaTim, namaTim)
			} else {
				fmt.Println("Pilihan cara mencari tidak ada.")
				break
			}

			if timKetemu != nil {
				fmt.Printf("Tim %s ketemu! Menang: %d, Kalah: %d, Poin: %d\n", timKetemu.NamaTim, timKetemu.Menang, timKetemu.Kalah, timKetemu.Poin)
			} else {
				fmt.Println("Maaf, tim yang kamu cari tidak ditemukan.")
			}

		case 3:
			fmt.Print("Nama tim baru: ")
			inputBaca.Scan()
			namaTim := inputBaca.Text()

			var jumlahMenang, jumlahKalah, jumlahPoin int
			fmt.Print("Berapa kali menang, kalah, dan poinnya (pisahkan dengan spasi): ")
			fmt.Scan(&jumlahMenang, &jumlahKalah, &jumlahPoin)

			semuaTim = tambahTim(semuaTim, Tim{namaTim, jumlahMenang, jumlahKalah, jumlahPoin})
			fmt.Println("Tim berhasil ditambahkan! Mantap.")

		case 4:
			fmt.Print("Nama tim yang mau diupdate: ")
			inputBaca.Scan()
			namaTim := inputBaca.Text()

			var menangBaru, kalahBaru, poinBaru int
			fmt.Print("Update menang, kalah, dan poinnya (pisahkan dengan spasi): ")
			fmt.Scan(&menangBaru, &kalahBaru, &poinBaru)

			var berhasil bool
			semuaTim, berhasil = updateInfoTim(semuaTim, namaTim, menangBaru, kalahBaru, poinBaru)
			if berhasil {
				fmt.Println("Info tim berhasil diperbarui!")
			} else {
				fmt.Println("Timnya nggak ketemu, coba cek lagi namanya ya.")
			}

		case 5:
			fmt.Print("Nama tim yang mau dihapus: ")
			inputBaca.Scan()
			namaTim := inputBaca.Text()

			var berhasil bool
			semuaTim, berhasil = hapusTim(semuaTim, namaTim)
			if berhasil {
				fmt.Println("Tim berhasil dihapus.")
			} else {
				fmt.Println("Timnya nggak ada di daftar.")
			}

		case 6:
			fmt.Print("Tim Pertama: ")
			inputBaca.Scan()
			timA := inputBaca.Text()

			fmt.Print("Tim Kedua: ")
			inputBaca.Scan()
			timB := inputBaca.Text()

			fmt.Print("Kapan pertandingannya (misal: 25 Mei 2025, 19:00): ")
			inputBaca.Scan()
			waktu := inputBaca.Text()

			semuaJadwal = tambahJadwal(semuaJadwal, JadwalTanding{timA, timB, waktu})
			fmt.Println("Jadwal pertandingan berhasil ditambahkan. Siap-siap nonton!")

		case 7:
			tampilJadwal(semuaJadwal)

		case 8:
			fmt.Println("Program selesai, terimakasihh! Sampai jumpa lagi, Nadiv.")
			return

		default:
			fmt.Println("Pilihan kamu nggak ada di menu, coba lagi ya.")
		}
	}
}