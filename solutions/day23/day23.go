package day23

import (
	"aoc_2024/utils"
	"slices"
	"strings"
)

type (
	graph     map[string][]string
	stringSet map[string]struct{}
)

// Run runs the day 23 challenge
func Run() error {
	dayPath := "23/input"
	body, err := utils.HTTPGet(dayPath)
	if err != nil {
		return err
	}

	part1Result := part1Puzzle(body)
	println("Part 1 result:", part1Result)

	part2Result := part2Puzzle(body)
	println("Part 2 result:", part2Result)

	return nil
}

func part1Puzzle(input string) int {
	graph := parseGraph(strings.Split(strings.TrimSpace(input), "\n"))
	return findCycles(graph)
}

func part2Puzzle(input string) string {
	graph := parseGraph(strings.Split(strings.TrimSpace(input), "\n"))
	return findPassword(graph)
}

func parseGraph(lines []string) graph {
	graph := make(graph)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		graph[parts[0]] = append(graph[parts[0]], parts[1])
		graph[parts[1]] = append(graph[parts[1]], parts[0])
	}
	return graph
}

func findCycles(graph graph) int {
	cycles := make(map[string][]string)
	for node := range graph {
		if node[0] != 't' {
			continue
		}
		for _, neighbor1 := range graph[node] {
			for _, neighbor2 := range graph[neighbor1] {
				if slices.Contains(graph[neighbor2], node) {
					cycle := []string{node, neighbor1, neighbor2}
					slices.Sort(cycle)
					key := strings.Join(cycle, "")
					cycles[key] = cycle
				}
			}
		}
	}
	return len(cycles)
}

func findPassword(graph graph) string {
	candidateNodes := make(stringSet)
	for node := range graph {
		candidateNodes[node] = struct{}{}
	}

	var maxClique []string
	var cliques [][]string
	graph.bronKerboschWithPivot(stringSet{}, candidateNodes, stringSet{}, &cliques)

	for _, clique := range cliques {
		if len(clique) > len(maxClique) {
			maxClique = clique
		}
	}

	slices.Sort(maxClique)
	return strings.Join(maxClique, ",")
}

func (g graph) bronKerboschWithPivot(r, p, x stringSet, cliques *[][]string) {
	if len(p) == 0 && len(x) == 0 {
		clique := make([]string, 0, len(r))
		for node := range r {
			clique = append(clique, node)
		}
		*cliques = append(*cliques, clique)
		return
	}

	pivot := g.selectPivot(p, x)
	for node := range p.difference(g.toSet(g[pivot])) {
		newR := r.copy()
		newR.add(node)
		newP := p.intersect(g.toSet(g[node]))
		newX := x.intersect(g.toSet(g[node]))
		g.bronKerboschWithPivot(newR, newP, newX, cliques)
		p.remove(node)
		x.add(node)
	}
}

func (g graph) selectPivot(p, x stringSet) string {
	union := p.union(x)
	for node := range union {
		return node
	}
	return ""
}

func (g graph) toSet(slice []string) stringSet {
	set := make(stringSet)
	for _, node := range slice {
		set[node] = struct{}{}
	}
	return set
}

func (s stringSet) add(value string) {
	s[value] = struct{}{}
}

func (s stringSet) remove(value string) {
	delete(s, value)
}

func (s stringSet) copy() stringSet {
	newSet := make(stringSet)
	for key := range s {
		newSet[key] = struct{}{}
	}
	return newSet
}

func (s stringSet) intersect(other stringSet) stringSet {
	result := make(stringSet)
	for key := range s {
		if _, exists := other[key]; exists {
			result[key] = struct{}{}
		}
	}
	return result
}

func (s stringSet) difference(other stringSet) stringSet {
	result := make(stringSet)
	for key := range s {
		if _, exists := other[key]; !exists {
			result[key] = struct{}{}
		}
	}
	return result
}

func (s stringSet) union(other stringSet) stringSet {
	result := s.copy()
	for key := range other {
		result[key] = struct{}{}
	}
	return result
}
