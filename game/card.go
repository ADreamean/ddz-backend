package game

import (
	"math/rand"
	"time"
	"sort"
)

var colors = []int{1, 2, 3, 4}                                        //0 无，红桃，黑桃，方片， 草花
var points = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15} //3,4,5,6,7,8,9,10,J,Q,K,A,2,小王，大王
var defaultCards Cards

func init() {
	defaultCards = make(Cards, 0, 54)
	defaultCards = append(defaultCards, Card{54, 15, 0})
	defaultCards = append(defaultCards, Card{53, 14, 0})
	id := 1
	for _, color := range colors {
		for _, point := range points {
			defaultCards = append(defaultCards, Card{id, point, color})
			id++
		}
	}
}

type Card struct {
	id, point, color int
}

type Cards []Card
type CardType int

func NewCards() Cards {
	return defaultCards.Copy()
}

func (c Cards) Copy() Cards {
	cp := make(Cards, 0, len(c))
	copy(cp, c)
	return cp
}

func (c Cards) Len() int {
	return len(c)
}

func (c Cards) Less(i, j int) bool {
	return c[i].point < c[j].point
}

func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

//洗牌
func (c Cards) Shuffle() Cards {
	rd := rand.New(rand.NewSource(time.Now().Unix()))
	cp := make(Cards, 0, 54)
	for _, i := range rd.Perm(54) {
		cp = append(cp, c[i])
	}
	return cp
}

func (c Cards) Resolve() (Cards, Cards, Cards, Cards) {
	cs := c.Shuffle()
	return cs[:17], cs[17:34], cs[34:51], cs[51:]
}

func (c Cards) Compare(cards Cards) bool {

	// c 为炸弹的时候
	if point, isJocker, ok := boom(c); ok {
		if isJocker {
			return true
		}
		if p, ij, ok := boom(c); ok {
			if ij {
				return false
			}

			return point > p
		}

		return true
	}

	// c 为顺子
	if point, _, ok := straight(c); ok {
		pt, _, _ := straight(cards)
		return point > pt
	}

	if point, _, _, _, ok := tack(c); ok {
		pt, _, _, _, _ := tack(cards)
		return point > pt
	}

	//单牌
	return c[0].point > cards[0].point
}

func count(cards Cards) map[int]int {
	count := make(map[int]int)
	for _, value := range cards {
		if v, ok := count[value.point]; ok {
			count[value.point] = v + 1
		} else {
			count[value.point] = 1
		}
	}

	return count
}

func straight(cards Cards) (point int, repeat int, ok bool) {
	if len(cards) < 5 {
		return
	}
	c := count(cards)

	for _, ct := range c {
		if repeat == 0 {
			repeat = ct
		} else {
			if repeat != ct {
				return
			}

			repeat = ct
		}
	}

	sort.Sort(cards)
	point = cards[len(cards)-1].point
	ok = true

	return
}

func tack(cards Cards) (point int, repeat int, tack int, tackRepeat int, ok bool) {
	if len(cards) < 4 {
		return
	}

	c := count(cards)
	m := 0

	for p, ct := range c {
		if ct == 3 || ct == 4 {
			if m != 0 && ct != m {
				return
			}
			if p > point {
				point = p
			}
			m = ct
			repeat++
			continue
		}

		if ct == 1 || ct == 2 {
			if tack != 0 && tack != ct {
				return
			}
			tack = ct
			tackRepeat++
			continue
		}

		return
	}
	tackRepeat = tackRepeat / repeat
	ok = true
	return
}

func boom(cards Cards) (point int, isJoker, ok bool) {
	if len(cards) == 2 {
		if cards[0].point+cards[1].point == 29 {
			return 0, true, true
		}

		return
	}

	if len(cards) != 4 {
		return
	}

	if cards[0].point == cards[1].point && cards[1].point == cards[2].point && cards[3].point == cards[2].point {
		return cards[0].point, false, true
	}

	return
}
