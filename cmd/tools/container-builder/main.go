package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/namsral/flag"

	"compile-and-run-sandbox/internal/sandbox"
)

var ENDCOLOR = "\033[0m"
var RED = "\033[31m"
var GREEN = "\033[32m"

func main() {
	if runtime.GOOS == "windows" {
		RED = ""
		ENDCOLOR = ""
		GREEN = ""
	}

	var (
		filterName string
		verbose    bool
	)

	flag.StringVar(&filterName, "clang", "", "")
	flag.BoolVar(&verbose, "v", false, "")

	flag.Parse()

	if strings.TrimSpace(filterName) != "" {
		c, ok := sandbox.Compilers[filterName]

		if !ok {
			log.Fatalf("language '%s' does not exist in supported compilers\n", filterName)
		}

		runDockerCommand(c, verbose)
		return
	}

	completed := map[string]bool{}

	for _, c := range sandbox.Compilers {
		if _, ok := completed[c.VirtualMachineName]; ok {
			continue
		}

		runDockerCommand(c, verbose)
	}
}

func runDockerCommand(compiler *sandbox.LanguageCompiler, verbose bool) {
	fmt.Printf("%sRunning language:%s %s%s%s\n", RED, ENDCOLOR, GREEN, compiler.Language, ENDCOLOR)

	path := fmt.Sprintf("./build/dockerfiles/%s.dockerfile", strings.ToLower(compiler.Dockerfile))
	cmd := exec.Command("docker", "build", "-f", path, "-t", compiler.VirtualMachineName)

	cmd.Stdout = nil
	cmd.Stderr = os.Stderr

	if verbose {
		cmd.Args = append(cmd.Args, "--progress=plain")
		cmd.Stdout = os.Stdout
	}

	cmd.Args = append(cmd.Args, ".")

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%sFinished language:%s %s%s%s\n", RED, ENDCOLOR, GREEN, compiler.Language, ENDCOLOR)
}
