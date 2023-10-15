package state_machine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (st *StateMachine) PrintMachine(name string) {
	fmt.Printf("----- %s -----\n", name)
	fmt.Println("Alphabet: " + string(st.alphabet))
	fmt.Println("Conditions: ", st.conditions, "( Start: ", st.startCondition, ")")
	fmt.Println("Finish conditions: ", st.finishConditions)
	fmt.Println("Transitions: \"-- {from} {to} {letter}\"")
	for i := 0; i < st.conditions; i++ {
		for j := 0; j < st.conditions; j++ {
			for k := range (*st.transitions)[i][j] {
				fmt.Println("-- ", i, " ", j, " ", k)
			}
		}
	}
	fmt.Println("-------------------------")
}

func (st *StateMachine) PrintMachineInFormat() {
	fmt.Printf("%d\n\n", st.startCondition)
	for finish_vert := range st.finishConditions {
		fmt.Println(finish_vert)
	}
	fmt.Print("\n")
	for i := 0; i < st.conditions; i++ {
		for j := 0; j < st.conditions; j++ {
			for k := range (*st.transitions)[i][j] {
				if k == eps {
					fmt.Printf("%d %d\n", i, j)
				} else {
					fmt.Printf("%d %d %s\n", i, j, string(k))
				}
			}
		}
	}
}

func (st *StateMachine) UserEnterMachine() {
	fmt.Printf("Enter the alphabet characters separated by a space. Don't use '%s' as a letter in your alphabet.\n", string(eps))
	fmt.Println("In the beginning of the line write the number of characters in the alphabet (e.g. \"3 a b c\"):")

	reader := bufio.NewReader(os.Stdin)
	abcSizeStr, err := reader.ReadString(' ')
	check(err)
	abcSize, err := strconv.Atoi(strings.TrimSpace(abcSizeStr))
	check(err)
	for i := 0; i < abcSize-1; i++ {
		alpha, err := reader.ReadString(' ')
		check(err)
		st.alphabet = append(st.alphabet, alpha[0])
	}
	alpha, err := reader.ReadString('\n')
	check(err)
	st.alphabet = append(st.alphabet, alpha[0])

	fmt.Println("Enter number z of conditions in your machine, then press Enter.\nIn the next step you will be able to use only conditions from 0 to z - 1:")

	condSizeStr, err := reader.ReadString('\n')
	check(err)

	st.conditions, err = strconv.Atoi(strings.TrimSpace(condSizeStr))
	check(err)

	st.makeTransitions()

	fmt.Println("Enter transitions in your machine in the following format: \"{condition} {letter or word} {new_condition}.\"")
	fmt.Printf("To enter epsilon transition use '%s' instead of letter.\n", string(eps))
	fmt.Println("Example: 1 a 0 . After every transition you have to press Enter.\nTo continue write \"-1\" then press Enter.")

	for {
		var q0, q1 int
		var alpha string
		fmt.Scanf("%d %s %d", &q0, &alpha, &q1)
		if q0 == -1 {
			break
		}

		alpha_len := len(alpha)
		if alpha_len == 1 {
			(*st.transitions)[q0][q1][alpha[0]] = exist
		} else {
			new_conditions := st.conditions + alpha_len - 1
			prev_conditions := st.conditions

			st.resizeTransitions(new_conditions)
			(*st.transitions)[q0][prev_conditions][alpha[0]] = exist

			for i := 0; i < alpha_len-2; i++ {
				(*st.transitions)[prev_conditions+i][prev_conditions+i+1][alpha[i+1]] = exist
			}
			(*st.transitions)[prev_conditions+alpha_len-2][q1][alpha[alpha_len-1]] = exist
		}
	}

	fmt.Println("Enter you start condition and then press Enter.")
	fmt.Scanf("%d", &st.startCondition)

	fmt.Println("Enter your finish conditions in the following format: \"{number of finishes} {finish 1} ... {finish n}\". Then press Enter.")
	finishSizeStr, err := reader.ReadString(' ')
	check(err)
	finishSize, err := strconv.Atoi(strings.TrimSpace(finishSizeStr))
	check(err)
	st.finishConditions = make(map[int]void)
	for i := 0; i < finishSize-1; i++ {
		condStr, err := reader.ReadString(' ')
		check(err)
		cond, err := strconv.Atoi(strings.TrimSpace(condStr))
		check(err)
		st.finishConditions[cond] = exist
	}
	condStr, err := reader.ReadString('\n')
	check(err)
	cond, err := strconv.Atoi(strings.TrimSpace(condStr))
	check(err)
	st.finishConditions[cond] = exist
}
