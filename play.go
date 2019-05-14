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
	borderImage      *ebiten.Image
	image            *ebiten.Image
	colour           *color.NRGBA
	unselectedColour *color.NRGBA
	selectedColour   *color.NRGBA
}

type cardStack struct {
	cardWidth       int
	cardHeight      int
	borderSize      int
	cards           []card
	maxSize         int
	selectedIndex   *int
	selectionActive bool
}

func (cs *cardStack) draw(screen *ebiten.Image, tx, ty float64, cardsWide int) {

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(tx, ty)

	for i, card := range cs.cards {

		if cs.selectionActive {
			if *cs.selectedIndex == i {
				card.borderImage.Fill(card.selectedColour)
			} else {
				card.borderImage.Fill(card.unselectedColour)
			}
		}

		screen.DrawImage(card.borderImage, opts)
		opts.GeoM.Translate(float64(cs.borderSize), float64(cs.borderSize))
		screen.DrawImage(card.image, opts)
		opts.GeoM.Translate(-float64(cs.borderSize), -float64(cs.borderSize))
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

func (cs *cardStack) addCard(cardToAdd card) error {

	if len(cs.cards) < cs.maxSize {
		cs.cards = append(cs.cards, cardToAdd)
		return nil
	}
	return errors.New("cardStack is full")
}

func (cs *cardStack) incrementSelected() {
	maxIndex := len(cs.cards) - 1
	if *cs.selectedIndex < maxIndex {
		*cs.selectedIndex++
	}
}

func (cs *cardStack) decrementSelected() {
	minIndex := 0
	if *cs.selectedIndex > minIndex {
		*cs.selectedIndex--
	}
}

func moveCard(source, destination *cardStack, index int) {
	cardToMove, err := source.removeCard(index)
	if err != nil {
		log.Println(err)
		return
	}
	err = destination.addCard(cardToMove)
	if err != nil {
		log.Println(err)
		source.addCard(cardToMove)
		return
	}
}

func newCardStack(width, height, maxSize, borderSize int) cardStack {

	defaultSelectedIndex := 0

	var cs cardStack

	cs.cardWidth = width
	cs.cardHeight = height
	cs.maxSize = maxSize
	cs.borderSize = borderSize
	cs.selectedIndex = &defaultSelectedIndex
	cs.selectionActive = false

	return cs
}

func newDeck(width, height, maxSize, borderSize int) cardStack {

	defaultSelectedIndex := 0

	var cs cardStack

	cs.cardWidth = width
	cs.cardHeight = height
	cs.maxSize = maxSize
	cs.borderSize = borderSize
	cs.selectedIndex = &defaultSelectedIndex
	cs.selectionActive = false

	//unselected border colour
	unselected := &color.NRGBA{0x44, 0x44, 0x44, 0xff}

	//selected border colour
	selected := &color.NRGBA{0xff, 0x00, 0xff, 0xff}

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
			borderImg, err := ebiten.NewImage(cs.cardWidth, cs.cardHeight, ebiten.FilterNearest)
			borderImg.Fill(unselected)
			img, err := ebiten.NewImage(cs.cardWidth-(cs.borderSize*2), cs.cardHeight-(cs.borderSize*2), ebiten.FilterNearest)

			img.Fill(colour)

			if err != nil {
				log.Println(err)
			}

			newCard := card{
				borderImage:      borderImg,
				image:            img,
				colour:           colour,
				unselectedColour: unselected,
				selectedColour:   selected,
			}

			cs.cards = append(cs.cards, newCard)
		}
	}
	return cs
}

func newHand(cardWidth, cardHeight, maxSize, borderSize int) cardStack {

	defaultSelectedIndex := 0

	var hand cardStack

	hand.cardWidth = cardWidth
	hand.cardHeight = cardHeight
	hand.maxSize = maxSize
	hand.borderSize = borderSize
	hand.selectedIndex = &defaultSelectedIndex
	hand.selectionActive = false
	return hand
}

func dealCards(dealer int, cardStock, playArea, player1Hand, player2Hand *cardStack) {

	var dealerHand, opponentHand *cardStack
	switch dealer {
	case 1:
		dealerHand = player1Hand
		opponentHand = player2Hand
	case 2:
		dealerHand = player2Hand
		opponentHand = player1Hand
	}

	//cards are dealt out in two rounds
	for round := 0; round < 2; round++ {

		// dealer deals 4 cards to opponent
		for i := 0; i < 4; i++ {
			moveCard(cardStock, opponentHand, len(cardStock.cards)-1)
		}

		// then dealer deals 4 cards to play area
		for i := 0; i < 4; i++ {
			moveCard(cardStock, playArea, len(cardStock.cards)-1)
		}

		// then dealer deals 4 cards to self
		for i := 0; i < 4; i++ {
			moveCard(cardStock, dealerHand, len(cardStock.cards)-1)
		}
	}
}

func initialisePlay() {

	defaultCardWidth := 30
	defaultCardHeight := 50
	defaultDeckSize := 48
	defaultBorderSize := 2

	cardStock = newDeck(defaultCardWidth, defaultCardHeight, defaultDeckSize, defaultBorderSize)
	cardStock.shuffle()

	player1Hand = newHand(defaultCardWidth, defaultCardHeight, 8, defaultBorderSize)
	player2Hand = newHand(defaultCardWidth, defaultCardHeight, 8, defaultBorderSize)

	playArea = newCardStack(defaultCardWidth, defaultCardHeight, 8, defaultBorderSize)

	player1DiscardPile = newCardStack(defaultCardWidth, defaultCardHeight, defaultDeckSize, defaultBorderSize)
	player2DiscardPile = newCardStack(defaultCardWidth, defaultCardHeight, defaultDeckSize, defaultBorderSize)

	dealer := 1 //TO DO pick cards to see who goes first

	dealCards(dealer, &cardStock, &playArea, &player1Hand, &player2Hand)
	switch dealer{
	case 1:
		activeHand = &player1Hand
	case 2:
		activeHand = &player2Hand 
	}
	activeHand.selectionActive = true

}

func updatePlay(screen *ebiten.Image) error {

	cardStock.draw(screen, 0, 200, 8)
	playArea.draw(screen, 200, 200, 4)
	player1DiscardPile.draw(screen, 450, 450, 8)
	player2DiscardPile.draw(screen, 450, 0, 8)
	player1Hand.draw(screen, 40, 450, 8)
	player2Hand.draw(screen, 40, 0, 8)

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		activeHand.incrementSelected()
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		activeHand.decrementSelected()
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		moveCard(&cardStock, &player1Hand, len(cardStock.cards)-1)
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		moveCard(&cardStock, &player2Hand, len(cardStock.cards)-1)
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		moveCard(&player2Hand, &player2DiscardPile, len(player2Hand.cards)-1)
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		moveCard(&player1Hand, &player1DiscardPile, len(player1Hand.cards)-1)
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
		return nil
	}

	return nil
}
