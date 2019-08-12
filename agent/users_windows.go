package agent

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

const WTS_CURRENT_SERVER_HANDLE uintptr = 0

type WTS_INFO_CLASS int

const (
	WTSInitialProgram WTS_INFO_CLASS = iota
	WTSApplicationName
	WTSWorkingDirectory
	WTSOEMId
	WTSSessionId
	WTSUserName
	WTSWinStationName
	WTSDomainName
	WTSConnectState
	WTSClientBuildNumber
	WTSClientName
	WTSClientDirectory
	WTSClientProductId
	WTSClientHardwareId
	WTSClientAddress
	WTSClientDisplay
	WTSClientProtocolType
	WTSIdleTime
	WTSLogonTime
	WTSIncomingBytes
	WTSOutgoingBytes
	WTSIncomingFrames
	WTSOutgoingFrames
	WTSClientInfo
	WTSSessionInfo
	WTSSessionInfoEx
	WTSConfigInfo
	WTSValidationInfo
	WTSSessionAddressV4
	WTSIsRemoteSession
)

var modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
var modWtsapi32 = windows.NewLazySystemDLL("Wtsapi32.dll")

var procWTSGetActiveConsoleSessionId = modkernel32.NewProc("WTSGetActiveConsoleSessionId")
var procWTSQuerySessionInformationA = modWtsapi32.NewProc("WTSQuerySessionInformationA")

func getSessionID() uintptr {
	sessionId, _, _ := procWTSGetActiveConsoleSessionId.Call()
	return sessionId
}

func getUsername(sessionId uintptr) (uintptr, string) {
	var result = make([]byte, 1024)
	var bytesReturned uint32
	ret, _, _ := procWTSQuerySessionInformationA.Call(WTS_CURRENT_SERVER_HANDLE, uintptr(sessionId),
		uintptr(WTSUserName), uintptr(unsafe.Pointer(&result)), uintptr(unsafe.Pointer(&bytesReturned)))

	username := string(result[:bytesReturned])

	return ret, username
}
