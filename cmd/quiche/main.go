package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/hscells/quickumlsrest"
	"net/url"
	"os"
	"sort"
)

type args struct {
	N      int    `help:"maximum number of candidates to cache for each term"`
	Input  string `help:"input for terms to cache from QuickUMLS"`
	Output string `help:"output to create the file containing the cache"`
	Url    string `help:"URL for QuickUMLSRest" arg:"required"`
}

func (args) Version() string {
	return "04.Dec.2018"
}

func (args) Description() string {
	return "QuickUMLS Cache Extractor"
}

func main() {
	var (
		args   args
		err    error
		input  *os.File
		output *os.File
		q      quickumlsrest.Client
		terms  []string
	)
	p := arg.MustParse(&args)

	// Check if the URL is complete.
	if _, err := url.Parse(args.Url); err != nil {
		p.Fail(fmt.Sprintf("could not parse QuickUMLS URL: %s", err.Error()))
	} else {
		q = quickumlsrest.NewClient(args.Url)
	}

	// Check if the input file is correct.
	if len(args.Input) == 0 {
		input = os.Stdin
	} else {
		input, err = os.OpenFile(args.Input, os.O_RDONLY, os.ModePerm)
		if err != nil {
			p.Fail(err.Error())
		}
	}

	// Check if the output file is correct.
	if len(args.Output) == 0 {
		output = os.Stdout
	} else {
		output, err = os.OpenFile(args.Output, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			p.Fail(err.Error())
		}
	}

	fmt.Println("loading terms")
	s := bufio.NewScanner(input)
	for s.Scan() {
		fmt.Print(".")
		terms = append(terms, s.Text())
	}
	fmt.Printf("\n%d terms for extraction\n", len(terms))

	fmt.Println("mapping terms in QuickUMLS")
	cache := make(quickumlsrest.Cache)
	gob.Register(cache)
	for _, term := range terms {
		c, err := q.Match(term)
		if err != nil {
			panic(err)
		}
		sort.Slice(c, func(i, j int) bool {
			return c[i].Similarity < c[j].Similarity
		})
		for i, candidate := range c {
			if i > args.N {
				break
			}
			cache[term] = append(cache[term], candidate)
		}
		fmt.Print(".")
	}
	fmt.Println("\nmapped terms in QuickUMLS")

	enc := gob.NewEncoder(output)
	err = enc.Encode(cache)
	if err != nil {
		panic(err)
	}
}
