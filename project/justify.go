package project

import (
	asc "ascii-art/utils"
	"fmt"
	"log"
	"os"
	"regexp"
)

var (
	align, _ = regexp.Compile("--align=")
	fileName = "standard.txt"
	option   = "reset"
)

// JustifyProject function is
func JustifyProject() {
	args := os.Args[1:]

	if len(args) < 1 || len(args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . something standard --align=right")
		os.Exit(0)
	}

	if len(args) == 1 {
		CallingTheProgram(os.Args[1], fileName, option)
	} else if len(args) == 2 {
		if !align.MatchString(os.Args[2]) {
			CallingTheProgram(os.Args[1], os.Args[2]+".txt", option)
		} else {
			CallingTheProgram(os.Args[1], fileName, os.Args[2][8:])
		}
	} else if align.MatchString(os.Args[3]) {
		CallingTheProgram(os.Args[1], os.Args[2]+".txt", os.Args[3][8:])
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . something standard --align=right")
		os.Exit(0)
	}
}

func CallingTheProgram(s, fileName, option string) {
	m, s := asc.SetAsciiArt(s, fileName)
	if option == "left" || option == "right" || option == "center" || option == "justify" || option == "reset" {
		fmt.Print(asc.ToString(asc.GetAsciiArt(m, s, option)))
	} else {
		log.Println("Undefined position. Try: left, right, center, justify.")
	}
}
