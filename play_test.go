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
			image:            img,
			colour:           colours[i],
			selectedColour:   &red,
			unselectedColour: &blue,
		}
		testCards = append(testCards, newCard)
	}
	return testCards, red, green, blue
}

func TestCardZoneAddCard(t *testing.T) {

	testCards, _, _, _ := makeTestCards()
	zone1 := cardZone{
		cards:       testCards,
		maxNumCards: 8,
		cardWidth:   30,
		cardHeight:  50,
	}

	testCards, _, _, _ = makeTestCards()
	zone2 := cardZone{
		cards:       testCards,
		maxNumCards: 3,
	}

	yellow := color.NRGBA{0xff, 0xff, 0x00, 0xff}
	img, _ := ebiten.NewImage(20, 40, ebiten.FilterNearest)

	newCard := card{
		image:            img,
		colour:           &yellow,
		selectedColour:   &yellow,
		unselectedColour: &yellow,
	}

	tests := []struct {
		cardZone     cardZone
		card          card
		handSizeAfter int
		err           error
	}{
		{zone1, newCard, 4, nil},
		{zone2, newCard, 3, errors.New("cardZone is full")},
	}

	for _, test := range tests {
		err := test.cardZone.addCard(test.card)
		assert.Equal(t, test.err, err)
		assert.Equal(t, test.handSizeAfter, len(test.cardZone.cards))
	}
}

func TestCardZoneRemoveCard(t *testing.T) {

	testCards, red, green, blue := makeTestCards()
	zone1 := cardZone{
		cards: testCards,
	}

	testCards, red, green, blue = makeTestCards()
	zone2 := cardZone{
		cards: testCards,
	}

	testCards, red, green, blue = makeTestCards()
	zone3 := cardZone{
		cards: testCards,
	}

	testCards, red, green, blue = makeTestCards()
	zone4 := cardZone{
		cards: testCards,
	}

	testCards, red, green, blue = makeTestCards()
	zone5 := cardZone{
		cards: testCards,
	}

	zone6 := cardZone{}

	tests := []struct {
		cardZone         cardZone
		index             int
		cardRemovedColour color.NRGBA
		zoneSizeAfter    int
		err               error
	}{
		{zone1, 0, red, 2, nil},
		{zone2, 1, green, 2, nil},
		{zone3, 2, blue, 2, nil},
		{zone4, 4, color.NRGBA{}, 3, errors.New("index out of range")},
		{zone5, -1, color.NRGBA{}, 3, errors.New("index out of range")},
		{zone6, 0, color.NRGBA{}, 0, errors.New("no removable cards")},
	}

	for _, test := range tests {
		card, err := test.cardZone.removeCard(test.index)
		assert.Equal(t, test.err, err)

		if err == nil {
			assert.Equal(t, test.cardRemovedColour, *card.colour)
		}
		assert.Equal(t, test.zoneSizeAfter, len(test.cardZone.cards))
	}
}
