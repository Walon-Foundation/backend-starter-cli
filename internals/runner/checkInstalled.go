package runner

import (
	"os/exec"
	"runtime"
)

// Map tool names to their actual binary
var toolBinaries = map[string]string{
	"Node":   "node",
	"Go":     "go",
	"Python": "python3",
}

var installMatrix = map[string]map[string][]string{
	"Node": {
		"apt":    {"sudo", "apt", "install", "-y", "nodejs", "npm"},
		"yum":    {"sudo", "yum", "install", "-y", "nodejs", "npm"},
		"dnf":    {"sudo", "dnf", "install", "-y", "nodejs", "npm"},
		"pacman": {"sudo", "pacman", "-S", "--noconfirm", "nodejs", "npm"},
		"zypper": {"sudo", "zypper", "install", "-y", "nodejs", "npm"},
		"brew":   {"brew", "install", "node"},
		"winget": {"winget", "install", "OpenJS.NodeJS"},
		"choco":  {"choco", "install", "nodejs"},
	},
	"Go": {
		"apt":    {"sudo", "apt", "install", "-y", "golang"},
		"yum":    {"sudo", "yum", "install", "-y", "golang"},
		"dnf":    {"sudo", "dnf", "install", "-y", "golang"},
		"pacman": {"sudo", "pacman", "-S", "--noconfirm", "go"},
		"zypper": {"sudo", "zypper", "install", "-y", "go"},
		"brew":   {"brew", "install", "go"},
		"winget": {"winget", "install", "GoLang.Go"},
		"choco":  {"choco", "install", "golang"},
	},
	"Python": {
		"apt":    {"sudo", "apt", "install", "-y", "python3"},
		"yum":    {"sudo", "yum", "install", "-y", "python3"},
		"dnf":    {"sudo", "dnf", "install", "-y", "python3"},
		"pacman": {"sudo", "pacman", "-S", "--noconfirm", "python"},
		"zypper": {"sudo", "zypper", "install", "-y", "python3"},
		"brew":   {"brew", "install", "python"},
		"winget": {"winget", "install", "Python.Python.3"},
		"choco":  {"choco", "install", "python"},
	},
}

// Detect OS
func DectectOs() string {
	switch runtime.GOOS {
	case "darwin":
		return "macos"
	case "windows":
		return "windows"
	case "linux":
		return "linux"
	default:
		return "unknown"
	}
}

// Check if tool is installed
func IsInstalled(tool string) bool {
	binary, ok := toolBinaries[tool]
	if !ok {
		return false
	}
	_, err := exec.LookPath(binary)
	return err == nil
}

// Detect package manager
func detectPackageManager(osName string) string {
	switch osName {
	case "linux":
		managers := []string{"apt", "dnf", "yum", "pacman", "zypper"}
		for _, m := range managers {
			if _, err := exec.LookPath(m); err == nil {
				return m
			}
		}
	case "windows":
		managers := []string{"winget", "choco"}
		for _, m := range managers {
			if _, err := exec.LookPath(m); err == nil {
				return m
			}
		}
	case "macos":
		if _, err := exec.LookPath("brew"); err == nil {
			return "brew"
		}
	}
	return ""
}

// Return install command as []string
func InstallTool(toolName, osName string) []string {
	packageManager := detectPackageManager(osName)
	if cmd, ok := installMatrix[toolName][packageManager]; ok {
		return cmd
	}
	return []string{} 
}
