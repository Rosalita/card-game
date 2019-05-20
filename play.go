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
	selectedBorderImage   *ebiten.Image
	unselectedBorderImage *ebiten.Image
	image                 *ebiten.Image
	colour                *color.NRGBA
	unselectedColour      *color.NRGBA
	selectedColour        *color.NRGBA
}

type cardZone struct {
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

type cardZoneInput struct {
	width       int
	height      int
	maxNumCards int
	borderSize  int
	tx          int
	ty          int
	cardsWide   int
}

func (cz *cardZone) draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(cz.tx), float64(cz.ty))

	for i, card := range cz.cards {

		screen.DrawImage(card.unselectedBorderImage, opts)

		if cz.selectionActive {
			if *cz.selectedIndex == i {
				screen.DrawImage(card.selectedBorderImage, opts)
			}
		}

		opts.GeoM.Translate(float64(cz.borderSize), float64(cz.borderSize))

		screen.DrawImage(card.image, opts)
		opts.GeoM.Translate(-float64(cz.borderSize), -float64(cz.borderSize))
		opts.GeoM.Translate(float64(cz.cardWidth), 0)

		if (i+1)%cz.cardsWide == 0 {
			opts.GeoM.Translate(-float64(cz.cardWidth*cz.cardsWide), float64(cz.cardHeight))
		}
	}
}
func (cz *cardZone) shuffle() {

	if len(cz.cards) < 2 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cz.cards), func(i, j int) { cz.cards[i], cz.cards[j] = cz.cards[j], cz.cards[i] })

	return
}

func (cz *cardZone) removeCard(index int) (card, error) {

	totalCards := len(cz.cards)
	if totalCards < 1 {
		return card{}, errors.New("no removable cards")
	}

	if index > (totalCards-1) || index < 0 {
		return card{}, errors.New("index out of range")
	}

	removedCard := cz.cards[index]

	cz.cards = append(cz.cards[:index], cz.cards[index+1:]...)

	return removedCard, nil
}

func (cz *cardZone) addCard(cardToAdd card) error {

	if len(cz.cards) < cz.maxNumCards {

		selectedBorderImg, err := ebiten.NewImage(cz.cardWidth, cz.cardHeight, ebiten.FilterNearest)
		if err != nil {
			return err
		}
		selectedBorderImg.Fill(cardToAdd.selectedColour)
		unselectedBorderImg, err := ebiten.NewImage(cz.cardWidth, cz.cardHeight, ebiten.FilterNearest)
		if err != nil {
			return err
		}
		unselectedBorderImg.Fill(cardToAdd.unselectedColour)
		img, err := ebiten.NewImage(cz.cardWidth-(cz.borderSize*2), cz.cardHeight-(cz.borderSize*2), ebiten.FilterNearest)
		img.Fill(cardToAdd.colour)

		cardToAdd.selectedBorderImage = selectedBorderImg
		cardToAdd.unselectedBorderImage = unselectedBorderImg
		cardToAdd.image = img

		cz.cards = append(cz.cards, cardToAdd)
		return nil
	}
	return errors.New("cardZone is full")
}

func (cz *cardZone) incrementSelected() {
	maxIndex := len(cz.cards) - 1
	if *cz.selectedIndex < maxIndex {
		*cz.selectedIndex++
	}
}

func (cz *cardZone) decrementSelected() {
	minIndex := 0
	if *cz.selectedIndex > minIndex {
		*cz.selectedIndex--
	}
}

func moveCard(source, destination *cardZone, index int) {
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

func newCardZone(input cardZoneInput) cardZone {

	defaultSelectedIndex := 0

	var cz cardZone

	cz.tx = input.tx
	cz.ty = input.ty
	cz.cardWidth = input.width
	cz.cardHeight = input.height
	cz.maxNumCards = input.maxNumCards
	cz.cardsWide = input.cardsWide
	cz.borderSize = input.borderSize
	cz.selectedIndex = &defaultSelectedIndex
	cz.selectionActive = false
	return cz
}

func initialiseCards(cz *cardZone) {

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
			selectedBorderImg, err := ebiten.NewImage(cz.cardWidth, cz.cardHeight, ebiten.FilterNearest)
			selectedBorderImg.Fill(selected)
			unselectedBorderImg, err := ebiten.NewImage(cz.cardWidth, cz.cardHeight, ebiten.FilterNearest)
			unselectedBorderImg.Fill(unselected)
			img, err := ebiten.NewImage(cz.cardWidth-(cz.borderSize*2), cz.cardHeight-(cz.borderSize*2), ebiten.FilterNearest)
			img.Fill(colour)

			if err != nil {
				log.Println(err)
			}

			newCard := card{
				selectedBorderImage:   selectedBorderImg,
				unselectedBorderImage: unselectedBorderImg,
				image:                 img,
				colour:                colour,
				unselectedColour:      unselected,
				selectedColour:        selected,
			}

			cz.cards = append(cz.cards, newCard)
		}
	}
}

func dealCards(dealer int, deck, playArea, player1Hand, player2Hand *cardZone) {

	var dealerHand, opponentHand *cardZone
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
			moveCard(deck, opponentHand, len(deck.cards)-1)
		}

		// then dealer deals 4 cards to play area
		for i := 0; i < 4; i++ {
			moveCard(deck, playArea, len(deck.cards)-1)
		}

		// then dealer deals 4 cards to self
		for i := 0; i < 4; i++ {
			moveCard(deck, dealerHand, len(deck.cards)-1)
		}
	}
}

func initialisePlay() {

	defaultCardWidth := 90
	defaultCardHeight := 150
	defaultDeckSize := 48
	defaultBorderSize := 2

	deckInput := cardZoneInput{
		tx:          0,
		ty:          200,
		width:       defaultCardWidth,
		height:      10,
		maxNumCards: defaultDeckSize,
		cardsWide:   1,
		borderSize:  defaultBorderSize,
	}

	deck = newCardZone(deckInput)

	initialiseCards(&deck)
	deck.shuffle()

	p1HandInput := cardZoneInput{
		tx:          40,
		ty:          600,
		width:       defaultCardWidth,
		height:      defaultCardHeight,
		maxNumCards: 8,
		cardsWide:   8,
		borderSize:  defaultBorderSize,
	}

	p2HandInput := cardZoneInput{
		tx:          40,
		ty:          0,
		width:       defaultCardWidth,
		height:      defaultCardHeight,
		maxNumCards: 8,
		cardsWide:   8,
		borderSize:  defaultBorderSize,
	}

	player1Hand = newCardZone(p1HandInput)
	player2Hand = newCardZone(p2HandInput)

	playAreaInput := cardZoneInput{
		tx:          200,
		ty:          200,
		width:       defaultCardWidth,
		height:      defaultCardHeight,
		maxNumCards: 8,
		cardsWide:   4,
		borderSize:  defaultBorderSize,
	}

	playArea = newCardZone(playAreaInput)

	p1DiscardInput := cardZoneInput{
		tx:          450,
		ty:          450,
		width:       defaultCardWidth,
		height:      defaultCardHeight,
		maxNumCards: defaultDeckSize,
		cardsWide:   8,
		borderSize:  defaultBorderSize,
	}

	p2DiscardInput := cardZoneInput{
		tx:          450,
		ty:          0,
		width:       defaultCardWidth,
		height:      defaultCardHeight,
		maxNumCards: defaultDeckSize,
		cardsWide:   8,
		borderSize:  defaultBorderSize,
	}

	player1DiscardPile = newCardZone(p1DiscardInput)
	player2DiscardPile = newCardZone(p2DiscardInput)

	dealer := 1 //TO DO pick cards to see who goes first

	dealCards(dealer, &deck, &playArea, &player1Hand, &player2Hand)
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
		moveCard(&deck, &player1Hand, len(deck.cards)-1)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		moveCard(&deck, &player2Hand, len(deck.cards)-1)
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

	deck.draw(screen)
	playArea.draw(screen)
	player1DiscardPile.draw(screen)
	player2DiscardPile.draw(screen)
	player1Hand.draw(screen)
	player2Hand.draw(screen)

	return nil
}
