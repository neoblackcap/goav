package main

import (
	"testing"

	"os"
)

func TestTranscodeAudio(t *testing.T) {

	inputFile := "bbb480.mp4"
	outputFile := "out.mp4"
	os.Remove(outputFile)
	if Timelapse(inputFile, outputFile) != 0 {
		t.Errorf("failed to transcode")
	}
	f, err := os.Stat(outputFile)
	if err != nil {
		t.Errorf("failed to read %s, err=%v", outputFile, err)
	}
	if f.Size() < 14_000_000 {
		t.Errorf("%s seems too short", outputFile)
	}
	os.Remove(outputFile)
}
