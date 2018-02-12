package ascii

import "fmt"

func IterateOverASCIIStringLiteral(sl string) {
	// Kode for Oppgave 4a
	length := len(sl)
	for i := 0; i < length; i++ {
		hexValue := sl[i]
		fmt.Printf("%X\t %c\t %b\n", hexValue, hexValue, hexValue)



		//for i := 0x80; i < 0x9F; i++{ //Eksempel som Wiklem viste
		//fmt.Printf("%X\t %c\t %b\n", i, i, i)


	}
}



// Kode for Oppgave 4b
func ExtendedASCIIText() {

}