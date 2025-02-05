package cmd

import (
	"bytes"
	"os"
	"testing"

	"golang.org/x/tools/go/expect"
)

func captureStandardOutput() (func () string) {
	oldStdout := os.Stdout
	oldStderr := os.Stderr

	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut
	os.Stderr = wOut

	return func() string {
		wOut.Close()

		var buf bytes.Buffer
		_, _ = buf.ReadFrom(rOut)

		rOut.Close()

		os.Stdout = oldStdout
		os.Stderr = oldStderr

		return buf.String()
	}
}

// Function to create temporary directories and files to test the tree structure
func createTempDirsAndFiles() {
	// Create temporary directories and files
	_ = os.MkdirAll("test_dir", 0755)
	_ = os.WriteFile("test_dir/test_file", []byte("test content"), 0644)
	_ = os.MkdirAll("test_dir/test_dir_lvl_1", 0755)
	_ = os.WriteFile("test_dir/test_dir_lvl_1/test_file", []byte("test content"), 0644)
	_ = os.MkdirAll("test_dir/test_dir_lvl_1/test_dir_lvl_2", 0755)
	_ = os.WriteFile("test_dir/test_dir_lvl_1/test_dir_lvl_2/test_file", []byte("test content"), 0644)
	_ = os.MkdirAll("test_dir/test_dir_lvl_1/test_dir_lvl_2/test_dir_lvl_3", 0755)
	_ = os.WriteFile("test_dir/test_dir_lvl_1/test_dir_lvl_2/test_dir_lvl_3/test_file", []byte("test content"), 0644)
	_ = os.MkdirAll("test_dir/test_dir_lvl_1/test_dir_lvl_2/test_dir_lvl_3/test_dir_lvl_4", 0755)
	_ = os.MkdirAll("test_dir/test_dir_lvl_1/test_dir_lvl_2/test_dir_lvl_3/test_dir_lvl_4_1", 0755)
	_ = os.MkdirAll("test_dir/test_dir_lvl_1/test_dir_lvl_2/test_dir_lvl_3/test_dir_lvl_4_2", 0755)
	_ = os.MkdirAll("test_dir/test_dir_lvl_1/test_dir_lvl_2/test_dir_lvl_3/test_dir_lvl_4_3", 0755)
}


func Test_print_tree_structure(t *testing.T){

	createTempDirsAndFiles()

	tests := []struct {
		args []string
		expectedOutput string
	}{
		{
			args: []string{"test_dir"},
			expectedOutput: `test_dir
├── test_dir_lvl_1
│   ├── test_dir_lvl_2
│   │   ├── test_dir_lvl_3
│   │   │   ├── test_file
│   │   │   ├── test_dir_lvl_4
│   │   │   ├── test_dir_lvl_4_1
│   │   │   ├── test_dir_lvl_4_2
│   │   │   └── test_dir_lvl_4_3
│   │   └── test_file
│   └── test_file
└── test_file

7 directories, 4 files
`,
		},
		{
			args: []string{"-f", "test_dir"},
			expectedOutput: `test_dir
├── test_dir/test_dir_lvl_1
│   ├── test_dir/test_dir_lvl_1/test_dir_lvl_2
│   │   ├── test_dir/test_dir_lvl_1/test_dir_lvl_2/test_dir_lvl_3
│   │   │   ├── test_dir/test_dir_lvl_1/test_dir_lvl_2/test_dir_lvl_3/test_file
│   │   │   ├── test_dir/test_dir_lvl_1/test_dir_lvl_2/test_dir_lvl_3/test_dir_lvl_4
│   │   │   ├── test_dir/test_dir_lvl_1/test_dir_lvl_2/test_dir_lvl_3/test_dir_lvl_4_1
│   │   │   ├── test_dir/test_dir_lvl_1/test_dir_lvl_2/test_dir_lvl_3/test_dir_lvl_4_2
│   │   │   └── test_dir/test_dir_lvl_1/test_dir_lvl_2/test_dir_lvl_3/test_dir_lvl_4_3
│   │   └── test_dir/test_dir_lvl_1/test_dir_lvl_2/test_file
│   └── test_dir/test_dir_lvl_1/test_file
└── test_dir/test_file

7 directories, 4 files
`,
		},
	}

	for _, test := range tests {
		capture := captureStandardOutput()

		rootCmd.SetArgs(test.args)
		_ = rootCmd.Execute()
		
		stdout := capture()

		expect.Equal(t, stdout, test.expectedOutput)
	}
}