package day15

import (
	"aoc_2024/utils"
	"fmt"
	"strings"
)

type pair struct {
	R, C int
}

type direction struct {
	Dx, Dy int
}

var directions = map[byte]direction{
	'^': {0, -1},
	'v': {0, 1},
	'<': {-1, 0},
	'>': {1, 0},
}

type warehouse struct {
	MoveSeq string
	Boxes   map[pair]struct{}
	Robot   pair
	Walls   map[pair]struct{}
	Width   int
	Height  int
}

type bigBox struct {
	Left  pair
	Right pair
}

type bigWarehouse struct {
	MoveSeq  string
	Boxes    map[bigBox]struct{}
	BoxParts map[pair]bigBox
	Robot    pair
	Walls    map[pair]struct{}
	Width    int
	Height   int
}

// Run runs the day 15 challenge
func Run() error {
	fmt.Println("Day 15:")

	dayPath := "15/input"
	body, err := utils.HTTPGet(dayPath)
	if err != nil {
		return err
	}

	part1Result := part1Puzzle(body)
	fmt.Println("Part 1 result:", part1Result)

	part2Result := part2Puzzle(body)
	fmt.Println("Part 2 result:", part2Result)

	return nil
}

func part1Puzzle(input string) int {
	warehouse := parseWarehouse(input)
	moves := []byte(warehouse.MoveSeq)
	for _, moveDir := range moves {
		move(&warehouse, moveDir)
	}

	score := 0
	for box := range warehouse.Boxes {
		score += 100*box.R + box.C
	}
	return score
}

func part2Puzzle(input string) int {
	warehouse := parseWarehouse(input)
	bigWarehouse := newBigWarehouse(warehouse)
	moves := []byte(bigWarehouse.MoveSeq)
	for _, move := range moves {
		bigMove(&bigWarehouse, move)
	}

	score := 0
	for box := range bigWarehouse.Boxes {
		score += 100*box.Left.R + box.Left.C
	}
	return score
}

func parseWarehouse(input string) warehouse {
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")
	moves := strings.ReplaceAll(parts[1], "\n", "")

	w := warehouse{
		MoveSeq: moves,
		Boxes:   make(map[pair]struct{}),
		Walls:   make(map[pair]struct{}),
	}

	w.Height = len(lines)
	w.Width = len(lines[0])

	for r, line := range lines {
		for c, ch := range line {
			p := pair{r, c}
			switch ch {
			case '#':
				w.Walls[p] = struct{}{}
			case 'O':
				w.Boxes[p] = struct{}{}
			case '@':
				w.Robot = p
			}
		}
	}

	return w
}

func newBigWarehouse(w warehouse) bigWarehouse {
	bw := bigWarehouse{
		MoveSeq:  w.MoveSeq,
		Boxes:    make(map[bigBox]struct{}),
		BoxParts: make(map[pair]bigBox),
		Walls:    make(map[pair]struct{}),
		Height:   w.Height,
		Width:    w.Width * 2,
	}

	// Convert robot position
	bw.Robot = pair{w.Robot.R, w.Robot.C * 2}

	// Convert walls
	for wall := range w.Walls {
		wall1 := pair{wall.R, wall.C * 2}
		wall2 := pair{wall1.R, wall1.C + 1}
		bw.Walls[wall1] = struct{}{}
		bw.Walls[wall2] = struct{}{}
	}

	// Convert boxes
	for box := range w.Boxes {
		left := pair{box.R, box.C * 2}
		right := pair{box.R, left.C + 1}
		bigBox := bigBox{Left: left, Right: right}
		bw.BoxParts[left] = bigBox
		bw.BoxParts[right] = bigBox
		bw.Boxes[bigBox] = struct{}{}
	}

	return bw
}

func move(w *warehouse, dir byte) {
	next := getNextPair(w.Robot, dir)
	_, isWall := w.Walls[next]
	if isWall {
		return
	}
	_, isBox := w.Boxes[next]
	if isBox && canBoxMove(w, next, dir) {
		moveBoxes(w, next, dir)
		w.Robot = next
	} else if !isBox {
		w.Robot = next
	}
}

func canBoxMove(w *warehouse, box pair, dir byte) bool {
	next := getNextPair(box, dir)
	_, isWall := w.Walls[next]
	if isWall {
		return false
	}
	_, isBox := w.Boxes[next]
	if isBox {
		return canBoxMove(w, next, dir)
	}
	return true
}

func moveBoxes(w *warehouse, box pair, dir byte) {
	next := getNextPair(box, dir)
	_, isBox := w.Boxes[next]
	if isBox {
		moveBoxes(w, next, dir)
	}
	delete(w.Boxes, box)
	w.Boxes[next] = struct{}{}
}

func bigMove(w *bigWarehouse, dir byte) {
	next := getNextPair(w.Robot, dir)
	_, isWall := w.Walls[next]
	if isWall {
		return
	}
	_, isBox := w.BoxParts[next]
	if isBox && canBigBoxMove(w, next, dir) {
		bigBoxMove(w, next, dir)
		w.Robot = next
	} else if !isBox {
		w.Robot = next
	}
}

func canBigBoxMove(w *bigWarehouse, side pair, dir byte) bool {
	bb := w.BoxParts[side]
	left, right := bb.Left, bb.Right
	leftNext := getNextPair(left, dir)
	rightNext := getNextPair(right, dir)

	_, lWall := w.Walls[leftNext]
	_, rWall := w.Walls[rightNext]
	if lWall || rWall {
		return false
	}

	if dir == '<' {
		_, lBox := w.BoxParts[leftNext]
		if lBox {
			return canBigBoxMove(w, leftNext, dir)
		}
		return true
	}

	if dir == '>' {
		_, rBox := w.BoxParts[rightNext]
		if rBox {
			return canBigBoxMove(w, rightNext, dir)
		}
		return true
	}

	bbL, lBox := w.BoxParts[leftNext]
	bbR, rBox := w.BoxParts[rightNext]

	canMove := true
	if lBox {
		canMove = canMove && canBigBoxMove(w, leftNext, dir)
	}
	if rBox && bbL != bbR {
		canMove = canMove && canBigBoxMove(w, rightNext, dir)
	}
	return canMove
}

func bigBoxMove(w *bigWarehouse, side pair, dir byte) {
	bb := w.BoxParts[side]
	left, right := bb.Left, bb.Right
	leftNext, rightNext := getNextPair(left, dir), getNextPair(right, dir)

	if dir == '<' {
		_, lBox := w.BoxParts[leftNext]
		if lBox {
			bigBoxMove(w, leftNext, dir)
		}
		delete(w.Boxes, bb)
		delete(w.BoxParts, left)
		delete(w.BoxParts, right)
		bb.Right = bb.Left
		bb.Left = leftNext
		w.Boxes[bb] = struct{}{}
		w.BoxParts[left] = bb
		w.BoxParts[leftNext] = bb
		return
	}

	if dir == '>' {
		_, rBox := w.BoxParts[rightNext]
		if rBox {
			bigBoxMove(w, rightNext, dir)
		}
		delete(w.Boxes, bb)
		delete(w.BoxParts, left)
		delete(w.BoxParts, right)
		bb.Left = bb.Right
		bb.Right = rightNext
		w.Boxes[bb] = struct{}{}
		w.BoxParts[right] = bb
		w.BoxParts[rightNext] = bb
		return
	}

	bbL, lBox := w.BoxParts[leftNext]
	bbR, rBox := w.BoxParts[rightNext]

	if lBox {
		bigBoxMove(w, leftNext, dir)
	}
	if rBox && bbL != bbR {
		bigBoxMove(w, rightNext, dir)
	}

	delete(w.Boxes, bb)
	delete(w.BoxParts, left)
	delete(w.BoxParts, right)
	bb.Left = leftNext
	bb.Right = rightNext
	w.Boxes[bb] = struct{}{}
	w.BoxParts[leftNext] = bb
	w.BoxParts[rightNext] = bb
}

func getNextPair(p pair, dir byte) pair {
	d := directions[dir]
	return pair{p.R + d.Dy, p.C + d.Dx}
}
