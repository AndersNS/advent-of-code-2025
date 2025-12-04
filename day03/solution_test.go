package day03

import "testing"

func TestPart1(t *testing.T) {
	res, err := Part1("input_test.txt")
	if err != nil {
		t.Fatalf("Part1 failed: %v", err)
	}
	expected := "357"
	if res != expected {
		t.Errorf("got %s, want %s", res, expected)
	}
}

func TestPart2(t *testing.T) {
	res, err := Part2("input_test.txt")
	if err != nil {
		t.Fatalf("Part2 failed: %v", err)
	}

	expected := "3121910778619"
	if res != expected {
		t.Errorf("got %s, want %s", res, expected)
	}
}

func TestPart2_real(t *testing.T) {
	res, err := Part2("input.txt")
	if err != nil {
		t.Fatalf("Part2 failed: %v", err)
	}

	expected := "169077317650774"
	if res != expected {
		t.Errorf("got %s, want %s", res, expected)
	}
}
