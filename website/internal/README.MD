Cod privat al modulelor/aplicației. Acesta este cod pe care nu vrei alții să îl
importe în aplicațiile/modulele lor. Această structură este forțată de
compilatorul Go. Vezi Go 1.4 release notes pentru mai multe detalii. Nu ești
însă limitat de un singur director top-level internal. Poți avea mai multe
directoare internal la orice nivel în cadrul proiectului tău.

Poți adăuga o structură adițională modulelor tale interne pentru a separa codul
partajat de cel nepartajat. Nu este necesar, dar este bune să ai indicii vizuale
pentru a arăta scopul de folosire al modulelor. Codul actual al aplicației poate
fi în directorul /internal/app (e.g., /internal/app/myapp) iar codul partajat de
acele aplicații în /internal/pkg (e.g., /internal/pkg/myprivlib).
