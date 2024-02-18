package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var _________testOk1 = `4
2
TS TC
AD AH
3
2H 3H
9S 9C
4D QS
3
4C 7H
4H 4D
6S 6H
3
2S 3H
2C 2D
3C 3D
`

var _________testOkResult1 = `2
TD
TH
0
3
7S
7C
7D
0
`

func TestPoker1(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(_________testOk1))
	out := new(bytes.Buffer)
	err := poker(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != _________testOkResult1 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _________testOkResult1)
	}
}

var _________testOk2 = `1
2
TS TC
AD AH`

var _________testOkResult2 = `2
TH
TD
`

func TestPoker2(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(_________testOk2))
	out := new(bytes.Buffer)
	err := poker(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != _________testOkResult2 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _________testOkResult2)
	}
}

var _________testOk3 = `1
3
2H 3H
9S 9C
4D QS`

var _________testOkResult3 = `0
`

func TestPoker3(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(_________testOk3))
	out := new(bytes.Buffer)
	err := poker(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != _________testOkResult3 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _________testOkResult3)
	}
}

var _________testOk4 = `1
3
4C 7H
4H 4D
6S 6H`

var _________testOkResult4 = `3
7S
7D
7C
`

func TestPoker4(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(_________testOk4))
	out := new(bytes.Buffer)
	err := poker(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != _________testOkResult4 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _________testOkResult4)
	}
}

var _________testOk5 = `1
3
2S 3H
2C 2D
3C 3D`

var _________testOkResult5 = `0
`

func TestPoker5(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(_________testOk5))
	out := new(bytes.Buffer)
	err := poker(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != _________testOkResult5 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _________testOkResult5)
	}
}

var _________testOk6 = `1
7
AS AC
AD AH
KS JH
9D 9C
5H 5D
3C 3S
TC TH`

var _________testOkResult6 = `30
2S
2C
2D
2H
4S
4C
4D
4H
6S
6C
6D
6H
7S
7C
7D
7H
8S
8C
8D
8H
JS
JC
JD
QS
QC
QD
QH
KC
KD
KH
`

func TestPoker6(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(_________testOk6))
	out := new(bytes.Buffer)
	err := poker(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != _________testOkResult6 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _________testOkResult6)
	}
}

func TestWinner_OnlyFirstHas_Set(t *testing.T) {
	firstPlayerCards := []Card{Card{Rank: Ace, Color: Spades}, Card{Rank: Ace, Color: Hearts}}
	secondPlayerCards := []Card{Card{Rank: Ace, Color: Clubs}, Card{Rank: King, Color: Hearts}}
	thirdCard := Card{Rank: Ace, Color: Diamonds}
	p1 := Player{cards: [3]Card{firstPlayerCards[0], firstPlayerCards[1], thirdCard}}
	p2 := Player{cards: [3]Card{secondPlayerCards[0], secondPlayerCards[1], thirdCard}}

	firstWins := p1.Wins(&p2)
	if !firstWins {
		t.Errorf("test for OK failed - results not match\n %v \n %v", firstWins, true)
	}

}

func TestWinner_OnlySecondHas_Set(t *testing.T) {
	firstPlayerCards := []Card{Card{Rank: Ace, Color: Spades}, Card{Rank: Ace, Color: Hearts}}
	secondPlayerCards := []Card{Card{Rank: King, Color: Clubs}, Card{Rank: King, Color: Hearts}}
	thirdCard := Card{Rank: King, Color: Diamonds}
	p1 := Player{cards: [3]Card{firstPlayerCards[0], firstPlayerCards[1], thirdCard}}
	p2 := Player{cards: [3]Card{secondPlayerCards[0], secondPlayerCards[1], thirdCard}}

	firstWins := p1.Wins(&p2)
	if firstWins {
		t.Errorf("test for OK failed - results not match\n %v \n %v", firstWins, true)
	}

}

func TestWinner_OnlyFirst_Has_Pair(t *testing.T) {
	firstPlayerCards := []Card{Card{Rank: Ace, Color: Spades}, Card{Rank: Ace, Color: Hearts}}
	secondPlayerCards := []Card{Card{Rank: King, Color: Clubs}, Card{Rank: Queen, Color: Hearts}}
	thirdCard := Card{Rank: Ten, Color: Diamonds}
	p1 := Player{cards: [3]Card{firstPlayerCards[0], firstPlayerCards[1], thirdCard}}
	p2 := Player{cards: [3]Card{secondPlayerCards[0], secondPlayerCards[1], thirdCard}}

	firstWins := p1.Wins(&p2)
	if !firstWins {
		t.Errorf("test for OK failed - results not match\n %v \n %v", firstWins, true)
	}

}

func TestWinner_Both_Have_Pair(t *testing.T) {
	firstPlayerCards := []Card{Card{Rank: Ace, Color: Spades}, Card{Rank: Ace, Color: Hearts}}
	secondPlayerCards := []Card{Card{Rank: King, Color: Clubs}, Card{Rank: King, Color: Hearts}}
	thirdCard := Card{Rank: Ten, Color: Diamonds}
	p1 := Player{cards: [3]Card{firstPlayerCards[0], firstPlayerCards[1], thirdCard}}
	p2 := Player{cards: [3]Card{secondPlayerCards[0], secondPlayerCards[1], thirdCard}}

	firstWins := p1.Wins(&p2)
	if !firstWins {
		t.Errorf("test for OK failed - results not match\n %v \n %v", firstWins, true)
	}

}

func TestWinner_Second_Has_Pair(t *testing.T) {
	firstPlayerCards := []Card{Card{Rank: Ace, Color: Spades}, Card{Rank: Ten, Color: Hearts}}
	secondPlayerCards := []Card{Card{Rank: King, Color: Clubs}, Card{Rank: Queen, Color: Hearts}}
	thirdCard := Card{Rank: King, Color: Diamonds}
	p1 := Player{cards: [3]Card{firstPlayerCards[0], firstPlayerCards[1], thirdCard}}
	p2 := Player{cards: [3]Card{secondPlayerCards[0], secondPlayerCards[1], thirdCard}}

	firstWins := p1.Wins(&p2)
	if firstWins {
		t.Errorf("test for OK failed - results not match\n %v \n %v", firstWins, false)
	}
}
