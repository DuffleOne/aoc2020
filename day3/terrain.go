package main

type Terrain struct {
	X   int
	Y   int
	Map [][]int
}

func (t *Terrain) Extend() {
	for i := range t.Map {
		t.Map[i] = append(t.Map[i], t.Map[i]...)
	}
}

func (t *Terrain) Navigate(x, y int) {
	t.X += x
	t.Y += y
}

func (t *Terrain) GetPosition() (int, int, int) {
	if t.Y >= len(t.Map) {
		return 0, 0, -1
	}

	if t.X >= len(t.Map[t.Y]) {
		t.Extend()
	}

	return t.X, t.Y, t.Map[t.Y][t.X]
}

func (t *Terrain) ResetPosition() {
	t.X = 0
	t.Y = 0
}
