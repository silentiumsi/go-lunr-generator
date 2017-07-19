package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path"

	"html"

	"encoding/json"

	"bufio"

	"github.com/ericaro/frontmatter"
)

type metadata struct {
	Title   string `yaml:"title"`
	URL     string `yaml:"url"`
	Slug    string `yaml:"slug"`
	Tags    string `yaml:"tags"`
	Content string `fm:"content" yaml:"-"`
}

type document struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Tags    string `json:"tags"`
	Content string `json:"content"`
}

func main() {
	inputPtr := flag.String("input", ".", "the directory of the content to index")
	outputPtr := flag.String("output", ".", "the output directory to store the generated index")
	flag.Parse()
	index(*inputPtr, *outputPtr)
}

func index(input string, output string) {
	files, err := ioutil.ReadDir(input)
	checkError(err)

	var documents []document

	for _, file := range files {
		fileBuffer, _ := ioutil.ReadFile(path.Join(input, file.Name()))
		fileContent := html.EscapeString(string(fileBuffer))

		if path.Ext(file.Name()) != ".md" && path.Ext(file.Name()) != ".html" {
			continue
		}

		v := new(metadata)
		frontMatterError := frontmatter.Unmarshal(([]byte)(fileContent), v)
		checkError(frontMatterError)

		document := createDocument(v)
		documents = append(documents, *document)

		documentsJSON, err := json.MarshalIndent(documents, "", "  ")
		checkError(err)

		outputDirErr := os.MkdirAll(output, os.ModePerm)
		checkError(outputDirErr)

		outputFile, createFileErr := os.Create(path.Join(output, "index.json"))
		checkError(createFileErr)
		defer outputFile.Close()

		w := bufio.NewWriter(outputFile)
		w.Write(documentsJSON)

		outputFile.Sync()

		w.Flush()
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func createDocument(metadata *metadata) *document {
	document := &document{
		ID:      metadata.Slug,
		Title:   metadata.Title,
		URL:     metadata.URL,
		Tags:    metadata.Tags,
		Content: metadata.Content,
	}
	return document
}
