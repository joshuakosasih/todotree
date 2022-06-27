package user

import "joshuako/todotree/tree"

type User struct {
	id    int
	trees []*tree.Tree
}
