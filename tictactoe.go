package main

import "fmt"

// 27 nov---> inizio progetto
//implementato: matrice, matriceViewer
//da implementare:
//posizioneOccupata(una funziona che ti chiede il valore finché non ne inserisci una valido)
//posizioneOccupata è valida se è [1-9] oppure non è già stata occupata (magari con una funzione che usa *)
//mettere X - O al posto di 0(libero), 1(p1), 2(p2)
//8 modi in cui si può vinceere, oppure se il turno supera 9 la partita è finita
func matriceViewer(tris [3][3]int) {
	for _, r := range tris {
		for _, q := range r {
			fmt.Print(q, " ")
		}
		fmt.Println()
	}
}
func main() {
	tris := [3][3]int{}
	var p int
	game := true
	turni := 1
	alternato := false //false = gioatore1, true = giocatore2
	fmt.Println("iniziamo il gioco")
	fmt.Scanf("%d", &p)
	p--
	for game {
		matriceViewer(tris)
		fmt.Println("vai: ")
		fmt.Scanf("%d", &p)
		p--
		//check se è occupato
		// faccio % per vedere in che riga, e / per vedere in che colonna es.(7 ---> (7%3=1) (7/3= 2))
		// 0 = prima griglia, 1 = seconda griglia...
		if alternato {
			//turno giocatore 1
			tris[p/3][p%3] = 1
			alternato = false

		} else {
			//turno giocatore 2
			tris[p/3][p%3] = 2
			alternato = true
		}
		if turni == 10 {
			fmt.Println("gioco finito")
			break
		}
		fmt.Println(" ")
		fmt.Println("turno #", turni)
		turni++
	}
}
