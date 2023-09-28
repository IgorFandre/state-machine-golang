package state_machine

func (st *StateMachine) makeTransitions() {
	transitions := make([][]map[byte]void, st.conditions)
	for i := 0; i < st.conditions; i++ {
		transitions[i] = make([]map[byte]void, st.conditions)
		for j := 0; j < st.conditions; j++ {
			transitions[i][j] = make(map[byte]void)
		}
	}
	st.transitions = &transitions
}

func (st *StateMachine) resizeTransitions(new_conditions int) {
	new_transitions := make([][]map[byte]void, new_conditions)
	for i := 0; i < new_conditions; i++ {
		new_transitions[i] = make([]map[byte]void, new_conditions)
		for j := 0; j < new_conditions; j++ {
			new_transitions[i][j] = make(map[byte]void)
			if i >= st.conditions || j >= st.conditions {
				continue
			}
			for a := range (*st.transitions)[i][j] {
				new_transitions[i][j][a] = exist
			}
		}
	}
	st.transitions = &new_transitions
	st.conditions = new_conditions
}
