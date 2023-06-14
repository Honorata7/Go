package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func collatz(n int) int {

	result := []int{n}

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
			result = append(result, n)
		} else {
			n = 3*n + 1
			result = append(result, n)
		}
	}
	//return result, len(result)
	return len(result)
}

func longestCollatz(start, end int) (int, int) {

	var max int
	var maxIndex int

	for i := start; i < end; i++ {
		length := collatz(i)
		if length > max {
			max = length
			maxIndex = i
		}
	}
	return maxIndex, max
}

// Calculate the observed frequencies of leading digits in the data
func calculateLeadingDigitFrequencies(data []float64) map[int]float64 {
	frequencies := make(map[int]float64)

	for _, value := range data {
		leadingDigit := getLeadingDigit(value)
		frequencies[leadingDigit]++
	}

	total := float64(len(data))
	for digit := 1; digit <= 9; digit++ {
		frequencies[digit] /= total
	}

	return frequencies
}

// Get the leading digit of a number
func getLeadingDigit(value float64) int {
	str := strconv.FormatFloat(value, 'f', -1, 64)
	str = strings.Replace(str, ".", "", 1)
	leadingDigit, _ := strconv.Atoi(string(str[0]))
	return leadingDigit
}

// Calculate the expected frequencies of leading digits based on Benford's Law
func calculateBenfordFrequencies() map[int]float64 {
	frequencies := make(map[int]float64)

	for digit := 1; digit <= 9; digit++ {
		frequencies[digit] = math.Log10(1 + 1/float64(digit))
	}

	return frequencies
}

//Drawing plot using gnuplot

func drawPlot(xData, yData []float64, title string) {
	// temporary file for the gnuplot script
	scriptFile, err := os.CreateTemp("", "gnuplot-script")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(scriptFile.Name())
	defer scriptFile.Close()

	script := fmt.Sprintf("set term png\n set output %q\n plot '-' with linespoints\n", title)
	// gnuplot commands
	scriptFile.WriteString(script)

	// writing the data to the temporary file
	for i := 0; i < len(xData); i++ {
		scriptFile.WriteString(strconv.FormatFloat(xData[i], 'f', -1, 64) + " " + strconv.FormatFloat(yData[i], 'f', -1, 64) + "\n")
	}
	scriptFile.WriteString("e\n")

	// executing the gnuplot script
	cmd := exec.Command("gnuplot", scriptFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Plot saved as plot.png")
}

func main() {

	fmt.Println(collatz(47))
	fmt.Println(longestCollatz(10, 100))
	//longest between 10 and 100 => 97, 119
	//longest between 100 and 1000 => 871, 179
	//longest between 1000 and 2000 => 1161, 182
	//longest between 2000 and 3000 => 2919, 217
	//longest between 3000 and 4000 => 3711, 238
	//longest betweeen 4000 and 5000 => 4379, 215
	//longest between 5000 and 6000 => 5561, 236
	//fmt.Println(longestCollatz(5000, 6000))

	file, err := os.OpenFile("collatz.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	var xData []float64
	var yData []float64

	for i := 1; i <= 10000; i++ {
		sequence := collatz(i)
		fmt.Fprintf(file, "%d: %v\n", i, sequence)
		xData = append(xData, float64(i))
		yData = append(yData, float64(sequence))
	}

	drawPlot(xData, yData, "plot100.png")

	//calculating mean
	var sum int
	for _, v := range yData {
		sum += int(v)
	}
	mean := float64(sum) / float64(len(yData))
	fmt.Println("Mean: ", mean)

	//calculating median
	median := yData[len(yData)/2]
	fmt.Println("Median: ", median)

	observedFreq := calculateLeadingDigitFrequencies(yData)

	// Calculate the expected frequencies based on Benford's Law
	expectedFreq := calculateBenfordFrequencies()

	for i := 1; i <= 9; i++ {
		fmt.Printf("Digit %d: Observed: %.2f%%, Expected: %.2f%%\n", i, observedFreq[i]*100, expectedFreq[i]*100)
	}

}
