package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Rank int

type Color int

func (p *Player) IsSet() (Rank, bool) {
	isSet := p.cards[0].Rank == p.cards[1].Rank && p.cards[1].Rank == p.cards[2].Rank
	if isSet {
		return p.cards[0].Rank, true
	} else {
		return Undefined, false
	}
}

func (p *Player) IsPair() (Rank, bool) {
	if p.cards[0].Rank == p.cards[1].Rank {
		return p.cards[0].Rank, true
	} else if p.cards[1].Rank == p.cards[2].Rank {
		return p.cards[1].Rank, true
	} else if p.cards[0].Rank == p.cards[2].Rank {
		return p.cards[0].Rank, true
	} else {
		return Undefined, false
	}
}

func (p *Player) GetHighest() Rank {
	if p.cards[0].Rank > p.cards[1].Rank {
		if p.cards[0].Rank > p.cards[2].Rank {
			return p.cards[0].Rank
		} else {
			return p.cards[2].Rank
		}
	} else {
		if p.cards[1].Rank > p.cards[2].Rank {
			return p.cards[1].Rank
		} else {
			return p.cards[2].Rank
		}
	}
}

func GetColor(card string) Color {
	switch card[1] {
	case 'S':
		return Spades
	case 'H':
		return Hearts
	case 'D':
		return Diamonds
	case 'C':
		return Clubs
	default:
		panic("unknown color")
	}
}

func GetRank(card string) Rank {
	switch card[0] {
	case '2':
		return Two
	case '3':
		return Three
	case '4':
		return Four
	case '5':
		return Five
	case '6':
		return Six
	case '7':
		return Seven
	case '8':
		return Eight
	case '9':
		return Nine
	case 'T':
		return Ten
	case 'J':
		return Jack
	case 'Q':
		return Queen
	case 'K':
		return King
	case 'A':
		return Ace
	default:
		panic("unknown rank")

	}
}

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
	Undefined = -1
)

const (
	Spades Color = iota
	Hearts
	Diamonds
	Clubs
)

func parseNumber7(item string) (int, error) {
	number, err := strconv.Atoi(item)
	if err != nil {
		fmt.Println("unable to parse int")
		return 0, fmt.Errorf("unable to parse int")
	}
	return number, nil
}

func poker(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)

	in.Scan()
	dataInputsAsStrings := in.Text()

	dataInputs, err := parseNumber7(dataInputsAsStrings)
	if err != nil {
		return fmt.Errorf("unable to parse data inputs length")
	}

	var i = 0
	for i < dataInputs {

		var playerCards = make(map[int][]Card)

		//		from 0 - 12 - spades
		//		from 13 - 25 - hearts
		//		from 26 - 38 - diamonds
		//		from 39 - 51 - clubs
		//	 value - means player number
		var state [52]int

		in.Scan()
		playersLengthAsString := in.Text()

		playersLength, err := parseNumber7(playersLengthAsString)
		if err != nil {
			fmt.Println(playersLengthAsString)
			return fmt.Errorf("unable to parse players length")
		}

		var currentPlayer = 1
		for currentPlayer < playersLength+1 {

			in.Scan()
			playerCardsText := in.Text()

			cards := strings.Split(playerCardsText, " ")

			for _, card := range cards {

				color := GetColor(card)
				rank := GetRank(card)

				state[int(rank)+int(color*13)] = currentPlayer

				playerCards[currentPlayer] = append(playerCards[currentPlayer], Card{Rank: rank, Color: color})
			}

			currentPlayer++

		}

		var result []Card

		for i := 0; i < 52; i++ {
			c := Card{Rank: Rank(i % 13), Color: Color(i / 13)}
			if state[i] == 0 && FirstWins(c, playerCards, playersLength) {
				result = append(result, c)
			}
		}

		fmt.Fprintln(output, len(result))
		for _, card := range result {
			fmt.Fprintf(output, "%s%s\n", RankToString(card.Rank), ColorToString(card.Color))
		}

		i++
	}

	return nil
}

func ColorToString(c Color) string {
	switch c {
	case Spades:
		return "S"
	case Hearts:
		return "H"
	case Diamonds:
		return "D"
	case Clubs:
		return "C"
	default:
		panic("unknown color")
	}
}

func RankToString(r Rank) string {
	switch r {
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "T"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		panic("unknown rank")
	}
}

func FirstWins(c Card, playerCards map[int][]Card, playersCount int) bool {
	firstPlayerCards := playerCards[1]
	firstPlayer := Player{cards: [3]Card{firstPlayerCards[0], firstPlayerCards[1], c}}
	loser := false
	i := 2
	for !loser && i < playersCount+1 {
		loser = !firstPlayer.Wins(&Player{cards: [3]Card{playerCards[i][0], playerCards[i][1], c}})
		i++
	}

	return !loser
}

type Card struct {
	Rank  Rank
	Color Color
}

type Player struct {
	cards [3]Card
}

func (p *Player) Wins(other *Player) bool {
	firstRank, firstHasSet := p.IsSet()
	secondRank, secondHasSet := other.IsSet()

	if firstHasSet {
		return true
	}

	if secondHasSet {
		return false
	}

	firstRank, hasPair := p.IsPair()
	secondRank, hasOtherPair := other.IsPair()

	if (hasOtherPair && !hasPair) || (hasPair && hasOtherPair && firstRank < secondRank) {
		return false
	}

	if (hasPair && !hasOtherPair) || (hasPair && hasOtherPair && firstRank >= secondRank) {
		return true
	}

	firstHighestRank := p.GetHighest()
	otherHighestRank := other.GetHighest()

	return firstHighestRank >= otherHighestRank
}
