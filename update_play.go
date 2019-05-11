package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil" // required for isKeyJustPressed
)

type card struct {
	image  *ebiten.Image
	colour *color.NRGBA
}

type deck struct {
	cards []card
}

func (d *deck) draw(screen *ebiten.Image, cardsWide int) {
	opts := &ebiten.DrawImageOptions{}

	w, h := d.cards[0].image.Size()

	for i, crd := range d.cards {

		crd.image.Fill(crd.colour)

		screen.DrawImage(crd.image, opts)
		opts.GeoM.Translate(float64(w), 0)

		if (i+1)%cardsWide == 0 {
			opts.GeoM.Translate(-float64(w*cardsWide), float64(h))
		}
	}
}

func updatePlay(screen *ebiten.Image) error {

	deck := newDeck(50, 70)

	deck.draw(screen, 6)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
		return nil
	}

	return nil
}

func newDeck(width, height int) deck {

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

	var deck deck

	for _, colour := range monthColours {

		for i := 0; i < cardsPerMonth; i++ {
			img, err := ebiten.NewImage(width, height, ebiten.FilterNearest)

			if err != nil {
				log.Println(err)
			}

			newCard := card{
				image:  img,
				colour: colour,
			}

			deck.cards = append(deck.cards, newCard)
		}
	}
	return deck
}
