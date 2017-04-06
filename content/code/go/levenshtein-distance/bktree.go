package ld

type BKNode struct {
	Str      string
	Children map[int]*BKNode
}

type BKTree struct {
	root         *BKNode
	DistanceFunc func(string, string) int
}

func NewBKTree() *BKTree {
	return &BKTree{
		root:         nil,
		DistanceFunc: EditDistance,
	}
}

func (b *BKTree) Add(s string) {
	if b.root == nil {
		b.root = &BKNode{Str: s, Children: make(map[int]*BKNode)}
		return
	}

	current := b.root
	for {
		d := b.DistanceFunc(s, current.Str)
		if target, ok := current.Children[d]; ok {
			current = target
		} else {
			current.Children[d] = &BKNode{Str: s, Children: make(map[int]*BKNode)}
			break
		}
	}
}

func (b *BKTree) Search(s string, radius int) []*BKNode {
	var resultList []*BKNode
	if b.root == nil {
		return resultList
	}

	var candidates []*BKNode
	candidates = append(candidates, b.root)
	for len(candidates) > 0 {
		candidate := candidates[0]
		candidates = candidates[1:]

		d := b.DistanceFunc(s, candidate.Str)
		if d <= radius {
			resultList = append(resultList, candidate)
		}

		low := d - radius
		high := d + radius
		for distance, node := range candidate.Children {
			if distance >= low && distance <= high {
				candidates = append(candidates, node)
			}
		}
	}

	return resultList
}
