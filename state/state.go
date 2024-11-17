package state

var currentworkspace string

func GetCurrentWorkspace() string {
	return currentworkspace
}

func SetCurrentWorkspace(workspace string) {
	currentworkspace = workspace
}
