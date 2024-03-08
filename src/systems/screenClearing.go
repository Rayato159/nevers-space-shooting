package systems

import (
	"os"
	"os/exec"
	"runtime"
)

var clear = map[string]func(){
	"linux": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
	"windows": func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
}

func ClearScreen() {
	clear, ok := clear[runtime.GOOS]
	if ok {
		clear()
	} else {
		panic("Error: Your platform is unsupported to clear terminal screen.")
	}
}
