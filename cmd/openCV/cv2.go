package main

import (
	"gocv.io/x/gocv"
)

/**
set CGO_CXXFLAGS="--std=c++11"
set CGO_CPPFLAGS=-ID:\tools\opencv\build\include
set CGO_LDFLAGS=-LD:\tools\Mingw\lib -lopencv_core410 -lopencv_face410 -lopencv_videoio410 -lopencv_imgproc410 -lopencv_highgui410 -lopencv_imgcodecs410 -lopencv_objdetect410 -lopencv_features2d410 -lopencv_video410 -lopencv_dnn410 -lopencv_xfeatures2d410 -lopencv_plot410 -lopencv_tracking410 -lopencv_img_hash410
*/
import "C"

func main() {
	webcam, _ := gocv.OpenVideoCapture(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
