package sample

import (
	"github.com/castorhrio/grpcpcbook/grpcbook"
	"github.com/golang/protobuf/ptypes"
)

func NewKeyboard() *grpcbook.Keyboard {
	keyboard := &grpcbook.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

func NewCPU() *grpcbook.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)
	cpu := &grpcbook.CPU{
		Brand:         brand,
		Name:          name,
		NumberCors:    uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

func NewGPU() *grpcbook.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)

	memory := &grpcbook.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  grpcbook.Memory_GIGABYTE,
	}

	gpu := &grpcbook.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return gpu
}

func NewRAM() *grpcbook.Memory {
	ram := &grpcbook.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  grpcbook.Memory_GIGABYTE,
	}

	return ram
}

func NewSSD() *grpcbook.Storage {
	ssd := &grpcbook.Storage{
		Driver: grpcbook.Storage_SSD,
		Memory: &grpcbook.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  grpcbook.Memory_GIGABYTE,
		},
	}

	return ssd
}

func NewHDD() *grpcbook.Storage {
	ssd := &grpcbook.Storage{
		Driver: grpcbook.Storage_HDD,
		Memory: &grpcbook.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  grpcbook.Memory_TERAVYTE,
		},
	}

	return ssd
}

func NewScreen() *grpcbook.Screen {
	screen := &grpcbook.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Mutitouch:  randomBool(),
	}

	return screen
}

func NewLaptop() *grpcbook.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	taptop := &grpcbook.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Gpus:     []*grpcbook.GPU{NewGPU()},
		Memory:   NewRAM(),
		Storages: []*grpcbook.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &grpcbook.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},

		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdateAt:    ptypes.TimestampNow(),
	}

	return taptop
}
