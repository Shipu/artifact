package generate

import (
	"log"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var MakeCommand = &cobra.Command{
	Use:  "make:command",
	Args: cobra.MinimumNArgs(2),
	RunE: makeCommand,
}

func makeCommand(cmd *cobra.Command, args []string) error {
	PackageName = args[0]
	name := args[1]
	crudStyle := args[2]

	PackageRoot = "art"

	hasDir, _ := afero.DirExists(afero.NewOsFs(), PackageRoot)

	if !hasDir {
		afero.NewOsFs().Mkdir(PackageRoot, 0755)
	}

	fs := afero.NewBasePathFs(afero.NewOsFs(), PackageRoot)

	createCMDFolders(fs, name)
	createCmdFiles(fs, name, crudStyle)

	log.Println("Successfully generated Command file for " + Title(name))

	return nil
}

func createCMDFolders(fs afero.Fs, name string) {
	fs.Mkdir(name, 0755)
	fs.Mkdir(name+"/commands", 0755)
}

func createCmdFiles(fs afero.Fs, name string, crudStyle string) {
	stubPath := "stubs"
	if crudStyle == "relational" {
		stubPath = "stubs/relational"
	}
	createFile(fs, name, stubPath+"/cmd.stub", name+"/commands/"+name+"Command.go")
}
