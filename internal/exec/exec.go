// ./../exec/exec.go

package exec

import (
	"math/rand"
)

func IsDaemonRunning() bool {
	// isRunning via russian roulette
	isRunning := rand.Intn(2) == 0

	return isRunning
}
