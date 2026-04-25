package logger

import "go.uber.org/zap"

var Log *zap.Logger

func Init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	Log = l
}
