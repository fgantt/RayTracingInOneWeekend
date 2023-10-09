package main

import (
	"fmt"
	"log"

	"vec3/vec3"
)

func main() {
	//t1 := vec3.New(3, 4, 5)
	//fmt.Println("t1 = ", t1)
	//
	//t2 := vec3.New(6, 7, 8)
	//fmt.Println("t2 = ", t2)
	//
	//t1.Inv()
	//fmt.Println("t1 = ", t1)
	//
	//t2.Add(t1)
	//fmt.Println("t2 = ", t2)
	//
	//t2.Mul(2)
	//fmt.Println("t2 = ", t2)
	//
	//t2.Div(2)
	//fmt.Println("t2 = ", t2)
	//
	//fmt.Println("t2.LengthSquared = ", t2.LengthSquared())
	//fmt.Println("t2.Length = ", t2.Length())
	//
	//c1 := vec3.NewColor(3, 4, 5)
	//fmt.Println(c1.Write())

	// Image

	imageWidth, imageHeight := 256, 256

	// Render

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := 0; j < imageHeight; j++ {
		log.Printf("\rScanlines remaining: %d", imageHeight-j)
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / float64(imageWidth-1)
			g := float64(j) / float64(imageHeight-1)

			color := vec3.NewColor(r, g, 0)
			fmt.Print(color.Write())
		}
	}
	log.Println("\rDone.")
}
