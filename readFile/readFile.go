package readFile

import(
	"os"
	"bufio"
	"strings"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func SplitByWhiteSpaceToInts(input []string) ([][]int) {
	var output [][]int
    for _, row := range input {
		parsedLine :=  strings.Fields(row)
		var convertedLine []int
		for _, i := range parsedLine {
			convertedInt, _ := strconv.Atoi(i)
			convertedLine = append(convertedLine, convertedInt)
		}
		output = append(output, convertedLine)
	}
	return output
}