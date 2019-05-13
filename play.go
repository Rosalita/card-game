package main

import (
	"errors"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil" // required for isKeyJustPressed
)

type card struct {
	image  *ebiten.Image
	colour *color.NRGBA
}

type cardStack struct {
	cardWidth  int
	cardHeight int
	cards      []card // index 0 is the bottom of the deck, representing the last card to be drawn from the deck
}

func (cs *cardStack) draw(screen *ebiten.Image, tx, ty float64, cardsWide int) {

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(tx, ty)

	for i, card := range cs.cards {

		card.image.Fill(card.colour)

		screen.DrawImage(card.image, opts)
		opts.GeoM.Translate(float64(cs.cardWidth), 0)

		if (i+1)%cardsWide == 0 {
			opts.GeoM.Translate(-float64(cs.cardWidth*cardsWide), float64(cs.cardHeight))
		}
	}
}
func (cs *cardStack) shuffle() {

	if len(cs.cards) < 2 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cs.cards), func(i, j int) { cs.cards[i], cs.cards[j] = cs.cards[j], cs.cards[i] })

	return
}

func (cs *cardStack) removeCard(index int) (card, error) {

	totalCards := len(cs.cards)
	if totalCards < 1 {
		return card{}, errors.New("no removable cards")
	}

	if index > (totalCards-1) || index < 0 {
		return card{}, errors.New("index out of range")
	}

	removedCard := cs.cards[index]

	cs.cards = append(cs.cards[:index], cs.cards[index+1:]...)

	return removedCard, nil
}

func (cs *cardStack) addCard(cardToAdd card) {
	cs.cards = append(cs.cards, cardToAdd)
}

func newCardStack(width, height int) cardStack {
	var cs cardStack
	cs.cardWidth = width
	cs.cardHeight = height
	return cs
}

func newDeck(width, height int) cardStack {

	var cs cardStack

	cs.cardWidth = width
	cs.cardHeight = height

	//January colour light frosty blue
	janCol := &color.NRGBA{0x81, 0xc1, 0xff, 0xff}

	//February colour light pink
	febCol := &color.NRGBA{0xfc, 0xb5, 0xb3, 0xff}

	//March colour light yellow green
	marCol := &color.NRGBA{0xaf, 0xfe, 0x88, 0xff}

	// april pastel yellow
	aprCol := &color.NRGBA{0xfe, 0xff, 0x99, 0xff}

	// may lavender
	mayCol := &color.NRGBA{0xc3, 0xa8, 0xff, 0xff}

	// june turquoise blue green
	junCol := &color.NRGBA{0x7a, 0xff, 0xd5, 0xff}

	// july dark blue
	julCol := &color.NRGBA{0x22, 0x3e, 0xcd, 0xff}

	// august purple
	augCol := &color.NRGBA{0x9f, 0x22, 0xcd, 0xff}

	//september golden
	sepCol := &color.NRGBA{0xff, 0xd0, 0x00, 0xff}

	//october orange
	octCol := &color.NRGBA{0xff, 0x8d, 0x28, 0xff}

	// november red
	novCol := &color.NRGBA{0xff, 0x46, 0x28, 0xff}

	// december white
	decCol := &color.NRGBA{0xff, 0xff, 0xff, 0xff}

	monthColours := []*color.NRGBA{
		janCol, febCol, marCol, aprCol, mayCol, junCol, julCol, augCol, sepCol, octCol, novCol, decCol,
	}

	cardsPerMonth := 4

	for _, colour := range monthColours {

		for i := 0; i < cardsPerMonth; i++ {
			img, err := ebiten.NewImage(cs.cardWidth, cs.cardHeight, ebiten.FilterNearest)

			if err != nil {
				log.Println(err)
			}

			newCard := card{
				image:  img,
				colour: colour,
			}

			cs.cards = append(cs.cards, newCard)
		}
	}
	return cs
}

type hand struct {
	originDeck  *cardStack
	discardPile *cardStack
	cards       []card // index 0 is first card drawn
	maxSize     int
}

// TO DO
// need an interface for card stacks, hands, play areas etc
// something that can remove cards and add a card
// move a card function removes from source, adds to destination, accepts interface
// have started writing add and remove card methods for hand and card stacks

func (h *hand) cardDraw() { //draws a card from the origin deck and adds it to the hand

	numDeckCards := len(h.originDeck.cards)
	if numDeckCards < 1 {
		return
	}

	if len(h.cards) >= h.maxSize {
		return
	}

	drawnCard := h.originDeck.cards[numDeckCards-1]

	h.originDeck.cards = append(h.originDeck.cards[:numDeckCards-1])

	h.cards = append(h.cards, drawnCard)

}

func (h *hand) cardDiscard() { //draws a card from the origin deck and adds it to the hand

	numHandCards := len(h.cards)
	if numHandCards < 1 {
		return
	}

	discardedCard := h.cards[numHandCards-1]

	h.cards = append(h.cards[:numHandCards-1])

	h.discardPile.cards = append(h.discardPile.cards, discardedCard)

}

func (h *hand) removeCard(index int) (card, error) {

	totalCards := len(h.cards)
	if totalCards < 1 {
		return card{}, errors.New("no removable cards")
	}

	if index > (totalCards-1) || index < 0 {
		return card{}, errors.New("index out of range")
	}

	removedCard := h.cards[index]

	h.cards = append(h.cards[:index], h.cards[index+1:]...)

	return removedCard, nil
}

func (h *hand) addCard(cardToAdd card) error {

	if len(h.cards) < h.maxSize {
		h.cards = append(h.cards, cardToAdd)
		return nil
	}
	return errors.New("hand is full")
}

func (h *hand) draw(screen *ebiten.Image, tx, ty float64) {

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(tx, ty)

	cardWidth := h.originDeck.cardWidth

	for _, card := range h.cards {

		card.image.Fill(card.colour)

		screen.DrawImage(card.image, opts)
		opts.GeoM.Translate(float64(cardWidth), 0)

	}
}

func newHand(maxSize int, originDeck *cardStack, discardPile *cardStack) hand {
	var hand hand
	hand.maxSize = maxSize
	hand.originDeck = originDeck
	hand.discardPile = discardPile
	return hand
}

// func dealCards(dealer int, cardStock, playArea *cardStack, player1Hand, player2Hand *hand) {

// 	var dealerHand, opponentHand *hand
// 	switch dealer {
// 	case 1:
// 		dealerHand = player1Hand
// 		opponentHand = player2Hand
// 	case 2:
// 		dealerHand = player2Hand
// 		opponentHand = player1Hand
// 	}

// 	// dealer deals 4 cards to opponent
// 	// then dealer deals 4 cards to play area
// 	// then dealer deals 4 cards to self
// 	// then repeat this once

// }

func initialisePlay() {

	defaultCardWidth := 20
	defaultCardHeight := 40

	cardStock = newDeck(defaultCardWidth, defaultCardHeight)
	cardStock.shuffle()

	player1Hand = newHand(6, &cardStock, &player1DiscardPile)
	player2Hand = newHand(6, &cardStock, &player2DiscardPile)

	playArea = newCardStack(defaultCardWidth, defaultCardHeight)

	player1DiscardPile = newCardStack(defaultCardWidth, defaultCardHeight)
	player2DiscardPile = newCardStack(defaultCardWidth, defaultCardHeight)

	//dealer := 1 //to do pick cards to see who goes first

	//dealCards(dealer, &cardStock, player1Hand, player2Hand, playArea)

}

func updatePlay(screen *ebiten.Image) error {

	cardStock.draw(screen, 0, 150, 8)
	player1DiscardPile.draw(screen, 450, 450, 8)
	player2DiscardPile.draw(screen, 450, 0, 8)
	player1Hand.draw(screen, 40, 450)
	player2Hand.draw(screen, 40, 0)

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {

		player1Hand.cardDraw()
		// card, err := cardStock.removeCard()

		// if err == nil{
		// 	player1DiscardPile.addCard(card)
		// }

		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		player2Hand.cardDraw()
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		player2Hand.cardDiscard()
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		player1Hand.cardDiscard()
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
		return nil
	}

	return nil
}
