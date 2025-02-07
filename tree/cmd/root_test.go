package cmd

import (
	"bytes"
	"os"
	"testing"
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
func createTempDirsAndFiles() (func ()) {
	// Create temporary directories and files
	_ = os.MkdirAll("src", 0755)
	_ = os.MkdirAll("src/main", 0755)
	_ = os.MkdirAll("src/main/java", 0755)
	_ = os.MkdirAll("src/main/java/in", 0755)
	_ = os.MkdirAll("src/main/java/in/one2n", 0755)
	_ = os.MkdirAll("src/main/java/in/one2n/exercise", 0755)
	_ = os.WriteFile("src/main/java/in/one2n/exercise/Grade.java", []byte("test content"), 0644)
	_ = os.WriteFile("src/main/java/in/one2n/exercise/Grader.java", []byte("test content"), 0644)
	_ = os.WriteFile("src/main/java/in/one2n/exercise/Student.java", []byte("test content"), 0644)
	_ = os.MkdirAll("src/main/resources", 0755)
	_ = os.MkdirAll("src/test", 0755)
	_ = os.MkdirAll("src/test/java", 0755)
	_ = os.MkdirAll("src/test/java/in", 0755)
	_ = os.MkdirAll("src/test/java/in/one2n", 0755)
	_ = os.MkdirAll("src/test/java/in/one2n/exercise", 0755)
	_ = os.WriteFile("src/test/java/in/one2n/exercise/GraderTest.java", []byte("test content"), 0644)
	_ = os.MkdirAll("src/test/resources", 0755)
	_ = os.WriteFile("src/test/resources/grades.csv", []byte("test content"), 0644)


	return func() {
		// _ = os.RemoveAll("src")
	}
}


func Test_print_tree_structure(t *testing.T){

	cleanup := createTempDirsAndFiles()
	defer cleanup()

	tests := []struct {
		args []string
		expectedOutput string
	}{
		{
			args: []string{"src"},
			expectedOutput: `src
├── main
│   ├── java
│   │   └── in
│   │       └── one2n
│   │           └── exercise
│   │               ├── Grade.java
│   │               ├── Grader.java
│   │               └── Student.java
│   └── resources
└── test
    ├── java
    │   └── in
    │       └── one2n
    │           └── exercise
    │               └── GraderTest.java
    └── resources
        └── grades.csv

12 directories, 5 files
`,
		},
		{
			args: []string{"-f", "src"},
			expectedOutput: `src
├── src/main
│   ├── src/main/java
│   │   └── src/main/java/in
│   │       └── src/main/java/in/one2n
│   │           └── src/main/java/in/one2n/exercise
│   │               ├── src/main/java/in/one2n/exercise/Grade.java
│   │               ├── src/main/java/in/one2n/exercise/Grader.java
│   │               └── src/main/java/in/one2n/exercise/Student.java
│   └── src/main/resources
└── src/test
    ├── src/test/java
    │   └── src/test/java/in
    │       └── src/test/java/in/one2n
    │           └── src/test/java/in/one2n/exercise
    │               └── src/test/java/in/one2n/exercise/GraderTest.java
    └── src/test/resources
        └── src/test/resources/grades.csv

12 directories, 5 files
`},
			{
			args: []string{"-d", "src"},
			expectedOutput: `src
├── main
│   ├── java
│   │   └── in
│   │       └── one2n
│   │           └── exercise
│   └── resources
└── test
    ├── java
    │   └── in
    │       └── one2n
    │           └── exercise
    └── resources

12 directories
`},
		{
			args: []string{"-L", "3", "src"},
			expectedOutput: `src
├── main
│   ├── java
│   │   └── in
│   └── resources
└── test
    ├── java
    │   └── in
    └── resources
        └── grades.csv

8 directories, 1 file
`},
		{
			args: []string{"-p", "src"},
			expectedOutput: `src
├── [-rwxr-xr-x]  main
│   ├── [-rwxr-xr-x]  java
│   │   └── [-rwxr-xr-x]  in
│   │       └── [-rwxr-xr-x]  one2n
│   │           └── [-rwxr-xr-x]  exercise
│   │               ├── [-rw-r--r--]  Grade.java
│   │               ├── [-rw-r--r--]  Grader.java
│   │               └── [-rw-r--r--]  Student.java
│   └── [-rwxr-xr-x]  resources
└── [-rwxr-xr-x]  test
    ├── [-rwxr-xr-x]  java
    │   └── [-rwxr-xr-x]  in
    │       └── [-rwxr-xr-x]  one2n
    │           └── [-rwxr-xr-x]  exercise
    │               └── [-rw-r--r--]  GraderTest.java
    └── [-rwxr-xr-x]  resources
        └── [-rw-r--r--]  grades.csv

12 directories, 5 files
`,},
		}

	for _, test := range tests {
		capture := captureStandardOutput()

		relativePath, directoryOnly, pemission = false, false, false
		directoriesCount, filesCount, nestedLevel = 0, 0, 0

		rootCmd.SetArgs(test.args)
		_ = rootCmd.Execute()
		
		stdout := capture()

		if stdout != test.expectedOutput {
			t.Errorf("Expected output: <<%s>>, but got: <<%s>>", test.expectedOutput, stdout)
		}
	}
}