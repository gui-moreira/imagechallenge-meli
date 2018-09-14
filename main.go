package main

import (
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
		panic(err.Error())
	}
	defer infile.Close()

	src, err := png.Decode(infile)
	if err != nil {
		panic(err.Error())
	}

	allSquares = getSquares(src)

	formedImage[0][0] = findFirst()
	fillSquares(0, 1)
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
	if i == -1 || j == -1 || i == 20 || j == 20 {
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

	if i == 19 && j == 19 {
		//fmt.Println(allSquares[9].used)
		formedImage[19][19] = allSquares[296]
		writeToFile(generateImage())
		return true
	}

	for _, s := range allSquares {
		if !s.used && fits(i, j, s) {
			//fmt.Println("new fit", i, j)
			formedImage[i][j] = s
			s.used = true

			if fillSquares(i+1, j-1) {
				return true
			}
		}
		formedImage[i][j] = nil
		s.used = false
	}

	return false
}

func fits(i, j int, s *square) bool {

	if i == 12 && j == 19 && s == allSquares[9] {
		return true
	}
	if i == 19 && j == 12 && s == allSquares[264] {
		return true
	}

	// fill first line
	if i == 0 {
		previousLeft := formedImage[i][j-1]
		if j == 19 {
			if isBlackOrWhite(s.getUpperLeftColor()) &&
				isBlackOrWhite(s.getUpperRightColor()) &&
				!isBlackOrWhite(s.getBottomLeftColor()) &&
				isBlackOrWhite(s.getBottomRightColor()) &&
				s.getBottomLeftColor() == previousLeft.getBottomRightColor() {
				return true
			}
		} else {
			if isBlackOrWhite(s.getUpperLeftColor()) &&
				isBlackOrWhite(s.getUpperRightColor()) &&
				!isBlackOrWhite(s.getBottomLeftColor()) &&
				!isBlackOrWhite(s.getBottomRightColor()) &&
				s.getBottomLeftColor() == previousLeft.getBottomRightColor() {
				return true
			}
		}

	}

	// fill first column
	if j == 0 {
		previousTop := formedImage[i-1][j]
		if i == 19 {
			if isBlackOrWhite(s.getUpperLeftColor()) &&
				isBlackOrWhite(s.getBottomLeftColor()) &&
				!isBlackOrWhite(s.getUpperRightColor()) &&
				isBlackOrWhite(s.getBottomRightColor()) &&
				s.getUpperRightColor() == previousTop.getBottomRightColor() {
				return true
			}
		} else {
			if isBlackOrWhite(s.getUpperLeftColor()) &&
				isBlackOrWhite(s.getBottomLeftColor()) &&
				!isBlackOrWhite(s.getUpperRightColor()) &&
				!isBlackOrWhite(s.getBottomRightColor()) &&
				s.getUpperRightColor() == previousTop.getBottomRightColor() {
				return true
			}
		}
	}

	// fill rest
	if i > 0 && j > 0 {
		previousLeft := formedImage[i][j-1]
		previousTop := formedImage[i-1][j]

		if s.getUpperLeftColor() == previousTop.getBottomLeftColor() &&
			s.getUpperRightColor() == previousTop.getBottomRightColor() &&
			s.getBottomLeftColor() == previousLeft.getBottomRightColor() {
			return true
		}
	}

	return false
}

func generateImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 980, 980))

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
	outputFile, err := os.Create("solution.png")
	if err != nil {
		panic(err)
	}

	png.Encode(outputFile, img)

	outputFile.Close()
}
