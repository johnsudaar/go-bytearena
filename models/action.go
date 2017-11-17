package models

const (
	// TODO: SteerAction
	SteerAction = "steer"
	ShootAction = "shoot"
)

type Action struct {
	Method    string      `json:"method"`
	Arguments interface{} `json:"arguments"`
}

type Actions struct {
	Actions []Action
}

func (a *Actions) Shoot(direction Vector2) {
	a.Actions = append(a.Actions, Action{
		Method:    ShootAction,
		Arguments: direction,
	})
}

func (a *Actions) Steer(direction Vector2) {
	a.Actions = append(a.Actions, Action{
		Method:    SteerAction,
		Arguments: direction,
	})
}
