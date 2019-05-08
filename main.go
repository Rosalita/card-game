package main

import (
	"image/color"
	"os"

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
)

func alignMenu(menu *lm.ListMenu, screenWidth float64, screenHeight float64) {

	menu.Tx = lerp(0.0, float64(screenWidth), 0.5) - float64(menu.Width/2)
	menu.Ty = lerp(0.0, float64(screenHeight), 0.7) - float64(menu.Height/2)

}

func update(screen *ebiten.Image) error {

	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})

	if state == titleScreen {

		activeMenu = mainMenu

		w, h := screen.Size()
		alignMenu(&activeMenu, float64(w), float64(h))

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

		w, h := screen.Size()
		alignMenu(&activeMenu, float64(w), float64(h))

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
		w, h := screen.Size()
		alignMenu(&activeMenu, float64(w), float64(h))

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

			switch activeMenu.GetSelectedItem() {
			case "640x480":
				ebiten.SetScreenSize(640, 480)
			case "800x600":
				ebiten.SetScreenSize(800, 600)
			case "1024x768":
				ebiten.SetScreenSize(1024, 768)
			case "1280x720":
				ebiten.SetScreenSize(1280, 720)
			case "1336x768":
				ebiten.SetScreenSize(1336, 768)
			case "1440x1080":
				ebiten.SetScreenSize(1440, 1080)
			case "1600x900":
				ebiten.SetScreenSize(1024, 768)
			case "1600x1200":
				ebiten.SetScreenSize(1600, 1200)
			case "1920x1080":
				ebiten.SetScreenSize(1920, 1080)
			case "1920x1200":
				ebiten.SetScreenSize(1920, 1200)
			}

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
