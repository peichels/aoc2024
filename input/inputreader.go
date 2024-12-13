package input

import (
	"bufio"
	"os"
)

func ReadInputFileAsStringList(inputFileName string) ([]string, error) {
	file, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var stringList []string
	for scanner.Scan() {
		line := scanner.Text()
		stringList = append(stringList, line)

	}
	return stringList, nil
}
