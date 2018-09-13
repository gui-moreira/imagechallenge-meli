package main

import (
	"image"
	"image/color"
	"image/draw"
)

type square struct {
	img image.Image
}

func (s *square) getUpperLeftColor() color.Color {
	return s.img.At(0, 0)
}

func (s *square) getUpperRightColor() color.Color {
	return s.img.At(47, 0)
}

func (s *square) getBottomLeftColor() color.Color {
	return s.img.At(0, 47)
}

func (s *square) getBottomRightColor() color.Color {
	return s.img.At(47, 47)
}

func getSquares(img image.Image) []*square {
	cursorX := 0
	cursorY := 0

	squares := []*square{}

	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			m := image.NewRGBA(image.Rect(0, 0, 49, 49))
			draw.Draw(m, image.Rect(0, 0, 49, 49), img, image.Point{cursorX + 1, cursorY + 1}, draw.Src)

			squares = append(squares, &square{img: m})

			cursorX += 50
		}
		cursorX = 0
		cursorY += 50
	}

	return squares
}

func isBlackOrWhite(c color.Color) bool {
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	return c == black || c == white
}
