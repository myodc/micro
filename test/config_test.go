// +build integration

package test

import (
	"os/exec"
	"testing"
)

func TestConfig(t *testing.T) {
	serv := newServer(t)
	serv.launch()
	defer serv.close()

	getCmd := exec.Command("micro", "config", "get", "somevalue")
	outp, err := getCmd.CombinedOutput()
	if err == nil {
		t.Fatalf("Config get should fail: %v", string(outp))
	}
	if string(outp) != "not found\n" {
		t.Fatalf("Expected 'not found', got: '%v'", string(outp))
	}

	setCmd := exec.Command("micro", "config", "set", "somevalue", "val1")
	outp, err = setCmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	if string(outp) != "" {
		t.Fatalf("Expected no output, got: %v", string(outp))
	}

	getCmd = exec.Command("micro", "config", "get", "somevalue")
	outp, err = getCmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	if string(outp) != "val1\n" {
		t.Fatalf("Expected 'val1', got: '%v'", string(outp))
	}
}
