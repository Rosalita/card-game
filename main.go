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
	screensize
)

var (
	state          gameState
	mainMenu       lm.ListMenu
	optionsMenu    lm.ListMenu
	screensizeMenu lm.ListMenu
	white          = &color.NRGBA{0xff, 0xff, 0xff, 0xff}
	pink           = &color.NRGBA{0xff, 0x69, 0xb4, 0xff}
)

func update(screen *ebiten.Image) error {

	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})

	if state == titleScreen {

		//ebitenutil.DebugPrint(screen, "Title screen")

		// call to draw needs to take in location to draw
		mainMenu.Draw(screen)

		//fmt.Println(getCentre(screen))

		// opts := &ebiten.DrawImageOptions{}
		// opts.GeoM.Translate(200, 24)

		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			mainMenu.DecrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			mainMenu.IncrementSelected()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			switch mainMenu.GetSelectedItem() {
			case "options":
				state = options
			case "quit":
				os.Exit(0)
			}
			return nil
		}

	}
	if state == options {
		optionsMenu.Draw(screen)

		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			optionsMenu.DecrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			optionsMenu.IncrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			state = titleScreen
			return nil
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			switch optionsMenu.GetSelectedItem() {
			case "screensize":
				state = screensize
				return nil
			}
			return nil
		}
	}
	if state == screensize {
		screensizeMenu.Draw(screen)

		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			screensizeMenu.DecrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			screensizeMenu.IncrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			state = options
			return nil
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			switch screensizeMenu.GetSelectedItem() {
			case "400x300":
				ebiten.SetScreenSize(400, 300)
			case "600x400":
				ebiten.SetScreenSize(600, 400)
			case "800x600":
				ebiten.SetScreenSize(800, 600)
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
