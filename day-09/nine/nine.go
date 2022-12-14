package nine

import (
	"strconv"
	"strings"
)

type Nine bool

type axis int

const (
	X = iota
	Y = iota
)

type LongRope struct {
	Knots []*Knot
}

func NewLongRope(l int) LongRope {
	r := LongRope{}
	for i := 0; i < l; i++ {
		k := NewKnot()
		k.ID = i
		if i > 0 {
			r.Knots[i-1].Next = &k
		}
		r.Knots = append(r.Knots, &k)
	}
	return r
}

type Knot struct {
	ID       int
	Position Point
	Visited  map[[2]int]bool
	Next     *Knot
}

func (k *Knot) Sync() {
	k.Visited[[2]int{k.Position.X, k.Position.Y}] = true
}

func (k *Knot) React(e knotMoveEvent) {
	if k.Position.Touch(e.current) {
		return
	}

	// When a move is required, the magnitude of the PointDelta
	// will be {0, 2} or {1, 2}, but we should only move 1 step
	// to get back into contact. so {0, 2} is clamped to {0, 1}.
	// When the lead Knot moves diagonally away, resulting in a
	// delta like {1, 2}. These should be clamped to {1, 1}.
	delta := e.current.Delta(k.Position)
	v := [2]int{}
	for i, j := range [2]int{delta.X, delta.Y} {
		if j > 1 {
			v[i] = j - 1
			continue
		}
		if j < -1 {
			v[i] = j + 1
			continue
		}
		v[i] = j
	}
	nd := PointDelta{X: v[0], Y: v[1]}
	//fmt.Println("new delta:", nd)
	k.MoveTo(k.Position.Add(nd))
}

func (k *Knot) moveWrapper(f func()) {
	p := k.Position
	f()
	k.Sync()
	if k.Next != nil {
		ev := knotMoveEvent{fromID: k.ID, current: k.Position, previous: p}
		//fmt.Println("ID:", k.ID, "current:", ev.current, "previous:", ev.previous)
		k.Next.React(ev)
	}
}

func (k *Knot) MoveTo(p Point) {
	k.moveWrapper(func() {
		k.Position = p
	})
}

func (k *Knot) Move(a axis, c int) {
	if c > 0 {
		for i := 0; i < c; i++ {
			k.moveWrapper(func() {
				if a == X {
					k.Position.X++
				}
				if a == Y {
					k.Position.Y++
				}
			})
		}
	} else {
		for i := 0; i > c; i-- {
			k.moveWrapper(func() {
				if a == X {
					k.Position.X--
				}
				if a == Y {
					k.Position.Y--
				}
			})
		}
	}
}

func NewKnot() Knot {
	k := Knot{}
	k.Position = Point{0, 0}
	k.Visited = map[[2]int]bool{}
	k.Sync()
	return k
}

type knotMoveEvent struct {
	fromID   int
	previous Point
	current  Point
}

type Point struct {
	X int
	Y int
}

func (p Point) Touch(op Point) bool {
	deltas := [2]int{
		p.X - op.X,
		p.Y - op.Y,
	}
	for _, d := range deltas {
		if d < 0 {
			d = d * -1
		}
		if d > 1 {
			return false
		}
	}
	return true
}

func (p Point) Add(d PointDelta) Point {
	return Point{
		X: p.X + d.X,
		Y: p.Y + d.Y,
	}
}

func (p Point) Delta(op Point) PointDelta {
	return PointDelta{
		X: p.X - op.X,
		Y: p.Y - op.Y,
	}
}

type PointDelta struct {
	X int
	Y int
}

type mover interface {
	Move(axis, int)
}

func ExecuteMoves(m mover, b [][]byte) {
	for _, bb := range b {
		move := string(bb)
		//fmt.Println("----->###", move, "###<------")
		dir, c, _ := strings.Cut(move, " ")
		count, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		switch dir {
		case "U":
			m.Move(Y, count)
		case "D":
			m.Move(Y, count*-1)
		case "R":
			m.Move(X, count)
		case "L":
			m.Move(X, count*-1)
		}
	}
}
