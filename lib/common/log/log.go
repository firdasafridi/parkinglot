package log

import "fmt"

// Debug will print the message without new line
func Debugln(args ...interface{}) {
	defaultLog.Debug().Timestamp().Msg(fmt.Sprint(args...))
}

// Debugf will print the message with formater
func Debugf(format string, v ...interface{}) {
	defaultLog.Debug().Timestamp().Msgf(format, v...)
}

// DebugWithFields will print the message with key value
func DebugWithFields(msg string, kv map[string]interface{}) {
	defaultLog.Debug().Timestamp().Fields(kv).Msg(msg)
}

// Print will print the message without new line
func Println(args ...interface{}) {
	defaultLog.Info().Timestamp().Msg(fmt.Sprint(args...))
}

// Printf will print the message with formater
func Printf(format string, v ...interface{}) {
	defaultLog.Info().Timestamp().Msgf(format, v...)
}

// PrintWithFields will print the message with key value
func PrintWithFields(msg string, kv map[string]interface{}) {
	defaultLog.Info().Timestamp().Fields(kv).Msg(msg)
}

// Info will print the message without new line
func Infoln(args ...interface{}) {
	defaultLog.Info().Timestamp().Msg(fmt.Sprint(args...))
}

// Infof will print the message with formater
func Infof(format string, v ...interface{}) {
	defaultLog.Info().Timestamp().Msgf(format, v...)
}

// InfoWithFields will print the message with key value
func InfoWithFields(msg string, kv map[string]interface{}) {
	defaultLog.Info().Timestamp().Fields(kv).Msg(msg)
}

// Error will print the message without new line
func Errorln(args ...interface{}) {
	defaultLog.Error().Timestamp().Msg(fmt.Sprint(args...))
}

// Errorf will print the message with formater
func Errorf(format string, v ...interface{}) {
	defaultLog.Error().Timestamp().Msgf(format, v...)
}

// ErrorWithFields will print the message with key value
func ErrorWithFields(msg string, kv map[string]interface{}) {
	defaultLog.Error().Timestamp().Fields(kv).Msg(msg)
}

// Fatal will print the message without new line
func Fatalln(args ...interface{}) {
	defaultLog.Fatal().Timestamp().MsgFatalln(fmt.Sprint(args...))
}

// Fatalf will print the message with formater
func Fatalf(format string, v ...interface{}) {
	defaultLog.Fatal().Timestamp().MsgFatalf(format, v...)
}

// FatalWithFields will print the message with key value
func FatalWithFields(msg string, kv map[string]interface{}) {
	defaultLog.Fatal().Timestamp().Fields(kv).MsgFatalln(msg)
}

// Panic will print the message without new line
func Panicln(args ...interface{}) {
	defaultLog.Panic().Timestamp().MsgFatalln(fmt.Sprint(args...))
}

// Panicf will print the message with formater
func Panicf(format string, v ...interface{}) {
	defaultLog.Panic().Timestamp().MsgFatalf(format, v...)
}

// PanicWithFields will print the message with key value
func PanicWithFields(msg string, kv map[string]interface{}) {
	defaultLog.Panic().Timestamp().Fields(kv).MsgFatalln(msg)
}

// NoLevel will print the message without new line
func NoLevelln(args ...interface{}) {
	defaultLog.NoLevel().Timestamp().Msg(fmt.Sprint(args...))
}

// NoLevelf will print the message with formater
func NoLevelf(format string, v ...interface{}) {
	defaultLog.NoLevel().Timestamp().Msgf(format, v...)
}

// NoLevelWithFields will print the message with key value
func NoLevelWithFields(msg string, kv map[string]interface{}) {
	defaultLog.NoLevel().Timestamp().Fields(kv).Msg(msg)
}
