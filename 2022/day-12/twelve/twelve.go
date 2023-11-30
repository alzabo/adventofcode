package twelve

type WalkerHist map[[2]int]bool

func (wh WalkerHist) Clone() WalkerHist {
	newHist := WalkerHist{}
	for k, v := range wh {
		newHist[k] = v
	}
	return newHist
}

type Map struct {
	Grid  [][]int
	Start [2]int
	Goal  [2]int
}

type HeightMap map[rune]int

func MakeHeightMap() HeightMap {
	m := HeightMap{}
	i := 1
	for r := 'a'; r <= 'z'; r++ {
		m[r] = i
		i++
	}
	return m
}

func MakeMapGrid(input []string, hm HeightMap) Map {
	m := Map{
		Grid: [][]int{},
	}
	for y, i := range input {
		row := []int{}
		for x, key := range i {
			val := 0
			switch key {
			case 'S':
				m.Start = [2]int{x, y}
				val = hm['a']
			case 'E':
				m.Goal = [2]int{x, y}
				val = hm['z']
			default:
				val = hm[key]
			}
			row = append(row, val)
		}
		m.Grid = append(m.Grid, row)
	}
	return m
}
