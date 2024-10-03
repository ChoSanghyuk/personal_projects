package bot

import "testing"

func TestBot(t *testing.T) {
	Temp()
}

func TestSendOptions(t *testing.T) {
	tele, err := NewTeleBot("", 0)
	if err != nil {
		t.Error(err)
	}
	tele.sendOptions("HI", []string{"덱1", "덱2"}, []string{"1", "2"})
}
