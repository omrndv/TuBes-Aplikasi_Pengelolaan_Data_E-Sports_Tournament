package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Tim struct {
	NamaTim string
	Menang, Kalah, Poin int
}

type JadwalTanding struct {
	TimSatu, TimDua, WaktuTanding string
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
		fmt.Printf("%d. %s - Menang: %d, Kalah: %d, Poin: %d\n", i+1,
			tim.NamaTim, tim.Menang, tim.Kalah, tim.Poin)
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
		for j >= 0 && strings.ToLower(daftarTim[j].NamaTim) >
			strings.ToLower(kunci.NamaTim) {
			daftarTim[j+1] = daftarTim[j]
			j--
		}
		daftarTim[j+1] = kunci
	}
}

func tambahJadwal(daftarJadwal []JadwalTanding, jadwalBaru JadwalTanding) []JadwalTanding {
	return append(daftarJadwal, jadwalBaru)
}

func updateJadwal(daftarJadwal []JadwalTanding, index int, timSatuBaru, timDuaBaru, waktuTandingBaru string) ([]JadwalTanding, bool) {
	if index < 0 || index >= len(daftarJadwal) {
		return daftarJadwal, false
	}
	daftarJadwal[index] = JadwalTanding{timSatuBaru, timDuaBaru, waktuTandingBaru}
	return daftarJadwal, true
}

func hapusJadwalByIndex(daftarJadwal []JadwalTanding, index int) ([]JadwalTanding, bool) {
	if index < 0 || index >= len(daftarJadwal) {
		return daftarJadwal, false
	}
	return append(daftarJadwal[:index], daftarJadwal[index+1:]...), true
}

func tampilJadwal(daftarJadwal []JadwalTanding) {
	if len(daftarJadwal) == 0 {
		fmt.Println("Belum ada jadwal pertandingan yang tersedia kak.")
		return
	}
	fmt.Println("\n--- Jadwal Pertandingan ---")
	for i, j := range daftarJadwal {
		fmt.Printf("%d. %s vs %s - %s\n", i+1, j.TimSatu, j.TimDua, j.WaktuTanding)
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
	fmt.Println("8. Perbarui Jadwal Pertandingan")
	fmt.Println("9. Hapus Jadwal Pertandingan")
	fmt.Println("10. Keluar")
	fmt.Print("Pilih opsi (angka aja ya kak): ")
}

func main() {
	semuaTim := []Tim{
		{"ONIC Esports", 9, 3, 27},
		{"EVOS Legends", 7, 5, 21},
		{"Alter Ego", 6, 6, 18},
		{"RRQ Hoshi", 10, 2, 30},
		{"Bigetron Alpha", 5, 7, 15},
	}

	semuaJadwal := []JadwalTanding{
		{"ONIC Esports", "EVOS Legends", "2 Juni 2025, 19:00"},
		{"RRQ Hoshi", "Alter Ego", "3 Juni 2025, 20:00"},
		{"Bigetron Alpha", "ONIC Esports", "4 Juni 2025, 18:30"},
		{"EVOS Legends", "RRQ Hoshi", "5 Juni 2025, 21:00"},
	}

	var pilihan int
	inputBaca := bufio.NewScanner(os.Stdin)

	for {
		tampilMenu()
		_, err := fmt.Scanln(&pilihan)
		if err != nil {
			fmt.Println("Input tidak valid, masukkan hanya angka ya.")
			inputBaca.Scan()
			continue
		}

		switch pilihan {
		case 1:
			tampilKlasemen(semuaTim)
		case 2:
			var caraCari int
			fmt.Println("\nMau cari pakai cara apa?")
			fmt.Println("1. Sequential Search")
			fmt.Println("2. Binary Search (Diurutkan)")
			fmt.Print("Pilih caranya, 1 atau 2: ")
			_, err := fmt.Scanln(&caraCari)
			if err != nil || (caraCari != 1 && caraCari != 2) {
				fmt.Println("Input tidak valid, masukkan angka aja ya dan cuman 1-2.")
				break
			}

			fmt.Print("Ketik nama tim yang mau dicari: ")
			inputBaca.Scan()
			namaTim := inputBaca.Text()

			var timKetemu *Tim
			if caraCari == 1 {
				timKetemu = SequentialSearch(semuaTim, namaTim)
			} else if caraCari == 2 {
				InsertionSortNama(semuaTim)
				timKetemu = BinarySearch(semuaTim, namaTim)
			} else {
				fmt.Println("Pilihan cara mencari tidak ada.")
				break
			}

			if timKetemu != nil {
				fmt.Printf("Tim %s ketemu! Menang: %d, Kalah: %d, Poin: %d\n", timKetemu.NamaTim, timKetemu.Menang, timKetemu.Kalah, timKetemu.Poin)
			} else {
				fmt.Println("Maaf kak, tim yang kamu cari tidak ditemukan.")
			}
		case 3:
			fmt.Print("Masukkan nama tim baru: ")
			inputBaca.Scan()
			namaTim := inputBaca.Text()

			var jumlahMenang, jumlahKalah, jumlahPoin int
			fmt.Print("Berapa kali menang, kalah, dan poinnya (pisahkan dengan spasi yaa): ")
			_, err := fmt.Scanln(&jumlahMenang, &jumlahKalah, &jumlahPoin)
			if err != nil || jumlahMenang < 0 || jumlahKalah < 0 || jumlahPoin < 0 {
				fmt.Println("Input tidak valid, masukkan angka aja ya dan harus positif.")
				break
			}

			semuaTim = tambahTim(semuaTim, Tim{namaTim, jumlahMenang, jumlahKalah, jumlahPoin})
			fmt.Println("Tim berhasil ditambahkan! Mantap.")
		case 4:
			fmt.Print("Masukkan nama tim yang mau diupdate: ")
			inputBaca.Scan()
			namaTim := inputBaca.Text()

			var menangBaru, kalahBaru, poinBaru int
			fmt.Print("Update menang, kalah, dan poinnya (pisahkan dengan spasi yaa): ")
			_, err := fmt.Scanln(&menangBaru, &kalahBaru, &poinBaru)
			if err != nil || menangBaru < 0 || kalahBaru < 0 || poinBaru < 0 {
				fmt.Println("Input tidak valid, masukkan angka aja ya dan harus positif.")
				break
			}

			var berhasil bool
			semuaTim, berhasil = updateInfoTim(semuaTim, namaTim, menangBaru, kalahBaru, poinBaru)
			if berhasil {
				fmt.Println("Info tim berhasil diperbarui! Keren.")
			} else {
				fmt.Println("Timnya nggak ketemu, coba cek lagi namanya ya.")
			}
		case 5:
			tampilKlasemen(semuaTim)
			fmt.Print("\nMasukkan nama tim yang mau dihapus: ")
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
			fmt.Println("Jadwal pertandingan berhasil ditambahkan. Siap-siap nonton! gasss")
		case 7:
			tampilJadwal(semuaJadwal)
		case 8:
			tampilJadwal(semuaJadwal)
			if len(semuaJadwal) == 0 {
				fmt.Println("Tidak ada jadwal untuk diperbarui.")
				break
			}

			var idJadwal int
			fmt.Print("Masukkan ID jadwal yang ingin diperbarui: ")
			_, err := fmt.Scanln(&idJadwal)
			if err != nil || idJadwal <= 0 || idJadwal > len(semuaJadwal) {
				fmt.Println("Input tidak valid. Masukkan ID jadwal yang ada dan berupa angka positif.")
				break
			}

			indexToUpdate := idJadwal - 1

			fmt.Print("Masukkan Tim Pertama yang baru: ")
			inputBaca.Scan()
			tim1Baru := inputBaca.Text()

			fmt.Print("Masukkan Tim Kedua yang baru: ")
			inputBaca.Scan()
			tim2Baru := inputBaca.Text()

			fmt.Print("Masukkan Waktu Pertandingan yang baru (misal: 25 Mei 2025, 19:00): ")
			inputBaca.Scan()
			waktuBaru := inputBaca.Text()

			var berhasil bool
			semuaJadwal, berhasil = updateJadwal(semuaJadwal, indexToUpdate, tim1Baru, tim2Baru, waktuBaru)
			if berhasil {
				fmt.Println("Jadwal pertandingan berhasil diperbarui!")
			} else {
				fmt.Println("ID jadwal tidak valid atau jadwal tidak ditemukan.")
			}
			tampilJadwal(semuaJadwal)
		case 9:
			tampilJadwal(semuaJadwal)
			if len(semuaJadwal) == 0 {
				fmt.Println("Tidak ada jadwal untuk dihapus.")
				break
			}

			var idJadwal int
			fmt.Print("Masukkan ID jadwal yang ingin dihapus: ")
			_, err := fmt.Scanln(&idJadwal)
			if err != nil || idJadwal <= 0 || idJadwal > len(semuaJadwal) {
				fmt.Println("Input tidak valid. Masukkan ID jadwal yang ada dan berupa angka positif.")
				break
			}

			var berhasilHapus bool
			semuaJadwal, berhasilHapus = hapusJadwalByIndex(semuaJadwal, idJadwal-1)
			if berhasilHapus {
				fmt.Println("Jadwal pertandingan berhasil dihapus.")
			} else {
				fmt.Println("ID jadwal tidak valid atau jadwal tidak ditemukan.")
			}
			tampilJadwal(semuaJadwal)
		case 10:
			fmt.Println("Program selesai, terimakasihh! Sampai jumpa lagi.")
			return
		default:
			fmt.Println("Pilihan kamu nggak ada di menu, coba lagi ya.")
		}
	}
}