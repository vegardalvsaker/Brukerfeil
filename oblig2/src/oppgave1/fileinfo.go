package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	filepath := os.Args[1]

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	info, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}

	//Leser, regner ut og printer størrelsen på filen i mb, kb osv
	conv := info.Size()
	bytes := float64(conv)
	kb := bytes / 1024
	mb := kb / 1024
	gb := mb / 1024
	fmt.Println("File is this large:\n", bytes, "Bytes\n", kb, "Kilobytes\n", mb, "Megabytes\n", gb, "Gigabytes\n")

	//Sjekker om filen er en directory eller ikke
	if info.IsDir() {
		fmt.Println("This is a directory file.")
	} else {
		fmt.Println("This is not a directory file.")
	}

	//Sjekker om filen er en "regular file"
	// For å gjøre dette må vi gjøre statsen(info) til en variabel av datatypen "FileMode" først
	infomode := info.Mode()
	if infomode.IsRegular() {
		fmt.Println("This is a regular file.")
	} else {
		fmt.Println("This is not a regular file.")
	}

	//Sjekker om filen har Unix permission bits -rwxrwxrwx (0666 er -rw-rw-rw- mens 0777 er -rwxrwxrwx i 8-tallsystem)
	permBits := infomode.Perm()

	if permBits == 0777 {
		fmt.Println("This file has Unixcode", permBits, ", this means that everyone can access the file.")
	} else {
		fmt.Println("This file does not have Unixcode -rwxrwxrwx, it has the following code:\n", permBits,
			", this means that the file has limited permissions.")
	}

	//Sjekker om filen er append only
	/* Switch case er en mulighet.
	switch mode := info.Mode(); {
	case mode&os.ModeAppend != 0:
		fmt.Println("This file is append only")
	case mode&os.ModeAppend == 0:
		fmt.Println("This file is not append only")
	}
	*/
	test := info.Mode()
	if test&os.ModeAppend == 0 {
		fmt.Println("This file is not append only.")
	} else {
		fmt.Println("This file is append only.")
	}


	//Sjekker om det er en device file
	if test&os.ModeDevice == 0 {
		fmt.Println("This file is not a device file.")
	} else {
		fmt.Println("This file is a device file.")
	}

	//Sjekker om det er en symbolic link
	if test&os.ModeSymlink == 0 {
		fmt.Println("This file is not a symbolic link.")
	} else {
		fmt.Println("This file is a symbolic link.")
	}
}