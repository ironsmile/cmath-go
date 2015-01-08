package math

/*
#cgo LDFLAGS: -lm
#include <math.h>
*/
import "C"

/* Constants */

// Returned by ilogb
var (
	FpIlogb0   int = int(C.FP_ILOGB0)
	FpIlogbNaN int = int(C.FP_ILOGBNAN)
)

/* Power functions */

func Pow(a, b float64) float64 {
	return float64(C.pow(C.double(a), C.double(b)))
}

func Sqrt(a float64) (float64, error) {
	cDouble, err := C.sqrt(C.double(a))
	return float64(cDouble), err
}

func Cbrt(a float64) float64 {
	return float64(C.cbrt(C.double(a)))
}

func Hypot(a, b float64) (float64, error) {
	cDouble, err := C.hypot(C.double(a), C.double(b))
	return float64(cDouble), err
}

/* Some trigonometric functions */

func Cos(a float64) float64 {
	return float64(C.cos(C.double(a)))
}

func Sin(a float64) float64 {
	return float64(C.sin(C.double(a)))
}

func Tan(a float64) float64 {
	return float64(C.tan(C.double(a)))
}

/* Some exponential and logarithmic functions */

func Exp(a float64) float64 {
	return float64(C.exp(C.double(a)))
}

func Log(a float64) (float64, error) {
	cDouble, err := C.log(C.double(a))
	return float64(cDouble), err
}

func Log10(a float64) (float64, error) {
	cDouble, err := C.log10(C.double(a))
	return float64(cDouble), err
}

func Log2(a float64) (float64, error) {
	cDouble, err := C.log2(C.double(a))
	return float64(cDouble), err
}

func Ilogb(a float64) (int, error) {
	cInt, err := C.ilogb(C.double(a))
	return int(cInt), err
}

/* Some rounding and remainder functions */

func Ceil(a float64) float64 {
	return float64(C.ceil(C.double(a)))
}

func Floor(a float64) float64 {
	return float64(C.floor(C.double(a)))
}

// Probably the only useful wrapper since golang math does not have Round
func Round(a float64) float64 {
	return float64(C.round(C.double(a)))
}

/* Other functions */

func Fabs(a float64) float64 {
	return float64(C.fabs(C.double(a)))
}
