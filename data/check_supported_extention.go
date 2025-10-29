package data

func checkSupportedExtention(path string) string {
	var supported []string = []string{".py", ".js", ".lua", ".cs", ".c", ".go", ".java", ".rb", ".php", ".html", ".css", ".ts", ".rs", ".swift", ".kt", ".m", ".sh", ".pl", ".r", ".dart", ".scala", ".hs", ".erl", ".ex", ".exs", ".jl", ".groovy", ".vb", ".fs", ".fsi", ".fsx", ".lisp", ".clj", ".cljs", ".coffee", ".elm", ".nim", ".v", ".zig"}

	for _, ext := range supported {
		if len(path) >= len(ext) && path[len(path)-len(ext):] == ext {
			return ext
		}
	}
	return ""

}

