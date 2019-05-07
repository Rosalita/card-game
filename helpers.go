package main

import (
	"github.com/hajimehoshi/ebiten"
)

// Linear interpolation helper
func lerp(a float64, b float64, pct float64) (result float64) {
	return a + pct*(b-a)
}

func getCentre(screen *ebiten.Image) (x int, y int) {
	x, y = screen.Size()
	return x / 2, y / 2
}
