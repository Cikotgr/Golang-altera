package main

import "fmt"

func main(){
	var panjangAtas, panjangBawah, tinggi float64

	fmt.Print("Masukan panjang sisi atas trapesium: ")
	fmt.Scanln(&panjangAtas)

	fmt.Print("Masukan panjang sisi bawah trapesium: ")
	fmt.Scanln(&panjangBawah)

	fmt.Print("Masukan tinggi trapesium: ")
	fmt.Scanln(&tinggi)

	luasTrapesium(panjangAtas,panjangBawah,tinggi)

}

func luasTrapesium(panjangAtas, panjangBawah, tinggi float64) {
	
	luas := 0.5 * (panjangAtas + panjangBawah) * tinggi 
	
	fmt.Printf("Luas trapesium dengan panjang sisi atas %.2f, panjang sisi bawah %.2f, dan tinggi %.2f adalah %.2f\n", panjangAtas, panjangBawah, tinggi, luas)
}