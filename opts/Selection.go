package opts

type Selection struct {
	RuleSet RuleSet
}

func New(r RuleSet) Selection {
	return Selection{RuleSet: r}
}

func (s Selection) StringSlice() []string {
	keys := []string{}
	for key, value := range s.RuleSet.Selections {
		if value == true {
			keys = append(keys, key)
		}
	}
	return keys
}

func (s Selection) FindCoflictPair(a string /*, m *map[string]bool  */) ([]string, bool) {
	pairs := []string{}
	for _, value := range s.RuleSet.Confs {
		if value.first == a {
			pairs = append(pairs, value.second)
		} else if value.second == a {
			pairs = append(pairs, value.first)

		}
	}
	if len(pairs) > 0 {
		return pairs, true
	}

	return []string{}, false
}

func (s Selection) ToggleHelper(a string) {
	var node *Node
	exStack := []*Node{}
	toggled := map[string]bool{}
	stack := []*Node{}
	visited := make(map[string]bool)

	selected, _ := s.RuleSet.Selections[a]
	if selected {
		node = s.RuleSet.ExMap[a]
	} else {
		node = s.RuleSet.InsertMap[a]
	}

	stack = append(stack, node)
	for len(stack) > 0 {
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		visited[node.value] = true
		selected, _ = s.RuleSet.Selections[node.value]
		if selected {
			s.DisableOption(node, visited, &stack)
		} else {
			s.EnableOption(node, toggled, &exStack, &stack, visited)
		}

		s.CleanUpConflict(exStack, visited)
	}

}

func (s *Selection) EnableOption(node *Node, toggled map[string]bool, exStack *[]*Node, stack *[]*Node, visited map[string]bool) {

	s.RuleSet.Selections[node.value] = true
	pairs, err := s.FindCoflictPair(node.value)
	for _, eachPair := range pairs {

		p, _ := toggled[eachPair]
		if err == true && !p {
			selection, _ := s.RuleSet.Selections[eachPair]
			if selection {
				*exStack = append(*exStack, s.RuleSet.ExMap[eachPair])
			}
		}
		toggled[eachPair] = true
	}
	for _, dep := range node.deps {
		if !visited[dep.value] {
			*stack = append(*stack, dep)
		}
	}

}

func (s *Selection) DisableOption(node *Node, visited map[string]bool, stack *[]*Node) {

	s.RuleSet.Selections[node.value] = false
	for _, dep := range node.deps {
		if !visited[dep.value] {
			*stack = append(*stack, dep)
		}
	}
}

func (s *Selection) CleanUpConflict(exStack []*Node, visited map[string]bool) {
	var node *Node
	for len(exStack) > 0 {
		node, exStack = exStack[len(exStack)-1], exStack[:len(exStack)-1]
		visited[node.value] = true
		s.RuleSet.Selections[node.value] = false
		for _, dep := range node.deps {
			if !visited[dep.value] {
				exStack = append(exStack, dep)
			}
		}
	}
}

func (s Selection) Toggle(a string) {
	s.ToggleHelper(a)
}
