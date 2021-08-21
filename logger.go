package main

import (
	"os"
	"log"
	"strings"
)

var LogLevelTypes = map[string]int{
	"DEBUG": 4,
	"INFO":  3,
	"WARN":  2,
	"ERROR": 1,
}

var Debug = getBoolEnvVarHelper("DEBUG")

func GetLogLevel() (int, string) {
	if Debug {
		return LogLevelTypes["DEBUG"], "DEBUG"
	}

	var logLevel = getEnvVarOrDefault("LOG_LEVEL", "INFO")
	if value, exist := LogLevelTypes[logLevel]; exist {
		return value, logLevel
	}

	return LogLevelTypes["INFO"], "INFO"
}

func GetLogLevelValue() int {
	value, _ := GetLogLevel()
	return value
}

func Logger(logLevel string, stringList ...string) {
	// Do not call this Logger in a functino that `SendSlackMessage` uses

	logBuilder := SimpleLogger(logLevel, stringList...)

	value, exist := LogLevelTypes[logLevel]
	if exist && GetLogLevelValue() >= value && value != LogLevelTypes["DEBUG"] {
		SendSlackMessage(logBuilder.String())
	}
}

func SimpleLogger(logLevel string, stringList ...string) strings.Builder {
	var logBuilder strings.Builder

	value, exist := LogLevelTypes[logLevel]
	if exist && GetLogLevelValue() >= value {
		var prefix string

		if value == LogLevelTypes["DEBUG"] {
			prefix = "🐛 DEBUG: "
		} else if value == LogLevelTypes["INFO"] {
			prefix = "ℹ️ INFO: "
		} else if value == LogLevelTypes["WARN"] {
			prefix = "🟠 WARN: "
		} else if value == LogLevelTypes["ERROR"] {
			prefix = "🛑 ERROR: "
		}

		logBuilder.WriteString(prefix)
		for _, v := range stringList {
			logBuilder.WriteString(v)
		}

		if value == LogLevelTypes["ERROR"] {
			log.Fatalf(logBuilder.String())
			os.Exit(1)
		} else {
			log.Println(logBuilder.String())
		}
	}

	return logBuilder
}