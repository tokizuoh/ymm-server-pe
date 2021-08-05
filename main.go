package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func extractCSV(args []string) (string, error) {
	if len(args) != 1 {
		return "", &notExistError{}
	}

	arg := args[0]

	if !strings.HasSuffix(arg, ".csv") {
		return "", &notExistError{}
	}

	return arg, nil
}

type scoreLog struct {
	playerId int
	score    int
}

func parseScoreLogs(lines [][]string) ([]scoreLog, error) {
	var sls []scoreLog

	for _, line := range lines {
		if len(line) != 2 {
			return nil, &invalidElementsCount{n: 2}
		}

		playerId, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, err
		}

		score, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}

		sl := scoreLog{playerId: playerId, score: score}
		sls = append(sls, sl)
	}

	return sls, nil
}

func main() {
	flag.Parse()
	args := flag.Args()

	fileName, err := extractCSV(args)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// ヘッダーを読み込む
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	var lines [][]string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, line)
	}

	sls, err := parseScoreLogs(lines)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range sls {
		log.Println(s.playerId, s.score)
	}

}
