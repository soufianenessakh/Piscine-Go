package main

import("fmt")

type Point struct {
    X int
    Y int
}
func setPoint(ptr *Point){
	ptr.X = 42
	ptr.Y = 21
}

func main() {
	points := &Point{}

	setPoint(points)

	fmt.Printf("X = %d, Y = %d\n",points.X, points.Y)
}