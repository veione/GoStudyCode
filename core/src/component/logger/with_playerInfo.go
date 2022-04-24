package logger

import "fmt"

func getPlayerInfo(playerInfo interface{}) string {
	return fmt.Sprintf("playerInfo: %v ;", playerInfo)
}

func ErrorFPlayer(playerInfo interface{}, format string, args ...interface{}) {
	format = getPlayerInfo(playerInfo) + format
	Errorf(format, args)
}

func DebugFPlayer(playerInfo interface{}, format string, args ...interface{}) {
	format = getPlayerInfo(playerInfo) + format
	Debugf(format, args)
}

func InfoFPlayer(playerInfo interface{}, format string, args ...interface{}) {
	format = getPlayerInfo(playerInfo) + format
	Infof(format, args)
}
