package make

import (
	"fmt"
	"github.com/liqian-spec/practice/pkg/console"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var version string

var CmdMakeAPIController = &cobra.Command{
	Use:     "apicontroller",
	Short:   "Create api controller",
	Run:     runMakeAPIController,
	Args:    cobra.ExactArgs(1),
	Example: "make apicontroller user --v=2",
}

func init() {
	CmdMakeAPIController.PersistentFlags().StringVarP(&version, "version", "v", "", "version")
}

func runMakeAPIController(command *cobra.Command, args []string) {

	dirs := strings.Split(args[0], "/")
	if len(dirs) > 3 {
		console.Exit("The directory supports up to three levels only, And the last name is the file name")
	}

	version = fmt.Sprintf("%s", version)

	dirPath := fmt.Sprintf("app/http/controllers/api/%s", version)
	if len(dirs) > 1 {
		for i := 0; i < len(dirs)-1; i++ {
			dirPath += dirs[i] + "/"
		}
	}

	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		console.ExitIf(err)
	}

	model := makeModelFromString(dirs[len(dirs)-1])

	fileName := fmt.Sprintf("%s_contoller.go", model.PackageName)

	createFileFromStub(dirPath+fileName, "apicontroller", model, map[string]string{"{{version}}": version})
}
