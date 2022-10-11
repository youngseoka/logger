package stackdriver

import (
	"context"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	if err := os.Setenv("DEPLOY_MODE", ""); err != nil {
		t.Fatal(err)
	}

	l, err := New(context.Background(), "test1", "v1", "")
	if err != nil {
		t.Fatal(err)
	}

	if l.deployMode != "dev" {
		t.Errorf("deploy mode not match. expected: %v, got: %v", "dev", l.deployMode)
	}

	if err := os.Setenv("DEPLOY_MODE", "prod"); err != nil {
		t.Fatal(err)
	}

	l, err = New(context.Background(), "test2", "v1", "")
	if err != nil {
		t.Fatal(err)
	}

	if l.deployMode != "prod" {
		t.Errorf("deploy mode not match. expected: %v, got: %v", "prod", l.deployMode)
	}
}
