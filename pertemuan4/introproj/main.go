package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// Sistem Manajemen Pesanan Restoran

type Pesanan struct {
	ID string
	Menu string
	Meja int
	Jumlah int
}

//TODO:
//Pesanan dapat dimasukkan secara draft dan banyak
//Data Pesanan disimpan dalam bentuk JSON
var ListPesanan []Pesanan

func TambahPesanan() {
	inputanUser := bufio.NewReader(os.Stdin)

	// menuPelanggan := ""
	mejaPelanggan := 0
	jumlahPesananPelanggan := 0

	fmt.Println("=================================")
	fmt.Println("Tambah Pesanan")
	fmt.Println("=================================")

	draftPesanan := []Pesanan{}

	for {

		fmt.Print("Silahkan Masukkan Menu : ")
	
		menuPelanggan, err := inputanUser.ReadString('\r')
	
		// _, err := fmt.Scanln(&menuPelanggan)
	
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}
	
		menuPelanggan = strings.Replace(menuPelanggan, "\n", "", 1)
		menuPelanggan = strings.Replace(menuPelanggan, "\r", "", 1)
	
		fmt.Print("Silahkan Masukkan Meja : ")
		_, err = fmt.Scanln(&mejaPelanggan)
	
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}
	
		fmt.Print("Silahkan Masukkan Jumlah : ")
		_, err = fmt.Scanln(&jumlahPesananPelanggan)
	
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}
	
		// Simpan ID dan Tanggal
		draftPesanan = append(draftPesanan, Pesanan{
			ID: fmt.Sprintf("PSN-%d", time.Now().Unix()),
			Menu: menuPelanggan,
			Meja: mejaPelanggan,
			Jumlah: jumlahPesananPelanggan,
		})
	
		var pilihanMenuPesanan = 0
		fmt.Println("Ketik 1 untuk tambah pesanan, ketik 0 untuk keluar ")
		_, err = fmt.Scanln(&pilihanMenuPesanan)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		if pilihanMenuPesanan == 0 {
			break
		}
	}

	fmt.Println("Menambah Pesanan...")
	_ = os.Mkdir("pesanan", 0777)

	ch := make(chan Pesanan)

	wg := sync.WaitGroup{}

	jumlahPelayan := 5 

	// Menjalankan receiver/pemroses data
	for i := 0; i < jumlahPelayan; i++ {
		wg.Add(1)
		go simpanPesanan(ch, &wg, i)
	}

	// Mengirimkan data ke channel
	for _, pesanan := range draftPesanan {
		ch <- pesanan
	}

	close(ch)

	wg.Wait()

	fmt.Println("Berhasil Menambah Pesanan!")
}

func simpanPesanan(ch <- chan Pesanan, wg *sync.WaitGroup, noPelayan int)  {
	for pesanan := range ch {
		dataJson, err := json.Marshal(pesanan)
		if err != nil {
			fmt.Println("Terjadi error: ", err)
		}

		err = os.WriteFile(fmt.Sprintf("pesanan/%s.json",pesanan.ID), dataJson, 0644)
		if err != nil {
			fmt.Println("Terjadi error: ", err)
		}

		fmt.Printf("Pelayan No %d Memproses Pesanan ID : %s!\n", noPelayan, pesanan.ID)
	}

	wg.Done()
}

func lihatPesanan(ch <- chan string, chPesanan chan Pesanan, wg *sync.WaitGroup)  {
	var pesanan Pesanan
	for idPesanan := range ch {
		dataJSON, err := os.ReadFile(fmt.Sprintf("pesanan/%s", idPesanan))
		if err != nil {
			fmt.Printf("Terjadi error :", err)
		}

		err = json.Unmarshal(dataJSON, &pesanan)
		if err != nil {
			fmt.Printf("Terjadi error :", err)
		}

		chPesanan <- pesanan
	}
	wg.Done()
}

func LiatPesanan()  {
	fmt.Println("=================================")
	fmt.Println("Lihat Pesanan")
	fmt.Println("=================================")
	fmt.Println("Memuat data ...")

	ListPesanan = []Pesanan{}

	listJsonPesanan, err := os.ReadDir("pesanan")
	if err != nil {
		fmt.Println("Terjadi error: ",err)
	}

	wg := sync.WaitGroup{}

	ch := make(chan string)
	chPesanan := make(chan Pesanan, len(listJsonPesanan))

	jumlahPelayan := 5

	for i := 0; i < jumlahPelayan; i++ {
		wg.Add(1)
		go lihatPesanan(ch, chPesanan, &wg)
	}

	for _, filePesanan := range listJsonPesanan {
		ch <- filePesanan.Name()
	}

	close(ch)
	
	wg.Wait()

	close(chPesanan)

	for dataPesanan := range chPesanan {
		ListPesanan = append(ListPesanan, dataPesanan)
	}

	for urutan, pesanan := range ListPesanan {
		fmt.Printf("%d. Nama Menu : %s, Meja : %d\n", 
		urutan+1, 
		pesanan.Menu, 
		pesanan.Meja)
	}
}

func HapusPesanan()  {
	var urutanPesanan int

	fmt.Println("=================================")
	fmt.Println("Hapus Pesanan")
	fmt.Println("=================================")
	LiatPesanan()
	fmt.Println("=================================")
	fmt.Print("Masukkan Urutan Pesanan : ")
	_, err := fmt.Scanln(&urutanPesanan)
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}

	if (urutanPesanan - 1) < 0 || (urutanPesanan - 1) > len(ListPesanan) {
		fmt.Println("Urutan Pesanan Tidak Sesuai")
		HapusPesanan()
		return
	}

	err = os.Remove(fmt.Sprintf("pesanan/%s.json", ListPesanan[urutanPesanan-1].ID))
	if err != nil {
		fmt.Printf("Terjadi error: ",err)
	}

	fmt.Println("Pesanan Berhasil Dihapus!")
}

func main()  {
	// pilihanMenu := 0
	var pilihanMenu int

	fmt.Println("=================================")
	fmt.Println("Sistem Manajemen Pesanan Restoran")
	fmt.Println("=================================")
	fmt.Println("Silahkan Pilih : ")
	fmt.Println("1. Tambah Pesanan")
	fmt.Println("2. Lihat Pesanan")
	fmt.Println("3. Hapus Pesanan")
	fmt.Println("4. Keluar")
	fmt.Println("=================================")
	fmt.Print("Masukkan Pilihan : ")
	_, err := fmt.Scanln(&pilihanMenu)
	if err != nil {
		fmt.Println("Terjadi error : ", err)
	}

	switch pilihanMenu {
	case 1:
		TambahPesanan()
	case 2:
		LiatPesanan()
	case 3:
		HapusPesanan()
	case 4:
		os.Exit(0)
	}
	main()
}