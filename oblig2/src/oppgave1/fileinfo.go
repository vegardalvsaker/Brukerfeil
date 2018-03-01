package main

import (
		"fmt"
		"os"
        "log"
)

func main(){
	filnavn := os.Args[1]

	//os.Lstat returnerer en FileInfo av filen
	fi, err := os.Lstat(filnavn)
		if err != nil {
			log.Fatal(err)
		}

	//fi.Mode() returner en FileMode av FileInfo.
	mode := fi.Mode();

	//Filnavn
	fmt.Println("Information about file <", fi.Name(), ">")

	//Size
	intBytes := fi.Size()
	bytes := float64(intBytes) //Gjør int om til float64

	KB := bytes / 1000
	MB := KB / 1000
	GB := MB / 1000
	fmt.Println("Size:", fi.Name(), "in bytes, KB, MB and GB:")
	fmt.Println(bytes, "bytes,", KB, "KB,", MB, "MB,", GB, "GB")

	//Directory
	if mode.IsDir() != true {
		fmt.Println("Is not a directory")
	} else {
		fmt.Println("Is a directory")
	}
	//Regular
	if mode.IsRegular() != true {
		fmt.Println("Is not a regular file")
	} else {
		fmt.Println("Is a regular file")
	}
	//Unix permissions
	//unixrep := mode&os.ModeSocket
	fmt.Println("Has Unix permission bits:", mode)

	//Append
	if mode&os.ModeAppend == 0 {
		fmt.Println("Is not append only")
	} else {
		fmt.Println("Is append only")
	}

	//Device file
	if mode&os.ModeDevice == 0 {
		fmt.Println("Is not a device file")
	} else {
		fmt.Println("Is a device file")
	}

	//Unix character
	if mode&os.ModeCharDevice == 0 {
		fmt.Println("Is not a Unix character device")
	} else {
		fmt.Println("Is a Unix character device")
	}

	//Unix block
	//if mode&os.

	//Symbolic link
	if mode&os.ModeSymlink == 0 {
		fmt.Println("Is not a symbolic link")
	} else {
		fmt.Println("Is a symbolic link")
	}
}

/*func main() {
	    filnavn := `C:\Users\Vegard\Documents\dev\Go\is-105\Brukerfeil\oblig2\src\oppgave1\filnavn.txt`
		fi, err := os.Lstat(filnavn)
		if err != nil {
			log.Fatal(err)
		}
		//Filnavn
		mode := fi.Mode();
		fmt.Println("Information about file <", fi.Name(), ">")

        //Size
        intBytes := fi.Size()
        bytes := float64(intBytes) //Gjør int om til float64

		KB := bytes / 1000
		MB := KB / 1000
		GB := MB / 1000
		fmt.Println("Size:", fi.Name(), "in bytes, KB, MB and GB:")
        fmt.Println(bytes, "bytes,", KB, "KB,", MB, "MB,", GB, "GB")

        //Directory
        if mode.IsDir() != true {
        	fmt.Println("Is not a directory")
		} else {
			fmt.Println("Is a directory")
		}
		//Regular
		if mode.IsRegular() != true {
			fmt.Println("Is not a regular file")
		} else {
			fmt.Println("Is a regular file")
		}
		//Unix permissions
		//unixrep := mode&os.ModeSocket
		fmt.Println("Has Unix permission bits:", mode)

	    //Append
	    if mode&os.ModeAppend == 0 {
	    	fmt.Println("Is not append only")
		} else {
			fmt.Println("Is append only")
		}

		//Device file
		if mode&os.ModeDevice == 0 {
			fmt.Println("Is not a device file")
		} else {
			fmt.Println("Is a device file")
		}

		//Unix character
		if mode&os.ModeCharDevice == 0 {
			fmt.Println("Is not a Unix character device")
		} else {
			fmt.Println("Is a Unix character device")
		}

		//Unix block
		//if mode&os.

		//Symbolic link
		if mode&os.ModeSymlink == 0 {
			fmt.Println("Is not a symbolic link")
		} else {
			fmt.Println("Is a symbolic link")
		}
}
*/