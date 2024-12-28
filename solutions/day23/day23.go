package day23

import (
	"aoc_2024/utils"
	"slices"
	"strings"
)

type (
	Graph     map[string][]string
	StringSet map[string]struct{}
)

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

func parseGraph(lines []string) Graph {
	graph := make(Graph)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		graph[parts[0]] = append(graph[parts[0]], parts[1])
		graph[parts[1]] = append(graph[parts[1]], parts[0])
	}
	return graph
}

func findCycles(graph Graph) int {
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

func findPassword(graph Graph) string {
	candidateNodes := make(StringSet)
	for node := range graph {
		candidateNodes[node] = struct{}{}
	}

	var maxClique []string
	var cliques [][]string
	graph.bronKerboschWithPivot(StringSet{}, candidateNodes, StringSet{}, &cliques)

	for _, clique := range cliques {
		if len(clique) > len(maxClique) {
			maxClique = clique
		}
	}

	slices.Sort(maxClique)
	return strings.Join(maxClique, ",")
}

func (g Graph) bronKerboschWithPivot(r, p, x StringSet, cliques *[][]string) {
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

func (g Graph) selectPivot(p, x StringSet) string {
	union := p.union(x)
	for node := range union {
		return node
	}
	return ""
}

func (g Graph) toSet(slice []string) StringSet {
	set := make(StringSet)
	for _, node := range slice {
		set[node] = struct{}{}
	}
	return set
}

func (s StringSet) add(value string) {
	s[value] = struct{}{}
}

func (s StringSet) remove(value string) {
	delete(s, value)
}

func (s StringSet) copy() StringSet {
	newSet := make(StringSet)
	for key := range s {
		newSet[key] = struct{}{}
	}
	return newSet
}

func (s StringSet) intersect(other StringSet) StringSet {
	result := make(StringSet)
	for key := range s {
		if _, exists := other[key]; exists {
			result[key] = struct{}{}
		}
	}
	return result
}

func (s StringSet) difference(other StringSet) StringSet {
	result := make(StringSet)
	for key := range s {
		if _, exists := other[key]; !exists {
			result[key] = struct{}{}
		}
	}
	return result
}

func (s StringSet) union(other StringSet) StringSet {
	result := s.copy()
	for key := range other {
		result[key] = struct{}{}
	}
	return result
}
