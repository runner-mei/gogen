package gengen

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/aryann/difflib"
)

func getGogen() string {
	for _, pa := range filepath.SplitList(os.Getenv("GOPATH")) {
		dir := filepath.Join(pa, "src/github.com/runner-mei/gogen")
		if st, err := os.Stat(dir); err == nil && st.IsDir() {
			return dir
		}
	}
	return ""
}

func TestGenerate(t *testing.T) {
	wd := getGogen()

	t.Run("gogen", func(t *testing.T) {
		for _, name := range []string{"test"} {
			t.Log("=====================", name)
			os.Remove(filepath.Join(wd, "gentest", name+".gogen.go"))
			// fmt.Println(filepath.Join(wd, "gentest", name+".gobatis.go"))

			var gen = WebServerGenerator{
				GeneratorBase: GeneratorBase{
					//ext:      ".echogen.go",
					buildTag: "echo",
				},
				config: "@echo",
			}
			if err := gen.Run([]string{filepath.Join(wd, "gentest", name+".go")}); err != nil {
				fmt.Println(err)
				t.Error(err)
				continue
			}

			actual := readFile(filepath.Join(wd, "gentest", name+".gogen.go"))
			excepted := readFile(filepath.Join(wd, "gentest", name+".gogen.txt"))
			if !reflect.DeepEqual(actual, excepted) {
				results := difflib.Diff(excepted, actual)
				for _, result := range results {
					if result.Delta == difflib.Common {
						continue
					}
					t.Error(result)
				}
			}
		}
	})

	t.Run("beegen", func(t *testing.T) {
		for _, name := range []string{"test"} {
			t.Log("=====================", name)
			os.Remove(filepath.Join(wd, "gentest", name+".beegen.go"))
			// fmt.Println(filepath.Join(wd, "gentest", name+".gobatis.go"))

			var gen = WebServerGenerator{
				GeneratorBase: GeneratorBase{
					ext:      ".beegen.go",
					buildTag: "beego",
				},
				config: "@beego",
			}
			if err := gen.Run([]string{filepath.Join(wd, "gentest", name+".go")}); err != nil {
				fmt.Println(err)
				t.Error(err)
				continue
			}

			actual := readFile(filepath.Join(wd, "gentest", name+".beegen.go"))
			excepted := readFile(filepath.Join(wd, "gentest", name+".beegen.txt"))
			if !reflect.DeepEqual(actual, excepted) {
				results := difflib.Diff(excepted, actual)
				for _, result := range results {
					if result.Delta == difflib.Common {
						continue
					}
					t.Error(result)
				}
			}
		}
	})

	t.Run("gingen", func(t *testing.T) {
		for _, name := range []string{"test"} {
			t.Log("=====================", name)
			os.Remove(filepath.Join(wd, "gentest", name+".gingen.go"))
			// fmt.Println(filepath.Join(wd, "gentest", name+".gobatis.go"))

			var gen = WebServerGenerator{
				GeneratorBase: GeneratorBase{
					ext:      ".gingen.go",
					buildTag: "gin",
				},
				config: "@gin",
			}
			if err := gen.Run([]string{filepath.Join(wd, "gentest", name+".go")}); err != nil {
				fmt.Println(err)
				t.Error(err)
				continue
			}

			actual := readFile(filepath.Join(wd, "gentest", name+".gingen.go"))
			excepted := readFile(filepath.Join(wd, "gentest", name+".gingen.txt"))
			if !reflect.DeepEqual(actual, excepted) {
				results := difflib.Diff(excepted, actual)
				for _, result := range results {
					if result.Delta == difflib.Common {
						continue
					}
					t.Error(result)
				}
			}
		}
	})

	t.Run("loong", func(t *testing.T) {
		for _, name := range []string{"test"} {
			t.Log("=====================", name)
			os.Remove(filepath.Join(wd, "gentest", name+".loonggen.go"))
			// fmt.Println(filepath.Join(wd, "gentest", name+".gobatis.go"))

			var gen = WebServerGenerator{
				GeneratorBase: GeneratorBase{
					ext:      ".loonggen.go",
					buildTag: "loong",
				},
				config: "@loong",
			}
			if err := gen.Run([]string{filepath.Join(wd, "gentest", name+".go")}); err != nil {
				fmt.Println(err)
				t.Error(err)
				continue
			}

			actual := readFile(filepath.Join(wd, "gentest", name+".loonggen.go"))
			excepted := readFile(filepath.Join(wd, "gentest", name+".loonggen.txt"))
			if !reflect.DeepEqual(actual, excepted) {
				results := difflib.Diff(excepted, actual)
				for _, result := range results {
					if result.Delta == difflib.Common {
						continue
					}
					t.Error(result)
				}
			}
		}
	})

	t.Run("client", func(t *testing.T) {
		for _, name := range []string{"test"} {
			t.Log("=====================", name)
			os.Remove(filepath.Join(wd, "gentest", name+".clientgen.go"))
			// fmt.Println(filepath.Join(wd, "gentest", name+".clientgen.go"))

			var gen = WebClientGenerator{
				GeneratorBase: GeneratorBase{
					ext: ".clientgen.go",
				},
				//config: "@echo",
			}
			gen.Flags(flag.NewFlagSet("", flag.PanicOnError)).Parse([]string{})
			gen.config.HasWrapper = false

			if err := gen.Run([]string{filepath.Join(wd, "gentest", name+".go")}); err != nil {
				fmt.Println(err)
				t.Error(err)
				continue
			}

			actual := readFile(filepath.Join(wd, "gentest", name+".clientgen.go"))
			excepted := readFile(filepath.Join(wd, "gentest", name+".clientgen.txt"))
			if !reflect.DeepEqual(actual, excepted) {
				results := difflib.Diff(excepted, actual)
				for _, result := range results {
					if result.Delta == difflib.Common {
						continue
					}
					t.Error(result)
				}
			}
		}
	})

	t.Run("loongclient", func(t *testing.T) {
		for _, name := range []string{"test"} {
			t.Log("=====================", name)
			os.Remove(filepath.Join(wd, "gentest", name+".loongclientgen.go"))
			// fmt.Println(filepath.Join(wd, "gentest", name+".clientgen.go"))

			var gen = WebClientGenerator{
				GeneratorBase: GeneratorBase{
					ext: ".loongclientgen.go",
				},
				//config: "@echo",
			}
			gen.Flags(flag.NewFlagSet("", flag.PanicOnError)).Parse([]string{
				"-has-wrapper", "true",
			})
			gen.config.HasWrapper = true

			if err := gen.Run([]string{filepath.Join(wd, "gentest", name+".go")}); err != nil {
				fmt.Println(err)
				t.Error(err)
				continue
			}

			actual := readFile(filepath.Join(wd, "gentest", name+".loongclientgen.go"))
			excepted := readFile(filepath.Join(wd, "gentest", name+".loongclientgen.txt"))
			if !reflect.DeepEqual(actual, excepted) {
				results := difflib.Diff(excepted, actual)
				for _, result := range results {
					if result.Delta == difflib.Common {
						continue
					}
					t.Error(result)
				}
			}
		}
	})

	// for _, name := range []string{"interface"} {
	// 	t.Log("===================== fail/", name)
	// 	os.Remove(filepath.Join(wd, "gentest", "fail", name+".gogen.go"))
	// 	// fmt.Println(filepath.Join(wd, "gentest", "fail", name+".gobatis.go"))

	// 	var gen = Generator{}
	// 	if err := gen.Run([]string{filepath.Join(wd, "gentest", "fail", name+".go")}); err != nil {
	// 		fmt.Println(err)
	// 		t.Error(err)
	// 		continue
	// 	}

	// 	actual := readFile(filepath.Join(wd, "gentest", "fail", name+".gogen.go"))
	// 	excepted := readFile(filepath.Join(wd, "gentest", "fail", name+".gogen.txt"))
	// 	if !reflect.DeepEqual(actual, excepted) {
	// 		results := difflib.Diff(excepted, actual)
	// 		for _, result := range results {
	// 			if result.Delta == difflib.Common {
	// 				continue
	// 			}

	// 			t.Error(result)
	// 		}
	// 	}
	// 	os.Remove(filepath.Join(wd, "gentest", "fail", name+".gogen.go"))
	// }
}

func readFile(pa string) []string {
	bs, err := ioutil.ReadFile(pa)
	if err != nil {
		panic(err)
	}

	return splitLines(bs)
}

func splitLines(txt []byte) []string {
	//r := bufio.NewReader(strings.NewReader(s))
	s := bufio.NewScanner(bytes.NewReader(txt))
	var ss []string
	for s.Scan() {
		ss = append(ss, s.Text())
	}
	return ss
}
