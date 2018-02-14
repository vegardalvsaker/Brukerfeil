package ascii

import (
	"fmt"
)

func IterateOverASCIIStringLiteral() {
	// Kode for Oppgave 4a
	const hexString = "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9A\x9B\x9C\x9D\x9E\x9F" +
		"\xA0\xA1\xA2\xA3\xA4\xA5\xA6\xA7\xA8\xA9\xAA\xAB\xAC\xAD\xAE\xAF\xB0\xB1\xB2\xB3\xB4\xB5\xB6\xB7\xB8\xB9\xBA\xBB\xBC\xBD\xBE\xBF" +
		"\xC0\xC1\xC2\xC3\xC4\xC5\xC6\xC7\xC8\xC9\xCA\xCB\xCC\xCD\xCE\xCF\xD0\xD1\xD2\xD3\xD4\xD5\xD6\xD7\xD8\xD9\xDA\xDB\xDC\xDD\xDE\xDF" +
		"\xE0\xE1\xE2\xE3\xE4\xE5\xE6\xE7\xE8\xE9\xEA\xEB\xEC\xED\xEE\xEF\xF0\xF1\xF2\xF3\xF4\xF5\xF6\xF7\xF8\xF9\xFA\xFB\xFC\xFD\xFE\xFF"

	length := len(hexString) //GO kjenner igjen hexaverdiene selv om de er skrevet som string, derfor kan vi finne lengden av stringen 'length'.

	for i := 0; i < length; i++ {
		hexValue := hexString[i]
		fmt.Printf("%X\t %c\t %b\n", hexValue, hexValue, hexValue)
		/////////////////////////////////////////////////////////////
		//for i := 0x80; i < 0x9F; i++{ //Eksempel som Wiklem viste//
		//fmt.Printf("%X\t %c\t %b\n", i, i, i)                    //
		/////////////////////////////////////////////////////////////
	}
}

// Kode for Oppgave 4b
func ExtendedASCIIText4b()string {
	//Verdiene skal i utgangspunktet representere " € ÷ ¾ dollar " (anførselstegn inkludert), men extended ascii-en er av type 8859-1
	tekstbytes := []byte{0x22, 0x20, 0x80, 0x20, 0xF7, 0x20, 0xBE, 0x20, 0x64, 0x6F, 0x6C, 0x6C, 0x61, 0x72, 0x20, 0x22}
    print(tekstbytes)
	i := len(tekstbytes)
	var nyTekst string

	for j := 0; j < i; j++ {
		if (tekstbytes[j] >= 0x80) && (tekstbytes[j] <= 0xFF )  {
			temp := string(tekstbytes[j])
			nyTekst = nyTekst + temp
		}
	}
	return nyTekst
}
/*
Egen funksjon for å printe siden ExtendedASCIIText returnerer, og ikke kan printe.
 */
func print(b  []byte){
	fmt.Printf("%c", b)
}

// Kode for Oppgave 4c
func ExtendedASCIIText(s string)string {
	//Verdiene skal i utgangspunktet representere " € ÷ ¾ dollar " (anførselstegn inkludert), men extended ascii-en er av type 8859-1
	tekstbytes := []byte(s)
	print(tekstbytes)
	i := len(tekstbytes)
	var nyTekst string
	for j := 0; j < i; j++ {
		if (tekstbytes[j] >= 0x80) && (tekstbytes[j] <= 0xFF )  {
			temp := string(tekstbytes[j])
			nyTekst = nyTekst + temp
		}
	}
	return nyTekst
}

