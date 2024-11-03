package dotnet

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strings"
)

var (
	dotnetVersionRe  = regexp.MustCompile(`([.\d_]+)`)
	runtimeVersionRe = regexp.MustCompile(`\b\d+\.\d+\.\d+\b`)
)

func getDotnetPath() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return findDotnetWindows()
	case "linux", "darwin", "freebsd", "netbsd", "openbsd":
		return findDotnetUnix()
	default:
		return "", fmt.Errorf("unsupported platform")
	}
}

// findDotnetWindows queries the Windows Registry to find the .NET installation path
func findDotnetWindows() (string, error) {
	// Query the registry to find the .NET SDK installation path
	// (This example assumes Go 1.19 or later with support for `exec.Command`)
	var arch string
	switch runtime.GOARCH {
	case "amd64":
		arch = "x64"
	case "arm64":
		arch = "arm64"
	default:
		arch = "x64"
	}
	cmd := exec.Command("reg", "query", fmt.Sprintf(`HKLM\SOFTWARE\dotnet\Setup\InstalledVersions\%s\sharedhost`, arch), "/v", "Version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to query the registry: %v", err)
	}

	// Parse the registry output to find the installation path
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Version") {
			// Assuming the path follows the version information in the registry
			versionPath := filepath.Join("C:\\Program Files\\dotnet", "dotnet.exe")
			return versionPath, nil
		}
	}

	return "", fmt.Errorf(".NET SDK not found in registry")
}

// findDotnetUnix searches common Unix-based paths for the .NET SDK
func findDotnetUnix() (string, error) {
	// Define common paths where dotnet might be installed
	paths := []string{
		"/usr/bin/dotnet",
		"/usr/local/bin/dotnet",
		"/usr/local/share/dotnet/dotnet",
		"/usr/share/dotnet/dotnet",
		"/opt/dotnet/dotnet",
		"/usr/local/dotnet/dotnet",
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf(".NET SDK not found in common directories")
}

func GetRuntimeVersions(config *Config) (runtimeVersions []string, err error) {
	var (
		out bytes.Buffer
	)

	_, err = CheckDotnetVersionInfo(config)
	if config.SdkPath == "" {
		err = fmt.Errorf(".Net SDK not found")
		return
	}
	cmd := exec.Command(config.SdkPath, "--list-runtimes")
	cmd.Stdout = &out
	cmd.Stderr = &out
	result := cmd.Run()
	if result != nil {
		err = fmt.Errorf("failed to check dotnet version: %v", err)
		return
	}

	for _, line := range strings.Split(out.String(), "\n") {
		version := parseRuntimeVersionOutput(line)
		if slices.Contains(runtimeVersions, version) != true {
			runtimeVersions = append(runtimeVersions, version)
		}
	}
	return

}
func CheckDotnetVersionInfo(config *Config) (version string, err error) {
	var (
		out        bytes.Buffer
		dotnetPath string
	)

	if config.SdkPath != "" {
		dotnetPath = config.SdkPath
	} else {
		absPath, fail := getDotnetPath()
		if fail != nil {
			fail = fmt.Errorf("failed to find dotnet SDK: %v", fail)
			return
		}
		dotnetPath = absPath
		config.SdkPath = dotnetPath
	}
	cmd := exec.Command(dotnetPath, "--version")
	cmd.Stdout = &out
	cmd.Stderr = &out
	result := cmd.Run()
	if result != nil {
		err = fmt.Errorf("failed to check dotnet version: %v", err)
		return
	}

	version = parseDotnetVersionOutput(out.String())
	return
}

func parseRuntimeVersionOutput(listRuntimesString string) (runtimeVersion string) {
	if match := runtimeVersionRe.FindString(listRuntimesString); match != "" {
		runtimeVersion = match
	}
	return
}

func parseDotnetVersionOutput(infoString string) string {
	versionString := strings.TrimSpace(infoString)

	if match := dotnetVersionRe.FindStringSubmatch(versionString); len(match) == 2 {
		versionString = match[1]
	}

	return versionString
}
