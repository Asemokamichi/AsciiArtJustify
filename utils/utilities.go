package asc

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// SetAsciiArt function is
func SetAsciiArt(s, banner string) (map[rune][]string, string) {
	s = strings.ReplaceAll(s, "\n", "\\n")
	content, err := ioutil.ReadFile(".//banner//" + banner)
	if err != nil {
		log.Println("Undefined banner. \n\nTry: standard, thinkertoy, shadow")
		os.Exit(0)
	}
	for _, r := range s {
		if r < 32 || r > 126 {
			log.Println("Undefined character.")
			os.Exit(0)
		}
	}
	arr := SplitLines(strings.Split(string(content), "\n"))
	m := make(map[rune][]string)
	for i, w := range arr {
		m[rune(i+32)] = w
	}
	return m, s
}

// GetAsciiArt function is
func GetAsciiArt(m map[rune][]string, s, option string) [][8]string {
	result := [][8]string{}
	lines := [8]string{}
	switch option {
	case "center":
		return Display(m, s, s, result, lines, false, center)
	case "left":
		return Display(m, s, s, result, lines, false, left)
	case "right":
		return Display(m, s, s, result, lines, false, right)
	case "justify":
		return Display(m, s, s, result, lines, false, justify)
	}
	return Display(m, s, s, result, lines, false, reset)
}

// ToString function is
func ToString(arr [][8]string) string {
	s := ""
	for _, w := range arr {
		for _, q := range w {
			if len(q) == 0 {
				s += "\n"
				break
			}
			s += q + "\n"
		}
	}
	return s
}

// IsNewline function is
func IsNewline(symbol byte, s string) bool {
	if symbol == '\\' && len(s) != 1 && s[1] == 'n' {
		return true
	}
	return false
}

// SplitLine funtion is
func SplitLines(lines []string) [][]string {
	symbol := []string{}
	symbols := [][]string{}
	for i, line := range lines {
		if line != "" {
			symbol = append(symbol, line)
		}
		if (line == "" || i == len(lines)-1) && len(symbol) > 0 {
			symbols = append(symbols, symbol)
			symbol = []string{}
		}
	}
	return symbols
}

type alignFunc func([8]string, string, map[rune][]string) [8]string

// Display function is
func Display(m map[rune][]string, s, str string, result [][8]string, lines [8]string, flag bool, f alignFunc) [][8]string {
	if len(s) == 0 {
		return result
	}
	linesOriginal := lines
	if !IsNewline(s[0], s) {
		for j := range lines {
			lines[j] += m[rune(s[0])][j]
		}
		flag = true
	} else if IsNewline(s[0], s) && Align(lines) {
		result = append(result, f(lines, str[:len(str)-len(s)], m))
		lines = [8]string{}
		if len(s) == 2 && flag {
			result = append(result, f(lines, "", m))
		}
		flag = false
		s = s[1:]
		str = s[1:]
	}

	if len(s) == 1 && flag && Align(lines) {
		result = append(result, f(lines, str, m))
	}

	if Align(lines) {
		return Display(m, s[1:], str, result, lines, flag, f)
	}
	result = append(result, f(linesOriginal, str[:len(str)-len(s)], m))
	lines = [8]string{}
	return Display(m, s, s, result, lines, flag, f)
}

// DisplayLength function is
func DisplayLength() int {
	out, err1 := exec.Command("tput", "cols").Output()
	out1 := strings.Replace(string(out), "\n", "", -1)
	num, err2 := strconv.Atoi(string(out1))
	if err1 != nil || err2 != nil {
		log.Println("Something went wrong while determining terminal length.")
		os.Exit(3)
	}
	return num
}

// Align function is
func Align(line [8]string) bool {
	return len(line[0]) <= DisplayLength()
}

func left(lines [8]string, str string, m map[rune][]string) [8]string {
	if len(lines[0]) == 0 {
		return lines
	}
	var s string
	for len(lines[0])+len(s) < DisplayLength() {
		s += " "
	}
	for i := 0; i < len(lines); i++ {
		lines[i] += s
	}
	return lines
}

func right(lines [8]string, str string, m map[rune][]string) [8]string {
	if len(lines[0]) == 0 {
		return lines
	}
	var s string
	for len(lines[0])+len(s) < DisplayLength() {
		s += " "
	}
	for i := 0; i < len(lines); i++ {
		lines[i] = s + lines[i]
	}
	return lines
}

func center(lines [8]string, str string, m map[rune][]string) [8]string {
	if len(lines[0]) == 0 {
		return lines
	}
	var s string
	for len(lines[0])+len(s) < DisplayLength() {
		s += " "
	}
	for i := 0; i < len(lines); i++ {
		lines[i] = s[:len(s)/2] + lines[i] + s[len(s)/2:]
	}
	return lines
}

func justify(lines [8]string, str string, m map[rune][]string) [8]string {
	if len(lines[0]) == 0 {
		return lines
	}
	var spaceIndex []int
	var index int = 0
	for _, w := range str {
		index += len(m[w][0])
		if w == ' ' {
			spaceIndex = append(spaceIndex, index-1)
		}

	}
	if len(spaceIndex) == 0 {
		return left(lines, str, m)
	}
	for i := 0; len(lines[0]) < DisplayLength(); i++ {
		if i == len(spaceIndex) {
			i = 0
		}
		for j := 0; j < len(lines); j++ {
			lines[j] = lines[j][:spaceIndex[i]] + " " + lines[j][spaceIndex[i]:]
		}
		spaceIndex[i] += i + 1
	}
	return lines
}

func reset(lines [8]string, str string, m map[rune][]string) [8]string {
	return lines
}
