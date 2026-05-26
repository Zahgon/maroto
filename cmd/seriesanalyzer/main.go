package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("docs/assets/text/benchmark.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	stringContent := string(bytes)

	var values []float64
	var sum float64
	lines := strings.Split(stringContent, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
		values = append(values, value)
		sum += value
	}

	values = MergeFloat64(values)

	fmt.Printf("min: %f, max: %f, avg: %f", values[0], values[len(values)-1], sum/float64(len(values)))
}

func MergeFloat64(arr []float64) []float64 { _ = "STUB: not implemented"; return nil }

func mergeArrays(a []float64, b []float64) []float64 { _ = "STUB: not implemented"; return nil }
