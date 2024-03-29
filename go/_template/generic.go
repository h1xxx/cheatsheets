package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"os/user"
	str "strings"
	"unicode/utf8"
)

func main() {
	isText("fs.md")
	u, _ := user.Current()
	fmt.Printf("%+v\n\n", u)

	parseEnv()
}


func parseEnv() {
	s := `AXXX='zxc' BB="xx zz" CC="musl-gcc -lulz -static" CLFAGS=-O2`
	env := getIniEnv(s)

	for _, e := range env {
	fmt.Printf("%+v\n", e)
	}
}

func getIniEnv(s string) []string {
        var env []string
        var v string
        var isQuoted bool

        fields := str.Split(s, " ")
        for _, f := range fields {
                if str.Contains(f, "=\"") || str.Contains(f, "='") {
                        isQuoted = true
                        v = f
			if str.HasSuffix(f, "\"") || str.HasSuffix(f, "'") {
				env = append(env, f)
				isQuoted = false
			}
			continue
		}

                if !isQuoted {
                        env = append(env, f)
			continue
		}

                if str.HasSuffix(f, "\"") || str.HasSuffix(f, "'") {
                        isQuoted = false
                        v += " " + f
			env = append(env, v)
			v = ""
                } else {
                        v += " " + f
                }
        }
	return env
}

func isText(file string) {
	fmt.Println("****** "+file)
	fd, _ := os.Open(file)
	buff := make([]byte, 512)
	_, _ = fd.Read(buff)
	fmt.Println(utf8.Valid(buff))
}

func walkDir(rootDir string) ([]string, error) {
	var names []string
	err := filepath.Walk(rootDir,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				names = append(names, path)
			}
			return nil
		})
	return names, err
}

func fileExists(fPath string) bool {
	p, err := os.Stat(fPath)
	fmt.Printf("%+v\n", p.IsDir())
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func Cp(src, dest string) {
        bb := "/bin/busybox"
        _ = bb + " ls -l " + "/* " + "/home/*"

        cmd := exec.Command(bb, "sh", "-c", bb , " ls -l " + "/* " + "/home/*")
        out, err := cmd.Output()
	fmt.Println(string(out), err)
}
