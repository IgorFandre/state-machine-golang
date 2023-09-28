package state_machine

import (
	"sort"
)

func (st *StateMachine) DoTask() *StateMachine {
	st.DeleteEpsilonTransitions()
	return st.MakeDeterministicMachine()
}

func (st *StateMachine) DeleteEpsilonTransitions() {
	var epsTransitions []edge
	for i := 0; i < st.conditions; i++ {
		for j := 0; j < st.conditions; j++ {
			if _, exists := (*st.transitions)[i][j][eps]; exists {
				epsTransitions = append(epsTransitions, edge{i, j})
			}
		}
	}
	for times := 0; times <= st.conditions; times++ {
		for _, e := range epsTransitions {
			if e.to == e.from {
				continue
			}
			for finish_vert := range st.finishConditions {
				if e.to == finish_vert {
					st.finishConditions[e.from] = exist
				}
			}
			for i := 0; i < st.conditions; i++ {
				for alpha := range (*st.transitions)[i][e.from] {
					if alpha == eps {
						continue
					}
					(*st.transitions)[i][e.to][alpha] = exist
				}
			}
			for j := 0; j < st.conditions; j++ {
				for alpha := range (*st.transitions)[e.to][j] {
					if alpha == eps {
						continue
					}
					(*st.transitions)[e.from][j][alpha] = exist
				}
			}
		}
	}
	for _, e := range epsTransitions {
		delete((*st.transitions)[e.from][e.to], eps)
	}
}

func (st *StateMachine) MakeDeterministicMachine() *StateMachine {
	var new_st StateMachine
	new_st.alphabet = st.alphabet
	new_st.conditions = 1
	new_st.startCondition = 0
	new_st.finishConditions = make(map[int]void)
	new_st.makeTransitions()

	for finish_vert := range st.finishConditions {
		if st.startCondition == finish_vert {
			new_st.finishConditions[new_st.startCondition] = exist
		}
	}

	queue := make([][]int, 1)
	queue[0] = append(queue[0], st.startCondition)

	set := make(map[string]int)
	set[getSliceString(queue[0])] = 0

	for len(queue) > 0 {
		prev_cond := queue[0]
		queue = queue[1:]

		for _, alpha := range new_st.alphabet {
			new_conds := make(map[int]void)

			for _, cond := range prev_cond {
				for to := 0; to < st.conditions; to++ {
					if _, exists := (*st.transitions)[cond][to][alpha]; exists {
						new_conds[to] = exist
					}
				}
			}
			if len(new_conds) == 0 {
				continue
			}
			var new_conds_array []int
			for cond := range new_conds {
				new_conds_array = append(new_conds_array, cond)
			}
			sort.Ints(new_conds_array)
			new_conds_string := getSliceString(new_conds_array)
			if _, exists := set[new_conds_string]; !exists {
				queue = append(queue, new_conds_array)
				set[new_conds_string] = len(set)
				new_st.resizeTransitions(len(set))
				for _, cond := range new_conds_array {
					if _, exists_fin := st.finishConditions[cond]; exists_fin {
						new_st.finishConditions[set[new_conds_string]] = exist
					}
				}
			}
			(*new_st.transitions)[set[getSliceString(prev_cond)]][set[new_conds_string]][alpha] = exist
		}
	}
	return &new_st
}
