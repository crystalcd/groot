package versionutil

import (
	"fmt"
	"strconv"
	"strings"
)

func GetNewVersion(version string) (string, error) {
	parts := strings.Split(version, ".")
	if len(parts) != 4 {
		return "", fmt.Errorf("invalid version format: %s", version)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", err
	}
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", err
	}
	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", err
	}
	build, err := strconv.Atoi(parts[3])
	if err != nil {
		return "", err
	}

	// Increment the build number to create a larger version
	build++

	// Format the new version
	newVersion := fmt.Sprintf("%d.%d.%d.%d", major, minor, patch, build)
	return newVersion, nil
}
