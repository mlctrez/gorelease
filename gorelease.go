package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func git(args ...string) (output string, err error) {
	cmd := exec.Command("git", args...)
	out, err := cmd.CombinedOutput()
	if out != nil {
		output = string(out)
		fmt.Println(output)
	}
	return output, err
}

func cleanupVersion(version string) (result string, err error) {
	result = version
	if result == "" {
		err = errors.New("version not provided")
		return
	}
	if strings.HasPrefix(result, "v") {
		result = result[1:]
	}

	parts := strings.Split(result, ".")
	if len(parts) != 3 {
		err = errors.New("version is not three part semantic version")
		return
	}

	var major, minor, patch int

	if major, err = strconv.Atoi(parts[0]); err != nil {
		return
	}
	if minor, err = strconv.Atoi(parts[1]); err != nil {
		return
	}
	if patch, err = strconv.Atoi(parts[2]); err != nil {
		return
	}

	return fmt.Sprintf("v%d.%d.%d", major, minor, patch), nil
}

func main() {

	var version string
	var comment string
	flag.StringVar(&version, "version", "", "the version to release in semver format")
	flag.StringVar(&comment, "comment", "", "the comment which will trigger a git add and commit")
	flag.Parse()

	version, err := cleanupVersion(version)
	if err != nil {
		log.Fatal(err)
	}
	if comment != "" {
		git("add", ".")
		git("commit", "-m", comment)
	}

	git("tag", "-a", version, "-m", version)
	git("push", "origin", "HEAD")
	git("push", "origin", version)

}
