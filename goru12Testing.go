package main

import (
		"testing"
		)
func Average(t []float64) float64 {
	var sum float64
	for _, value := range t {
		sum += value
	}
	return sum/float64(len(t))
}

func TestAverage(t *testing.T) {
    var v float64
    v = Average([]float64{1,2})
    if v != 1.5 {
        t.Error("Expected 1.5, got ", v)
    }
}
func main() {

}

//go test
