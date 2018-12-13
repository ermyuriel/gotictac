package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httputil"
	"os"
)

const (
	player = "X"
	ai     = "O"
	draw   = "DRAW"
	nval   = ""
)

//Game defines the current state of a board, intended to be parsed from JSON request from frontend
type Game struct {
	Board  []string `json:"board"`
	Winner string   `json:"winner"`
}

//main starts AI API on port from .env
func main() {

	http.HandleFunc(os.Getenv("REACT_APP_API_ENDPOINT"), moveHandler)
	http.ListenAndServe(":"+os.Getenv("REACT_APP_API_PORT"), nil)

}

//moveHandler parses JSON body from request as Game object and calls returnMove with next state from AI perspective, returning errors if necessary

func moveHandler(w http.ResponseWriter, r *http.Request) {

	x, _ := httputil.DumpRequest(r, true)

	fmt.Println(string(x))

	g := &Game{}

	d := json.NewDecoder(r.Body)

	err := d.Decode(g)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	nm := returnMove(g)
	j, err := json.Marshal(nm)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//returnMove executes minimax algorithm from AI perspective if game is not over, returns new game board state. State is evaluated again after finding best move to determine if game should be terminated

func returnMove(g *Game) *Game {

	if g.Winner != nval {
		return g
	}

	if wins(g.Board, ai) {
		g.Winner = ai
		return g
	}

	if wins(g.Board, player) {
		g.Winner = player
		return g
	}

	bm := minimax(g.Board, ai)[0]

	if bm == -1 {

		g.Winner = draw
		return g
	}

	g.Board[bm] = ai

	if wins(g.Board, ai) {
		g.Winner = ai
	}

	return g

}

//minimax determines best move by alternatively maximizing and minimizing function over all possible subsequent states

func minimax(nb []string, p string) []int {

	if wins(nb, player) {

		return []int{-1, -10}
	}

	if wins(nb, ai) {

		return []int{-1, 10}
	}

	np := notPlayedIndexes(nb)

	if len(np) == 0 {
		return []int{-1, 0}
	}

	ms := make([][]int, 0)
	for _, v := range np {

		m := make([]int, 2)

		m[0] = v

		nbc := make([]string, len(nb))

		copy(nbc, nb)

		nbc[v] = p

		r := make([]int, 2)

		if p == ai {
			r = minimax(nbc, player)
		} else {

			r = minimax(nbc, ai)

		}
		m[1] = r[1]

		ms = append(ms, m)

	}

	var bm []int
	if p == ai {

		bs := int(math.MinInt32)

		for _, v := range ms {

			if v[1] > bs {
				bs = v[1]
				bm = v
			}

		}

	} else {

		bs := int(math.MaxInt32)

		for _, v := range ms {

			if v[1] < bs {
				bs = v[1]
				bm = v
			}

		}

	}

	return bm

}

//wins evaluates winning condition by static index validation
func wins(board []string, p string) bool {
	if (board[0] == p && board[1] == p && board[2] == p) ||
		(board[3] == p && board[4] == p && board[5] == p) ||
		(board[6] == p && board[7] == p && board[8] == p) ||
		(board[0] == p && board[3] == p && board[6] == p) ||
		(board[1] == p && board[4] == p && board[7] == p) ||
		(board[2] == p && board[5] == p && board[8] == p) ||
		(board[0] == p && board[4] == p && board[8] == p) ||
		(board[2] == p && board[4] == p && board[6] == p) {
		return true
	}
	return false
}

//notPlayedIndexes returns possible plays
func notPlayedIndexes(b []string) []int {

	is := make([]int, 0)

	for i, v := range b {
		if v == nval {
			is = append(is, i)
		}
	}

	return is

}

//pb is an auxiliary method to print a board in readable format for testing
func pb(b []string) {

	for i := 1; i < 10; i++ {
		if b[i-1] == nval {
			fmt.Printf("|%v|", "_")
		} else {
			fmt.Printf("|%v|", b[i-1])

		}

		if i%3 == 0 {
			fmt.Println()
		}

	}

	fmt.Println("--------")

}
