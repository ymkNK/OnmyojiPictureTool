package cmd

import "os"

func Execute() {
	err := onmyojiCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
