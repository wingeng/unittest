package ut_util

import (
	"strings"
	"testing"
)

func TEST(t *testing.T, b bool) {
	t.Helper()

	if !b {
		t.Error("Failed Test")
	}
}

func TEST_EQ(t *testing.T, a interface{}, b interface{}) {
	t.Helper()

	if a != b {
		t.Error(a, "is not equal to", b, "\n")
	}
}

func TEST_NEQ(t *testing.T, a interface{}, b interface{}) {
	t.Helper()

	if a == b {
		t.Error(a, " equal to ", b)
	}
}

func TEST_STR(t *testing.T, s1 string, s2 string) {
	t.Helper()

	s1 = strings.Trim(s1, "\n ")
	s2 = strings.Trim(s2, "\n ")
	if s1 != s2 {
		t.Errorf("\n\nstrings don't match\ns1: '%s'\n\ns2: '%s'",
			s1, s2)
	}
}
