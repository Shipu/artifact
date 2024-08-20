package generate

import (
	"fmt"
	pluralize "github.com/gertd/go-pluralize"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"log"
	"path"
	"runtime"
	"strings"
)

var PackageName string

var PackageRoot string = "src"

var CrudCmd = &cobra.Command{
	Use:  "crud",
	Args: cobra.MinimumNArgs(2),
	RunE: crud,
}

func crud(cmd *cobra.Command, args []string) error {
	PackageName = args[0]
	name := args[1]
	crudStyle := args[2]

	hasDir, _ := afero.DirExists(afero.NewOsFs(), PackageRoot)
	if !hasDir {
		afero.NewOsFs().Mkdir(PackageRoot, 0755)
	}

	fs := afero.NewBasePathFs(afero.NewOsFs(), PackageRoot+"/")

	createFolders(fs, name)
	createFiles(fs, name, crudStyle)

	log.Println("Successfully generated CRUD for " + Title(name))

	return nil
}

func createFolders(fs afero.Fs, name string) {
	fs.Mkdir(name, 0755)
	fs.Mkdir(name+"/controllers", 0755)
	fs.Mkdir(name+"/models", 0755)
	fs.Mkdir(name+"/routes", 0755)
	fs.Mkdir(name+"/services", 0755)
	fs.Mkdir(name+"/dto", 0755)
}

func createFiles(fs afero.Fs, name string, crudStyle string) {
	stubPath := "stubs"
	if crudStyle == "relational" {
		stubPath = "stubs/relational"
	}
	createFile(fs, name, stubPath+"/controller.stub", name+"/controllers/"+name+"_controller.go")
	createFile(fs, name, stubPath+"/model.stub", name+"/models/"+name+".go")
	createFile(fs, name, stubPath+"/route.stub", name+"/routes/api.go")
	createFile(fs, name, stubPath+"/dto.stub", name+"/dto/"+name+"_dto.go")
	createFile(fs, name, stubPath+"/service.stub", name+"/services/"+name+"_service.go")
}

func createFile(fs afero.Fs, name string, stubPath, filePath string) {
	fs.Create(filePath)

	_, filename, _, _ := runtime.Caller(1)
	stubPath = path.Join(path.Dir(filename), stubPath)

	contents, _ := fileContents(stubPath)
	contents = replaceStub(contents, name)

	err := overwrite(PackageRoot+"/"+filePath, contents)
	if err != nil {
		fmt.Println(err)
	}
}

func fileContents(file string) (string, error) {
	a := afero.NewOsFs()
	contents, err := afero.ReadFile(a, file)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func overwrite(file string, message string) error {
	a := afero.NewOsFs()
	return afero.WriteFile(a, file, []byte(message), 0666)
}

func replaceStub(content string, name string) string {
	content = strings.Replace(content, "{{TitleName}}", Title(name), -1)
	content = strings.Replace(content, "{{PluralLowerName}}", Lower(Plural(name)), -1)
	content = strings.Replace(content, "{{SingularLowerName}}", Lower(Singular(name)), -1)
	content = strings.Replace(content, "{{PackageName}}", PackageName, -1)
	content = strings.Replace(content, "{{PackageRoot}}", PackageRoot, -1)
	return content
}

func Plural(name string) string {
	pluralize := pluralize.NewClient()

	return pluralize.Plural(name)
}

func Singular(name string) string {
	pluralize := pluralize.NewClient()
	return pluralize.Singular(name)
}

func Lower(name string) string {
	return strings.ToLower(name)
}

func Title(name string) string {
	return strings.Title(Lower(name))
}
