/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"crane/cmd"
	_ "crane/init"
	"k8s.io/klog/v2"
)

func init() {
	klog.LogToStderr(false)
}

func main() {
	cmd.Execute()
}
