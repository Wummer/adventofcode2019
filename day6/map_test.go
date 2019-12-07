package day6

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func TestSolution1(t *testing.T) {
	expected := 621125
	orbits, err := readFile("input")
	require.NoError(t, err)

	var total int
	for key, _ := range orbits {
		total += getOrbits(key, orbits)
	}
	assert.Equal(t, expected, total)
}

func TestSolution2(t *testing.T) {
	orbits, err := readFile("input")
	require.NoError(t, err)

	actual, ok := shortestTransfer("YOU", "SAN", orbits)
	require.True(t, ok)
	assert.Equal(t, 0, actual)
}

func traverse(from string, orbits map[string]string) []string {
	if next, ok := orbits[from]; ok {
		return append([]string{from}, traverse(next, orbits)...)
	} else {
		return []string{from}
	}
}

func getOrbits(key string, orbits map[string]string) int {
	val, ok := orbits[key]
	if !ok {
		return 0
	}
	return getOrbits(val, orbits) + 1
}

func shortestTransfer(you string, san string, orbits map[string]string) (int, bool) {
	pathYou := traverse(you, orbits)
	pathSan := traverse(san, orbits)
	for i, node := range pathSan {
		if length, ok := getIndex(node, pathYou); ok {
			return length + i - 2, true
		}
	}
	return 0, false
}

func getIndex(element string, data []string) (int, bool) {
	for index, value := range data {
		if element == value {
			return index, true
		}
	}
	return 0, false
}

func readFile(fileName string) (map[string]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	orbits := make(map[string]string)
	for _, entry := range input {
		split := strings.Split(entry, ")")

		orbits[split[1]] = split[0]
	}

	return orbits, nil
}
