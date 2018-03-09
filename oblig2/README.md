## Brukerfeil
# Oblig2

Vegard Alvsaker, Ingve Fosse, Espen Oftedal og Filip Østrem

## Oppgave 1
Oppgave 1 ligger i oppgave1-mappen med navnet fileinfo.go. Den bygde filen av programmet ligger i src/bin.

Programmet tar en fil som et argument.

## Oppgave 2
Oppgave 2 ligger i opggave2-mappen med navnet filecount.go. Den bygde filen av programmet ligger i src/bin.

Litt trøbbel med filepath i IDE vs Exe.
Filepathen må se slik ut i IDE : `oblig2/src/files/text.txt`
og slik i Exe: `../files/text.txt`. Dette avhenger av hvilken mappe du åpner som GoLand-prosjekt. GoLand tar
 utgangspunkt at du ligger i den øverste mappen, som i vårt tilfelle er mappen "Brukerfeil"

Programmet tar ingen argumenter.

## Oppgave 3

### a)
Oppgave 3a) ligger i oppgave3-mappen med navnet addup.go. Den bygde filen av programmet ligger i src/bin.

### b)
Oppgave 3b) ligger i oppgave3-mappen med navnene addtofile.go og sumfromfile.go.
Den bygde filen av programmet ligger i src/bin og tar to tall som argument.
For å kjøre go filene må du skrive ``go run addtofile.go sumfromfile.go <tall1> <tall2> ``.

### c)
Vi har lagt til feilhåndtering ved alle I/O-operasjoner. Hvis det kommer en error som != nil kjører vi `log.Fatal()`.
Denne printer erroren til loggen og avslutter programmet med ``os.Exit(1)``.

### d)
Programmene printer "Sigint received. Shutting down safely" hvis de mottar en SIGINT.

### e)
Alle de bygde filene ligger i src/bin.
