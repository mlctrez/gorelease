package main

import (
	"testing"
)

func Test_cleanupVersion(t *testing.T) {
	if _, err := cleanupVersion(""); err == nil {
		t.Fatal("no version check failed")
	}
	if _, err := cleanupVersion("v1.1"); err == nil {
		t.Fatal("valid semver check failed")
	}
	if _, err := cleanupVersion("1.1"); err == nil {
		t.Fatal("valid semver check failed")
	}
	if sv, err := cleanupVersion("1.1.1"); err != nil {
		t.Fatal("correct semver check failed")
	} else {
		if sv != "v1.1.1" {
			t.Fatal("cleanup of semver failed")
		}
	}
	if sv, err := cleanupVersion("01.02.03"); err != nil {
		t.Fatal("correct semver check failed")
	} else {
		if sv != "v1.2.3" {
			t.Fatal("cleanup of semver failed")
		}
	}

}
