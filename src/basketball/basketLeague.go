package basketball

import (
    "io"
    "sort"
    "fmt"
)

type Team struct {
    teamName string
    playerNames []string
}

type League struct {
    Teams []string
    Wins map[string]int
}

type Ranker interface {
    Ranking() []string
}

func (l League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
    // I want to update the Wins field
    if score1 > score2 {
        l.Wins[team1]++
    }else if score1 < score2 {
        l.Wins[team2]++
    }
}

func (l League) Ranking() []string{
    // Need to sort the keys of the map by the value of keys in the map
    type pair struct {
        str string
        val int
    }
    var sortList = []pair{}
    for k,v := range l.Wins {
        sortList = append(sortList, pair{str: k, val: v})
    }
    sort.Slice(sortList, func (i int, j int) bool {
       return sortList[i].val > sortList[j].val
    })
    var winOrder = []string{}
    for _, v := range sortList {
        winOrder = append(winOrder, v.str)
    }
    return winOrder
}

func RankPrinter(ranker Ranker, w io.Writer) {
    league, ok := ranker.(League) // type assertion (actually redundant here, since Ranker interface makes sure)
    if !ok {
        fmt.Errorf("The `ranker` input in RankPrinter function is not of `League` type\n")
    }
    var rankings []string = league.Ranking()
    var bufString string = ""
    for i, v := range rankings {
        bufString += v
        if i < len(rankings) - 1 {
            bufString += "\n"
        }
    }
    io.WriteString(w, bufString)
}



