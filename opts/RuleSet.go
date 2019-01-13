package opts

type RuleSet struct {
	InsertMap  map[string]*Node
	ExMap      map[string]*Node
	CohMap     map[string]*Node
	Confs      []Tuple
	Selections map[string]bool
}

func (r *RuleSet) AddNewNode(value string, m map[string]*Node) *Node {
	node, pres := m[value]
	if pres {
		return node
	}
	node = &Node{value: value, deps: []*Node{}}
	m[value] = node
	return node

}

func (r *RuleSet) AddDep(a, b string) {

	source := r.AddNewNode(a, r.InsertMap)
	target := r.AddNewNode(b, r.InsertMap)
	source.deps = append(source.deps, target)

	source = r.AddNewNode(a, r.ExMap)
	target = r.AddNewNode(b, r.ExMap)
	target.deps = append(target.deps, source)

	source = r.AddNewNode(a, r.CohMap)
	target = r.AddNewNode(b, r.CohMap)
	target.deps = append(target.deps, source)

	_, e := r.Selections[a]
	if !e {
		r.Selections[a] = false
	}
	_, e = r.Selections[b]
	if !e {
		r.Selections[b] = false
	}

}

func (r *RuleSet) AddConflict(a, b string) {
	r.Confs = append(r.Confs, Tuple{first: a, second: b})

	r.AddNewNode(a, r.CohMap)
	r.AddNewNode(b, r.CohMap)

	r.AddNewNode(a, r.InsertMap)
	r.AddNewNode(b, r.InsertMap)

	r.AddNewNode(a, r.ExMap)
	r.AddNewNode(b, r.ExMap)

}

func (r *RuleSet) CheckConflict(a string, b string) bool {
	visited := make(map[string]bool)
	source := r.CohMap[a]
	stack := []*Node{}

	stack = append(stack, source)
	for len(stack) > 0 {
		var node *Node
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		visited[node.value] = true
		for _, dep := range node.deps {
			if !visited[dep.value] {
				stack = append(stack, dep)
			}
		}
	}

	visitB := make(map[string]bool)
	stack = append(stack, r.CohMap[b])
	for len(stack) > 0 {
		var node *Node
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		visitB[node.value] = true
		val, pres := visited[node.value]
		if pres && val {
			return true
		}
		for _, dep := range node.deps {
			val, pres = visitB[dep.value]
			if !pres {
				stack = append(stack, dep)
				continue
			}
			if pres && !val {
				stack = append(stack, dep)
				continue
			}
		}
	}

	return false

}

func (r *RuleSet) IsCoherent() bool {
	for _, t := range r.Confs {
		if r.CheckConflict(t.first, t.second) || r.CheckConflict(t.second, t.first) {
			return false
		}
	}
	return true
}

func NewRuleSet() RuleSet {
	return RuleSet{InsertMap: map[string]*Node{}, ExMap: map[string]*Node{}, CohMap: map[string]*Node{}, Selections: map[string]bool{}, Confs: []Tuple{}}
}
