package main

import "fmt"

const NMax int = 100

type Kandidat struct {
    Nama       string
    Posisi     string
    Pengalaman int
    Gaji       int
    Status     string
}

type arrKandidat [NMax]Kandidat

func main() {
    var ats arrKandidat
    var nKandidat int
    var pilihan int

    nKandidat = 10
    ats[0] = Kandidat{"Dimas", "Backend_Dev", 3, 15000000, "Interview"}
    ats[1] = Kandidat{"Dabi", "Frontend_Dev", 7, 8000000, "Screening"}
    ats[2] = Kandidat{"Hans", "Mobile_Dev", 5, 9000000, "Hired"}
    ats[3] = Kandidat{"Anggun", "Backend_Dev", 6, 30000000, "Interview"}
    ats[4] = Kandidat{"Naufal", "Frontend_Dev", 8, 35000000, "Screening"}
    ats[5] = Kandidat{"Cayu", "Mobile_Dev", 4, 25000000, "Hired"}
    ats[6] = Kandidat{"Rekal", "Backend_Dev", 2, 5000000, "Interview"}
    ats[7] = Kandidat{"Pasha", "Frontend_Dev", 9, 40000000, "Screening"}
    ats[8] = Kandidat{"Fathan", "Mobile_Dev", 10, 37000000, "Hired"}
    ats[9] = Kandidat{"Joji", "Mobile_Dev", 5, 20000000, "Hired"}

    pilihan = 0
    for pilihan != 10 {
        fmt.Println("\n=== Recruiter Applicant Tracking System ===")
        fmt.Println("1. Tambah Data Kandidat Baru")
        fmt.Println("2. Update Status/Data Kandidat")
        fmt.Println("3. Hapus Data Kandidat")
        fmt.Println("4. Lihat Semua Kandidat")
        fmt.Println("5. Cari Posisi Dilamar (Sequential Search)")
        fmt.Println("6. Cari Nama Kandidat (Binary Search)")
        fmt.Println("7. Urutkan Pengalaman Senior ke Junior (Selection Sort)")
        fmt.Println("8. Urutkan Nama Abjad A-Z (Insertion Sort)")
        fmt.Println("9. Cari Ekspektasi Gaji Tertinggi & Terendah")
        fmt.Println("10. Tutup Sistem")
        fmt.Print("Pilih menu: ")
        fmt.Scan(&pilihan)

        if pilihan == 1 {
            tambahKandidat(&ats, &nKandidat)
        } else if pilihan == 2 {
            ubahKandidat(&ats, nKandidat)
        } else if pilihan == 3 {
            hapusKandidat(&ats, &nKandidat)
        } else if pilihan == 4 {
            tampilData(ats, nKandidat)
        } else if pilihan == 5 {
            cariSequentialPosisi(ats, nKandidat)
        } else if pilihan == 6 {
            cariBinaryNama(ats, nKandidat)
        } else if pilihan == 7 {
            urutkanSelectionPengalaman(ats, nKandidat)
        } else if pilihan == 8 {
            urutkanInsertionNama(ats, nKandidat)
        } else if pilihan == 9 {
            cariEkstrim(ats, nKandidat)
        } else if pilihan == 10 {
            fmt.Println("Sistem ditutup. Waktunya HRD PushRank!")
        } else {
            fmt.Println("Pilihan ga valid, cek lagi menunya!")
        }
    }
}

func tambahKandidat(data *arrKandidat, n *int) {
    var k Kandidat

    if *n < NMax {
        fmt.Print("Masukkan Nama Kandidat: ")
        fmt.Scan(&k.Nama)
        fmt.Print("Masukkan Posisi Dilamar (Gunakan_underscore): ")
        fmt.Scan(&k.Posisi)
        fmt.Print("Masukkan Pengalaman (Tahun): ")
        fmt.Scan(&k.Pengalaman)
        fmt.Print("Masukkan Ekspektasi Gaji (Angka): ")
        fmt.Scan(&k.Gaji)
        fmt.Print("Masukkan Status (Screening/Test/Interview/Hired/Rejected): ")
        fmt.Scan(&k.Status)

        data[*n] = k
        *n = *n + 1
        fmt.Println("Kandidat berhasil dimasukkan ke sistem!")
    } else {
        fmt.Println("Kapasitas sistem penuh!")
    }
}

func ubahKandidat(data *arrKandidat, n int) {
    var nama string
    var i int
    var ketemu bool

    fmt.Print("Masukkan Nama Kandidat yang mau diupdate: ")
    fmt.Scan(&nama)

    ketemu = false
    i = 0
    for i < n && ketemu == false {
        if data[i].Nama == nama {
            fmt.Print("Update Posisi: ")
            fmt.Scan(&data[i].Posisi)
            fmt.Print("Update Pengalaman (Tahun): ")
            fmt.Scan(&data[i].Pengalaman)
            fmt.Print("Update Ekspektasi Gaji: ")
            fmt.Scan(&data[i].Gaji)
            fmt.Print("Update Status: ")
            fmt.Scan(&data[i].Status)
            fmt.Println("Data kandidat berhasil diupdate!")
            ketemu = true
        }
        i = i + 1
    }

    if ketemu == false {
        fmt.Println("Kandidat ga ketemu di sistem!")
    }
}

func hapusKandidat(data *arrKandidat, n *int) {
    var nama string
    var i, j int
    var ketemu bool

    fmt.Print("Masukkan Nama Kandidat yang mau dihapus: ")
    fmt.Scan(&nama)

    ketemu = false
    i = 0
    for i < *n && ketemu == false {
        if data[i].Nama == nama {
            ketemu = true
            j = i
            for j < *n-1 {
                data[j] = data[j+1]
                j = j + 1
            }
            *n = *n - 1
            fmt.Println("Data kandidat berhasil dihapus dari sistem!")
        }
        i = i + 1
    }

    if ketemu == false {
        fmt.Println("Kandidat ga ketemu di sistem!")
    }
}

func tampilData(data arrKandidat, n int) {
    var i int

    fmt.Println("--------------------------------------------------------------------------------")
    fmt.Printf("%-15s | %-15s | %-10s | %-12s | %-10s\n", "Nama Kandidat", "Posisi", "Pengalaman", "Eksp. Gaji", "Status")
    fmt.Println("--------------------------------------------------------------------------------")
    
    i = 0
    for i < n {
        fmt.Printf("%-15s | %-15s | %-10d | %-12d | %-10s\n", data[i].Nama, data[i].Posisi, data[i].Pengalaman, data[i].Gaji, data[i].Status)
        i = i + 1
    }
}

func cariSequentialPosisi(data arrKandidat, n int) {
    var posisi string
    var i int
    var ketemu bool

    fmt.Print("Masukkan Posisi yang dicari: ")
    fmt.Scan(&posisi)

    ketemu = false
    fmt.Println("\n[Hasil Pencarian Sequential]")
    
    i = 0
    for i < n {
        if data[i].Posisi == posisi {
            fmt.Printf("Nama: %s | Posisi: %s | Pengalaman: %d Tahun | Gaji: %d | Status: %s\n", data[i].Nama, data[i].Posisi, data[i].Pengalaman, data[i].Gaji, data[i].Status)
            ketemu = true
        }
        i = i + 1
    }
    
    if ketemu == false {
        fmt.Println("Belum ada kandidat untuk posisi tersebut.")
    }
}

func cariBinaryNama(data arrKandidat, n int) {
    var nama string
    var temp arrKandidat
    var i, j int
    var key Kandidat
    var low, high, mid int
    var ketemu bool

    fmt.Print("Masukkan Nama Kandidat: ")
    fmt.Scan(&nama)

    i = 0
    for i < n {
        temp[i] = data[i]
        i = i + 1
    }

    i = 1
    for i < n {
        key = temp[i]
        j = i - 1
        for j >= 0 && temp[j].Nama > key.Nama {
            temp[j+1] = temp[j]
            j = j - 1
        }
        temp[j+1] = key
        i = i + 1
    }

    low = 0
    high = n - 1
    ketemu = false

    for low <= high && ketemu == false {
        mid = (low + high) / 2

        if temp[mid].Nama == nama {
            fmt.Println("\n[Data Ditemukan via Binary Search!]")
            fmt.Printf("Nama: %s | Posisi: %s | Pengalaman: %d Tahun | Gaji: %d | Status: %s\n", temp[mid].Nama, temp[mid].Posisi, temp[mid].Pengalaman, temp[mid].Gaji, temp[mid].Status)
            ketemu = true
        } else if temp[mid].Nama < nama {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    if ketemu == false {
        fmt.Println("Kandidat belum ada di sistem.")
    }
}

func urutkanSelectionPengalaman(data arrKandidat, n int) {
    var temp arrKandidat
    var i, j, maxIdx int
    var tukar Kandidat

    i = 0
    for i < n {
        temp[i] = data[i]
        i = i + 1
    }

    i = 0
    for i < n-1 {
        maxIdx = i
        j = i + 1
        for j < n {
            if temp[j].Pengalaman > temp[maxIdx].Pengalaman {
                maxIdx = j
            }
            j = j + 1
        }
        tukar = temp[i]
        temp[i] = temp[maxIdx]
        temp[maxIdx] = tukar
        
        i = i + 1
    }

    fmt.Println("\nLeaderboard Kandidat (Selection Sort - Paling Senior):")
    tampilData(temp, n)
}

func urutkanInsertionNama(data arrKandidat, n int) {
    var temp arrKandidat
    var i, j int
    var key Kandidat

    i = 0
    for i < n {
        temp[i] = data[i]
        i = i + 1
    }

    i = 1
    for i < n {
        key = temp[i]
        j = i - 1
        for j >= 0 && temp[j].Nama > key.Nama {
            temp[j+1] = temp[j]
            j = j - 1
        }
        temp[j+1] = key
        i = i + 1
    }

    fmt.Println("\nData Diurutkan A-Z (Insertion Sort):")
    tampilData(temp, n)
}

func cariEkstrim(data arrKandidat, n int) {
    var maxGaji, minGaji Kandidat
    var i int

    if n == 0 {
        fmt.Println("Sistem masih kosong, belum ada kandidat!")
    } else {
        maxGaji = data[0]
        minGaji = data[0]

        i = 1
        for i < n {
            if data[i].Gaji > maxGaji.Gaji {
                maxGaji = data[i]
            }
            if data[i].Gaji < minGaji.Gaji {
                minGaji = data[i]
            }
            i = i + 1
        }

        fmt.Println("\n[Statistik Ekstrim Ekspektasi Gaji]")
        fmt.Printf("Ekspektasi Gaji Paling Tinggi: %s (%s) minta %d\n", maxGaji.Nama, maxGaji.Posisi, maxGaji.Gaji)
        fmt.Printf("Ekspektasi Gaji Paling Rendah: %s (%s) minta %d\n", minGaji.Nama, minGaji.Posisi, minGaji.Gaji)
    }
}
