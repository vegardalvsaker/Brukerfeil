# Brukerfeil

##Oppgave 3
```go
package vegard

import (
	"fmt"
	"bufio"
	"os"
	//"strings"
)

func SignalFunc() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "kake" {
			fmt.Println("The kake shut the program down")
			os.Exit(0)
		}
		if text == "pølse" {
			fmt.Println("The pølse shut the program down")
			os.Exit(0)
		}
		if text == "hest"{
			fmt.Println("The hest shut the program down")
			os.Exit(0)
		}
	}

}

```
<p>'main'-prossessen bruker 1,0 MB av ram, og go bruker 5,8 MB ram. Begger bruker 0% av CPUen.</p>
<p>Programmet reagerer på 3 ulike avslutningssignaler, 'kake', 'pølse' og 'hest'. Hvert signal har en forskjellig avslutningsmelding.</p>

## Oppgave 4a

```
80	 	 10000000
81	 	 10000001
82	 	 10000010
83	 	 10000011
84	 	 10000100
85	 	 10000101
86	 	 10000110
87	 	 10000111
88	 	 10001000
89	 	 10001001
90	 	 10010000
91	 	 10010001
92	 	 10010010
93	 	 10010011
94	 	 10010100
95	 	 10010101
96	 	 10010110
97	 	 10010111
98	 	 10011000
99	 	 10011001
9A	 	 10011010
9B	 	 10011011
9C	 	 10011100
9D	 	 10011101
9E	 	 10011110
9F	 	 10011111
A0	  	 10100000
A1	 ¡	 10100001
A2	 ¢	 10100010
A3	 £	 10100011
A4	 ¤	 10100100
A5	 ¥	 10100101
A6	 ¦	 10100110
A7	 §	 10100111
A8	 ¨	 10101000
A9	 ©	 10101001
AA	 ª	 10101010
AB	 «	 10101011
AC	 ¬	 10101100
AD	 ­	 10101101
AE	 ®	 10101110
AF	 ¯	 10101111
B0	 °	 10110000
B1	 ±	 10110001
B2	 ²	 10110010
B3	 ³	 10110011
B4	 ´	 10110100
B5	 µ	 10110101
B6	 ¶	 10110110
B7	 ·	 10110111
B8	 ¸	 10111000
B9	 ¹	 10111001
BA	 º	 10111010
BB	 »	 10111011
BC	 ¼	 10111100
BD	 ½	 10111101
BE	 ¾	 10111110
BF	 ¿	 10111111
C0	 À	 11000000
C1	 Á	 11000001
C2	 Â	 11000010
C3	 Ã	 11000011
C4	 Ä	 11000100
C5	 Å	 11000101
C6	 Æ	 11000110
C7	 Ç	 11000111
C8	 È	 11001000
C9	 É	 11001001
CA	 Ê	 11001010
CB	 Ë	 11001011
CC	 Ì	 11001100
CD	 Í	 11001101
CE	 Î	 11001110
CF	 Ï	 11001111
D0	 Ð	 11010000
D1	 Ñ	 11010001
D2	 Ò	 11010010
D3	 Ó	 11010011
D4	 Ô	 11010100
D5	 Õ	 11010101
D6	 Ö	 11010110
D7	 ×	 11010111
D8	 Ø	 11011000
D9	 Ù	 11011001
DA	 Ú	 11011010
DB	 Û	 11011011
DC	 Ü	 11011100
DD	 Ý	 11011101
DE	 Þ	 11011110
DF	 ß	 11011111
E0	 à	 11100000
E1	 á	 11100001
E2	 â	 11100010
E3	 ã	 11100011
E4	 ä	 11100100
E5	 å	 11100101
E6	 æ	 11100110
E7	 ç	 11100111
E8	 è	 11101000
E9	 é	 11101001
EA	 ê	 11101010
EB	 ë	 11101011
EC	 ì	 11101100
ED	 í	 11101101
EE	 î	 11101110
EF	 ï	 11101111
F0	 ð	 11110000
F1	 ñ	 11110001
F2	 ò	 11110010
F3	 ó	 11110011
F4	 ô	 11110100
F5	 õ	 11110101
F6	 ö	 11110110
F7	 ÷	 11110111
F8	 ø	 11111000
F9	 ù	 11111001
FA	 ú	 11111010
FB	 û	 11111011
FC	 ü	 11111100
FD	 ý	 11111101
FE	 þ	 11111110
FF	 ÿ	 11111111
```

<p>Extended ASCII-symbolene fra 80 til 9F vises ikke på en Windows-maskin.</p>

## Oppgave 4b
