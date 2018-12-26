package main

import (
	"bytes"
	"flag"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const filename = "/tmp/fperf_main.go"

type option struct {
	output  string
	verbose bool
}

func gobuild(o *option, imports []string) error {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("package main\n")
	buf.WriteString(`import "github.com/fperf/fperf"` + "\n")
	for _, imp := range imports {
		buf.WriteString(`import _ "` + imp + `"` + "\n")
	}
	buf.WriteString(`
	func main() {
		fperf.Main()
	}
	`)

	if err := ioutil.WriteFile(filename, buf.Bytes(), 0655); err != nil {
		log.Fatalln(err)
	}
	defer os.Remove(filename)

	args := []string{"build", "-o", o.output, filename}
	if o.verbose {
		args = append(args, "-v")
	}
	cmd := exec.Command("go", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func main() {
	o := &option{}
	flag.StringVar(&o.output, "o", "fperf", "build output")
	flag.BoolVar(&o.verbose, "v", false, "print the names of packages as they are compiled.")
	flag.Parse()

	paths := flag.Args()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	if len(paths) == 0 {
		paths = append(paths, cwd)
	}

	imports := make([]string, len(paths))
	for i := range paths {
		path := paths[i]
		if filepath.IsAbs(path) {
			var err error
			path, err = filepath.Rel(cwd, path)
			if err != nil {
				log.Fatalln(err)
			}
		}
		p, err := build.Import(path, cwd, 0)
		if err != nil {
			log.Fatalln(err)
		}
		imports[i] = p.ImportPath
	}
	if err := gobuild(o, imports); err != nil {
		log.Fatalln(err)
	}
}
