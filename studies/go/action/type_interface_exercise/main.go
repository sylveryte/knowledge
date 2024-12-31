package main

import (
	"fmt"
	"sort"
)

type Team struct {
	name    string
	players []string
}

type League struct {
	Wins map[string]int
}

func (l *League) MatchResult(firstTeam string, firstTeamScore int, secondTeam string, secondTeamScore int) {
	// update wins
	if firstTeamScore < secondTeamScore {
		l.Wins[secondTeam] = l.Wins[secondTeam] + 1
	} else {
		l.Wins[firstTeam] = l.Wins[firstTeam] + 1
	}
}

func (l *League) Print() {
	fmt.Println(l.Wins)
}

func (l *League) Ranking() []string {
	// returns teams in order of wins

	values := []string{}

	for k, _ := range l.Wins {
		values = append(values, k)
	}

	fmt.Println(values)

  // this is new 🔷
	sort.SliceStable(values, func(i, j int) bool { return l.Wins[values[i]] < l.Wins[values[j]] })

	return values
}

func main() {
	s, k := Team{
		name:    "San",
		players: []string{"sandip", "rajiv"},
	}, Team{
		name:    "Kam",
		players: []string{"kandip", "kajiv"},
	}

	l := League{
		Wins: map[string]int{},
	}

	l.MatchResult("rivani", 2, k.name, 1)
	l.MatchResult("rivani", 2, k.name, 1)
	l.MatchResult(s.name, 23, k.name, 43)
	l.MatchResult("jivani", 2, k.name, 1)
	l.MatchResult("jivani", 2, k.name, 1)
	l.MatchResult("jivani", 2, k.name, 1)
	l.MatchResult(s.name, 45, k.name, 4)
	l.MatchResult(s.name, 45, "jivani", 8)
	l.Print()
	fmt.Println(l.Ranking())

}
