package main

import (
	"fmt"
	"image"
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

	for i := 0; i < 20; i++ { // line
		for j := 0; j < 20; j++ { //column
			var foundSquare *square
			if i == 0 && j == 0 {
				foundSquare = findFirst()
			} else {
				foundSquare = findNext(i, j)
			}

			formedImage[i][j] = foundSquare
		}
	}

	writeToFile(formedImage[1][2].img)
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

func findNext(i, j int) *square {
	// fill first line
	if i == 0 {
		previousLeft := formedImage[i][j-1]
		for _, s := range allSquares {
			if isBlackOrWhite(s.getUpperLeftColor()) &&
				isBlackOrWhite(s.getUpperRightColor()) &&
				s.getBottomLeftColor() == previousLeft.getBottomRightColor() {
				return s
			}
		}
	}

	// fill first column
	if j == 0 {
		previousTop := formedImage[i-1][j]
		for _, s := range allSquares {
			if isBlackOrWhite(s.getUpperLeftColor()) &&
				isBlackOrWhite(s.getBottomLeftColor()) &&
				s.getUpperRightColor() == previousTop.getBottomRightColor() {
				return s
			}
		}
	}

	// fill rest
	if i > 0 && j > 0 {
		previousLeft := formedImage[i][j-1]
		previousTop := formedImage[i-1][j]
		fmt.Println(i, j)
		for _, s := range allSquares {
			if s.getUpperLeftColor() == previousTop.getBottomLeftColor() &&
				s.getUpperRightColor() == previousTop.getBottomRightColor() &&
				s.getBottomLeftColor() == previousLeft.getBottomRightColor() {
				fmt.Println("found next")
				return s
			}
		}

		fmt.Println("not found")
	}

	return &square{}
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
