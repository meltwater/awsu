package main

import (
	"fmt"
	"math/rand"
)

var randAlphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// CommandError is the checked exception thrown on runtime errors
type CommandError struct {
	msg string // description of error
	err error  // inner error
}

func (e *CommandError) Error() string { return e.msg }

// Panics with a message if the given error isn't nil
func check(err error, a ...interface{}) {
	if err != nil {
		var msg string
		if len(a) > 0 {
			msg = fmt.Sprintf("%s (%s)", fmt.Sprintf(a[0].(string), a[1:]...), err)
		} else {
			msg = fmt.Sprintf("%s", err)
		}

		panic(&CommandError{msg, err})
	}
}

// Panics with a message if the given condition isn't true
func assertThat(condition bool, msg string, a ...interface{}) {
	if !condition {
		panic(&CommandError{fmt.Sprintf(msg, a...), nil})
	}
}

func defaults(a ...string) string {
	for _, item := range a {
		if len(item) > 0 {
			return item
		}
	}

	return ""
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = randAlphabet[rand.Intn(len(randAlphabet))]
	}

	return string(b)
}
