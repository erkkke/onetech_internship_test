package full_outer_join

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

func FullOuterJoin(f1Path, f2Path, resultPath string) {
	f1, f2 := getContentOfFile(f1Path), getContentOfFile(f2Path)

	content := append(f1, f2...)
	content = removeDuplicates(content)
	contentLength := len(content)

	resultString := new(strings.Builder)
	for i, v := range content {
		resultString.WriteString(v)
		if i < contentLength-1 {
			resultString.WriteString("\n")
		}
	}

	writeFile(resultString.String(), resultPath)
}

func writeFile(data, fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(data); err != nil {
		panic(err)
	}
}

func getContentOfFile(path string) []string {
	f, err := os.OpenFile(path, os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	content := make([]string, 0)

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return content
}

func removeDuplicates(content []string) []string {
	wordsCnt := make(map[string]int)
	res := make([]string, 0)

	for _, item := range content {
		wordsCnt[item] += 1
	}

	for k, v := range wordsCnt {
		if v == 1 {
			res = append(res, k)
		}
	}

	sort.Strings(res)

	return res
}
