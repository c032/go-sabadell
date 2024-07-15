package sabadell

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Line struct {
	OperatingDate string
	Concept       string
	ValueDate     string
	Amount        string
	Balance       string
	Ref1          string
	Ref2          string
}

func ParseTXT(r io.Reader) ([]Line, error) {
	var (
		err   error
		lines []Line
	)

	sc := bufio.NewScanner(r)
	for i := 1; sc.Scan(); i++ {
		rawLine := sc.Text()
		rawLine = strings.TrimSpace(rawLine)

		parts := strings.Split(rawLine, "|")
		const expectedLength = 7
		if gotLength := len(parts); gotLength != expectedLength {
			return nil, fmt.Errorf("error at line %d (1-indexed): expected %d items in the line, but got %d", i, expectedLength, gotLength)
		}

		line := Line{
			OperatingDate: parts[0],
			Concept:       parts[1],
			ValueDate:     parts[2],
			Amount:        parts[3],
			Balance:       parts[4],
			Ref1:          parts[5],
			Ref2:          parts[6],
		}

		lines = append(lines, line)
	}

	err = sc.Err()
	if err != nil {
		return nil, fmt.Errorf("error reading lines: %w", err)
	}

	return lines, nil
}

func ParseTXTFile(file string) ([]Line, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer f.Close()

	return ParseTXT(f)
}
