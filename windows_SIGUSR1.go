package main

import (
	"log"
	"strconv"

	"golang.org/x/sys/windows"
)

/*
Standard Access Rights:
----------------------------
EVENT_ALL_ACCESS (0x1F0003)
EVENT_MODIFY_STATE (0x0002)
----------------------------
Reference: https://msdn.microsoft.com/en-us/library/windows/desktop/ms686670(v=vs.85).aspx
*/

const (
	EVENT_ALL_ACCESS   = "1F0003"
	EVENT_MODIFY_STATE = "0002"

	EVENT_NAME = "SIGUSR1" // Event name must match the value defined in Traefik::server_signals_windows.go
)

func main() {
	namep, _ := windows.UTF16PtrFromString(EVENT_NAME)
	// Convert Standard Access Rights hex code to integer
	i, err := strconv.ParseUint(EVENT_ALL_ACCESS, 16, 32)
	desiredAccess := uint32(i)
	// Open existing event handle
	handle, err := windows.OpenEvent(desiredAccess, false, namep)
	if err != nil {
		log.Fatalf("Failed to create Windows event object: %+v", err)
	}
	if handle == windows.InvalidHandle {
		log.Fatalf("Invalid Windows event object handle")
	}
	lastErr := windows.GetLastError()
	if lastErr == windows.ERROR_ALREADY_EXISTS {
		log.Println("Windows event object already exists")
	}
	// Signal event
	err = windows.SetEvent(handle)
	if err != nil {
		log.Fatalf("Error signaling event: %+v", err)
	}
	windows.CloseHandle(handle)
}
