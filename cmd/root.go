package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

// 供主程序调用
func Execute() error {
	return rootCmd.Execute()
}
