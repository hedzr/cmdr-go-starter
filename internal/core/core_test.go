package core

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}

func TestCode(t *testing.T) {
	v := Code()
	assertTrue(t, v == 1)
	assertFalse(t, v == 0)
}

//

//

//

func assertTrue(t testing.TB, cond bool, msg ...any) {
	if cond {
		return
	}

	var mesg string
	if len(msg) > 0 {
		if format, ok := msg[0].(string); ok {
			mesg = fmt.Sprintf(format, msg[1:]...)
		} else {
			mesg = fmt.Sprint(msg...)
		}
	}

	t.Fatalf("assertTrue failed: %s", mesg)
}

func assertFalse(t testing.TB, cond bool, msg ...any) {
	if !cond {
		return
	}

	var mesg string
	if len(msg) > 0 {
		if format, ok := msg[0].(string); ok {
			mesg = fmt.Sprintf(format, msg[1:]...)
		} else {
			mesg = fmt.Sprint(msg...)
		}
	}

	t.Fatalf("assertFalse failed: %s", mesg)
}
