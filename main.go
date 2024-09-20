package main

import (
	"fmt"
)

var levelNodes = [17]uint64{0, 5, 15, 25, 35, 45, 55, 65, 75, 85, 95, 110, 125, 140, 160, 180, 200}

func getLevelFromNodes(nodes uint64) uint64 {
	var level uint64
	if nodes <= 200 {
		var i uint64 = 17
		for ; i > 0; i-- {
			if nodes >= levelNodes[i-1] {
				level = i
				break
			}
		}
	} else if nodes < 1000 {
		level = 17 + ((nodes - 200) / 40)
	} else {
		level = 37 + ((nodes - 1000) / 50)
	}
	return level
}

func getNodesFromLevel(level uint64) uint64 {
	if level < 18 {
		return levelNodes[level-1]
	}
	if level < 38 {
		return 200 + ((level - 17) * 40)
	}
	return 1000 + ((level - 37) * 50)
}

func getMinimum(nodes uint64, level uint64) uint64 {
	if level == 1 {
		return 1
	} else if level == 2 {
		return uint64(float64(nodes) * 0.8)
	}
	eightyPercent := uint64(float64(nodes) * 0.8)
	if eightyPercent < getNodesFromLevel(level-2) {
		return getNodesFromLevel(level - 2)
	} else {
		return eightyPercent
	}
}

func main() {
	var nodes uint64
	for {
		fmt.Print("Nodes: ")
		fmt.Scanln(&nodes)
		level := getLevelFromNodes(nodes)
		getMinimumNodes := getMinimum(nodes, level)
		fmt.Print("Minimum Nodes: ", getMinimumNodes)
		fmt.Println("\n")
	}
}
