package main

import (
	"flag"
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
)

func main() {
	eventName := flag.String("eventId", "1234", "name of the Windows event object to signal")
	flag.Parse()
	namep, _ := windows.UTF16PtrFromString(*eventName)
	// Convert Standard Access Rights hex code to integer
	i, err := strconv.ParseUint(EVENT_ALL_ACCESS, 16, 32)
	desiredAccess := uint32(i)
	// Open existing event handle
	handle, err := windows.OpenEvent(desiredAccess, false, namep)
	if err != nil {
		log.Fatalf("Failed to create Windows event object: %+v", err)
	}
	defer windows.CloseHandle(handle)
	if handle == windows.InvalidHandle {
		log.Fatalf("Invalid Windows event object handle")
	}
	lastErr := windows.GetLastError()
	if lastErr == windows.ERROR_ALREADY_EXISTS {
		log.Println("nothing to see here kids...")
	}
	// Signal event
	err = windows.SetEvent(handle)
	if err != nil {
		log.Fatalf("Error signaling event: %+v", err)
	}
}
