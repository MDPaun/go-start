Cod librărie care poate fi folosit de aplicațiile externe (e.g.,
/pkg/mypubliclib). Alte proiecte vor importa aceste module și se vor aștepta ca
ele să funcționeze, așa că gândește-te de 2 ori înainte de a pune ceva aici :)
Directorul internal este o modalitate mai bună de a fi sigur că modulele tale no
sunt importabile deoarece acest fapt este forțat de Go. Directorul /pkg este în
continuare un mod bun de a comunica explicit codul ca fiind sigur de folosit de
către ceilalți. Postarea I'll take pkg over internal oferă o prezentare generală
a directoarelor pkg și internal.

Este o metodă bună să grupezi codul Go într-un singur loc atunci când directorul
root al aplicației conține multe componente no-Go, (cum este menționat în aceste
prezentări: Best Practices for Industrial Programming from GopherCon EU 2018,
GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps and GoLab 2018 -
Massimiliano Pippi - Project layout patterns in Go).

Vezi /pkg dacă vrei să vezi care repo-uri populare Go folosesc această structură
de proiect. Aceasta este o structură comună de proiect însă nu este universală
și nu o recomand în comunitatea Go.
