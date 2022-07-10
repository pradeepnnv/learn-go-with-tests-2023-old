package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spy := &SpySleeper{}
	Countdown(buffer, spy)

	got := buffer.String()

	want := `3
2
1
Go!
`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spy.Calls != 3 {
		t.Errorf("not enough calls to sleeper. got %d want %d", spy.Calls, 3)
	}

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v but got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v ,but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
