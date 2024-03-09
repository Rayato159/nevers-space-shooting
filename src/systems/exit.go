package systems

import (
	"os"

	"github.com/nsf/termbox-go"
)

func IsExit(event termbox.Event) {
	switch event.Type {
	case termbox.EventKey:
		switch event.Key {
		case termbox.KeyEsc:
			termbox.Close()
			os.Exit(0)
		}
	case termbox.EventError:
		panic(event.Err)
	}
}
