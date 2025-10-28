package automation

import "log"
import "os"
import "fmt"
import "path/filepath"
import "strings"
import "github.com/denormal/go-gitignore"

var ignores []*gitignore.GitIgnore = []*gitignore.GitIgnore{}

// write a function that checks if path is ignored by any of the ignore clients


func ListFilesWithPath(workspace string) []string {
	var filesWithPath []string

	err := filepath.Walk(workspace, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(path, ".gitignore") {

			ignore, err := gitignore.NewRepository(
				strings.Replace(path, ".gitignore", "", -1),
			)
			ignores = append(
				ignores,
				&ignore,
			)

			if err != nil {
				log.Println(err)
			}

		}

		if !info.IsDir() {

			// call the function here
			log.Println(path + info.Name())
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	log.Println("Ignored patterns:", filesWithPath)
	return filesWithPath
}
