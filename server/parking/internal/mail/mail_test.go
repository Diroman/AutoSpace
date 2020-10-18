package mail

import "testing"

func TestSend(t *testing.T) {
	if err := Send("", 800, "Ghbdtn"); err != nil {
		t.Error(err)

	}
}