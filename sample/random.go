package sample

import (
	"github.com/google/uuid"
	"github.com/menyasosali/grpc-pcbook/pb"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomStringFromSet(set ...string) string {
	n := len(set)
	if n == 0 {
		return ""
	}

	return set[rand.Intn(n)]
}

func randomCPUName(brand string) string {
	if strings.ToLower(brand) == "intel" {
		return randomStringFromSet(
			"Intel Core i9-11900H",
			"Intel Core i7-11800H",
			"Intel Core i7-1165G7",
			"Intel Core i5-1135G7",
		)
	}
	return randomStringFromSet(
		"AMD Ryzen 9 5900HS",
		"AMD Ryzen 9 5900HX",
		"AMD Ryzen 7 5800H",
		"AMD Ryzen 7 5800U",
	)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomGPUBrand() string {
	return randomStringFromSet("Nvidia", "AMD", "Intel")
}

func randomGPUName(brand string) string {
	if strings.ToLower(brand) == "nvidia" {
		return randomStringFromSet(
			"NVIDIA GeForce RTX 3080 Ti 175W",
			"NVIDIA GeForce RTX 3080 Ti 150W",
			"NVIDIA GeForce RTX 3080 165W",
			"NVIDIA GeForce RTX 3080 150W",
		)
	} else if strings.ToLower(brand) == "amd" {
		return randomStringFromSet(
			"AMD Radeon RX 6850M XT",
			"AMD Radeon RX 6700M",
			"AMD Radeon RX 6700S",
			"AMD Radeon RX 6600M",
		)
	}
	return randomStringFromSet(
		"Intel Arc A730M",
		"Intel Arc A550M",
		"Intel Arc A550M",
		"Intel Iris Xe Max",
	)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
	return resolution
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Huawei", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch strings.ToLower(brand) {
	case "apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "huawei":
		return randomStringFromSet("Huawei MateBook D14", "Huawei MateBook 14 2021")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
	}
}
