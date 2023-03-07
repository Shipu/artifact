package generate

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"log"
)

var MakeCommand = &cobra.Command{
	Use:  "make:command",
	Args: cobra.MinimumNArgs(2),
	RunE: makeCommand,
}

func makeCommand(cmd *cobra.Command, args []string) error {
	PackageName = args[0]
	name := args[1]

	PackageRoot = "art"

	hasDir, _ := afero.DirExists(afero.NewOsFs(), PackageRoot)

	if !hasDir {
		afero.NewOsFs().Mkdir(PackageRoot, 0755)
	}

	fs := afero.NewBasePathFs(afero.NewOsFs(), PackageRoot)

	createCMDFolders(fs, name)
	createCmdFiles(fs, name)

	log.Println("Successfully generated Command file for " + Title(name))

	return nil
}

func createCMDFolders(fs afero.Fs, name string) {
	fs.Mkdir(name, 0755)
	fs.Mkdir(name+"/commands", 0755)
}

func createCmdFiles(fs afero.Fs, name string) {
	createFile(fs, name, "stubs/cmd.stub", name+"/commands/"+name+"Command.go")
}
