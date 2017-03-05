package main

import (
	"github.com/Sirupsen/logrus"
	colorable "github.com/mattn/go-colorable"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true, TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp: true})
	logrus.SetOutput(colorable.NewColorableStdout())
}
func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Println("logrus")
	logrus.Debug("debug....default.")

	newLogger := logrus.New()
	newLogger.Formatter = &logrus.TextFormatter{
		ForceColors: true, TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp: true}

	newLogger.Level = logrus.DebugLevel
	newLogger.Out = colorable.NewColorableStdout()
	newLogger.Infoln("info info.")
	newLogger.Println("p....")
	newLogger.Debug("debug...")
	newLogger.Error("err...")

	newLogger.WithField("hello", "world.")
	logrus.WithField("log hello", "world.").Info()
}
