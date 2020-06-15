package clear

import (
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Clear() {
	f, ok := clear[runtime.GOOS]
	if ok {
		f()
	} else {
		panic("unsupported OS")
	}

}
