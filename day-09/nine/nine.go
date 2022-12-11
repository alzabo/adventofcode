package nine

import (
	"errors"
	"strconv"
	"strings"
)

type Nine bool

type axis int

const (
	X = iota
	Y = iota
)

type Rope struct {
	Head Head
	Tail Tail
}

func (r *Rope) Move(a axis, c int) {
	r.Head.Move(a, c)
}

func NewRope() Rope {
	r := Rope{}
	r.Head = NewHead()
	r.Tail = NewTail()
	r.Head.Tail = &r.Tail
	r.Tail.Head = &r.Head
	return r
}

type Head struct {
	Position Point
	Visited  map[[2]int]bool
	Tail     *Tail
}

func (h *Head) SavePosition(p Point) {
	h.Visited[[2]int{p.X, p.Y}] = true
}

func (h *Head) Sync() {
	h.SavePosition(h.Position)
}

func (h *Head) Move(a axis, c int) {
	if c > 0 {
		for i := 0; i < c; i++ {
			if a == X {
				h.Position.X++
			}
			if a == Y {
				h.Position.Y++
			}
			h.Sync()
		}
	} else {
		for i := 0; i > c; i-- {
			if a == X {
				h.Position.X--
			}
			if a == Y {
				h.Position.Y--
			}
			h.Sync()
		}
	}
}

type Tail struct {
	Position Point
	Visited  map[[2]int]bool
	Head     *Head
}

func (t *Tail) SavePosition(p Point) {
	t.Visited[[2]int{p.X, p.Y}] = true
}

func (t *Tail) Sync() {
	t.SavePosition(t.Position)
}

func (t *Tail) EndsTouch() (bool, error) {
	if t.Head == nil {
		return false, errors.New("relation to end cannot be computed when Head is not set")
	}
	deltas := []int{
		t.Position.X - t.Head.Position.X,
		t.Position.Y - t.Head.Position.Y,
	}
	for _, d := range deltas {
		if d < 0 {
			d = d * -1
		}
		if d > 1 {
			return false, nil
		}
	}
	return true, nil

}

type Point struct {
	X int
	Y int
}

func NewHead() Head {
	h := Head{}
	h.Position.X = 0
	h.Position.Y = 0
	h.Visited = map[[2]int]bool{}
	h.SavePosition(h.Position)
	return h
}

func NewTail() Tail {
	t := Tail{}
	t.Position.X = 0
	t.Position.Y = 0
	t.Visited = map[[2]int]bool{}
	t.SavePosition(t.Position)
	return t
}

func ExecuteMoves(r *Rope, b [][]byte) {
	for _, bb := range b {
		move := string(bb)
		dir, c, _ := strings.Cut(move, " ")
		count, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		switch dir {
		case "U":
			r.Move(Y, count)
		case "D":
			r.Move(Y, count*-1)
		case "R":
			r.Move(X, count)
		case "L":
			r.Move(X, count*-1)
		}
	}
}
