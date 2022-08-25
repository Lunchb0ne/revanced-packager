package utils

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func JavaPreflightCheck() error {
	// check if we have atleast java 17
	out, err := exec.Command("java", "-version").CombinedOutput()

	if err != nil {
		return err
	}

	// check if we have atleast java 17
	re := regexp.MustCompile(`version\s"(?P<version>(?:\d+\.{0,1})+)"`)
	// get the correct capture group
	match := re.FindStringSubmatch(string(out))
	if match == nil {
		return fmt.Errorf("could not find java version")
	}

	// version check
	version := strings.Split(
		// select the capture group we want
		match[re.SubexpIndex("version")],
		".")[0]
	version_int, err := strconv.Atoi(version)
	if err != nil {
		return err
	}
	if version_int < 17 {
		return fmt.Errorf("java version is less than 17")
	}

	// check if we have zulu flavour of java
	zulu_match, err := regexp.MatchString("[zZ]ulu", string(out))
	if !zulu_match || err != nil {
		return fmt.Errorf("java version is not zulu")
	}

	return nil
}
