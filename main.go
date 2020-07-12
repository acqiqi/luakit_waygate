package main

import (
	"golang.org/x/exp/io/spi"
)

func main() {
	dev, err := spi.Open(&spi.Devfs{
		Dev:      "/dev/spidev0.1",
		Mode:     spi.Mode3,
		MaxSpeed: 500000,
	})
	if err != nil {
		panic(err)
	}
	defer dev.Close()

	if err := dev.Tx([]byte{
		0, 0, 0, 0,
		0xff, 200, 0, 200,
		0xff, 200, 0, 200,
		0xe0, 200, 0, 200,
		0xff, 200, 0, 200,
		0xff, 8, 50, 0,
		0xff, 200, 0, 0,
		0xff, 0, 0, 0,
		0xff, 200, 0, 200,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff,
	}, nil); err != nil {
		panic(err)
	}
}
