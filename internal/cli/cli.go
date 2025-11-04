// ./../cli/cli.go

package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/bxavaby/dorei/internal/conf"
	"github.com/bxavaby/dorei/internal/exec"
	"github.com/bxavaby/dorei/internal/noti"
)

func Logo() string {
	logo := `
    .___                   .__
  __| _/___________   ____ |__|
 / __ |/  _ \_  __ \_/ __ \|  |
/ /_/ (  <_> )  | \/\  ___/|  |
\____ |\____/|__|    \___  >__|
     \/                  \/

:::::::::::::::::::::::::::::::
      >_ ARR bxavaby 2025     +
:::::::::::::::::::::::::::::::

+-----------------------------+
|    lesser-cron scheduler    |
|      command-line tool      |
+-----------------------------+
`
	return logo
}

func Help() string {
	help := `
Usage: dorei [options]

Options:
  -h, --help          Display this help message
  -v, --version       Display the version number

  -a, --add           Add an entry to the config
                      interactively with editor
  -d, --dry-run       Print all scheduled commands
                      without running
  -m, --matrix        Toggle notifications on/off
                      in [matrix] settings
`
	return help
}

func Version() string {
	version := "v0.1.0"

	return version
}

// Add task to config
func AddTask() {
	return
}

func Run() int {
	// Parse config
	configPath, _ := conf.ConfigPath()
	cfg, err := conf.ParseConfig(configPath)
	if err != nil {
		ohNoes("Error loading config: %s", err)
		return 1
	}

	// Initialize notifier
	notifier, err := noti.New(
		cfg.Matrix.UserID,
		cfg.Matrix.RoomID,
		cfg.Matrix.HomeServer,
		cfg.Matrix.AccessToken,
		cfg.Matrix.Enabled,
	)
	if err != nil {
		ohNoes("Error initializing notifier: %s", err)
		return 1
	}

	if len(os.Args) < 2 {
		// No args == YesOrNo() daemon
		running := exec.IsDaemonRunning()
		var question string
		if running {
			question = "Do you want to stop dorei?"
		} else {
			question = "Do you want to start dorei?"
		}
		if YesOrNo(question) {
			if running {
				singleWell("Stopping dorei...")
				// Call daemon stopper
			} else {
				singleWell("Starting dorei...")
				// Call daemon starter
			}
		} else {
			singleWell("No action taken.")
		}

		return 0
	}

	if len(os.Args) > 2 {
		tripleWell("Use only one argument at a time!")
	}

	arg := strings.ToLower(os.Args[1])

	switch arg {
	case "-h", "--help", "help":
		fmt.Println(Logo())
		fmt.Println(Help())
		return 0
	case "-v", "--version", "version":
		fmt.Println(Version())
		return 0
	case "-a", "--add", "add":
		singleWell("Opening your config...")
		// Logic for "adding", aka opening the config LOL
		AddTask()
		return 0
	case "-d", "--dry-run", "dry-run":
		singleWell("Dry running your set tasks...")
		// Logic for "dry-running", aka printing every cmd
		// in a VERY readable format LOL
		return 0
	case "-m", "--matrix", "matrix":
		// Logic to toggle [matrix] enabled true/flase
		// 1st check its state
		currentState := notifier.IsEnabled()

		// Then reverse it
		newState := !currentState
		notifier.Enable(newState)

		if newState {
			singleWell("Notifications enabled!")
		} else {
			singleWell("Notifications disabled!")
		}
		return 0
	default:
		fmt.Println("Unknown argument: ", os.Args[1])
		fmt.Println(Help())
		return 1
	}
}
