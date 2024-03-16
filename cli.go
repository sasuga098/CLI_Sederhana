package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"time"
	"path/filepath"
	"github.com/go-pdf/fpdf"
)

type Buku struct {
	KodeBuku      string
	JudulBuku     string
	Pengarang     string
	Penerbit      string
	JumlahHalaman int
	TahunTerbit   int
}

var DaftarBuku []Buku


func TambahBuku() {
    inputanUser := bufio.NewReader(os.Stdin)
    var bukuBaru Buku

    fmt.Println("=================================")
    fmt.Println("Tambah Buku Baru")
    fmt.Println("=================================")

    fmt.Print("Kode Buku: ")
    kodeBuku, _ := inputanUser.ReadString('\n')
    kodeBuku = strings.TrimSpace(kodeBuku)

    // Periksa apakah kode buku sudah digunakan
    for _, buku := range DaftarBuku {
        if buku.KodeBuku == kodeBuku {
            fmt.Println("Kode Buku sudah digunakan. Silakan gunakan kode buku lain.")
            return
        }
    }

    bukuBaru.KodeBuku = kodeBuku

    fmt.Print("Judul Buku: ")
    bukuBaru.JudulBuku, _ = inputanUser.ReadString('\n')
    bukuBaru.JudulBuku = strings.TrimSpace(bukuBaru.JudulBuku)

    fmt.Print("Pengarang: ")
    bukuBaru.Pengarang, _ = inputanUser.ReadString('\n')
    bukuBaru.Pengarang = strings.TrimSpace(bukuBaru.Pengarang)

    fmt.Print("Penerbit: ")
    bukuBaru.Penerbit, _ = inputanUser.ReadString('\n')
    bukuBaru.Penerbit = strings.TrimSpace(bukuBaru.Penerbit)

    fmt.Print("Jumlah Halaman: ")
    fmt.Scanln(&bukuBaru.JumlahHalaman)

    fmt.Print("Tahun Terbit: ")
    fmt.Scanln(&bukuBaru.TahunTerbit)

    DaftarBuku = append(DaftarBuku, bukuBaru)

    // Simpan buku ke file JSON
    fileName := fmt.Sprintf("books/book-%s.json", bukuBaru.KodeBuku)
    bukuJSON, err := json.Marshal(bukuBaru)
    if err != nil {
        fmt.Println("Error encoding book to JSON:", err)
        return
    }

    err = os.WriteFile(fileName, bukuJSON, 0644)
    if err != nil {
        fmt.Println("Error writing book JSON to file:", err)
        return
    }

    fmt.Println("Buku berhasil ditambahkan dan disimpan!")
}



func LihatSemuaBuku() {
	fmt.Println("=================================")
	fmt.Println("Daftar Buku di Perpustakaan")
	fmt.Println("=================================")
	for i, buku := range DaftarBuku {
		fmt.Printf("%d. Kode: %s, Judul: %s, Pengarang: %s, Penerbit: %s, Jumlah Halaman: %d, Tahun Terbit: %d\n",
			i+1, buku.KodeBuku, buku.JudulBuku, buku.Pengarang, buku.Penerbit, buku.JumlahHalaman, buku.TahunTerbit)
	}
}

func HapusBuku() {
    fmt.Println("=================================")
    fmt.Println("Hapus Buku")
    fmt.Println("=================================")

    LihatSemuaBuku()

    fmt.Print("Masukkan Kode Buku yang akan dihapus: ")
    inputanUser := bufio.NewReader(os.Stdin)
    kodeBuku, _ := inputanUser.ReadString('\n')
    kodeBuku = strings.TrimSpace(kodeBuku)

    for i, buku := range DaftarBuku {
        if buku.KodeBuku == kodeBuku {
            // Hapus file JSON terkait
            fileName := fmt.Sprintf("books/book-%s.json", buku.KodeBuku)
            err := os.Remove(fileName)
            if err != nil {
                fmt.Println("Error deleting book file:", err)
                return
            }

            // Hapus buku dari slice DaftarBuku
            DaftarBuku = append(DaftarBuku[:i], DaftarBuku[i+1:]...)
            fmt.Println("Buku berhasil dihapus!")
            return
        }
    }

    fmt.Println("Buku dengan Kode Buku tersebut tidak ditemukan.")
}

func EditBuku() {
    fmt.Println("=================================")
    fmt.Println("Edit Buku")
    fmt.Println("=================================")

    LihatSemuaBuku()

    fmt.Print("Masukkan Kode Buku yang akan diubah: ")
    inputanUser := bufio.NewReader(os.Stdin)
    kodeBuku, _ := inputanUser.ReadString('\n')
    kodeBuku = strings.TrimSpace(kodeBuku)

    for i := range DaftarBuku {
        if DaftarBuku[i].KodeBuku == kodeBuku {
            fmt.Println("Masukkan informasi baru:")

            fmt.Print("Judul Buku: ")
            judul, _ := inputanUser.ReadString('\n')
            DaftarBuku[i].JudulBuku = strings.TrimSpace(judul)

            fmt.Print("Pengarang: ")
            pengarang, _ := inputanUser.ReadString('\n')
            DaftarBuku[i].Pengarang = strings.TrimSpace(pengarang)

            fmt.Print("Penerbit: ")
            penerbit, _ := inputanUser.ReadString('\n')
            DaftarBuku[i].Penerbit = strings.TrimSpace(penerbit)

            fmt.Print("Jumlah Halaman: ")
            fmt.Scanln(&DaftarBuku[i].JumlahHalaman)

            fmt.Print("Tahun Terbit: ")
            fmt.Scanln(&DaftarBuku[i].TahunTerbit)

            // Update file JSON
            fileName := fmt.Sprintf("books/book-%s.json", DaftarBuku[i].KodeBuku)
            bukuJSON, err := json.Marshal(DaftarBuku[i])
            if err != nil {
                fmt.Println("Error encoding book to JSON:", err)
                return
            }

            err = os.WriteFile(fileName, bukuJSON, 0644)
            if err != nil {
                fmt.Println("Error writing book JSON to file:", err)
                return
            }

            fmt.Println("Buku berhasil diubah!")
            return
        }
    }

    fmt.Println("Buku dengan Kode Buku tersebut tidak ditemukan.")
}

func MuatDataBukuDariFile() {
    files, err := os.ReadDir("books")
    if err != nil {
        fmt.Println("Error reading books directory:", err)
        return
    }

    for _, file := range files {
        if filepath.Ext(file.Name()) == ".json" {
            filePath := filepath.Join("books", file.Name())
            data, err := os.ReadFile(filePath)
            if err != nil {
                fmt.Println("Error reading file:", err)
                continue
            }

            var buku Buku
            err = json.Unmarshal(data, &buku)
            if err != nil {
                fmt.Println("Error unmarshalling JSON data:", err)
                continue
            }

            DaftarBuku = append(DaftarBuku, buku)
        }
    }
}

func GeneratePdfBuku(buku Buku) {
    pdf := fpdf.New("P", "mm", "A4", "")
    pdf.AddPage()

    pdf.SetFont("Arial", "", 12)
    pdf.SetLeftMargin(10)
    pdf.SetRightMargin(10)

    bukuText := fmt.Sprintf(
        "Kode Buku: %s\nJudul: %s\nPengarang: %s\nPenerbit: %s\nJumlah Halaman: %d\nTahun Terbit: %d",
        buku.KodeBuku, buku.JudulBuku, buku.Pengarang, buku.Penerbit, buku.JumlahHalaman, buku.TahunTerbit)

    pdf.MultiCell(0, 10, bukuText, "0", "L", false)
    pdf.Ln(5)

    // Tambahkan pencetakan waktu
    waktuCetak := time.Now().Format("2006-01-02 15:04:05")
    pdf.SetFont("Arial", "I", 10)
    pdf.CellFormat(0, 10, "Waktu Cetak: "+waktuCetak, "", 0, "R", false, 0, "")

    err := pdf.OutputFileAndClose(
        fmt.Sprintf("pdf/%s.pdf", buku.KodeBuku))

    if err != nil {
        fmt.Println("Error creating PDF file:", err)
    } else {
        fmt.Println("Buku berhasil dicetak ke file PDF.")
    }
}

func main() {
    MuatDataBukuDariFile()
	var pilihanMenuPrintBuku int
    var pilihanMenu int

    for {
        fmt.Println("=================================")
        fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
        fmt.Println("=================================")
        fmt.Println("Silahkan Pilih : ")
        fmt.Println("1. Tambah Buku")
        fmt.Println("2. Lihat Semua Buku")
        fmt.Println("3. Hapus Buku")
        fmt.Println("4. Edit Buku")
        fmt.Println("5. Keluar")
        fmt.Println("6. Print Buku")
        fmt.Println("=================================")
        fmt.Print("Masukkan Pilihan : ")
        fmt.Scanln(&pilihanMenu)

        switch pilihanMenu {
        case 1:
            TambahBuku()
        case 2:
            LihatSemuaBuku()
        case 3:
            HapusBuku()
        case 4:
            EditBuku()
        case 5:
            os.Exit(0)
        case 6:
			fmt.Println("=================================")
			fmt.Println("Print Buku")
			fmt.Println("=================================")
			fmt.Println("1. Print Satu Buku")
			fmt.Println("2. Print Semua Buku")
			fmt.Println("=================================")
			fmt.Print("Masukkan Pilihan : ")
			fmt.Scanln(&pilihanMenuPrintBuku)
		
			switch pilihanMenuPrintBuku {
			case 1:
				fmt.Print("Masukkan Kode Buku yang akan dicetak: ")
				inputanUser := bufio.NewReader(os.Stdin)
				kodeBuku, _ := inputanUser.ReadString('\n')
				kodeBuku = strings.TrimSpace(kodeBuku)
		
				for _, buku := range DaftarBuku {
					if buku.KodeBuku == kodeBuku {
						GeneratePdfBuku(buku)
						return
					}
				}
		
				fmt.Println("Buku dengan Kode Buku tersebut tidak ditemukan.")
			case 2:
				for _, buku := range DaftarBuku {
					GeneratePdfBuku(buku)
				}
				fmt.Println("Semua buku berhasil dicetak ke file PDF.")
			default:
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
        }
    }
}
