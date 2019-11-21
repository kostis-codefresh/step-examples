package main

import (
	"testing"
)

func TestMinimalPom(t *testing.T) {
	coords := readCoords("pom.xml")
	if coords.Group != "com.mycompany.app" {
		t.Errorf("Got = %s; want com.mycompany.app", coords.Group)
	}
	if coords.Artifact != "my-app" {
		t.Errorf("Got = %s; want my-app", coords.Artifact)
	}
	if coords.Version != "1" {
		t.Errorf("Got = %s; want 1", coords.Version)
	}
}

func TestRealPom(t *testing.T) {
	coords := readCoords("pom-example1.xml")
	if coords.Group != "com.javatpoint.application1" {
		t.Errorf("Got = %s; want com.javatpoint.application1", coords.Group)
	}
	if coords.Artifact != "my-application1" {
		t.Errorf("Got = %s; want my-application1", coords.Artifact)
	}
	if coords.Version != "2.0" {
		t.Errorf("Got = %s; want 2.0", coords.Version)
	}
}

func TestSnapshotPom(t *testing.T) {
	coords := readCoords("pom-example2.xml")
	if coords.Group != "com.saucelabs" {
		t.Errorf("Got = %s; want com.saucelabs", coords.Group)
	}
	if coords.Artifact != "sauce_appium_junit" {
		t.Errorf("Got = %s; want sauce_appium_junit", coords.Artifact)
	}
	if coords.Version != "0.0.1-SNAPSHOT" {
		t.Errorf("Got = %s; want 0.0.1-SNAPSHOT", coords.Version)
	}
}
