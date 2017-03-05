package util

import (
	"bufio"
	"io"
)

// ExtractTomlStr reciever an input of hugo Toml frontmatter file
// and returns only toml frontmatter
func ExtractTomlStr(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	delim := scanner.Text()

	var text string
	var tomlStr string
	for ; text != delim; text = scanner.Text() {
		scanner.Scan()
		if len(text) > 0 {
			tomlStr += text + "\n"
		}
	}
	return tomlStr, nil
}
