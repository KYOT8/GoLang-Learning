package main

import (
	"fmt"
	"sort"
)

// Testy wycinkow, czyli slice'ow, czyli dynamicznych tablic,
// ktore moga sie rozrastac w trakcie dzialania programu, w przeciwienstwie do zwyklych tablic,
// ktore maja staly rozmiar i nie moga sie zmieniac
//

func main() {

	var listagpuNVIDA = []string{"Nvidia GTX 1030", "Nvidia GTX 1060 3GB", "Nvidia GTX 1060 6GB", "Nvidia GTX 1070 8gb", "Nvidia GTX 1080 8GB"}

	var listagpuAMD = []string{"AMD RX 550 2GB", "AMD RX 560 4GB", "AMD RX 570 4GB", "AMD RX 580 8GB", "AMD RX 590 8GB"}

	fmt.Printf("Listy gpu NVIDA: %v\n", listagpuNVIDA)
	fmt.Printf("Listy gpu AMD: %v\n", listagpuAMD)

	var zarobki = make([]int, 12)

	zarobki[0] = 30200
	zarobki[1] = 40300
	zarobki[2] = 54000
	zarobki[3] = 603400
	zarobki[4] = 703300
	zarobki[5] = 80030
	zarobki[6] = 90300
	zarobki[7] = 103000
	zarobki[8] = 1334000
	zarobki[9] = 1000
	zarobki[10] = 13000
	zarobki[11] = 140300

	fmt.Printf("Zarobki: %v\n", zarobki)
	sort.Ints(zarobki) //sortujemy slice zarobki w porzadku rosnacym, funkcja sort.Ints sortuje slice typu int
	fmt.Printf("Zarobki posortowane od najmniejszych do najwiekszych: %v\n", zarobki)

	sort.Sort(sort.Reverse(sort.IntSlice(zarobki))) //sortujemy slice zarobki w porzadku malejacym, funkcja sort.Reverse sortuje slice w odwrotnym porzadku, a sort.IntSlice zamienia slice typu int na slice typu sort.IntSlice, ktory jest wymagany przez funkcje sort.Reverse
	fmt.Printf("Zarobki posortowane od najwiekszych do najmniejszych: %v\n", zarobki)

	var zarobki2 = []int{30200, 40300, 54000, 603400, 703300, 80030, 90300, 103000, 1334000, 1000, 13000, 140300} //mozemy tez od razu zainicjalizowac slice zarobki2 bez podawania jego rozmiaru, wtedy rozmiar zostanie automatycznie dopasowany do ilosci elementow w nawiasach klamrowych
	fmt.Printf("Zarobki2: %v\n", zarobki2)

	var cpu = []string{"Intel Core i3 10100F", "Intel Core i5 10400F", "Intel Core i7 10700K", "Intel Core i9 10900K", "AMD Ryzen 3 3100", "AMD Ryzen 5 3600", "AMD Ryzen 7 3700X", "AMD Ryzen 9 3900X"}
	fmt.Printf("CPU: %v\n", cpu) //print zwykle

	sort.Strings(cpu) //sortujemy slice cpu w porzadku alfabetycznym, funkcja sort.Strings sortuje slice typu string
	fmt.Printf("CPU posortowane alfabetycznie: %v\n", cpu)

}
