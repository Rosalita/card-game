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
	tx              int
	ty              int
	cardWidth       int
	cardHeight      int
	borderSize      int
	cards           []card
	maxNumCards     int
	cardsWide       int
	selectedIndex   *int
	selectionActive bool
}

type cardStackInput struct {
	width       int
	height      int
	maxNumCards int
	borderSize  int
	tx          int
	ty          int
	cardsWide   int
}

func (cs *cardStack) draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(cs.tx), float64(cs.ty))

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

		if (i+1)%cs.cardsWide == 0 {
			opts.GeoM.Translate(-float64(cs.cardWidth*cs.cardsWide), float64(cs.cardHeight))
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

	if len(cs.cards) < cs.maxNumCards {
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

func newCardStack(input cardStackInput) cardStack {

	defaultSelectedIndex := 0

	var cs cardStack

	cs.tx = input.tx
	cs.ty = input.ty
	cs.cardWidth = input.width
	cs.cardHeight = input.height
	cs.maxNumCards = input.maxNumCards
	cs.cardsWide = input.cardsWide
	cs.borderSize = input.borderSize
	cs.selectedIndex = &defaultSelectedIndex
	cs.selectionActive = false
	return cs
}

func initialiseCards(cs *cardStack) {

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

	cardStockInput := cardStackInput{
		tx:          0,
		ty:          200,
		width:       defaultCardWidth,
		height:      defaultCardHeight,
		maxNumCards: defaultDeckSize,
		cardsWide:   8,
		borderSize:  defaultBorderSize,
	}

	cardStock = newCardStack(cardStockInput)
	initialiseCards(&cardStock)
	cardStock.shuffle()

	p1HandInput := cardStackInput{
		tx:          40,
		ty:          450,
		width:       defaultCardWidth,
		height:      defaultCardHeight,
		maxNumCards: 8,
		cardsWide:   8,
		borderSize:  defaultBorderSize,
	}

	p2HandInput := cardStackInput{
		tx:          40,
		ty:          0,
		width:       defaultCardWidth,
		height:      defaultCardHeight,
		maxNumCards: 8,
		cardsWide:   8,
		borderSize:  defaultBorderSize,
	}

	player1Hand = newCardStack(p1HandInput)
	player2Hand = newCardStack(p2HandInput)

	playAreaInput := cardStackInput{
		tx:          200,
		ty:          200,
		width:       defaultCardWidth,
		height:      defaultCardHeight,
		maxNumCards: 8,
		cardsWide:   4,
		borderSize:  defaultBorderSize,
	}

	playArea = newCardStack(playAreaInput)

	p1DiscardInput := cardStackInput{
		tx:          450,
		ty:          450,
		width:       defaultCardWidth,
		height:      defaultCardHeight,
		maxNumCards: defaultDeckSize,
		cardsWide:   8,
		borderSize:  defaultBorderSize,
	}

	p2DiscardInput := cardStackInput{
		tx:          450,
		ty:          0,
		width:       defaultCardWidth,
		height:      defaultCardHeight,
		maxNumCards: defaultDeckSize,
		cardsWide:   8,
		borderSize:  defaultBorderSize,
	}

	player1DiscardPile = newCardStack(p1DiscardInput)
	player2DiscardPile = newCardStack(p2DiscardInput)

	dealer := 1 //TO DO pick cards to see who goes first

	dealCards(dealer, &cardStock, &playArea, &player1Hand, &player2Hand)
	switch dealer {
	case 1:
		activeHand = &player1Hand
	case 2:
		activeHand = &player2Hand
	}
	activeHand.selectionActive = true

}

func updatePlay(screen *ebiten.Image) error {

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		activeHand.incrementSelected()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		activeHand.decrementSelected()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		moveCard(&cardStock, &player1Hand, len(cardStock.cards)-1)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		moveCard(&cardStock, &player2Hand, len(cardStock.cards)-1)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		moveCard(&player2Hand, &player2DiscardPile, len(player2Hand.cards)-1)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		moveCard(&player1Hand, &player1DiscardPile, len(player1Hand.cards)-1)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
		return nil
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	cardStock.draw(screen)
	playArea.draw(screen)
	player1DiscardPile.draw(screen)
	player2DiscardPile.draw(screen)
	player1Hand.draw(screen)
	player2Hand.draw(screen)

	return nil
}
