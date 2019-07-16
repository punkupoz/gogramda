package schema

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestSchema(t *testing.T) {
	tempDir, err := ioutil.TempDir(".", "test")
	if err != nil {
		t.Error(err)
	}

	err = os.MkdirAll(filepath.Join(tempDir, "type"), 0777)
	if err != nil {
		t.Error(err)
	}

	ioutil.WriteFile(filepath.Join(tempDir, "schema.graphql"), []byte(`schema {
	query: Query
}`), 0777)

	ioutil.WriteFile(filepath.Join(tempDir, "query.graphql"), []byte(`
type Query {
	person(id: ID!): Person
}`), 0777)

	err = ioutil.WriteFile(filepath.Join(tempDir, "type", "query.graphql"), []byte(`type Person{
	id: ID!
	firstName: String!
	lastName: String
}`), 0777)
	if err != nil {
		t.Error(err)
	}

	expected := `
type Query {
	person(id: ID!): Person
}
schema {
	query: Query
}
type Person{
	id: ID!
	firstName: String!
	lastName: String
}
`

	if String(tempDir) != expected {
		t.Error("mismatch")
	}
	os.RemoveAll(tempDir)
}
