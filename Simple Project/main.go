package main

import "fmt"

type LogLevel int

const (
	levelTrace LogLevel = iota
	levelDebug
	levelInfo
)

var levelNames = []string{"Trace", "Debug", "Info"}

func (l LogLevel) String() string {
	if l < levelTrace || l > levelInfo {
		return "Unknown"
	}
	return levelNames[l]
}

func printLogLevel(level LogLevel){
	fmt.Printf("log Level: %d %s\n", level, level.String())
}

func main(){
	printLogLevel(levelInfo)
}