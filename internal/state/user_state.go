package state

type UserState int

const (
	StateInitial UserState = iota
	StateImageSent
)

// TODO: Implement GetUserState and SetUserState

func GetUserState(userId int64, userStates map[int64]UserState) UserState {
	if state, exists := userStates[userId]; exists {
		return state
	}
	return StateInitial
}

// SetUserState устанавливает состояние пользователя.
func SetUserState(userId int64, state UserState, userStates map[int64]UserState) {
	if userStates == nil {
		userStates = make(map[int64]UserState)
	}
	userStates[userId] = state
}
