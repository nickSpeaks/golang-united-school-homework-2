package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type customSideNum int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const (
	SidesCircle   customSideNum = 0
	SidesTriangle customSideNum = 3
	SidesSquare   customSideNum = 4
	Pi                          = 3.141592653589793
)

func CalcSquare(sideLen float64, sidesNum customSideNum) float64 {
	switch sidesNum {
	case SidesCircle:
		return Pi * sideLen * sideLen
	case SidesTriangle:
		return math.Sqrt(3) / 4 * sideLen * sideLen
	case SidesSquare:
		return sideLen * sideLen
	default:
		return 0
	}

}
