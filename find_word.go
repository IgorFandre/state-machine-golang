package state_machine

func (st *StateMachine) dfs(vert int, word string, depth int, max_depth int) bool {
	if depth > max_depth { // to prevent infinite recursion if there is epsilon cycle
		return false
	}
	found := false
	for i := 0; i < st.conditions; i++ {
		for alpha := range (*st.transitions)[vert][i] {
			if alpha == eps {
				found = found || st.dfs(i, word, depth+1, max_depth)
			}
			if len(word) > 0 && alpha == word[0] {
				found = found || st.dfs(i, word[1:], depth+1, max_depth)
			}
		}
		if found {
			return true
		}
	}
	if word == "" {
		for finish_vert := range st.finishConditions {
			if vert == finish_vert {
				return true
			}
		}
	}
	return false
}

// Check that Nondetermenistic State Machine language contains the given word
func (st *StateMachine) CheckWordInNonDeterministic(word string) bool {
	return st.dfs(st.startCondition, word, 0, len(word)*st.conditions+st.conditions)
}

// Check that Determenistic State Machine language contains the given word
func (st *StateMachine) CheckWordInDeterministic(word string) bool {
	vert := st.startCondition
	for idx := range word {
		found := false
		for i := 0; i < st.conditions; i++ {
			if _, exists := (*st.transitions)[vert][i][word[idx]]; exists {
				vert = i
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	for finish_vert := range st.finishConditions {
		if vert == finish_vert {
			return true
		}
	}
	return false
}
