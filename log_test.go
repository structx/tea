package teapot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

func TestLogger_Levels(t *testing.T) {
	var buf bytes.Buffer

	l := New(
		WithLevel(INFO),
		WithWriter(&buf),
	)

	l.printf("hello", DEBUG)
	if buf.Len() > 0 {
		t.Fatal("expected empty buffer")
	}

	l.printf("hello", INFO)
	if buf.Len() == 0 {
		t.Fatal("expected non empty buffer")
	}
}

func TestLogger_JSON(t *testing.T) {
	var buf bytes.Buffer

	l := New(
		WithLevel(INFO),
		WithWriter(&buf),
	)

	l.Info("hello",
		String("user", "alice"),
		StringSlice("x", []string{"hello", "world"}),
	)

	fmt.Println(buf.String())

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}

	if result["user"] != "alice" {
		t.Fatalf("unexpected user value %s expected alice", result["user"])
	}
}

func TestLogger_Error(t *testing.T) {
	var buf bytes.Buffer

	l := New(
		WithLevel(ERROR),
		WithWriter(&buf),
	)

	e := errors.New("something happened")

	l.Error("this is bad", Error(e))

	fmt.Println(buf.String())
}
