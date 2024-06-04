package dotnet

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var dotnetVersionCommand = []string{"dotnet", "--version"}

func DotnetVersionInfo() (version string, err error) {
	var out bytes.Buffer

	cmd := exec.Command(dotnetVersionCommand[0], dotnetVersionCommand[1:]...)
	cmd.Stdout = &out
	cmd.Stderr = &out
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("failed to check dotnet version: %v", err)
		return
	}

	version = parseDotnetVersionOutput(out.String())
	return
}

var (
	dotnetVersionRe = regexp.MustCompile(`([.\d_]+)`)
)

func parseDotnetVersionOutput(infoString string) string {
	versionString := strings.TrimSpace(infoString)

	if match := dotnetVersionRe.FindStringSubmatch(versionString); len(match) == 2 {
		versionString = match[1]
	}

	return versionString
}
