package pkg

import (
	"path/filepath"
)

func GetInputPath(fileName string, test bool) string {
	var inputDirPath string
	if test == true {
		inputDirPath = "/Users/Aniek/GolandProjects/advent_of_code_2022/input_test"
	} else {
		inputDirPath = "/Users/Aniek/GolandProjects/advent_of_code_2022/input"
	}
	return filepath.Join(inputDirPath, fileName)
}
