package cmd

import (
	"github.com/shipu/artifact/cmd/generate"
	"github.com/spf13/cobra"
	"log"
)

var RootCmd = &cobra.Command{
	Use:   "art",
	Short: "A simple artifact command",
	Long:  `A simple artifact command`,
	Run:   hello,
}

func hello(cmd *cobra.Command, args []string) {
	log.Println("art command")
}

// add command if necessary
func AddCommand(cmd *cobra.Command)  {
	RootCmd.AddCommand(cmd)
}

func init() {
	RootCmd.AddCommand(generate.CrudCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
