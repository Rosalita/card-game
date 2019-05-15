package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil" // required for isKeyJustPressed
)

func updateScreensize(screen *ebiten.Image) error {
	activeMenu = screensizeMenu
	activeMenu.SetScale(bestRatio)

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
		case "1366x768":
			newWidth = 1366
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

		widthRatio := float64(newWidth) / float64(activeMenu.Width)
		heightRatio := float64(newHeight) / float64(activeMenu.Height)

		bestRatio = math.Min(widthRatio, heightRatio)

		activeMenu.SetScale(bestRatio)

		ebiten.SetScreenSize(newWidth, newHeight)

		if ebiten.IsDrawingSkipped() {
			return nil
		}

		activeMenu.Draw(screen)

		return nil
	}

	return nil
}

func calcRatio(width int, height int) (ratio float64) {
	return float64(width) / float64(height)
}
