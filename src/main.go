package main

import (
	"fmt"
	"os"
)

func main() {
	var text, output, top, bottom string = "", "", " ", " "
	length := len(os.Args[1:])
	for i, arg := range os.Args[1:] {
		text += arg
		if (i != length - 1) {
			text += " "
		}
	}
	lines := (len(text) / 39) + 1
	if (lines == 1) {
		for i := -2; i < len(text); i++ {
			top += "_"
			bottom += "-"
		}
	} else {
		for i := -2; i < 39; i++ {
			top += "_"
			bottom += "-"
		}
	}
	output += top + "\n"
	if(lines == 1) {
		output += "< " + text + " >\n"
	} else {
		for i := 0; i < lines; i++ {
			var lineText string
			if(i == lines - 1) {
				lineText = text[i * 39:]
				for j := len(lineText); j < 39; j++ {
					lineText += " "
				}
			} else {
				lineText = text[i * 39:(i + 1) * 39]
			}
			if(i == 0) {
				output += "/ " + lineText + " \\\n"
			} else if (i == lines - 1) {
				output += "\\ " + lineText + " /\n"
			} else {
				output += "| " + lineText + " |\n"
			}
		}
	}
	output += bottom
	output += "\n        \\   ^__^\n         \\  (oo)\\_______\n            (__)\\       )" +
				"\\/\\\n                ||----w |\n                ||     ||"
	fmt.Println(output)
}
