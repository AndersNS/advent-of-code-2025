package day11

import "testing"

func TestPart1(t *testing.T) {
	res, err := Part1("input_test.txt")
	if err != nil {
		t.Fatalf("Part1 failed: %v", err)
	}

	expected := "5"
	if res != expected {
		t.Errorf("got %s, want %s", res, expected)
	}
}

func TestPart2(t *testing.T) {
	res, err := Part2("input_test_2.txt")
	if err != nil {
		t.Fatalf("Part2 failed: %v", err)
	}

	expected := "2"
	if res != expected {
		t.Errorf("got %s, want %s", res, expected)
	}
}
