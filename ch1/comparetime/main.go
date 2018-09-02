// For each function f(n) determine largest n, within given time intervals
// f(n): lg(n), sqrt(n), n, nlg(n), n**2, n**3, 2**n, n!
// Skipped nlg(n) and further
package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"text/tabwriter"
)

const (
	Sec     float64 = 1
	Min     float64 = Sec * 60
	Hour    float64 = Min * 60
	Day     float64 = Hour * 24
	Month   float64 = Day * 30
	Year    float64 = Month * 12
	Century float64 = Year * 100
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 10, 10, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "f(n)\tsecond\tminute\thour\tday\tmonth\tyear\tcentury\t")

	writeResult(w, "lg(n)", lgn)
	writeResult(w, "sqrt(n)", sqrtn)
	writeResult(w, "n", n)

	w.Flush()
}

func lgn(t float64) float64 {
	return math.Pow(2, t)
}

func sqrtn(t float64) float64 {
	return math.Pow(t, 2)
}

func n(t float64) float64 {
	return t
}

func writeResult(w io.Writer, name string, f func(float64) float64) {
	times := []float64{Sec, Min, Hour, Day, Month, Year, Century}
	res := [7]float64{}
	for i, t := range times {
		res[i] = f(t)
	}
	fmt.Fprintf(w, "%s\t%0.2e\t%0.2e\t%0.2e\t%0.2e\t%0.2e\t%0.2e\t%0.2e\t\n", name, res[0], res[1], res[2], res[3], res[4], res[5], res[6])
}
