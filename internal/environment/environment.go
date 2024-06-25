package environment

import (
	"os"
	"fmt"
)

type environment struct {
	environment string
}

var Env *environment

func init() {
	populateEnvironment()
}

// Populates the environment.
func populateEnvironment() {
	envType, exists := os.LookupEnv("env")
	if (!exists) {
		panic("No environment specified")
	}
	switch envType {
	case "dev":
		Env = &devEnvironment
	default:
		panic(fmt.Sprintf("Environment of type '%s' is invalid", envType))
	}
}

func (e *environment) ClearEnvironment() {
	Env = &environment{}
}
