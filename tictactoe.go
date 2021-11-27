package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 27 nov---> inizio progetto
//implementato: matrice, matriceViewer
//da implementare:
//posizioneOccupata(una funziona che ti chiede il valore finché non ne inserisci una valido) fatto
//posizioneOccupata è valida se è [1-9] oppure non è già stata occupata (magari con una funzione che usa *) fatto
//mettere X - O al posto di 0(libero), 1(p1), 2(p2) fatto
//8 modi in cui si può vinceere, oppure se il turno supera 9 la partita è finita fatto
func matriceViewer(tris [3][3]string) {
	for _, r := range tris {
		for _, q := range r {
			fmt.Print(q, " ")
		}
		fmt.Println()
	}
}
func inizio() {
	fmt.Println("Scivere un numero da 1-9, oppure Q per abbandonare")
}
func wcL(tris [3][3]string) bool { //vitorie laterali
	if (tris[0][0] == tris[0][1] && tris[0][1] == tris[0][2] && tris[0][2] != "-" && tris[0][1] != "-" && tris[0][0] != "-") || (tris[1][0] == tris[1][1] && tris[1][1] == tris[1][2] && tris[1][2] != "-" && tris[1][1] != "-" && tris[1][0] != "-") || (tris[2][0] == tris[2][1] && tris[2][1] == tris[2][2] && tris[2][2] != "-" && tris[2][1] != "-" && tris[2][0] != "-") {
		return true
	}
	return false
}
func wcV(tris [3][3]string) bool { //vitorie laterali
	if (tris[0][0] == tris[1][0] && tris[1][0] == tris[2][0] && tris[1][0] != "-" && tris[2][0] != "-" && tris[0][0] != "-") || (tris[0][1] == tris[1][1] && tris[1][1] == tris[2][1] && tris[0][1] != "-" && tris[1][1] != "-" && tris[2][1] != "-") || (tris[0][2] == tris[1][2] && tris[1][2] == tris[2][2] && tris[0][2] != "-" && tris[1][2] != "-" && tris[2][2] != "-") {
		return true
	}
	return false
}
func wcC(tris [3][3]string) bool {
	if (tris[0][0] == tris[1][1] && tris[1][1] == tris[2][2] && tris[0][0] != "-" && tris[1][1] != "-") || (tris[0][2] == tris[1][1] && tris[1][1] == tris[2][0] && tris[2][1] != "-" && tris[0][2] != "-") {
		return true
	}
	return false
}
func winCondition(s [3][3]string) bool {
	if wcL(s) || wcV(s) || wcC(s) {
		return true
	}
	return false
}
func main() {
	tris := [3][3]string{}
	//k := 1
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < len(tris); i++ {
		for j := 0; j < len(tris); j++ {
			//u := strconv.Itoa(k)
			tris[i][j] = "-"
			//k++
		}
	}
	game := true
	turno := 1

	fmt.Println("iniziamo il gioco")
	//matriceViewer(tris)
	for game {
		inizio()
		if turno == 10 {
			fmt.Println("gioco finito")
			break
		}
		matriceViewer(tris)
		fmt.Println("vai: ")
		scanner.Scan()
		text := scanner.Text()
		if text == "q" {
			break
		}
		for len(text) != 1 || text == "0" { //quando sbaglia l'input
			fmt.Println("riprova a mettere un input corretto: ")
			var y string
			fmt.Scanln(&y)
			text = y
		}
		p, err := strconv.Atoi(text)
		if err != nil { //riga inutile
			fmt.Println("sbagliato")
		}
		p--
		//check se è occupato
		for tris[p/3][p%3] == "X" || tris[p/3][p%3] == "O" {
			fmt.Println("occupata: riprova.")
			fmt.Scanln(&p)
			if err != nil {
				fmt.Println("sbagliato")
			}
			p--
		}
		// faccio % per vedere in che riga, e / per vedere in che colonna es.(7 ---> (7%3=1) (7/3= 2))
		// 0 = prima griglia, 1 = seconda griglia...
		if turno%2 != 0 {
			//turno giocatore 1
			tris[p/3][p%3] = "X"

		} else {
			//turno giocatore 2
			tris[p/3][p%3] = "O"

		}
		fmt.Println(" ")
		if winCondition(tris) {
			matriceViewer(tris)
			fmt.Println("GG")
			break
		}
		fmt.Println("turno #", turno)
		turno++
	}
}
