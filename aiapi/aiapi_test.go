package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//TestServerResponses validates response via status code for valid and invalid JSON objects sent to server
func TestServerResponses(t *testing.T) {

	s := httptest.NewServer(http.HandlerFunc(moveHandler))
	defer s.Close()

	empty := "{}"
	ej, err := json.Marshal(empty)
	if err != nil {
		t.Errorf("Error when posting:%v", err)
		t.Fail()
	}

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(ej))

	if resp.StatusCode != 400 {
		t.Errorf("Server accepting invalid board with code:%v", resp.StatusCode)
		t.Fail()
	}

	valid := "{\"board\":[\"O\",\"\",\"X\",\"X\",\"\",\"X\",\"\",\"O\",\"O\"],\"winner\":\"\"}"
	if err != nil {
		t.Errorf("Error when posting:%v", err)
		t.Fail()
	}

	vr, err := http.Post(s.URL, "application/json", bytes.NewBuffer([]byte(valid)))

	if vr.StatusCode != 200 {
		t.Errorf("Server not accepting valid board with code:%v", vr.StatusCode)
		t.Fail()
	}

}

//TestGameOver validates that server recognizes a game as being over and makes no further attempts to modify state
func TestGameOver(t *testing.T) {

	s := httptest.NewServer(http.HandlerFunc(moveHandler))
	defer s.Close()

	over := "{\"board\":[\"O\",\"\",\"X\",\"X\",\"\",\"X\",\"\",\"O\",\"O\"],\"winner\":\"X\"}"

	or, _ := http.Post(s.URL, "application/json", bytes.NewBuffer([]byte(over)))
	bb, _ := ioutil.ReadAll(or.Body)
	b := string(bb)
	if b != over {
		t.Error("Server playing after game is over")
		t.Fail()
	}

}

//TestValidMove checks that the server is returning a correct response and determining the winner correctly in it's response
func TestValidMove(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(moveHandler))
	defer s.Close()
	valid := "{\"board\":[\"O\",\"\",\"X\",\"X\",\"\",\"X\",\"\",\"O\",\"O\"],\"winner\":\"\"}"

	expected := "{\"board\":[\"O\",\"\",\"X\",\"X\",\"O\",\"X\",\"\",\"O\",\"O\"],\"winner\":\"O\"}"

	vr, _ := http.Post(s.URL, "application/json", bytes.NewBuffer([]byte(valid)))
	bb, _ := ioutil.ReadAll(vr.Body)
	b := string(bb)
	if b != expected {
		t.Error("Server playing incorrectly")
		t.Fail()
	}

}

//TestMinimax ensures minimax algorithm determines heuristically determined next move
func TestMinimax(t *testing.T) {

	testBoard := []string{ai, nval, player, player, nval, player, nval, ai, ai}

	if r := minimax(testBoard, ai); r[0] != 4 {

		t.Errorf("Wrong move, should be  %d got %d.", 4, r[0])
		t.Fail()

	}

}
