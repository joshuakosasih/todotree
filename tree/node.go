package tree

type statusType uint32

const (
	PENDING statusType = iota
	DONE
	CANCELLED
)

type node struct {
	id int

	title string
	desc  string

	status   statusType
	progress int // 0 - 100

	color string
	shape string

	parent *node
	child  []*node
}

func (n *node) forceDone() bool {
	if n == nil {
		return false
	}

	n.progress = 100
	n.status = DONE
	for i := range n.child {
		n.child[i].forceDone()
	}

	return true
}

func (n *node) markDone() bool {
	if n == nil {
		return false
	}

	for i := range n.child {
		if n.child[i].status == PENDING {
			return false
		}
	}

	n.progress = 100
	n.status = DONE
	return true
}

func (n *node) markCancelled() bool {
	if n == nil {
		return false
	}

	n.status = CANCELLED
	return true
}
