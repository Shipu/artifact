package cmd

import (
	"github.com/shipu/artifact/cmd/generate"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "art",
	Short: "A simple artifact command",
	Long:  `A simple artifact command`,
	Run:   hello,
}

func hello(cmd *cobra.Command, args []string) {
	log.Println("art command")
}

func init() {
	rootCmd.AddCommand(generate.CrudCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
