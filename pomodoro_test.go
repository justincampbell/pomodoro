package main

import (
	"strings"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

var emptyArgs []string

var newTime = make(chan time.Time)
var output = make(chan string, 1)

func Test_parseCommand_start(t *testing.T) {
	now, _ := time.Parse(time.Kitchen, "10:00AM")

	newTime, output := parseCommand(noTime, now, emptyArgs)

	expected, _ := time.Parse(time.Kitchen, "10:25AM")

	assert.T(t, expected.Equal(newTime))

	assert.Equal(t, "", output)
}

func Test_parseCommand_status(t *testing.T) {
	existingTime, _ := time.Parse(time.Kitchen, "10:25AM")
	now, _ := time.Parse(time.Kitchen, "10:05AM")

	newTime, output := parseCommand(existingTime, now, []string{"status"})

	assert.Equal(t, noTime, newTime)
	assert.Equal(t, "20", output)
}

func Test_readExistingTime(t *testing.T) {
	readExistingTime()
}

func Test_filePath(t *testing.T) {
	assert.T(t, strings.HasSuffix(filePath(), "/.pomodoro"))
}
