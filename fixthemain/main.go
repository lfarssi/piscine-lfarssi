package main

import "github.com/01-edu/z01" // Import z01 package for PrintRune

// Define Door type and states
type Door struct {
	state string
}

const (
	OPEN  = "open"
	CLOSE = "close"
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...\n")
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("Is the Door opened? \n")
	return ptrDoor.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("Is the Door closed? \n")
	return ptrDoor.state == CLOSE
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...\n")
	ptrDoor.state = OPEN
	return true
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
