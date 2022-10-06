package main

import (
	"class_main_service/src/cmd"
	"go.uber.org/zap"
	"log"
)

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln(err)
	}
	zap.ReplaceGlobals(logger)

}

func main() {
	root := cmd.GetRoot()
	err := root.Execute()
	if err != nil {
		zap.S().Error(err)
	}

}
