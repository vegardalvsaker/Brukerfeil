# Oppgave 1
#### Fyll ut manglende tall i tabell

Binære tall | Hexadesimaltall | Desimaltall
--- | --- | ---
1101 | 0xD | 13
1101 1110 1010 | 0xDEA | 3562
1010 1111 0011 0100 | 0xAF34 | 44852
1111 1111 1111 1111 | 0xFFFF | 65535
0001 0001 0111 1000 1010 | 0x1178A | 71562


## Oppgave 1 A
#### Beskriv kort metode for å gå fra binære tall til hexadesimale tall og motsatt. Beskriv kort metoden for å gå fra binære tall til desimaltall og motsatt.
Binære til hexadesimal:   
For å gå fra binære- til heksadesimaltall må man dele tallet opp i 4-deler. Så om man har tallet 101011 må man legge til 0’er i starten slik at man kan dele det i fire. Slik: 0010 1011. 
Etter dette må man kunne verdiene i hexadesimal systemet. 0010 = 2 i desimal og 2 i hexa. 1011 = 11 i desimal eller B i hexa. Så hexatallet vårt blir 0x2B  

Desimal til binær og motsatt:   
Del desimaltallet på 2, hvis rest -> 1, hvis ikke rest ->0, repeter. Begynn nederst når du lager binærtallet.
For å gå fra binær til desimal ser man på plassen til binærtallet, det representerer antall ganger man må gange med 2. Hvis man har binærtallet 1001 1101 så er det bakerste tallet verdt 1, det neste 2 og videre 4, 8, 16 osv. Man ganger tallet med to hver gang. Hver gang det er 1 i binærtallet må man plusse på denne verdien. I vårt eksempel blir det 1+4+8+16+128 = 157.


## Oppgave 1 B
#### Beskriv kort metoden for å gå fra hexadesimale tall til desimaltall og motsatt.
For å gå fra hexa- til desimaltall må man først og fremst vite at tallene 1-9 representerer de samme tallene i desimaltall. Videre har man A-F som er 10-15. Tilslutt må man se på tallets plass. Det bakerste tallet (x) skal regnes ut slik x * 16^0. På det neste tallet (z) regner vi z * 16^1 osv. Det eneste som stiger er tallet vi opphøyer 16 med.  

---

# Oppgave 2
**Oppgave 2A** *Se fil under: src/algorithms/sorting.go*  
**Oppgave 2B** *Se fil under: src/algorithms/sorting_test.go*


## Oppgave 2C
Big O handler om å se hvilke algoritmer som scaler best, en algoritme kan være god på å gjøre det den skal gjør om den for eksempel får få enkle tall, mens hvis den får mange tall kan den være treg. Det handler om å kunne si at en algoritme er spesialisert til en type oppgave.

Navn på benchmark | Func per 1sec | Nanosekund per operasjon
--- | --- | ---
BenchmarkBSortM100-8 | 100000 | 21827 ns/op
BenchmarkBSortM1000-8 | 1000 | 1887953 ns/op
BenchmarkBSortM10000-8 | 10 | 187135212 ns/op
|  |  |
BenchmarkBSort100-8 | 100000 | 21060 ns/op
BenchmarkBSort1000-8 | 1000 | 1313876 ns/op
BenchmarkBSort10000-8 | 10 | 186884906 ns/op
|  |  |
BenchmarkQ	Sort100-8 | 300000 | 4749 ns/op
BenchmarkQSort1000-8 | 30000 | 50128 ns/op
BenchmarkQSort10000-8 | 2000 | 639300 ns/op


Vi ser at alle algoritmene øker eksponentielt. I følge cheatsheetet skal både bubblesort og Qsort være O(n^2), men vi ser at Qsort er desidert den raskeste av de tre.

---

# Oppgave 3
*Se fil under: src/oppg3/oppg3.go*


#### Skriv et program som består av en evig løkke. Hvor mye minne og CPU bruker programmet når det kjører?
'main'-prossessen bruker 1,0 MB av ram, og go bruker 5,8 MB ram. Begger bruker 0% av CPUen.
 
#### Programmet skal skrive ut en avslutningsmelding når programmet mottar et SIGINT (Control + C) signal. Generer ulike avslutningssignaler til prosessen og dokumenter hvilke avslutningskommandoer programmet håndterer og som trigger avslutningsmeldingen.
Programmet reagerer på SIGINT (CTRL+C)og 3 andre signaler som du skriver inn i terminalen. Disse er : “sigint”, “sigterm”, og “sigquit”. Alle disse vil gjøre at prosessen avslutter (os.Exit). Koden for dette ligger i mappen og filen opp3

---

# Oppgave 4
## Oppgave 4A
*Se fil under src/ascii/ascii.go*  


#### Utfør programmet på alle gruppemedlemmers datamaskin og analyser
Vi ser at vi ikke får listet eller lest karakterene fra og med 80 til og med 9F på Windows i både cmd og powershell. I GoLand får vi samme resultat uavhengig av operativsystem, dette er fordi GoLand bruker UTF-8 by default.   
Forskjellige operativsystem bruker forskjellige karaktersett avhengig hvor du er og hvilket OS du bruker. Det varierer fordi folk trenger forskjellige karakterer.  
CMD og Powershell i Windows bruker Multilangual Latin 1 (Code page 850) i Vest-Europa.  

## Oppgave 4B
*Se fil under: src/ascii/ascii.go*

#### Utfør programmet på forskjellige plattformer(mac,windows,linux) eller forskjellig software(terminal, bash, powershell) og analyser resultater
På Windows i GoLand kom alle tegnene unntatt €, i cmd på Windows var det et ? istedenfor  €.  
På OSx(mac) i GoLand kom alle tegnene mens i Terminal skrev alt untatt €.  
Analyse:   
I Windows har CMD og Powershell det samme karaktersettet. Windows cmd og ps returnerer et spørsmåltegn i motsetning til GoLand som bare skriver whitespace.   
I OSx (mac) fikk vi riktig resultat i GoLand men ikke i terminal, den skrev whitespace.   
Vi ser at disse resultatene stemmer overens med det vi sjekket i Oppgave 4a.  

## Oppgave 4C
*Se filene under: src/ascii/iso_test.go og src/ascii/ascii.go*
