package state_machine

type void struct{}

var exist void

const eps byte = '0'

type StateMachine struct {
	alphabet         []byte
	conditions       int
	startCondition   int
	finishConditions map[int]void
	transitions      *[][]map[byte]void
}

type edge struct {
	from int
	to   int
}
