package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var enter string

const NMAX int = 100

type e_learning struct {
	tugas                         [NMAX]string
	quiz                          data_quiz
	nama_pendiskusi               [NMAX]string
	chat                          [NMAX]string
	status_simpan_permanent_tugas bool
	status_simpan_permanent_quiz  bool
}
type data_quiz struct {
	soal    [NMAX]string
	jawaban option
	kunjaw  [NMAX]string
}
type data_siswa struct {
	nama_siswa               [NMAX]string
	hasil_nilai_tugas        [NMAX]float64
	hasil_nilai_quiz         [NMAX]float64
	hasil_nilai_akhir        [NMAX]float64
	jawaban_quiz             [NMAX]data_jawaban_siswa
	jawaban_tugas            [NMAX]data_jawaban_siswa
	status_mengerjakan_quiz  [NMAX]bool
	status_mengerjakan_tugas [NMAX]bool
}
type data_jawaban_siswa struct {
	option_quiz  [NMAX]string
	option_tugas [NMAX]string
}
type option struct {
	a [NMAX]string
	b [NMAX]string
	c [NMAX]string
	d [NMAX]string
}

func main() {
	var ntugas, nquiz, nforum, pilihan_nomer, nsiswa int
	var nama_siswa string
	var pengguna int
	var course e_learning
	var murid_sekolah data_siswa

	tampilan_login_pengguna()
	fmt.Print("Pilih nomer: ")
	fmt.Scan(&pengguna)
	fmt.Println()
	for pengguna != 1 && pengguna != 2 && pengguna != 3 {
		fmt.Println("=========================================")
		fmt.Print("Pilih di antara login yang sudah ada: ")
		fmt.Scan(&pengguna)
		fmt.Println("=========================================")

		fmt.Println()
	}
	for pengguna != 3 {

		if pengguna == 1 {
			tampilan_elearning_guru()
			fmt.Print("Pilih nomer: ")
			fmt.Scan(&pilihan_nomer)
			fmt.Println()
			for pilihan_nomer > 6 || pilihan_nomer < 1 {
				fmt.Println("=========================================")
				fmt.Print("Pilih di antara nomor yang sudah ada: ")
				fmt.Scan(&pilihan_nomer)
				fmt.Println("=========================================")
				fmt.Println()
			}
			for pilihan_nomer != 6 {
				switch pilihan_nomer {
				case 1:
					tampilan_elearning_guru_tugas()
					fmt.Print("Pilih nomer: ")
					fmt.Scan(&pilihan_nomer)
					fmt.Println()
					for pilihan_nomer > 6 || pilihan_nomer < 1 {
						fmt.Println("===========================================")
						fmt.Print("Pilih di antara pilihan yang sudah ada : ")
						fmt.Scan(&pilihan_nomer)
						fmt.Println("===========================================")
						fmt.Println()
					}
					for pilihan_nomer != 6 || course.status_simpan_permanent_tugas == false {
						switch pilihan_nomer {
						case 1:
							tambah_tugas(&course, &ntugas)
						case 2:
							cek_tugas(course, ntugas)
						case 3:
							hapus_tugas(&course, &ntugas)
						case 4:
							ubah_tugas(&course, &ntugas)
						case 5:
							simpan_permanen_tugas(&course)
						}
						tampilan_elearning_guru_tugas()
						fmt.Print("Pilih nomer: ")
						fmt.Scan(&pilihan_nomer)
						fmt.Println()
						for pilihan_nomer > 6 || pilihan_nomer < 1 {
							fmt.Println("===========================================")
							fmt.Print("Pilih di antara pilihan yang sudah ada : ")
							fmt.Scan(&pilihan_nomer)
							fmt.Println("===========================================")
							fmt.Println()
						}
					}
				case 2:
					tampilan_elearning_guru_quiz()
					fmt.Print("Pilih nomer: ")
					fmt.Scan(&pilihan_nomer)
					fmt.Println()
					for pilihan_nomer > 6 || pilihan_nomer < 1 {
						fmt.Println("=========================================")
						fmt.Print("Pilih di antara nomor yang sudah ada: ")
						fmt.Scan(&pilihan_nomer)
						fmt.Println("=========================================")

						fmt.Println()
					}
					for pilihan_nomer != 6 || course.status_simpan_permanent_quiz == false {
						switch pilihan_nomer {
						case 1:
							tambah_quiz(&course, &nquiz)
						case 2:
							cek_quiz(course, nquiz)
						case 3:
							hapus_quiz(&course, &nquiz)
						case 4:
							ubah_quiz(&course, &nquiz)
						case 5:
							simpan_permanen_quiz(&course)
						}
						tampilan_elearning_guru_quiz()
						fmt.Print("Pilih nomer: ")
						fmt.Scan(&pilihan_nomer)
						fmt.Println()
						for pilihan_nomer > 6 || pilihan_nomer < 1 {
							fmt.Println("=========================================")
							fmt.Print("Pilih di antara nomor yang sudah ada: ")
							fmt.Scan(&pilihan_nomer)
							fmt.Println("=========================================")
							fmt.Println()
						}
					}
				case 3:
					nilai_tugas(&course, ntugas, nsiswa, &murid_sekolah)
				case 4:
					perhitungan_nilai_akhir(&murid_sekolah, nsiswa)
					tampilan_nilai_urutan_siswa()
					fmt.Print("Pilih nomer: ")
					fmt.Scan(&pilihan_nomer)
					fmt.Println()
					for pilihan_nomer > 3 && pilihan_nomer < 1 {
						fmt.Println("=========================================")
						fmt.Print("Pilih di antara nomor yang sudah ada: ")
						fmt.Scan(&pilihan_nomer)
						fmt.Println("=========================================")
						fmt.Println()
					}
					for pilihan_nomer != 3 {
						switch pilihan_nomer {
						case 1:
							selectionSort_mengurut_ascending(murid_sekolah, nsiswa)
						case 2:
							insertionSort_descending(murid_sekolah, nsiswa) //
						}
						tampilan_nilai_urutan_siswa()
						fmt.Print("Pilih nomer: ")
						fmt.Scan(&pilihan_nomer)
						fmt.Println()
						for pilihan_nomer > 3 && pilihan_nomer < 1 {
							fmt.Println("=========================================")
							fmt.Print("Pilih di antara nomor yang sudah ada: ")
							fmt.Scan(&pilihan_nomer)
							fmt.Println("=========================================")
							fmt.Println()
						}
					}

				case 5:
					diskusi_forum(&course, pengguna, nama_siswa, &nforum)

				}
				tampilan_elearning_guru()
				fmt.Print("Pilih nomer: ")
				fmt.Scan(&pilihan_nomer)
				fmt.Println()
				for pilihan_nomer > 6 && pilihan_nomer < 1 {
					fmt.Println("=========================================")
					fmt.Print("Pilih di antara nomor yang sudah ada: ")
					fmt.Scan(&pilihan_nomer)
					fmt.Println("=========================================")
					fmt.Println()
				}
			}

		} else if pengguna == 2 {
			fmt.Print("masukan nama siswa: ")
			fmt.Scan(&nama_siswa)
			simpan_data_siswa(&murid_sekolah, nama_siswa, &nsiswa)
			fmt.Println()
			tampilan_elearning_murid()
			fmt.Print("Pilih nomer: ")
			fmt.Scan(&pilihan_nomer)
			fmt.Println()
			for pilihan_nomer > 4 || pilihan_nomer < 1 {
				fmt.Println("=========================================")
				fmt.Print("Pilih di antara nomor yang sudah ada: ")
				fmt.Scan(&pilihan_nomer)
				fmt.Println("=========================================")
				fmt.Println()
			}
			for pilihan_nomer != 4 {
				switch pilihan_nomer {
				case 1:
					siswa_mengerjakan_tugas(course, ntugas, cari_indeks_siswa(murid_sekolah, nama_siswa, nsiswa), &murid_sekolah)
				case 2:
					siswa_mengerjakan_quiz(course, nquiz, cari_indeks_siswa(murid_sekolah, nama_siswa, nsiswa), &murid_sekolah)
				case 3:
					diskusi_forum(&course, pengguna, nama_siswa, &nforum)
				}
				tampilan_elearning_murid()
				fmt.Print("Pilih nomer: ")
				fmt.Scan(&pilihan_nomer)
				fmt.Println()
				for pilihan_nomer > 4 || pilihan_nomer < 1 {
					fmt.Println("=========================================")
					fmt.Print("Pilih di antara nomor yang sudah ada: ")
					fmt.Scan(&pilihan_nomer)
					fmt.Println("=========================================")
					fmt.Println()
				}

			}
		}
		tampilan_login_pengguna()
		fmt.Print("Pilih nomer: ")
		fmt.Scan(&pengguna)
		fmt.Println()
		for pengguna != 1 && pengguna != 2 && pengguna != 3 {
			fmt.Println("=========================================")
			fmt.Print("Pilih di antara nomor yang sudah ada: ")
			fmt.Scan(&pengguna)
			fmt.Println("=========================================")
			fmt.Println()
		}
	}

}
func tampilan_login_pengguna() {
	clearScreen()
	fmt.Println()
	fmt.Println("=========================================")
	fmt.Println("            -- L O G I N --            ")
	fmt.Println("=========================================")
	fmt.Println("silahkan pilih-pilihan yang diinginkan: ")
	fmt.Println("1). guru")
	fmt.Println("2). murid")
	fmt.Println("3). keluar")
	fmt.Println("=========================================")

}
func tampilan_elearning_guru() {
	clearScreen()
	fmt.Println()
	fmt.Println("=========================================")
	fmt.Println("            -- C O U R S E --            ")
	fmt.Println("=========================================")
	fmt.Println("silahkan pilih-pilihan yang diinginkan: ")
	fmt.Println("1). tugas")
	fmt.Println("2). quiz")
	fmt.Println("3). menilai tugas siswa")
	fmt.Println("4). nilai-nilai siswa")
	fmt.Println("5). diskusi forum")
	fmt.Println("6). keluar")
	fmt.Println("=========================================")

}
func tampilan_elearning_guru_tugas() {
	clearScreen()
	fmt.Println()
	fmt.Println("=========================================")
	fmt.Println("             -- T U G A S --             ")
	fmt.Println("=========================================")
	fmt.Println("silahkan pilih-pilihan yang diinginkan: ")
	fmt.Println("!!!    Lakukan lah simpan permanent   !!!")
	fmt.Println("!!! setelah simpan permanent anda bisa!!!")
	fmt.Println("!!!         kembali kemenu awal       !!!")
	fmt.Println("1). tambah tugas")
	fmt.Println("2). tampilkan tugas")
	fmt.Println("3). hapus tugas")
	fmt.Println("4). edit tugas")
	fmt.Println("5). simpan permanent tugas")
	fmt.Println("6). kembali kemenu awal")
	fmt.Println("=========================================")

}
func tampilan_elearning_guru_quiz() {
	clearScreen()
	fmt.Println()
	fmt.Println("=========================================")
	fmt.Println("              -- Q U I Z --              ")
	fmt.Println("=========================================")
	fmt.Println("silahkan pilih-pilihan yang diinginkan: ")
	fmt.Println("!!!    Lakukan lah simpan permanent   !!!")
	fmt.Println("!!! setelah simpan permanent anda bisa!!!")
	fmt.Println("!!!         kembali kemenu awal       !!!")
	fmt.Println("1). tambah quiz")
	fmt.Println("2). tampilkan quiz")
	fmt.Println("3). hapus quiz")
	fmt.Println("5). simpan permanent quiz")
	fmt.Println("6). kembali kemenu awal")
	fmt.Println("=========================================")
}

func tampilan_elearning_murid() {
	clearScreen()
	fmt.Println()
	fmt.Println("=========================================")
	fmt.Println("            -- C O U R S E --            ")
	fmt.Println("=========================================")
	fmt.Println("silahkan pilih-pilihan yang diinginkan: ")
	fmt.Println("1). mengerjakan tugas")
	fmt.Println("2). mengerjakan quiz")
	fmt.Println("3). diskusi forum")
	fmt.Println("4). keluar")
	fmt.Println("=========================================")
}

func tampilan_nilai_urutan_siswa() {
	clearScreen()
	fmt.Println("=============================================")
	fmt.Println("            -- NILAI-NILAI SISWA --          ")
	fmt.Println("=============================================")
	fmt.Println("silahkan pilih-pilihan yang diinginkan: ")
	fmt.Println("1). tampilan nilai siswa dari yang terkecil")
	fmt.Println("2). tampilan nilai siswa dari yang terbesar")
	fmt.Println("3). keluar")
	fmt.Println("=============================================")

}

func cek_tugas(A e_learning, n int) {
	clearScreen()
	fmt.Println("================================================")
	fmt.Println("Soal tugas: ")
	if n != 0 {
		for i := 0; i < n; i++ {
			fmt.Printf("%d. %s\n", i+1, A.tugas[i])
		}
	} else {

		fmt.Println("              Tugas kosong")

	}
	fmt.Println("================================================")
	fmt.Println("Tekan enter untuk melanjutkan")
	fmt.Scanln(&enter)

}

func tambah_tugas(A *e_learning, n *int) {
	clearScreen()
	var tugas string
	if A.status_simpan_permanent_tugas == false {
		fmt.Println("================================================")
		fmt.Println("//ketik stop untuk berhenti menambahkan tugas//")
		fmt.Print("Masukan soal tugas : ")
		fmt.Scan(&tugas)
		fmt.Println()
		for tugas != "stop" {
			A.tugas[*n] = tugas
			*n = *n + 1
			fmt.Print("Masukan soal tugas : ")
			fmt.Scan(&tugas)
			fmt.Println()
		}
		fmt.Println("================================================")
	} else {
		fmt.Println("===================================================")
		fmt.Println("      anda sudah melakukan simpan permanent")
		fmt.Println("===================================================")
	}
	fmt.Println("Tekan enter untuk melanjutkan")
	fmt.Scanln(&enter)
}

func hapus_tugas(A *e_learning, n *int) {
	clearScreen()
	var nomor int
	if A.status_simpan_permanent_tugas == false {
		fmt.Println("================================================")
		fmt.Print("Ketik nomor tugas yang ingin dihapus: ")
		fmt.Scan(&nomor)
		fmt.Println()
		for i := nomor - 1; i < *n; i++ {
			A.tugas[i] = A.tugas[i+1]
		}
		*n = *n - 1
		if *n < 0 {
			*n = 0
		}
		fmt.Println("================================================")
	} else {
		fmt.Println("===================================================")
		fmt.Println("      anda sudah melakukan simpan permanent")
		fmt.Println("===================================================")
	}
	fmt.Println("Tekan enter untuk melanjutkan")
	fmt.Scanln(&enter)

}

func ubah_tugas(A *e_learning, n *int) {
	clearScreen()
	var nomor int
	var tugas_pengganti string
	if A.status_simpan_permanent_tugas == false {
		fmt.Println("================================================")
		fmt.Print("Masukan nomor tugas yang ingin diubah: ")
		fmt.Scan(&nomor)
		fmt.Println("-----------------------------------------")
		fmt.Print("Masukan tugas pengganti: ")
		fmt.Scan(&tugas_pengganti)
		fmt.Println()
		A.tugas[nomor-1] = tugas_pengganti
		fmt.Println("================================================")
	} else {
		fmt.Println("===================================================")
		fmt.Println("      anda sudah melakukan simpan permanent")
		fmt.Println("===================================================")
	}
	fmt.Println("Tekan enter untuk melanjutkan")
	fmt.Scanln(&enter)

}

func nilai_tugas(A *e_learning, nT int, nS int, B *data_siswa) {
	var siswa int
	if nT != 0 {
		fmt.Println("================================================")
		fmt.Printf("%s %10s %10s", "Nomer", "Siswa ", "Status")
		for i := 0; i < nS; i++ {
			fmt.Printf("%d. %10s ", i+1, B.nama_siswa[i])
			if B.status_mengerjakan_tugas[i] == false {
				fmt.Printf("%10s\n", "siswa belum mengerjakan tugas")
			} else {
				fmt.Printf("%10s\n", "siswa sudah mengerjakan tugas")
			}
		}
		fmt.Println("================================================")
		fmt.Print("Pilih nomer siswa yang akan dinilai: ")
		fmt.Scan(&siswa)
		for B.status_mengerjakan_tugas[siswa-1] == false {
			fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			fmt.Println("         siswa belum mengerjakan tugas")
			fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			fmt.Println("Pilih nomer siswa yang akan dinilai: ")
			fmt.Scan(&siswa)
		}
		for i := 0; i < nT; i++ {
			fmt.Println("================================================")
			fmt.Printf("%d. %s\n", i+1, A.tugas[i])
			fmt.Println("jawaban: ", B.jawaban_tugas[siswa-1].option_tugas[i])
			fmt.Println("================================================")
		}
		fmt.Print("Masukan penilaian tugas: ")
		fmt.Scan(&B.hasil_nilai_tugas[siswa-1])
		fmt.Println("================================================")
	} else {
		fmt.Println("================================================")
		fmt.Println("             Tugas belum tersedia")
		fmt.Println("================================================")
	}
	fmt.Println("Tekan enter untuk melanjutkan")
	fmt.Scanln(&enter)
}

func cek_quiz(A e_learning, n int) {
	fmt.Println("================================================")
	if n != 0 {
		for i := 0; i < n; i++ {
			fmt.Printf("%d. %s\n", i+1, A.quiz.soal[i])
			fmt.Println(" a.", A.quiz.jawaban.a[i])
			fmt.Println(" b.", A.quiz.jawaban.b[i])
			fmt.Println(" c.", A.quiz.jawaban.c[i])
			fmt.Println(" d.", A.quiz.jawaban.d[i])
			fmt.Println("kunci jawaban: ")
			fmt.Println(A.quiz.kunjaw[i])
			fmt.Println("-----------------------------------------------")
		}
	} else {
		fmt.Println("                 Quiz kosong")

	}
	fmt.Println("================================================")
	fmt.Println("Tekan enter untuk melanjutkan")
	fmt.Scanln(&enter)

}
func tambah_quiz(A *e_learning, n *int) {
	var kuis, jawabankuis, kunci_jawaban string
	if A.status_simpan_permanent_quiz == false {
		fmt.Println("================================================")
		fmt.Println("//ketik stop untuk berhenti menambahkan quiz//")
		fmt.Print("Masukan Soal kuis: ")
		fmt.Scan(&kuis)
		fmt.Println()
		for kuis != "stop" {
			fmt.Print("Masukan option a:")
			fmt.Scan(&jawabankuis)
			A.quiz.jawaban.a[*n] = jawabankuis
			fmt.Print("Masukan option b:")
			fmt.Scan(&jawabankuis)
			A.quiz.jawaban.b[*n] = jawabankuis
			fmt.Print("Masukan option c:")
			fmt.Scan(&jawabankuis)
			A.quiz.jawaban.c[*n] = jawabankuis
			fmt.Print("Masukan option d:")
			fmt.Scan(&jawabankuis)
			A.quiz.jawaban.d[*n] = jawabankuis
			fmt.Println("<<<<<<<<<<>>>>>>>>>>")
			fmt.Print("Masukan kunci jawaban:")
			A.quiz.soal[*n] = kuis
			fmt.Scan(&kunci_jawaban)
			A.quiz.kunjaw[*n] = kunci_jawaban
			*n = *n + 1
			fmt.Println("-----------------------------------------------")
			fmt.Println("//ketik stop untuk berhenti menambahkan quiz//")
			fmt.Print("Masukan kuis: ")
			fmt.Scan(&kuis)
			fmt.Println()
		}
		fmt.Println("================================================")
	} else {
		fmt.Println("===================================================")
		fmt.Println("      anda sudah melakukan simpan permanent")
		fmt.Println("===================================================")
	}
	fmt.Println("Tekan enter untuk melanjutkan")
	fmt.Scanln(&enter)
}
func hapus_quiz(A *e_learning, n *int) {
	var nomor int
	if A.status_simpan_permanent_quiz == false {
		fmt.Println("================================================")
		fmt.Print("Masukan nomor quiz yang ingin dihapus: ")
		fmt.Scan(&nomor)
		fmt.Println()
		for i := nomor - 1; i < *n; i++ {
			A.quiz.soal[i] = A.quiz.soal[i+1]
			A.quiz.jawaban.a[i] = A.quiz.jawaban.a[i+1]
			A.quiz.jawaban.b[i] = A.quiz.jawaban.b[i+1]
			A.quiz.jawaban.c[i] = A.quiz.jawaban.c[i+1]
			A.quiz.jawaban.d[i] = A.quiz.jawaban.d[i+1]
			A.quiz.kunjaw[i] = A.quiz.kunjaw[i+1]
		}
		*n = *n - 1
		*n = *n - 1
		if *n < 0 {
			*n = 0
		}
		fmt.Println("================================================")
	} else {
		fmt.Println("===================================================")
		fmt.Println("      anda sudah melakukan simpan permanent")
		fmt.Println("===================================================")
	}
	fmt.Println("Tekan enter untuk melanjutkan")
	fmt.Scanln(&enter)
}
func ubah_quiz(A *e_learning, n *int) {
	var nomor int
	var quiz_pengganti, jawabankuis, kunci_jawaban string
	if A.status_simpan_permanent_quiz == false {
		fmt.Println("================================================")
		fmt.Print("Masukan nomor Quiz yang ingin diubah: ")
		fmt.Scan(&nomor)
		fmt.Println()
		fmt.Print("Masukan Quiz pengganti: ")
		fmt.Scan(&quiz_pengganti)
		fmt.Println()
		A.quiz.soal[nomor-1] = quiz_pengganti
		fmt.Print("Masukan option a:")
		fmt.Scan(&jawabankuis)
		A.quiz.jawaban.a[nomor-1] = jawabankuis
		fmt.Print("Masukan option b:")
		fmt.Scan(&jawabankuis)
		A.quiz.jawaban.b[nomor-1] = jawabankuis
		fmt.Print("Masukan option c:")
		fmt.Scan(&jawabankuis)
		A.quiz.jawaban.c[nomor-1] = jawabankuis
		fmt.Print("Masukan option d:")
		fmt.Scan(&jawabankuis)
		A.quiz.jawaban.d[nomor-1] = jawabankuis
		fmt.Println("<<<<<<<<<<>>>>>>>>>>")
		fmt.Print("Masukan kunci jawaban:")
		fmt.Scan(&kunci_jawaban)
		A.quiz.kunjaw[nomor-1] = kunci_jawaban
		fmt.Println("================================================")
	} else {
		fmt.Println("===================================================")
		fmt.Println("      anda sudah melakukan simpan permanent")
		fmt.Println("===================================================")
	}
	fmt.Println("Tekan enter untuk melanjutkan")
	fmt.Scanln(&enter)
}

func tampilan_nilai_siswa(B data_siswa, n int) {
	fmt.Println("======================================================================")
	fmt.Printf("%s %10s %10s %10s %10s\n", "No", "Nama_siswa", "Nilai_quiz", "Nilai_tugas", "Nilai_akhir")
	fmt.Println()
	for i := 0; i < n; i++ {

		fmt.Printf("%d). %10s %10.2f %10.2f %10.2f\n", i+1, B.nama_siswa[i], B.hasil_nilai_quiz[i], B.hasil_nilai_tugas[i], B.hasil_nilai_akhir[i])

	}
	if n == 0 {
		fmt.Println("             Belum ada siswa yang mengerjakan tugas")
	}
	fmt.Println("======================================================================")
	fmt.Println("Tekan enter untuk melanjutkan")
	fmt.Scanln(&enter)
}

func diskusi_forum(A *e_learning, user int, nama_pengguna_siswa string, n *int) {

	var pesan string
	fmt.Println("================================================")

	for i := 0; i < *n; i++ {
		fmt.Printf("%s: %s\n", A.nama_pendiskusi[i], A.chat[i])
	}
	fmt.Println("<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("//ketik stop untuk berhenti menambahkan tugas//")
	fmt.Print("masukan pesan: ")
	fmt.Scan(&pesan)

	//
	for pesan != "stop" {
		if pesan != "stop" {
			*n = *n + 1
			A.chat[*n-1] = pesan

			if user == 1 {
				A.nama_pendiskusi[*n-1] = "guru"
			} else if user == 2 {
				A.nama_pendiskusi[*n-1] = nama_pengguna_siswa
			}
		}
		fmt.Println("<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>")
		fmt.Println("//ketik stop untuk berhenti menambahkan tugas//")
		fmt.Print("masukan pesan: ")
		fmt.Scan(&pesan)
	}
	fmt.Println("================================================")
	fmt.Println("Tekan enter untuk melanjutkan")
	fmt.Scanln(&enter)
}

func simpan_data_siswa(A *data_siswa, nama string, n *int) {
	var ada_siswa bool
	ada_siswa = false
	for i := 0; i < *n; i++ {
		if nama == A.nama_siswa[i] {
			ada_siswa = true
		}
	}
	if !ada_siswa {
		A.nama_siswa[*n] = nama
		*n = *n + 1
	}
}

func cari_indeks_siswa(A data_siswa, nama string, n int) int {
	var idx int
	for i := 0; i < n; i++ {
		if nama == A.nama_siswa[i] {
			idx = i
		}
	}
	return idx
}

func siswa_mengerjakan_tugas(A e_learning, n int, j int, B *data_siswa) {
	fmt.Println("================================================")
	if n != 0 && B.status_mengerjakan_tugas[j] == false {
		for i := 0; i < n; i++ {
			fmt.Println("-----------------------------")
			fmt.Printf("%d). %s\n", i+1, A.tugas[i])
			fmt.Println("<<<<<<<<<<>>>>>>>>>>")
			fmt.Print("jawaban: ")
			fmt.Scan(&B.jawaban_tugas[j].option_tugas[i])
		}
		B.status_mengerjakan_tugas[j] = true
	} else if n == 0 && B.status_mengerjakan_tugas[j] == false {
		fmt.Println("         Tugas belum tersedia")
	} else {
		fmt.Println("Anda sudah mengerjakan tugas")
	}
	fmt.Println("================================================")
	fmt.Println("Tekan enter untuk melanjutkan")
	fmt.Scanln(&enter)

}

func siswa_mengerjakan_quiz(A e_learning, n int, j int, B *data_siswa) {

	if n != 0 && B.status_mengerjakan_quiz[j] == false {
		for i := 0; i < n; i++ {
			fmt.Printf("%d. %s\n", i+1, A.quiz.soal[i])
			fmt.Println(" a.", A.quiz.jawaban.a[i])
			fmt.Println(" b.", A.quiz.jawaban.b[i])
			fmt.Println(" c.", A.quiz.jawaban.c[i])
			fmt.Println(" d.", A.quiz.jawaban.d[i])
			fmt.Print("jawaban: ")
			fmt.Scan(&B.jawaban_quiz[j].option_quiz[i])
			if B.jawaban_quiz[j].option_quiz[i] == A.quiz.kunjaw[i] {
				B.hasil_nilai_quiz[j] = B.hasil_nilai_quiz[j] + 100/float64(n)
			}

		}
		B.status_mengerjakan_quiz[j] = true
	} else if n == 0 && B.status_mengerjakan_quiz[j] == false {
		fmt.Println("Quiz belum tersedia")
		fmt.Println("Tekan enter untuk melanjutkan")
		fmt.Scanln(&enter)
	} else {
		fmt.Println("Anda sudah mengerjakan quiz")

	}
	if n > 0 {
		fmt.Printf("Nilai Quiz: %.2f\n", B.hasil_nilai_quiz[j])
		fmt.Println("Tekan enter untuk melanjutkan")
		fmt.Scanln(&enter)

	}

}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd = exec.Command("clear")
	} else {
		fmt.Println("Platform tidak didukung.")
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func selectionSort_mengurut_ascending(T data_siswa, nS int) {
	var i, j, idx_min int
	var t data_siswa
	i = 1
	for i <= nS-1 {
		idx_min = i - 1
		j = i
		for j < nS {
			if T.hasil_nilai_akhir[idx_min] > T.hasil_nilai_akhir[j] {
				idx_min = j
			}
			j = j + 1
		}
		t.hasil_nilai_akhir[0] = T.hasil_nilai_akhir[idx_min]
		t.hasil_nilai_quiz[0] = T.hasil_nilai_quiz[idx_min]
		t.hasil_nilai_tugas[0] = T.hasil_nilai_tugas[idx_min]
		t.nama_siswa[0] = T.nama_siswa[idx_min]
		T.hasil_nilai_akhir[idx_min] = T.hasil_nilai_akhir[i-1]
		T.hasil_nilai_quiz[idx_min] = T.hasil_nilai_quiz[i-1]
		T.hasil_nilai_tugas[idx_min] = T.hasil_nilai_tugas[i-1]
		T.nama_siswa[idx_min] = T.nama_siswa[i-1]
		T.hasil_nilai_akhir[i-1] = t.hasil_nilai_akhir[0]
		T.hasil_nilai_quiz[i-1] = t.hasil_nilai_quiz[0]
		T.hasil_nilai_tugas[i-1] = t.hasil_nilai_tugas[0]
		T.nama_siswa[i-1] = t.nama_siswa[0]
		i = i + 1
	}
	tampilan_nilai_siswa(T, nS)

}

func insertionSort_descending(T data_siswa, nS int) {
	var i, j int
	var temp data_siswa
	i = 1
	for i <= nS-1 {
		j = i
		temp.hasil_nilai_akhir[0] = T.hasil_nilai_akhir[j]
		temp.hasil_nilai_quiz[0] = T.hasil_nilai_quiz[j]
		temp.hasil_nilai_tugas[0] = T.hasil_nilai_tugas[j]
		temp.nama_siswa[0] = T.nama_siswa[j]
		for j > 0 && temp.hasil_nilai_akhir[0] > T.hasil_nilai_akhir[j-1] {
			T.hasil_nilai_akhir[j] = T.hasil_nilai_akhir[j-1]
			T.hasil_nilai_quiz[j] = T.hasil_nilai_quiz[j-1]
			T.hasil_nilai_tugas[j] = T.hasil_nilai_tugas[j-1]
			T.nama_siswa[j] = T.nama_siswa[j-1]
			j = j - 1

		}
		T.hasil_nilai_akhir[j] = temp.hasil_nilai_akhir[0]
		T.hasil_nilai_quiz[j] = temp.hasil_nilai_quiz[0]
		T.hasil_nilai_tugas[j] = temp.hasil_nilai_tugas[0]
		T.nama_siswa[j] = temp.nama_siswa[0]
		i = i + 1
	}

	tampilan_nilai_siswa(T, nS)
}

func perhitungan_nilai_akhir(B *data_siswa, nS int) {
	for i := 0; i < nS; i++ {
		if B.status_mengerjakan_quiz[i] == false && B.status_mengerjakan_tugas[i] == true {
			B.hasil_nilai_akhir[i] = (B.hasil_nilai_tugas[i]) / 2
		} else if B.status_mengerjakan_tugas[i] == false && B.status_mengerjakan_quiz[i] == true {
			B.hasil_nilai_akhir[i] = (B.hasil_nilai_quiz[i]) / 2
		} else if B.status_mengerjakan_quiz[i] == true && B.status_mengerjakan_tugas[i] == true {
			B.hasil_nilai_akhir[i] = (B.hasil_nilai_quiz[i] + B.hasil_nilai_tugas[i]) / 2
		} else if B.status_mengerjakan_quiz[i] == false && B.status_mengerjakan_tugas[i] == false {
			B.hasil_nilai_akhir[i] = 0
		}

	}
}

func simpan_permanen_tugas(A *e_learning) {
	var konfirmasi int
	clearScreen()
	if A.status_simpan_permanent_tugas == false {

		fmt.Println("===================================================")
		fmt.Println("apakah anda yakin akan melakukan simpan permanent?")
		fmt.Println("1. iya")
		fmt.Println("2. tidak")
		fmt.Println("===================================================")
		fmt.Print("pilihan nomer: ")
		fmt.Scan(&konfirmasi)
		if konfirmasi == 1 {
			A.status_simpan_permanent_tugas = true
		}
		for konfirmasi > 2 || konfirmasi < 1 {
			fmt.Println("===========================================")
			fmt.Print("Pilih di antara pilihan yang sudah ada : ")
			fmt.Scan(&konfirmasi)
			fmt.Println("===========================================")
			fmt.Println()
		}
	} else {
		fmt.Println("===================================================")
		fmt.Println("      anda sudah melakukan simpan permanent")
		fmt.Println("===================================================")
	}
	fmt.Println("Tekan enter untuk melanjutkan")
	fmt.Scanln(&enter)
}

func simpan_permanen_quiz(A *e_learning) {
	var konfirmasi int
	clearScreen()
	if A.status_simpan_permanent_quiz == false {

		fmt.Println("===================================================")
		fmt.Println("apakah anda yakin akan melakukan simpan permanent?")
		fmt.Println("1. iya")
		fmt.Println("2. tidak")
		fmt.Println("===================================================")
		fmt.Print("pilihan nomer: ")
		fmt.Scan(&konfirmasi)
		if konfirmasi == 1 {
			A.status_simpan_permanent_quiz = true
		}
		for konfirmasi > 2 || konfirmasi < 1 {
			fmt.Println("===========================================")
			fmt.Print("Pilih di antara pilihan yang sudah ada : ")
			fmt.Scan(&konfirmasi)
			fmt.Println("===========================================")
			fmt.Println()
		}
	} else {
		fmt.Println("===================================================")
		fmt.Println("      anda sudah melakukan simpan permanent")
		fmt.Println("===================================================")
	}
	fmt.Println("Tekan enter untuk melanjutkan")
	fmt.Scanln(&enter)
}
