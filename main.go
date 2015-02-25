package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"path"
)

type convertCommand struct {
	dimensions  string
	transparent bool
	output      string
}

var convertStrings = []convertCommand{
	{
		dimensions: "144x144",
		output:     "apple-touch-icon-144x144.png",
	},
	{
		dimensions: "76x76",
		output:     "apple-touch-icon-76x76.png",
	},
	{
		dimensions: "60x60",
		output:     "apple-touch-icon.png",
	},
	{
		dimensions:  "144x144",
		output:      "favicon.ico",
		transparent: true,
	},
}

// convert ~/Desktop/fav-icon.png -resize 76x76 -background white -alpha remove apple-touch-icon-76x76.png
func convert(original, outputDir string, cs convertCommand) error {
	var cmd *exec.Cmd
	if cs.transparent {
		cmd = exec.Command("convert", original, "-resize", cs.dimensions, path.Join(outputDir, cs.output))
	} else {
		cmd = exec.Command("convert", original, "-resize", cs.dimensions, "-background", "white", "-alpha", "remove", path.Join(outputDir, cs.output))
	}
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	path, err := exec.LookPath("convert")
	if err != nil {
		log.Fatal("ImageMagick is required for this utility to work!")
	}

	var original, output string

	flag.StringVar(&original, "s", "", "Source for favicon generation")
	flag.StringVar(&output, "o", "", "Output for favicon generation")

	flag.Parse()

	for _, cs := range convertStrings {
		err := convert(original, output, cs)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("ImageMagick found at %s\n", path)
}
