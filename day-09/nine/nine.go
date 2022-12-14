package nine

import (
	"fmt"
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
	fmt.Println("ID:", k.ID, "received move from", e.fromID, "--", e.previous, "to:", e.current)
	if k.Position.Touch(e.current) {
		return
	}

	fmt.Println("ID:", k.ID, "current position:", k.Position, "didn't touch:", e.current)

	d := e.current.Delta(k.Position)
	fmt.Println("delta is:", d)
	if d.X != 0 && d.Y != 0 {
		v := [2]int{}
		for i, j := range [2]int{d.X, d.Y} {
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
		fmt.Println("new delta:", nd)
		k.MoveTo(k.Position.Add(nd))
		return
	}
	if d.X > 0 {
		k.Move(X, d.X-1)
		return
	}
	if d.Y != 0 {
		k.Move(Y, d.Y-1)
		return
	}
}

func (k *Knot) moveWrapper(f func()) {
	p := k.Position
	f()
	k.Sync()
	if k.Next != nil {
		ev := knotMoveEvent{fromID: k.ID, current: k.Position, previous: p}
		fmt.Println("ID:", k.ID, "current:", ev.current, "previous:", ev.previous)
		k.Next.React(ev)
	}
}

func (k *Knot) MoveTo(p Point) {
	k.moveWrapper(func() {
		k.Position = p
	})
}

func (k *Knot) Move(a axis, c int) {
	wrap := func(f func()) {
		p := k.Position
		f()
		k.Sync()
		if k.Next != nil {
			k.Next.React(knotMoveEvent{fromID: k.ID, current: k.Position, previous: p})
		}
	}

	if c > 0 {
		for i := 0; i < c; i++ {
			wrap(func() {
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
			wrap(func() {
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
	wrap := func(f func()) {
		p := h.Position
		f()
		h.Sync()
		if h.Tail != nil {
			h.Tail.React(headMoveEvent{current: h.Position, previous: p})
		}
	}

	if c > 0 {
		for i := 0; i < c; i++ {
			wrap(func() {
				if a == X {
					h.Position.X++
				}
				if a == Y {
					h.Position.Y++
				}
			})
		}
	} else {
		for i := 0; i > c; i-- {
			wrap(func() {
				if a == X {
					h.Position.X--
				}
				if a == Y {
					h.Position.Y--
				}
			})
		}
	}
}

type headMoveEvent struct {
	previous Point
	current  Point
}

type knotMoveEvent struct {
	fromID   int
	previous Point
	current  Point
}

type Tail struct {
	Position Point
	Visited  map[[2]int]bool
	Head     *Head
}

func (t *Tail) SavePosition(p Point) {
	t.Visited[[2]int{p.X, p.Y}] = true
}

func (t *Tail) React(e headMoveEvent) {
	if !t.Position.Touch(e.current) {
		//fmt.Println("tail position", t.Position, "previous head position", e.previous)
		t.Position = e.previous
		t.Sync()
	}
}

func (t *Tail) Sync() {
	t.SavePosition(t.Position)
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

type mover interface {
	Move(axis, int)
}

func ExecuteMoves(m mover, b [][]byte) {
	for _, bb := range b {
		move := string(bb)
		fmt.Println("----->###", move, "###<------")
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
