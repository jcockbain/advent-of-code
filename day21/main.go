package main

import (
	_ "embed"
	"regexp"
	"strconv"

	"fmt"

	"github.com/jcockbain/advent-of-code-2021/utils"
)

var (
	benchmark = false
	re        = regexp.MustCompile(`Player (\d+) starting position: (\d+)`)
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func getWrappedPos(p int) int  { return ((p - 1) % 10) + 1 }
func getWrappedDice(x int) int { return ((x - 1) % 1000) + 1 }

func part1() int {
	lines := utils.GetLines(input)
	p1Pos := toInt(re.FindStringSubmatch(lines[0])[2])
	p2Pos := toInt(re.FindStringSubmatch(lines[1])[2])
	dice, rolls, p1Score, p2Score := 1, 0, 0, 0
	for (p1Score < 1000) && (p2Score < 1000) {
		p1Pos = getWrappedPos(p1Pos + (dice * 3) + 3)
		dice = getWrappedDice(dice + 3)
		p1Score += p1Pos
		rolls += 3
		if p1Score >= 1000 {
			break
		}
		p2Pos = getWrappedPos(p2Pos + (dice * 3) + 3)
		dice = getWrappedDice(dice + 3)
		p2Score += p2Pos
		rolls += 3
	}
	if p1Score >= 1000 {
		return p2Score * rolls
	}
	return p1Score * rolls
}

const TARGET = 21

var combos = map[int]int{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

type player struct {
	score int
	pos   int
}

func newPlayer(score int, pos int) *player { return &player{score, pos} }

func (p *player) nextState(roll int) {
	p.pos = getWrappedPos(p.pos + roll)
	p.score += p.pos
}

type game struct {
	player1     player
	player2     player
	player1Turn bool
	universes   int
}

func newGameState(p1Start int, p2Start int) game {
	return game{
		player{0, p1Start},
		player{0, p2Start},
		true,
		1,
	}
}

func (g *game) nextState(d int) *game {
	newGs := game{
		player1:     *newPlayer(g.player1.score, g.player1.pos),
		player2:     *newPlayer(g.player2.score, g.player2.pos),
		player1Turn: g.player1Turn,
		universes:   g.universes,
	}
	if newGs.player1Turn {
		newGs.player1.nextState(d)
	} else {
		newGs.player2.nextState(d)
	}
	newGs.player1Turn = !newGs.player1Turn
	newGs.universes *= combos[d]
	return &newGs
}

type gameStack []game

func (gs *gameStack) pop() game {
	item := (*gs)[len(*gs)-1]
	*gs = (*gs)[:len(*gs)-1]
	return item
}

func (gs *gameStack) push(g game) {
	*gs = append(*gs, g)
}

func part2() int {
	lines := utils.GetLines(input)
	p1Pos := toInt(re.FindStringSubmatch(lines[0])[2])
	p2Pos := toInt(re.FindStringSubmatch(lines[1])[2])
	player1Wins, player2Wins := 0, 0
	stack := gameStack{newGameState(p1Pos, p2Pos)}
	for len(stack) > 0 {
		game := stack.pop()
		for d := 3; d <= 9; d++ {
			newState := *game.nextState(d)
			if (newState.player1.score >= TARGET) || (newState.player2.score >= TARGET) {
				if newState.player1.score >= newState.player2.score {
					player1Wins += newState.universes
				} else {
					player2Wins += newState.universes
				}
			} else {
				stack.push(newState)
			}
		}
	}
	return max(player1Wins, player2Wins)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func toInt(x string) int {
	i, err := strconv.Atoi(x)
	check(err)
	return i
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
