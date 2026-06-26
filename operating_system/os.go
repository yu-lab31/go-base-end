// Simple utils to get infomation from operating system on which the current service is running.
package operating_system

import "os"

// GetEnv gets environment variable by its `name`; returns `default` when no variable found.
func GetEnv(name, defaultValue string) string {
	if value := os.Getenv(name); value == "" {
		return defaultValue
	} else {
		return value
	}
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	// do not let service create new file with `path` to avoid unknown behavior
	return true
}
