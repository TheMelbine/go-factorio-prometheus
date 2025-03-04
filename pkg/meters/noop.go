package meters

import (
	"regexp"
	"strings"
)

var (
	_ Executor = &NoopExecutor{}
	_ Executor = &FakeExecutor{}
)

type NoopExecutor struct{}

// Execute implements Executor.
func (n *NoopExecutor) Execute(cmd string) (string, error) {
	return "", nil
}

type FakeExecutor struct{}

// Execute implements Executor.
func (f *FakeExecutor) Execute(cmd string) (string, error) {
	m := regexp.MustCompile("((1|,%s|%s)+)")
	l := m.FindString(cmd)

	return strings.ReplaceAll(l, "%s", "1"), nil
}
