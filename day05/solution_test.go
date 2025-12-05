package day05

import "testing"

func TestPart1(t *testing.T) {
	res, err := Part1("input_test.txt")
	if err != nil {
		t.Fatalf("Part1 failed: %v", err)
	}
	expected := "3"
	if res != expected {
		t.Errorf("got %s, want %s", res, expected)
	}
}

func TestPart2(t *testing.T) {
	res, err := Part2("input_test.txt")
	if err != nil {
		t.Fatalf("Part2 failed: %v", err)
	}

	expected := "14"
	if res != expected {
		t.Errorf("got %s, want %s", res, expected)
	}
}
