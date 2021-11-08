package commands

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type fileServerFileSystem struct {
	http.FileSystem
	isSinglePageApp bool
}

func (fs fileServerFileSystem) Open(name string) (http.File, error) {
	file, err := fs.FileSystem.Open(name)

	if os.IsNotExist(err) && filepath.Ext(name) == "" {
		if fs.isSinglePageApp {
			return fs.FileSystem.Open("/index.html")
		}

		return fs.FileSystem.Open(name + ".html")
	}

	return file, err
}

var serveCommand = &cobra.Command{
	Use:   "serve [directory]",
	Short: "Start a static file server for the current or provided directory",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		root := "."

		if len(args) != 0 {
			root = args[0]

			if _, err := os.Stat(root); os.IsNotExist(err) {
				fmt.Println("Error: provided directory could not be found")
				os.Exit(1)
			}
		}

		port, _ := cmd.Flags().GetString("port")
		isSinglePageApp, _ := cmd.Flags().GetBool("single")

		fileServer := http.FileServer(fileServerFileSystem{
			http.Dir(root),
			isSinglePageApp,
		})

		fmt.Printf("\n\033[1;32m•\033[0m http://localhost:%s\n\n", port)
		log.Fatal(http.ListenAndServe("localhost:"+port, fileServer))
	},
}

func init() {
	serveCommand.Flags().StringP("port", "p", "8080", "listen on specified port")
	serveCommand.Flags().BoolP("single", "s", false, "serve as single-page application")
}
