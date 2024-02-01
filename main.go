package main

import (
	"encoding/json"
	"fmt"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"io"
	"os"
	"time"
)

type Change struct {
	DateTime  time.Time       `json:"date_time"`
	ChangedBy string          `json:"changed_by"`
	Patch     jsonpatch.Patch `json:"patch"`
}

type BaseDoc struct {
	Hello   string   `json:"hello"`
	Author  string   `json:"author"`
	Foo     string   `json:"foo"`
	Qux     *string  `json:"qux"`
	Names   []string `json:"names"`
	Changes []Change `json:"changes"`
}

func readJsonData(file string) []byte {
	data, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	binData, _ := io.ReadAll(data)

	return binData
}

func toDoc(data []byte) BaseDoc {
	var doc BaseDoc

	err := json.Unmarshal(data, &doc)

	if err != nil {
		panic(err)
	}

	return doc
}

func (b *BaseDoc) addChange(patch jsonpatch.Patch, changedBy string) {
	b.Changes = append(b.Changes, Change{
		Patch:     patch,
		ChangedBy: changedBy,
		DateTime:  time.Now(),
	})
}

func (b *BaseDoc) toJSON() []byte {
	data, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		panic(err)
	}

	return data
}

func main() {
	baseDoc := readJsonData("basedoc.json")
	patchDoc := readJsonData("patch.json")

	patch, err := jsonpatch.DecodePatch(patchDoc)
	if err != nil {
		panic(err)
	}

	modified, err := patch.Apply(baseDoc)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Original: %s\n", baseDoc)
	fmt.Printf("Patched: %s\n", modified)

	fmt.Println()

	// Saving the new doc
	base := toDoc(baseDoc)
	newDoc := toDoc(modified)

	newDoc.addChange(patch, "Nick")

	fmt.Printf("Original struct: %s\n", base.toJSON())
	fmt.Printf("New doc: %s\n", newDoc.toJSON())

}
