package state_machine

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestFunctionalTransition(t *testing.T) {
	var st StateMachine
	st.conditions = 5
	st.makeTransitions()
	for i := 0; i < st.conditions; i++ {
		(*st.transitions)[i][i][eps] = exist
	}

	new_condition := st.conditions
	st.resizeTransitions(new_condition)
	for i := 0; i < st.conditions/2; i++ {
		if _, exists := (*st.transitions)[i][i][eps]; !exists {
			t.Fatal("resizeTransitions() have eaten conditions")
		}
	}
	for i := st.conditions / 2; i < st.conditions; i++ {
		(*st.transitions)[i][i][eps] = exist
	}
}

func TestEasyEpsDeletion(t *testing.T) {
	var st StateMachine
	st.alphabet = []byte{'a', 'b'}
	st.conditions = 2
	st.startCondition = 0
	st.finishConditions = map[int]void{1: {}}
	st.makeTransitions()

	(*st.transitions)[0][1]['a'] = exist
	(*st.transitions)[1][0]['b'] = exist
	(*st.transitions)[1][1][eps] = exist
	(*st.transitions)[1][0][eps] = exist

	st.DeleteEpsilonTransitions()

	// Check

	_, exists_1 := (*st.transitions)[0][1]['a']
	_, exists_2 := (*st.transitions)[1][0]['b']
	if !exists_1 || !exists_2 {
		t.Fatal("One of edges was deleted")
	}

	if _, exists_new := (*st.transitions)[0][0]['a']; !exists_new {
		t.Fatal("Func doesn't create new edges")
	}

	_, exists_1 = (*st.transitions)[1][1][eps]
	_, exists_2 = (*st.transitions)[1][0][eps]
	if exists_1 || exists_2 {
		t.Fatal("Func doesn't delete eps edges")
	}
}

func TestEpsDeletion(t *testing.T) {
	var st StateMachine
	st.alphabet = []byte{'a'}
	st.conditions = 3
	st.startCondition = 0
	st.finishConditions = map[int]void{1: {}}
	st.makeTransitions()

	(*st.transitions)[1][2]['b'] = exist
	(*st.transitions)[1][1][eps] = exist
	(*st.transitions)[0][1][eps] = exist

	st.DeleteEpsilonTransitions()

	// Check

	if _, exists_new := (*st.transitions)[0][2]['b']; !exists_new {
		t.Fatal("Func doesn't create new edges")
	}

	_, exists_1 := (*st.transitions)[1][1][eps]
	_, exists_2 := (*st.transitions)[0][1][eps]
	if exists_1 || exists_2 {
		t.Fatal("Func doesn't delete eps edges")
	}

	for finish_vert := range st.finishConditions {
		if finish_vert == 0 || finish_vert == 1 {
			continue
		}
		t.Fatal("Func doesn't create finish conditions")
	}
}

func TestHardEpsDeletion(t *testing.T) {
	var st StateMachine
	st.alphabet = []byte{'a', 'b', 'c'}
	st.conditions = 5
	st.startCondition = 0
	st.finishConditions = map[int]void{4: {}}
	st.makeTransitions()

	(*st.transitions)[0][0]['a'] = exist
	(*st.transitions)[0][2]['a'] = exist
	(*st.transitions)[0][1]['b'] = exist

	(*st.transitions)[1][2][eps] = exist
	(*st.transitions)[1][4]['c'] = exist

	(*st.transitions)[2][3][eps] = exist

	(*st.transitions)[4][2][eps] = exist
	(*st.transitions)[4][3]['a'] = exist

	st.DeleteEpsilonTransitions()

	// Check

	_, exists_1 := (*st.transitions)[0][2]['b']
	_, exists_2 := (*st.transitions)[0][3]['a']
	_, exists_3 := (*st.transitions)[0][3]['b']

	_, exists_4 := (*st.transitions)[1][2]['c']
	_, exists_5 := (*st.transitions)[1][3]['c']

	if !exists_1 || !exists_2 || !exists_3 || !exists_4 || !exists_5 {
		t.Fatal("Func doesn't create new edges")
	}

	_, exists_1 = (*st.transitions)[1][2][eps]
	_, exists_2 = (*st.transitions)[2][3][eps]
	_, exists_3 = (*st.transitions)[4][2][eps]
	if exists_1 || exists_2 || exists_3 {
		t.Fatal("Func doesn't delete eps edges")
	}
}

func TestEasyDeterministic(t *testing.T) {
	var st StateMachine
	st.alphabet = []byte{'a', 'b'}
	st.conditions = 2
	st.startCondition = 0
	st.finishConditions = map[int]void{1: {}}
	st.makeTransitions()

	(*st.transitions)[0][0]['a'] = exist
	(*st.transitions)[0][1]['a'] = exist

	new_st := st.MakeDeterministicMachine()

	// Check

	if new_st.conditions != 2 {
		t.Fatal("Expected 2 conditions")
	}
	if new_st.startCondition != 0 {
		t.Fatal("Expected 0 as start condition")
	}

	_, exists_1 := (*new_st.transitions)[0][1]['a']
	_, exists_2 := (*new_st.transitions)[1][1]['a']
	if !exists_1 || !exists_2 {
		t.Fatal("Func doesn't make right transition")
	}

	_, exists_1 = new_st.finishConditions[0]
	_, exists_2 = new_st.finishConditions[1]
	if exists_1 || !exists_2 {
		t.Fatal("Func make wrong finish conditions")
	}
}

func TestHardDeterministic(t *testing.T) {
	var st StateMachine
	st.alphabet = []byte{'a', 'b'}
	st.conditions = 2
	st.startCondition = 0
	st.finishConditions = map[int]void{1: {}}
	st.makeTransitions()

	(*st.transitions)[0][0]['a'] = exist
	(*st.transitions)[0][0]['b'] = exist

	(*st.transitions)[0][1]['a'] = exist
	(*st.transitions)[1][0]['b'] = exist

	(*st.transitions)[1][1]['b'] = exist

	new_st := st.MakeDeterministicMachine()

	// Check

	if new_st.conditions != 2 {
		t.Fatal("Expected 4 conditions")
	}
	if new_st.startCondition != 0 {
		t.Fatal("Expected 0 as start condition")
	}

	_, exists_1 := new_st.finishConditions[0]
	_, exists_2 := new_st.finishConditions[1]
	if exists_1 || !exists_2 {
		t.Fatal("Func makes wrong finish conditions")
	}

	_, exists_1 = (*new_st.transitions)[0][0]['b']
	_, exists_2 = (*new_st.transitions)[0][1]['a']
	_, exists_3 := (*new_st.transitions)[1][1]['a']
	_, exists_4 := (*new_st.transitions)[1][1]['b']
	total_len := len((*new_st.transitions)[0][0]) + len((*new_st.transitions)[0][1]) + len((*new_st.transitions)[1][0]) + len((*new_st.transitions)[1][1])
	if !exists_1 || !exists_2 || !exists_3 || !exists_4 || total_len != 4 {
		t.Fatal("Func makes wrong edges")
	}
}

// Test takes ~1 minute to pass if 1000
func TestStress(t *testing.T) {
	for i := 0; i < 50; i++ { // 1000
		alphabet := []byte{'a', 'b', 'c'}
		alphabet_2 := append(alphabet, eps)

		var st StateMachine
		var new_st_nondet StateMachine

		st.alphabet = alphabet
		new_st_nondet.alphabet = alphabet

		st.conditions = rand.Intn(5) + 1
		new_st_nondet.conditions = st.conditions

		st.makeTransitions()
		new_st_nondet.makeTransitions()

		trans_num := rand.Intn(10)
		a := 0
		b := rand.Intn(st.conditions)
		alpha := alphabet_2[rand.Intn(len(alphabet_2))]
		for i := 0; i <= trans_num; i++ {
			(*st.transitions)[a][b][alpha] = exist
			(*new_st_nondet.transitions)[a][b][alpha] = exist
			a = rand.Intn(st.conditions)
			b = rand.Intn(st.conditions)
			alpha = alphabet_2[rand.Intn(len(alphabet_2))]
		}

		finish := rand.Intn(st.conditions)

		st.startCondition = 0
		st.finishConditions = make(map[int]void)
		st.finishConditions[finish] = exist

		new_st_nondet.startCondition = 0
		new_st_nondet.finishConditions = make(map[int]void)
		new_st_nondet.finishConditions[finish] = exist
		new_st := new_st_nondet.DoTask()

		// Check

		for i := 0; i < 100; i++ {
			var word string
			for j := 0; j < rand.Intn(10); j++ {
				word = word + string(alphabet[rand.Intn(len(alphabet))])
			}
			bl_1 := new_st.CheckWordInDeterministic(word)
			bl_2 := st.CheckWordInNonDeterministic(word)

			if bl_1 != bl_2 {
				st.PrintMachine("Non deterministic")
				new_st.PrintMachine("Deterministic")
				fmt.Println("Word:", []byte(word), "\nDeterministic got: ", bl_1, "\nNondeterministic got: ", bl_2)
				t.Fatal("Stress test failed")
			}
		}
	}
}
