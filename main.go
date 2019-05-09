package main

import (
	"image/color"
	"os"
	"math"

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
	activeMenu     lm.ListMenu
	mainMenu       lm.ListMenu
	optionsMenu    lm.ListMenu
	screensizeMenu lm.ListMenu
	white          = &color.NRGBA{0xff, 0xff, 0xff, 0xff}
	pink           = &color.NRGBA{0xff, 0x69, 0xb4, 0xff}
	bestRatio = 1.0
)


func update(screen *ebiten.Image) error {

	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})

	if state == titleScreen {

		activeMenu = mainMenu
		activeMenu.SetScale(bestRatio)

		activeMenu.Draw(screen)

		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			activeMenu.DecrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			activeMenu.IncrementSelected()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			switch activeMenu.GetSelectedItem() {
			case "options":
				state = options
			case "quit":
				os.Exit(0)
			}
			return nil
		}

	}

	if state == options {
		activeMenu = optionsMenu
		activeMenu.SetScale(bestRatio)

		activeMenu.Draw(screen)

		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			activeMenu.DecrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			activeMenu.IncrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			state = titleScreen
			return nil
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			switch activeMenu.GetSelectedItem() {
			case "screensize":
				state = screensize
				return nil
			}
			return nil
		}
	}

	if state == screensize {
		activeMenu = screensizeMenu
		activeMenu.SetScale(bestRatio)

		activeMenu.Draw(screen)

		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			activeMenu.DecrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			activeMenu.IncrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			state = options
			return nil
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {

			var newWidth, newHeight int

			switch activeMenu.GetSelectedItem() {
			case "640x480":
				newWidth = 640
				newHeight = 480
			case "800x600":
				newWidth = 800
				newHeight = 600
			case "1024x768":
				newWidth = 1024
				newHeight = 768
			case "1280x720":
				newWidth = 1280
				newHeight = 720
			case "1336x768":
				newWidth = 1336
				newHeight = 768
			case "1440x1080":
				newWidth = 1440
				newHeight = 1080
			case "1600x900":
				newWidth = 1600
				newHeight = 900
			case "1600x1200":
				newWidth = 1600
				newHeight = 1200
			case "1920x1080":
				newWidth = 1920
				newHeight = 1080
			case "1920x1200":
				newWidth = 1920
				newHeight = 1200
			}

			widthRatio := float64(newWidth)/float64(activeMenu.Width)
			heightRatio := float64(newHeight)/float64(activeMenu.Height)

			bestRatio = math.Min(widthRatio, heightRatio)

			activeMenu.SetScale(bestRatio)
			ebiten.SetScreenSize(newWidth, newHeight)

			return nil
		}
	}

	return nil
}

func main() {

	initMenus()
	state = titleScreen

	if err := ebiten.Run(update, 1024, 768, 1, "Card Game"); err != nil {
		panic(err)
	}
}

func calcRatio(width int, height int) (ratio float64) {
	return float64(width) / float64(height)
}
