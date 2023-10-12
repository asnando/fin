package actions

type Action interface {
	Execute()
	GetActionName() string
}
