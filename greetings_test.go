package greetings_test

import (
	"regexp"
	"testing"

	"github.com/ksirrah13/greetings"
)

func TestGreeting(t *testing.T) {
	name := "Ben"
	result, err := greetings.GetGreeting(name)
	expected := regexp.MustCompile(".*" + name + ".*")
	if !expected.MatchString(result) || err != nil {
		t.Fatalf("GetGreeting(%q) = %q, %v, wanted match for %q, nil", name, result, err, expected)
	}
}

func TestGreetingNoName(t *testing.T) {
	name := ""
	result, err := greetings.GetGreeting(name)
	if err == nil || result != "" {
		t.Fatalf(`GetGreeting(%q) = %q, %v, wanted "", error`, name, result, err)
	}
}

func TestGreetings(t *testing.T) {
	names := []string{"Tom", "Dick", "Harry"}
	results, err := greetings.GetGreetings(names...)
	if err != nil {
		t.Fatalf("GetGreetings error = %v, wanted nil", err)
	}
	for _, name := range names {
		expected := regexp.MustCompile(".*" + name + ".*")
		if !expected.MatchString(results[name]) || err != nil {
			t.Fatalf("Greeting for %v = %q, wanted match for %q", name, results[name], expected)
		}
	}
}

func TestGreetingsWithNoName(t *testing.T) {
	names := []string{"Tom", "", "Harry"}
	results, err := greetings.GetGreetings(names...)
	if err == nil || results != nil {
		t.Fatalf("GetGreetings = %v, %v, wanted nil, error", results, err)
	}
}
