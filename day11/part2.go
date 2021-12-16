package day11

func (g OctopusGrid) IsSynchronous() bool {
	return len(g.flashed) == g.grid.Area()
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
