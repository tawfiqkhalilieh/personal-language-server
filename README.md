### Personalized Language Server

I'm building this project in university nights to have some fun, and get used to using Go even more. It is gonna be slowly built over time as I get more free time.


This is a personalized language server that provides autocompletion suggestions based on your own codebase. It scans through all your projects, extracts commonly used patterns, and serves them via a local server for enhanced coding efficiency.

--- 

## To Start we need all your projects cloned locally inside one parent directory.

```bash
# Make sure you have the GitHub CLI installed and authenticated
# You can install it from https://cli.github.com/

gh repo list --limit 4000 | while read -r repo _; do
  gh repo clone "$repo" "$repo"
done

```

---

## after cloning all of the repos we want to recursively go into each directory, and read all the non-ignored files and store the data in the following format:

```json
{
    "python": {
        start_with_dot: ["split", "join", "replace" ],
        words: ["def", "print(", "as", "None", "self" ], 
    },
    "javascript": {
        start_with_dot: ["map", "filter", "reduce" ],
        words: ["function", "console.log(", "let", "const", "this" ], 
    }
}
```

---


## finally we want to run a local server that can respond to requests for autocompletions based on the above data.

```go

r.GET(
    "/complete",
    func complete(c *gin.Context) { 
        lang := c.Query("lang")
        prefix := c.Query("prefix")

        completions := getCompletions(lang, prefix) // Implement this function to fetch completions based on the data structure

        c.JSON(200, gin.H{
            "completions": completions,
        })
    },
);
```


