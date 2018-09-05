package ut_util

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func test_error_msg(t *testing.T, skip int, msg string) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "???"
		line = 1
	}

	errMsg := fmt.Sprintf("\n\n\r%s:%d:\n\n\rError:%s\n", file, line, msg)

	fmt.Printf(errMsg)
	t.Fail()
}

func TEST(t *testing.T, b bool) {
	if !b {
		test_error_msg(t, 2, "Failed Test")
	}
}

func TEST_EQ(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		msg := fmt.Sprint(a, " not equal to ", b, "\n")
		test_error_msg(t, 2, msg)
	}
}

func TEST_STR(t *testing.T, s1 string, s2 string) {
	s1 = strings.Trim(s1, "\n ")
	s2 = strings.Trim(s2, "\n ")
	if s1 != s2 {
		msg := fmt.Sprintf("\n\nstrings don't match\ns1: '%s'\n\ns2: '%s'",
			s1, s2)
		test_error_msg(t, 2, msg)
	}
}
