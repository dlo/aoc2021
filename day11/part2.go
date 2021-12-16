package day11

func (g OctopusGrid) IsSynchronous() bool {
	return len(g.flashed) == len(g.grid)*len(g.grid[0])
}

func (g *OctopusGrid) FirstSynchronousStep() int {
	step := 0
	for {
		g.Step()
		step++

		if g.IsSynchronous() {
			return step
		}
	}
}
