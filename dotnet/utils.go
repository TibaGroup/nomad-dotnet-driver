package dotnet

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var dotnetVersionCommand = []string{"dotnet", "-version"}

func dotnetVersionInfo() (version, runtime, vm string, err error) {
	var out bytes.Buffer

	cmd := exec.Command(dotnetVersionCommand[0], dotnetVersionCommand[1:]...)
	cmd.Stdout = &out
	cmd.Stderr = &out
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("failed to check dotnet version: %v", err)
		return
	}

	version, runtime, vm = parseDotnetVersionOutput(out.String())
	return
}

var (
	dotnetVersionRe = regexp.MustCompile(`([.\d_]+)`)
)

func parseDotnetVersionOutput(infoString string) (version, runtime, vm string) {
	infoString = strings.TrimSpace(infoString)

	lines := strings.Split(infoString, "\n")
	if strings.Contains(lines[0], "Picked up _JAVA_OPTIONS") {
		lines = lines[1:]
	}

	if len(lines) < 3 {
		// unexpected output format, don't attempt to parse output for version
		return "", "", ""
	}

	versionString := strings.TrimSpace(lines[0])

	if match := dotnetVersionRe.FindStringSubmatch(versionString); len(match) == 2 {
		versionString = match[1]
	}

	return versionString, strings.TrimSpace(lines[1]), strings.TrimSpace(lines[2])
}
