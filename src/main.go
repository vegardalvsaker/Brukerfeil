package main

import (
	"./test"
    //"./ascii"
	"fmt"
)


func main() {
	//Oppgave 4a
	const hexString = "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9A\x9B\x9C\x9D\x9E\x9F" +
		"\xA0\xA1\xA2\xA3\xA4\xA5\xA6\xA7\xA8\xA9\xAA\xAB\xAC\xAD\xAE\xAF\xB0\xB1\xB2\xB3\xB4\xB5\xB6\xB7\xB8\xB9\xBA\xBB\xBC\xBD\xBE\xBF" +
		"\xC0\xC1\xC2\xC3\xC4\xC5\xC6\xC7\xC8\xC9\xCA\xCB\xCC\xCD\xCE\xCF\xD0\xD1\xD2\xD3\xD4\xD5\xD6\xD7\xD8\xD9\xDA\xDB\xDC\xDD\xDE\xDF" +
		"\xE0\xE1\xE2\xE3\xE4\xE5\xE6\xE7\xE8\xE9\xEA\xEB\xEC\xED\xEE\xEF\xF0\xF1\xF2\xF3\xF4\xF5\xF6\xF7\xF8\xF9\xFA\xFB\xFC\xFD\xFE\xFF"

    //ascii.IterateOverASCIIStringLiteral(hexString)
	vegard.GetInput()
 }


func IterateOverASCIIStringLiteral(sl string) {
	// Kode for Oppgave 4b
	length := len(sl)
	//for i := 0x80; i < 0x9F; i++{ //Eksempel som Wiklem viste
		//fmt.Printf("%X\t %c\t %b\n", i, i, i)
	//}
	for i := 0; i < length; i++ {
		hexValue := sl[i]
		fmt.Printf("%X\t %c\t %b\n", hexValue, hexValue, hexValue)

	}
}
/*
	 func signalgreia() {
		 ch := make(chan os.Signal, 1)
		 signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		 go func() {
			 for sig := range ch {

				 switch sig {
				 case syscall.SIGTERM:
					 fmt.Println("sigterm recieved")
					 os.Exit(0)
				 case syscall.SIGINT:
					 fmt.Println("sigint received. Shutting down")
					 os.Exit(0)
				 case syscall.SIGQUIT:
					 fmt.Println("sigquit received")
					 os.Exit(0)

				 }
			 }
		 }()
		 time.Sleep(time.Minute)

	 }
*/




