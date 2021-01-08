package main

import (
    "fmt"
    "time"
	"runtime"
	"math"
)

type Sphere struct {
	number int
	volume float64
	surface float64
	circularArea float64
	parts int
}

func volume(radius float64, volumeMath chan float64){
	powerOfThree := radius * radius * radius
	volume := float64((4/3) * math.Pi * powerOfThree)
	volumeMath <- volume
}

func surface(radius float64, surfaceMath chan float64){
	surface := 4.0 * math.Pi * math.Sqrt(radius)
	surfaceMath <- surface
}

func circularArea(radius float64, areaMath chan float64){
	area := math.Pi * math.Sqrt(radius)
	areaMath <- area
}

func allResults(ready chan string, volumeMath chan float64, areaMath chan float64, 
	surfaceMath chan float64, Sphere1 Sphere){
	for{
		select{
		case volume:= <-volumeMath:{
				Sphere1.volume = volume
				Sphere1.parts++
				fmt.Println("Volume is ", volume)
				
			}
		case area:= <-areaMath:{
			Sphere1.circularArea = area
			Sphere1.parts++
			fmt.Println("Area is ", area)
			
			}
		case surface:= <-surfaceMath:{
			Sphere1.surface = surface
			Sphere1.parts++
			fmt.Println("Surface is ", surface)
			
			}
		}
		if Sphere1.parts == 3{
		ready <- "All parts calculated"}
	}
}

func main() {
  runtime.GOMAXPROCS(4)
  start := time.Now()

  Sphere1 := new(Sphere)
  Sphere1.number = 0

  ready := make(chan string)
  areaMath := make(chan float64)
  volumeMath := make(chan float64)
  surfaceMath := make(chan float64)

  go allResults(ready, volumeMath, areaMath, surfaceMath, *Sphere1)
  go circularArea(2.5, areaMath)  
  go volume(2.5, volumeMath)
  go surface(2.5, surfaceMath)

  end:= <- ready

  fmt.Println(end)
  elapsedTime := time.Since(start)
  fmt.Println("Total Time For Execution: " + elapsedTime.String())
  time.Sleep(time.Second)
}