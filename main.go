package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"ymmerrs"
)

type scoreLog struct {
	playerId int
	score    int
}

type player struct {
	id        int
	meanScore int
	rank      int
}

// extractCSV returns CSV file name from []string.
func extractCSV(args []string) (string, error) {
	if len(args) != 1 {
		return "", &ymmerrs.NotExistError{}
	}

	arg := args[0]

	if !strings.HasSuffix(arg, ".csv") {
		return "", &ymmerrs.NotExistError{}
	}

	return arg, nil
}

// readCSV returns [][]string using *csv.Reader.
func readCSV(reader *csv.Reader, header bool) ([][]string, error) {
	var lines [][]string

	// ヘッダーを読み込む
	line, err := reader.Read()
	if err != nil {
		return nil, err
	}

	if header {
		lines = append(lines, line)
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		lines = append(lines, line)
	}

	return lines, nil
}

// parseScoreLogs returns parsing of [][]string to []scoreLog.
func parseScoreLogs(lines [][]string) ([]scoreLog, error) {
	var sls []scoreLog

	for _, line := range lines {
		if len(line) != 2 {
			return nil, &ymmerrs.InvalidElementsCountError{N: 2}
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

	lines, err := readCSV(reader, false)
	if err != nil {
		log.Fatal(err)
	}

	sls, err := parseScoreLogs(lines)
	if err != nil {
		log.Fatal(err)
	}

	// totalScoreMap. key: id, value: total score
	tsm := make(map[int]int)

	// scoreCountMap. key: id, value: amount of score-log
	scm := make(map[int]int)

	for _, sl := range sls {
		tsm[sl.playerId] += sl.score
		scm[sl.playerId] += 1
	}

	var meanScores [][]int // [[meanScore, id], [meanScore, id] ...]
	for key, value := range tsm {
		id := key
		meanScore := value / scm[id]
		ms := []int{meanScore, id}
		meanScores = append(meanScores, ms)
	}

	sort.Slice(meanScores, func(i, j int) bool {
		return meanScores[i][0] > meanScores[j][0]
	})

	var ps []player

	cr := 1  // current rank
	cs := -1 // current top mean-score
	for _, ms := range meanScores {

		if cs == -1 {
			cs = ms[0]
		} else if cs > ms[0] {
			if len(ps) >= 10 {
				break
			}
			cr += 1
			cs = ms[0]
		}

		p := player{id: ms[1], meanScore: ms[0], rank: cr}
		ps = append(ps, p)
	}

	// stdout result
	fmt.Println("rank,player_id,mean_score")
	for _, p := range ps {
		row := fmt.Sprintf("%v,%v,%v", p.rank, p.id, p.meanScore)
		fmt.Println(row)
	}
}
