// ./../conf/conf.go

package conf

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	Tasks  []Task
	Editor string
	Matrix MatrixNoti
}

type Task struct {
	Interval int
	Command  string
}

type MatrixNoti struct {
	Enabled     bool
	HomeServer  string
	UserID      string
	RoomID      string
	AccessToken string
}

// Default path to dorei.conf
func ConfigPath() (string, error) {
	usd, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(usd, ".config", "dorei", "dorei.conf"), nil
}

/*
Config parser:

Ignores "#", parses "[section]"

[tasks] contains t:cmd (or time:comand)
[editor] contains cmd=(command to invoke the editor)
[matrix] contains the fields needed by Notify
*/
func ParseConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	var currentSection string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue // Skip comments and empty lines
		}

		// Detect [section] headers
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = strings.ToLower(line[1 : len(line)-1])
			continue
		}

		switch currentSection {
		case "tasks":
			// Format: (timeSec):(cmd)
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				timeRaw := strings.TrimSpace(parts[0])
				cmd := strings.TrimSpace(parts[1])
				// Stupid, but we need to trim the "s"
				// which is there just as a reminder
				// of the granularity dorei interprets
				timeRaw = strings.TrimSuffix(timeRaw, "s")
				timeSec, err := strconv.Atoi(timeRaw)
				if err == nil && timeSec > 0 && cmd != "" {
					task := Task{Interval: timeSec, Command: cmd}
					config.Tasks = append(config.Tasks, task)
				}
			}

		case "editor":
			// Parse key=value: cmd=(cmd)
			if strings.HasPrefix(line, "cmd") {
				config.Editor = strings.TrimSpace(line[4:])
			}

		case "matrix":
			// Parse key=value for MatrixNoti
			if strings.Contains(line, "=") {
				kv := strings.SplitN(line, "=", 2)
				key := strings.ToLower(strings.TrimSpace(kv[0]))
				value := strings.TrimSpace(kv[1])
				switch key {
				case "enabled":
					config.Matrix.Enabled = (value == "true")
				case "home_server":
					config.Matrix.HomeServer = value
				case "user_id":
					config.Matrix.UserID = value
				case "room_id":
					config.Matrix.RoomID = value
				case "access_token":
					config.Matrix.AccessToken = value
				}
			}
		}
	}

	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return config, nil
}

func UpdateMatrixSection(configPath string, enabled bool) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()

	if err := scanner.Err(); err != nil {
		return err
	}

	var newLines []string
	inMatrixSection := false
	keyUpdated := false

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "[") && strings.HasSuffix(trimmedLine, "]") {
			inMatrixSection = (trimmedLine == "[matrix]")
		}

		if inMatrixSection && strings.HasPrefix(trimmedLine, "enabled=") {
			newLines = append(newLines, "enabled="+strconv.FormatBool(enabled))
			keyUpdated = true
		} else {
			newLines = append(newLines, line)
		}
	}

	if inMatrixSection && !keyUpdated {
		newLines = append(newLines, "enabled="+strconv.FormatBool(enabled))
	}

	return os.WriteFile(configPath, []byte(strings.Join(newLines, "\n")+"\n"), 0644)
}
