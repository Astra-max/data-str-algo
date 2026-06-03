package libs

import "github.com/01-edu/z01"

func PrintMemory(data [10]byte) {
	hex := "0123456789abcdef"

	for i:= 0; i<10; i++ {
		printHex(hex[data[i]/16])
		printHex(hex[data[i]%16])

		if i != 9 {
			z01.PrintRune(' ')
		}
		if i%4 == 3 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')
	printData(data)
}

func printData(data [10]byte) {
	for i:=0; i<10; i++ {
		if data[i] >= 32 && data[i] <= 126 {
			z01.PrintRune(rune(data[i]))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func printHex(data byte) {
	z01.PrintRune(rune(data))
}