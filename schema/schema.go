package schema

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

//String reads all "*.graphql" files in schema directory
func String(path string) (string, error) {
	buf := bytes.Buffer{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return nil
		}

		if r, _ := regexp.Compile(`.+\.graphql`); r.MatchString(info.Name()) && !info.IsDir() {
			b, _ := ioutil.ReadFile(path)
			buf.Write(b)
			buf.Write([]byte("\n"))
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
