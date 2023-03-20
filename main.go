package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/heesu0/bitrate-viewer-go/pkg/ffprobe"
)

func main() {
	inputFile := flag.String("input", "", "input file path.")
	outputFile := flag.String("output", "", "output file path.")

	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Input file is not entered.")
		return
	}
	if *outputFile == "" {
		fmt.Println("Output file is not entered.")
		return
	}

	progress := make(chan struct{})
	done := make(chan struct{})
	go progressIndicator(progress, done)

	probeData, err := ffprobe.GetProbeData(*inputFile)
	if err != nil {
		panic(err)
	}

	bitrates, iFrames, err := getPoints(probeData)
	if err != nil {
		panic(err)
	}

	err = plotGraph(bitrates, iFrames, *outputFile)
	if err != nil {
		panic(err)
	}

	close(progress)
	<-done

	fmt.Printf("Graph file saved as %v", *outputFile)
	fmt.Println()
}

func progressIndicator(progess, done chan struct{}) {
	progressChars := []string{"|", "/", "-", "\\"}
	var i int
	for {
		select {
		case <-progess:
			fmt.Printf("\rTask complete!\n")
			close(done)
			return
		default:
			fmt.Printf("\rProcessing... %s", progressChars[i])
			time.Sleep(100 * time.Millisecond)
			i = (i + 1) % len(progressChars)
		}
	}
}
