package state_machine

import (
	"testing"
)

func TestDeterministicFind(t *testing.T) {
	var st StateMachine
	st.alphabet = []byte{'a', 'b', 'c'}
	st.conditions = 3
	st.startCondition = 0
	st.finishConditions = map[int]void{1: {}}
	st.makeTransitions()

	(*st.transitions)[0][0]['a'] = exist
	(*st.transitions)[0][1]['b'] = exist
	(*st.transitions)[0][2]['b'] = exist
	(*st.transitions)[1][2]['b'] = exist

	// Check

	found_1 := st.CheckWordInDeterministic("")
	found_2 := st.CheckWordInDeterministic("aaa")
	found_3 := st.CheckWordInDeterministic("ab")
	found_4 := st.CheckWordInDeterministic("b")
	found_5 := st.CheckWordInDeterministic("aaaaaaaaaab")
	found_6 := st.CheckWordInDeterministic("aaaaaaaaaabc")
	if found_1 || found_2 || !found_3 || !found_4 || !found_5 || found_6 {
		t.Fatal("Func got wrong answer")
	}
}

func TestNonDeterministicFind(t *testing.T) {
	var st StateMachine
	st.alphabet = []byte{'a', 'b', 'c'}
	st.conditions = 3
	st.startCondition = 0
	st.finishConditions = map[int]void{1: {}}
	st.makeTransitions()

	(*st.transitions)[0][0][eps] = exist
	(*st.transitions)[0][1][eps] = exist
	(*st.transitions)[1][1]['a'] = exist
	(*st.transitions)[1][2]['b'] = exist
	(*st.transitions)[2][1]['b'] = exist
	(*st.transitions)[0][1]['b'] = exist

	// Check

	found_1 := st.CheckWordInNonDeterministic("")
	found_2 := st.CheckWordInNonDeterministic("aaa")
	found_3 := st.CheckWordInNonDeterministic("ab")
	found_4 := st.CheckWordInNonDeterministic("b")
	found_5 := st.CheckWordInNonDeterministic("aaaaaaaaaab")
	found_6 := st.CheckWordInNonDeterministic("aaaaaaaaaabb")
	found_7 := st.CheckWordInNonDeterministic("aaaaaacaaab")

	if !found_1 || !found_2 || found_3 || !found_4 || found_5 || !found_6 || found_7 {
		t.Fatal("Func got wrong answer")
	}
}
