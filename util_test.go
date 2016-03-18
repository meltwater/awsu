package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	// Handle checked errors nicely
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case *CommandError:
				assert.Equal(t, "Test Error (Inner Error)", fmt.Sprintf("%s", err))
			default:
				t.Errorf("Expected to catch a CommandError but got %v", err)
			}
		}
	}()

	check(errors.New("Inner Error"), "Test Error")
}

func TestAssert(t *testing.T) {
	// Handle checked errors nicely
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case *CommandError:
				assert.Equal(t, "Test Error", fmt.Sprintf("%s", err))
			default:
				t.Errorf("Expected to catch a CommandError but got %v", err)
			}
		}
	}()

	assertThat(false, "Test Error")
}

func TestDefaults(t *testing.T) {
	assert.Equal(t, "abc", defaults("abc", "123"))
	assert.Equal(t, "123", defaults("", "123"))
	assert.Equal(t, "", defaults("", ""))
	assert.Equal(t, "", defaults(""))
	assert.Equal(t, "", defaults())
}
