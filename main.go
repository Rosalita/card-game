package main

import (
	//"fmt"

	//"bytes"
	//"image"
	"image/color"
	//"log"
	"os"

	//im "github.com/Rosalita/ebiten-pkgs/imagemenu"
	lm "github.com/Rosalita/ebiten-pkgs/listmenu"
	"github.com/hajimehoshi/ebiten"

	//	"github.com/hajimehoshi/ebiten/ebitenutil" // required for debug text
	"github.com/hajimehoshi/ebiten/inpututil" // required for isKeyJustPressed
)

type gameState int

const (
	titleScreen gameState = iota
	options
	charCreation
	quit
)

var (
	state    gameState
	mainMenu lm.ListMenu
	white    = &color.NRGBA{0xff, 0xff, 0xff, 0xff}
	pink     = &color.NRGBA{0xff, 0x69, 0xb4, 0xff}
)

func update(screen *ebiten.Image) error {

	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})

	if state == titleScreen {

		//ebitenutil.DebugPrint(screen, "Title screen")

		// call to draw needs to take in location to draw
		mainMenu.Draw(screen)

		//	fmt.Println(getCentre(screen))

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(200, 24)

		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			mainMenu.DecrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			mainMenu.IncrementSelected()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			switch mainMenu.GetSelectedItem() {
			case "playButton":
				state = charCreation
			case "quitButton":
				os.Exit(0)
			}
			return nil
		}

	}

	return nil
}

func main() {

	initMenus()
	state = titleScreen

	if err := ebiten.Run(update, 400, 300, 2, "Card Game"); err != nil {
		panic(err)
	}
}
