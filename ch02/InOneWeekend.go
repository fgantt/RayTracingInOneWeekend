package main

import (
	"fmt"
	"log"
)

func main() {

	// Image

	imageWidth, imageHeight := 256, 256

	// Render

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := 0; j < imageHeight; j++ {
		log.Printf("\rScanlines remaining: %d", imageHeight-j)
		for i := 0; i < imageWidth; i++ {
			r := float32(i) / float32(imageWidth-1)
			g := float32(j) / float32(imageHeight-1)
			var b int

			ir := int(255.999 * float32(r))
			ig := int(255.999 * float32(g))
			ib := int(255.999 * float32(b))

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}

	log.Println("\rDone.")
}
