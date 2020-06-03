package events_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/whiteCcinn/events"
	"testing"
)

func TestDifferentEventHandlers(t *testing.T) {
	one := events.New()
	two := events.New()

	_ = one.On("one", func() {})
	assert.Equal(t, 1, one.EventCount())
	assert.Equal(t, 0, two.EventCount())
}
