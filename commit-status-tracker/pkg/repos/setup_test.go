package repos

import (
	"testing"

	"github.com/jenkins-x/go-scm/scm/factory"
)

func TestSetup(t *testing.T) {
	err := Setup("testdata/repos.yaml")
	if err != nil {
		t.Fatal(err)
	}

	driver, err := factory.DefaultIdentifier.Identify("example.com")
	if err != nil {
		t.Fatal(err)
	}

	if driver != "github" {
		t.Fatalf("got %q, want %q", driver, "github")
	}
}

func TestSetupMissingFile(t *testing.T) {
	err := Setup("testdata/unknown.yaml")
	if err == nil {
		t.Fatal("expected an error reading an unknown file")
	}
}

func TestSetupInvalidFile(t *testing.T) {
	err := Setup("setup_test.go")

	if err == nil {
		t.Fatal("expected an error reading an unknown file")
	}
}
