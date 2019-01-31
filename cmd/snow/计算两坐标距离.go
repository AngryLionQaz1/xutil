package snow

import (
	"fmt"
	"math"
)

func main() {

	lat1 := 108.398700
	lng1 := 31.173324
	lat2 := 108.40049589117558
	lng2 := 31.170160268820645
	fmt.Println(EarthDistance(lat1, lng1, lat2, lng2))
}

func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6371000.0 // 6378137
	rad := math.Pi / 180.0

	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius

}
