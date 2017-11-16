package models

const (
	SteerAction = "steer"
	ShootAction = "shoot"
)

type Action struct {
	Method    string      `json:"method"`
	Arguments interface{} `json:"arguments"`
}

type Actions []Action

func (a Actions) Shoot(direction Vector2) Actions {
	return append(a, Action{
		Method:    ShootAction,
		Arguments: direction,
	})
}

func (a Actions) Steer(direction Vector2) Actions {
	return append(a, Action{
		Method:    SteerAction,
		Arguments: direction,
	})
}
