package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const FREQ = 0.1

func rgb(freq float64, i int) (int, int, int) {
	return int(math.Sin(freq*float64(i)+0)*127 + 128),
		int(math.Sin(freq*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(freq*float64(i)+4*math.Pi/3)*127 + 128)
}

func print(output []rune) {
	for i, c := range output {
		r, g, b := rgb(FREQ, i)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, c)
	}

	fmt.Println()
}

func main() {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if fileInfo.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gorainbow")
		return
	}

	var output []rune
	reader := bufio.NewReader(os.Stdin)
	for {
		c, _, err := reader.ReadRune()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		output = append(output, c)
	}

	print(output)

}
