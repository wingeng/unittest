package ut_util

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

var (
	skipExtra = 0
)

func test_error_msg(t *testing.T, skip int, msg string) {
	skip += 2

	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "???"
		line = 1
	}

	errMsg := fmt.Sprintf("\n\n\r%s:%d:\n\n\rError:%s\n", file, line, msg)

	fmt.Printf(errMsg)
	t.Fail()
}

// Set this in case you wrap TEST, TEST_EQ inside
// a wrapper that needs to be skipped
// Example:
//    func myTestWrapper (s string, exp string) {
//       ut.SkipExtraPush()
//       TEST_STR(s, exp)
//       ut.SkipExtraPop()
//    }
//
//    myTestWrapper("1", "2")  << will report error here
//    myTestWrapper("2", "2")

func SkipExtraPush() {
	skipExtra += 1
}
func SkipExtraPop() {
	if skipExtra > 0 {
		skipExtra -= 1
	}
}

func TEST(t *testing.T, b bool) {
	if !b {
		test_error_msg(t, skipExtra, "Failed Test")
	}
}

func TEST_EQ(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		msg := fmt.Sprint(a, " not equal to ", b, "\n")
		test_error_msg(t, skipExtra, msg)
	}
}

func TEST_STR(t *testing.T, s1 string, s2 string) {
	s1 = strings.Trim(s1, "\n ")
	s2 = strings.Trim(s2, "\n ")
	if s1 != s2 {
		msg := fmt.Sprintf("\n\nstrings don't match\ns1: '%s'\n\ns2: '%s'",
			s1, s2)
		test_error_msg(t, skipExtra, msg)
	}
}
