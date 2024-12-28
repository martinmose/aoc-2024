package day24

import (
	"aoc_2024/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type gateInfo struct {
	operation int // 0=AND, 1=OR, 2=XOR
	inputs    []string
	output    string
}

// Run runs the day 24 challenge
func Run() error {
	dayPath := "24/input"
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

func part1Puzzle(input string) string {
	wires, gates := parseInput(input)

	for len(gates) > 0 {
		for wireName, gate := range gates {
			if canEvalGate(gate, wires) {
				wires[wireName] = evaluateGate(gate, wires)
				delete(gates, wireName)
			}
		}
	}

	var zWires []string
	for wire := range wires {
		if strings.HasPrefix(wire, "z") {
			zWires = append(zWires, wire)
		}
	}
	sort.Strings(zWires)

	result := 0
	for i := len(zWires) - 1; i >= 0; i-- {
		result = (result << 1) | wires[zWires[i]]
	}

	return strconv.Itoa(result)
}

func part2Puzzle(input string) string {
	_, gates := parseInput(input)
	var swapped []string
	var carry string

	var gateStrings []string
	for wireName, gate := range gates {
		gateStr := fmt.Sprintf("%s %s %s -> %s",
			gate.inputs[0],
			[]string{"AND", "OR", "XOR"}[gate.operation],
			gate.inputs[1],
			wireName)
		gateStrings = append(gateStrings, gateStr)
	}

	for i := 0; i < 45; i++ {
		n := fmt.Sprintf("%02d", i)
		var m1, n1, r1, z1, c1 string

		m1 = find(fmt.Sprintf("x%s", n), fmt.Sprintf("y%s", n), "XOR", gateStrings)
		n1 = find(fmt.Sprintf("x%s", n), fmt.Sprintf("y%s", n), "AND", gateStrings)

		if carry != "" {
			r1 = find(carry, m1, "AND", gateStrings)
			if r1 == "" {
				m1, n1 = n1, m1
				swapped = append(swapped, m1, n1)
				r1 = find(carry, m1, "AND", gateStrings)
			}

			z1 = find(carry, m1, "XOR", gateStrings)

			if strings.HasPrefix(m1, "z") {
				m1, z1 = z1, m1
				swapped = append(swapped, m1, z1)
			}
			if strings.HasPrefix(n1, "z") {
				n1, z1 = z1, n1
				swapped = append(swapped, n1, z1)
			}
			if strings.HasPrefix(r1, "z") {
				r1, z1 = z1, r1
				swapped = append(swapped, r1, z1)
			}

			c1 = find(r1, n1, "OR", gateStrings)
		}

		if strings.HasPrefix(c1, "z") && c1 != "z45" {
			c1, z1 = z1, c1
			swapped = append(swapped, c1, z1)
		}

		if carry == "" {
			carry = n1
		} else {
			carry = c1
		}
	}

	sort.Strings(swapped)
	return strings.Join(swapped, ",")
}

func parseInput(input string) (map[string]int, map[string]gateInfo) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	wires := make(map[string]int)
	gates := make(map[string]gateInfo)

	for _, line := range strings.Split(strings.TrimSpace(parts[0]), "\n") {
		parts := strings.Split(line, ": ")
		wires[parts[0]] = atoiNoErr(parts[1])
	}

	for _, line := range strings.Split(strings.TrimSpace(parts[1]), "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		inputs := strings.Split(parts[0], " ")

		var operation int
		var ins []string

		if len(inputs) == 3 {
			switch inputs[1] {
			case "AND":
				operation = 0
			case "OR":
				operation = 1
			case "XOR":
				operation = 2
			}
			ins = []string{inputs[0], inputs[2]}
		}

		gates[parts[1]] = gateInfo{
			operation: operation,
			inputs:    ins,
			output:    parts[1],
		}
	}

	return wires, gates
}

func evaluateGate(gate gateInfo, wires map[string]int) int {
	in1 := wires[gate.inputs[0]]
	in2 := wires[gate.inputs[1]]

	switch gate.operation {
	case 0:
		return in1 & in2
	case 1:
		return in1 | in2
	case 2:
		return in1 ^ in2
	}
	return 0
}

func canEvalGate(gate gateInfo, wires map[string]int) bool {
	_, hasIn1 := wires[gate.inputs[0]]
	_, hasIn2 := wires[gate.inputs[1]]
	return hasIn1 && hasIn2
}

func find(a, b, operator string, gates []string) string {
	for _, gate := range gates {
		if strings.HasPrefix(gate, fmt.Sprintf("%s %s %s", a, operator, b)) ||
			strings.HasPrefix(gate, fmt.Sprintf("%s %s %s", b, operator, a)) {
			parts := strings.Split(gate, " -> ")
			return parts[len(parts)-1]
		}
	}
	return ""
}

func atoiNoErr(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
