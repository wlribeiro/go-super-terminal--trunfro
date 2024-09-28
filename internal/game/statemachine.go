package game

// State interface
type State interface {
	Execute(sm *StateMachine, ctx *Context)
}

// Máquina de estados
type StateMachine struct {
	currentState State
}

// Função para criar uma nova máquina de estados
func NewStateMachine() *StateMachine {
	return &StateMachine{}
}

// Método para mudar o estado
func (sm *StateMachine) SetState(state State) {
	sm.currentState = state
}

// Método para processar o estado atual
func (sm *StateMachine) Process(ctx *Context) {
	//if sm.currentState != nil {
	//	sm.currentState.Execute(sm, ctx)
	//} else {
	//	fmt.Println("No current state set")
	//}
	for sm.currentState != nil {
		sm.currentState.Execute(sm, ctx)
	}
}

// Método para processar o estado atual
// func (sm *StateMachine) Process() {
//     for sm.currentState != nil {
//         sm.currentState.Execute(sm)
//     }
// }
