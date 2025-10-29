package data

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/exec"
)


func checkSupportedExtention(path string) string {
	var supported []string = []string{".py", ".js", ".lua", ".cs", ".c", ".go", ".java", ".rb", ".php", ".html", ".css", ".ts", ".rs", ".swift", ".kt", ".m", ".sh", ".pl", ".r", ".dart", ".scala", ".hs", ".erl", ".ex", ".exs", ".jl", ".groovy", ".vb", ".fs", ".fsi", ".fsx", ".lisp", ".clj", ".cljs", ".coffee", ".elm", ".nim", ".v", ".zig",};

	for _, ext := range supported {
		if len(path) >= len(ext) && path[len(path)-len(ext):] == ext {
			return ext 
		}
	}
	return ""

}


func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {

		return err == nil && !info.IsDir()
	}
}

func LoadData(c *gin.Context) {
	var paths = []string{}

	if !fileExists("scan.txt") {
		cmd := exec.Command("bash", "./myscript.sh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	file, err := os.Open("scan.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		paths = append(paths, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// close fi on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	log.Println(paths)
	for _, path := range paths {
		ext := checkSupportedExtention(path)
		// close fi on exit and check for its returned error
		if ext == "" {	
			continue
		}

		file, err := os.Open(path)
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			log.Printf("Read line: %s", line)
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		defer func() {
			if err := file.Close(); err != nil {
				panic(err)
			}
		}()

	}

	c.JSON(200, gin.H{"status": "data loaded"})
}
