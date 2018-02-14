Fyll ut manglende tall i tabell

Binære tall | Hexadesimaltall | Desimaltall
--- | --- | ---
1101 | 0xD | 13
1101 1110 1010 | 0xDEA | 3562
1010 1111 0011 0100 | 0xAF34 | 44852
1111 1111 1111 1111 | 0xFFFF | 65535
0001 0001 0111 1000 1010 | 0x1178A | 71562

---

# Oppgave A
## Beskriv kort metode for å gå fra binære tall til hexadesimale tall og motsatt. Beskriv kort metoden for å gå fra binære tall til desimaltall og motsatt.
Binære til hexadesimal:   
For å gå fra binære- til heksadesimaltall må man dele tallet opp i 4-deler. Så om man har tallet 101011 må man legge til 0’er i starten slik at man kan dele det i fire. Slik: 0010 1011. 
Etter dette må man kunne verdiene i hexadesimal systemet. 0010 = 2 i desimal og 2 i hexa. 1011 = 11 i desimal eller B i hexa. Så hexatallet vårt blir 0x2B  

Desimal til binær og motsatt:   
Del desimaltallet på 2, hvis rest -> 1, hvis ikke rest ->0, repeter. Begynn nederst når du lager binærtallet.
For å gå fra binær til desimal ser man på plassen til binærtallet, det representerer antall ganger man må gange med 2. Hvis man har binærtallet 1001 1101 så er det bakerste tallet verdt 1, det neste 2 og videre 4, 8, 16 osv. Man ganger tallet med to hver gang. Hver gang det er 1 i binærtallet må man plusse på denne verdien. I vårt eksempel blir det 1+4+8+16+128 = 157.

---

# Oppgave B
## Beskriv kort metoden for å gå fra hexadesimale tall til desimaltall og motsatt.
For å gå fra hexa- til desimaltall må man først og fremst vite at tallene 1-9 representerer de samme tallene i desimaltall. Videre har man A-F som er 10-15. Tilslutt må man se på tallets plass. Det bakerste tallet (x) skal regnes ut slik x * 16^0. På det neste tallet (z) regner vi z * 16^1 osv. Det eneste som stiger er tallet vi opphøyer 16 med.  
