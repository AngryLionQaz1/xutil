package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tshowimage [imgfile]")
		return
	}

	filename := os.Args[1]
	window := gocv.NewWindow("Hello")
	img := gocv.IMRead(filename, gocv.IMReadColor)
	grayImage := gocv.NewMat()
	defer grayImage.Close()

	gocv.CvtColor(img, &grayImage, gocv.ColorBGRToGray)
	destImage := gocv.NewMat()
	gocv.Threshold(grayImage, &destImage, 100, 255, gocv.ThresholdBinaryInv)
	resultImage := gocv.NewMatWithSize(500, 400, gocv.MatTypeCV8U)

	gocv.Resize(destImage, &resultImage, image.Pt(resultImage.Rows(), resultImage.Cols()), 0, 0, gocv.InterpolationCubic)
	gocv.Dilate(resultImage, &resultImage, gocv.NewMat())
	gocv.GaussianBlur(resultImage, &resultImage, image.Pt(5, 5), 0, 0, gocv.BorderWrap)
	results := gocv.FindContours(resultImage, gocv.RetrievalTree, gocv.ChainApproxSimple)
	imageForShowing := gocv.NewMatWithSize(resultImage.Rows(), resultImage.Cols(), gocv.MatChannels4)
	for index, element := range results {
		fmt.Println(index)
		gocv.DrawContours(&imageForShowing, results, index, color.RGBA{R: 0, G: 0, B: 255, A: 255}, 1)
		gocv.Rectangle(&imageForShowing,
			gocv.BoundingRect(element),
			color.RGBA{R: 0, G: 255, B: 0, A: 100}, 1)
	}

	if img.Empty() {
		fmt.Println("Error reading image from: %v", filename)
		return
	}

	for {
		window.IMShow(imageForShowing)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
