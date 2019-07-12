package schema

import (
	"path/filepath"
	"fmt"
	"os"
	"io/ioutil"
	"bytes"
	"regexp"
)

// Read all "*.graphql" files in schema directory

func String() string {
	buf := bytes.Buffer{}
	filepath.Walk("./schema", func (path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if r, _ := regexp.Compile(".+\\.graphql"); r.MatchString(info.Name()) && !info.IsDir() {
			b, _ := ioutil.ReadFile(path)
			buf.Write(b)
			buf.Write([]byte("\n"))
		}

		return nil
	})

	return buf.String()
}
