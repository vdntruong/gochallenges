package soldierranking

func Solution(ranks []int) (rs int) {
	// armyCenter with each rank has reports number
	armyCenter := make(map[int]int)

	for i := 0; i < len(ranks); i++ {
		if armyCenter[ranks[i]] == 0 {
			totalrps := 0
			for _, ortherRank := range ranks {
				if ortherRank == ranks[i]-1 {
					totalrps++
				}
			}
			armyCenter[ranks[i]] = totalrps
		}
	}

	for _, reportOfRank := range armyCenter {
		rs += reportOfRank
	}
	return
}
