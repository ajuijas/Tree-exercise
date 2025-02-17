# Tree Exercise

This is a learning project from the [One2N Go Bootcamp](https://one2n.io/go-bootcamp/go-projects/tree-in-go). It's a simple program that shows folders and files in a tree-like structure, similar to the `tree` command in Unix.

## What Can It Do?

This program can:
- Show all your files and folders in a nice tree structure
- Show the full path of each file if you want
- Show file permissions
- Show only folders (skip files)
- Let you choose how deep to show the folder structure
- Count how many files and folders it found

## How to Run It

First, make sure you have Go installed on your computer. Then:

1. Get the code:
```bash
git clone https://github.com/yourusername/Tree-exercise.git
```

2. Go to the project folder:
```bash
cd Tree-exercise
```

3. Build the program:
```bash
go build -o tree ./tree
```

## How to Use It

Here are some basic commands you can try:

```bash
# Show files and folders in current directory
./tree

# Show full paths
./tree --relative

# Show permissions
./tree --permission

# Show only folders
./tree --directory-only

# Only show folders up to a certain depth
./tree --level 2    # This will show only 2 levels deep
```

## For Developers

This project is built with:
- The [Cobra](https://github.com/spf13/cobra) package to handle commands
- Basic Go file operations

## Testing

To run the tests:
```bash
go test ./...
```

## License

Check the LICENSE file for license details.
