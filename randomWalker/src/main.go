package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func rwalk(x float32, y float32) float32 {
	return x + y
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func sim() {
	rand.Seed(time.Now().Unix())
	fmt.Printf("Float32: %f\n", rand.Float32())

	var walk float32
	var value float32
	var result []string

	value = 0.0

	a := makeRange(1, 1000)

	for i := range a {
		if rand.Float32() < 0.5 {
			walk = -1
		} else {
			walk = 1
		}

		value += rwalk(walk, rand.Float32()*walk)
		fmt.Println(i)
		fmt.Println(value)
		str := fmt.Sprint(value)
		result = append(result, str)

	}
	wr := csv.NewWriter(os.Stdout)
	wr.Write(result)
	wr.Flush()

}

func main() {
	fmt.Println("Hello, Walker")
	sim()
}
