package main

import (
	"errors"
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten"
	"github.com/stretchr/testify/assert"
)

func makeTestCards() (cards []card, red, green, blue color.NRGBA) {
	red = color.NRGBA{0xff, 0x00, 0x00, 0xff}
	green = color.NRGBA{0x00, 0xff, 0x00, 0xff}
	blue = color.NRGBA{0x00, 0x00, 0xff, 0xff}

	colours := []*color.NRGBA{&red, &green, &blue}

	var testCards []card
	for i := 0; i < 3; i++ {
		img, _ := ebiten.NewImage(20, 40, ebiten.FilterNearest)
		newCard := card{
			image:  img,
			colour: colours[i],
		}
		testCards = append(testCards, newCard)
	}
	return testCards, red, green, blue
}

func TestCardStackAddCard(t *testing.T) {

	testCards, _, _, _ := makeTestCards()
	stack1 := cardStack{
		cards:   testCards,
		maxSize: 8,
	}

	testCards, _, _, _ = makeTestCards()
	stack2 := cardStack{
		cards:   testCards,
		maxSize: 3,
	}

	yellow := color.NRGBA{0xff, 0xff, 0x00, 0xff}
	img, _ := ebiten.NewImage(20, 40, ebiten.FilterNearest)
	newCard := card{
		image:  img,
		colour: &yellow,
	}

	tests := []struct {
		cardStack          cardStack
		card          card
		handSizeAfter int
		err           error
	}{
		{stack1, newCard, 4, nil},
		{stack2, newCard, 3, errors.New("cardStack is full")},
	}

	for _, test := range tests {
		err := test.cardStack.addCard(test.card)
		assert.Equal(t, test.err, err)
		assert.Equal(t, test.handSizeAfter, len(test.cardStack.cards))
	}
}


func TestCardStackRemoveCard(t *testing.T) {

	testCards, red, green, blue := makeTestCards()
	stack1 := cardStack{
		cards: testCards,
	}

	testCards, red, green, blue = makeTestCards()
	stack2 := cardStack{
		cards: testCards,
	}

	testCards, red, green, blue = makeTestCards()
	stack3 := cardStack{
		cards: testCards,
	}

	testCards, red, green, blue = makeTestCards()
	stack4 := cardStack{
		cards: testCards,
	}

	testCards, red, green, blue = makeTestCards()
	stack5 := cardStack{
		cards: testCards,
	}

	stack6 := cardStack{}

	tests := []struct {
		cardStack         cardStack
		index             int
		cardRemovedColour color.NRGBA
		stackSizeAfter    int
		err               error
	}{
		{stack1, 0, red, 2, nil},
		{stack2, 1, green, 2, nil},
		{stack3, 2, blue, 2, nil},
		{stack4, 4, color.NRGBA{}, 3, errors.New("index out of range")},
		{stack5, -1, color.NRGBA{}, 3, errors.New("index out of range")},
		{stack6, 0, color.NRGBA{}, 0, errors.New("no removable cards")},
	}

	for _, test := range tests {
		card, err := test.cardStack.removeCard(test.index)
		assert.Equal(t, test.err, err)

		if err == nil {
			assert.Equal(t, test.cardRemovedColour, *card.colour)
		}
		assert.Equal(t, test.stackSizeAfter, len(test.cardStack.cards))
	}
}
