package sample

import (
	"math/rand"
	"time"

	"github.com/castorhrio/grpcpcbook/grpcbook"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() grpcbook.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return grpcbook.Keyboard_QWERTY
	case 2:
		return grpcbook.Keyboard_QWERTZ
	default:
		return grpcbook.Keyboard_AZERTY
	}
}
func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}

	return a[rand.Intn(n)]
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}

	return randomStringFromSet(
		"RyZen 7 PRO 2700U",
		"RyZen 5 PRO 23500U",
		"RyZen 3 PRO 3200GE",
	)
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-TI",
			"GTX 1070",
		)
	}

	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vaga-56",
	)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomScreenPanel() grpcbook.Screen_Panel {
	if rand.Intn(2) == 1 {
		return grpcbook.Screen_IPS
	}

	return grpcbook.Screen_OLED
}

func randomScreenResolution() *grpcbook.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9

	resolution := &grpcbook.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}

	return resolution
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenove")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")

	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
	}
}
