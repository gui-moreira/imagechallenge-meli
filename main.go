package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

var allSquares []*square
var formedImage [20][20]*square

func main() {
	infile, err := os.Open("challenge.png")
	if err != nil {
		// replace this with real error handling
		panic(err.Error())
	}
	defer infile.Close()

	// Decode will figure out what type of image is in the file on its own.
	// We just have to be sure all the image packages we want are imported.
	src, err := png.Decode(infile)
	if err != nil {
		// replace this with real error handling
		panic(err.Error())
	}

	allSquares = getSquares(src)

	formedImage[0][0] = findFirst()
	fillSquares(0, 1)

	writeToFile(generateImage())
	//fmt.Println(src.At(49, 51))
}

func findFirst() *square {
	for _, s := range allSquares {
		if isBlackOrWhite(s.getUpperLeftColor()) &&
			isBlackOrWhite(s.getUpperRightColor()) &&
			isBlackOrWhite(s.getBottomLeftColor()) {
			return s
		}
	}
	return &square{}
}

func fillSquares(i, j int) bool { // i = line, j = column
	if i == 0 && j == 0 {
		j++
	}

	// find next empty
	if i == -1 || j == -1 {
		for ix, line := range formedImage {
			found := false
			for jx, s := range line {
				if s == nil {
					i = ix
					j = jx
					found = true
					break
				}
			}

			if found {
				break
			}
		}
	}

	fmt.Println(i, j)

	for ix, s := range allSquares {
		if fits(i, j, s) {
			formedImage[i][j] = s
			allSquares = append(allSquares[:ix], allSquares[ix+1:]...)

			if fillSquares(i+1, j-1) {
				return true
			}
		}
	}

	fmt.Println("did not find fit")

	return false
}

func fits(i, j int, s *square) bool {

	// fill first line
	if i == 0 {
		previousLeft := formedImage[i][j-1]
		if previousLeft == nil {
			//fmt.Println("left null")
		}

		if isBlackOrWhite(s.getUpperLeftColor()) &&
			isBlackOrWhite(s.getUpperRightColor()) &&
			!isBlackOrWhite(s.getBottomLeftColor()) &&
			!isBlackOrWhite(s.getBottomRightColor()) &&
			s.getBottomLeftColor() == previousLeft.getBottomRightColor() {
			return true
		}
	}

	// fill first column
	if j == 0 {
		previousTop := formedImage[i-1][j]
		if previousTop == nil {
			//fmt.Println("top null")
		}
		if isBlackOrWhite(s.getUpperLeftColor()) &&
			isBlackOrWhite(s.getBottomLeftColor()) &&
			!isBlackOrWhite(s.getUpperRightColor()) &&
			!isBlackOrWhite(s.getBottomRightColor()) &&
			s.getUpperRightColor() == previousTop.getBottomRightColor() {
			return true
		}
	}

	// fill rest
	if i > 0 && j > 0 {
		previousLeft := formedImage[i][j-1]
		previousTop := formedImage[i-1][j]

		if previousLeft == nil {
			//fmt.Println("left null")
		}
		if previousTop == nil {
			//fmt.Println("top null")
		}

		if s.getUpperLeftColor() == previousTop.getBottomLeftColor() &&
			s.getUpperRightColor() == previousTop.getBottomRightColor() &&
			s.getBottomLeftColor() == previousLeft.getBottomRightColor() {
			return true
		}
	}

	return false
}

func generateImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 960, 960))

	for i := 0; i < 20; i++ { // line
		for j := 0; j < 20; j++ { //column
			s := formedImage[i][j]
			if s == nil {
				s = &square{img: image.NewRGBA(image.Rect(0, 0, 49, 49))}
			}

			draw.Draw(img, image.Rect(j*49, i*49, (j*49)+49, (i*49)+49), s.img, image.Point{0, 0}, draw.Src)
		}
	}

	return img
}

func writeToFile(img image.Image) {
	outputFile, err := os.Create("test.png")
	if err != nil {
		panic(err)
	}

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, img)

	// Don't forget to close files
	outputFile.Close()
}
