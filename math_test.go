package math

import (
	"math"
	"testing"
)

func floatsEqual(a, b float64) bool {
	eps := 0.000000001
	res := a - b
	return res < eps && res > -eps
}

func assertEqual(expected, found float64, t *testing.T) {
	if !floatsEqual(expected, found) {
		t.Errorf("Expected %.9f but found %.9f", expected, found)
	}
}

func TestPow(t *testing.T) {
	assertEqual(math.Pow(3.4567, 2.356), Pow(3.4567, 2.356), t)
}

func TestSqrt(t *testing.T) {
	expected := math.Sqrt(126)
	found, _ := Sqrt(126)
	assertEqual(expected, found, t)

	_, err := Sqrt(-1)
	if err == nil {
		t.Error("Calling for -1 did not return an error")
	}
}

func TestLog(t *testing.T) {
	expected := math.Log(126)
	found, _ := Log(126)
	assertEqual(expected, found, t)

	_, err := Log(-1)
	if err == nil {
		t.Error("Calling for -1 did not return an error")
	}
}

func TestLog10(t *testing.T) {
	expected := math.Log10(126)
	found, _ := Log10(126)
	assertEqual(expected, found, t)

	_, err := Log10(-1)
	if err == nil {
		t.Error("Calling for -1 did not return an error")
	}
}

func TestLog2(t *testing.T) {
	expected := math.Log2(126)
	found, _ := Log2(126)
	assertEqual(expected, found, t)

	_, err := Log2(-1)
	if err == nil {
		t.Error("Calling for -1 did not return an error")
	}
}

func TestIlogb(t *testing.T) {
	expected := math.Ilogb(126)
	found, _ := Ilogb(126)

	if expected != found {
		t.Errorf("Expected %d but found %s", expected, found)
	}

	val, _ := Ilogb(0)

	if val != FpIlogb0 {
		t.Error("Expected FpIlogb0 but got something different ")
	}

	val, _ = Ilogb(math.NaN())

	if val != FpIlogbNaN {
		t.Error("Expected FpIlogbNaN but got something different ")
	}
}

func TestCbrt(t *testing.T) {
	assertEqual(math.Cbrt(42), Cbrt(42), t)
}

func TestCos(t *testing.T) {
	assertEqual(math.Cos(42), Cos(42), t)
}

func TestSin(t *testing.T) {
	assertEqual(math.Sin(42), Sin(42), t)
}

func TestTan(t *testing.T) {
	assertEqual(math.Tan(42), Tan(42), t)
}

func TestExp(t *testing.T) {
	assertEqual(math.Exp(5), Exp(5), t)
}

func TestHypot(t *testing.T) {
	expected := math.Hypot(5.3, 2.6)
	found, _ := Hypot(5.3, 2.6)
	assertEqual(expected, found, t)

	_, err := Hypot(math.MaxFloat64, math.MaxFloat64)
	if err == nil {
		t.Error("Calling for hude values did not return an error")
	}
}

func TestCeil(t *testing.T) {
	assertEqual(math.Ceil(5.3), Ceil(5.3), t)
	assertEqual(math.Ceil(5.6), Ceil(5.6), t)
	assertEqual(math.Ceil(-5.3), Ceil(-5.3), t)
	assertEqual(math.Ceil(-5.6), Ceil(-5.6), t)
}

func TestFloor(t *testing.T) {
	assertEqual(math.Floor(5.3), Floor(5.3), t)
	assertEqual(math.Floor(5.6), Floor(5.6), t)
	assertEqual(math.Floor(-5.3), Floor(-5.3), t)
	assertEqual(math.Floor(-5.6), Floor(-5.6), t)
}

func TestRound(t *testing.T) {
	assertEqual(5, Round(5.3), t)
	assertEqual(6, Round(5.6), t)
	assertEqual(-5, Round(-5.3), t)
	assertEqual(-6, Round(-5.6), t)
}

func TestFabs(t *testing.T) {
	assertEqual(math.Abs(-5.3), Fabs(-5.3), t)
	assertEqual(math.Abs(5.3), Fabs(5.3), t)
	assertEqual(math.Abs(0), Fabs(0), t)
}
