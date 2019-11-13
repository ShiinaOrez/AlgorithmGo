package unionfindset

type UnionFindSet struct {
	Fa []int
}

func (set *UnionFindSet) Union(x, y int) {
	set.Fa[set.Find(x)] = y
	return
}

func (set *UnionFindSet) Find(v int) int {
	u := v
	if set.Fa[u] == u {
		set.Fa[v] = set.Find(set.Fa[u])
	} else {
		return u
	}
	return set.Fa[v]
}
