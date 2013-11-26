package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tmpl *template.Template
var types = []string{
	"int", "int8", "int16", "int32", "int64",
	"uint", "uint8", "uint16", "uint32", "uint64",
}

func init() {
	var err error
	tmpl, err = template.New("tmpl").Parse(`package incremental

import (
	"sync"
)

type {{.Upper}} struct {
	increment {{.Lower}}
	lock      sync.Mutex
}

// Next returns with an integer that is exactly one higher as the previous call to Next() for this {{.Upper}}
func (i *{{.Upper}}) Next() {{.Lower}} {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.increment++
	return i.increment
}

// Last returns the number ({{.Lower}}) that was returned by the most recent call to this instance's Next()
func (i *{{.Upper}}) Last() {{.Lower}} {
	return i.increment
}
`)
	if err != nil {
		log.Fatal(err)
	}
}

type data struct {
	Upper string
	Lower string
}

func main() {
	// loop over integer types
	for _, t := range types {
		// create file for type
		file, err := os.Create(t + ".go")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// create data with upper and lower names for the type
		d := &data{
			Upper: strings.ToUpper(t[0:1]) + t[1:],
			Lower: t,
		}

		// execute template, write directly to file
		err = tmpl.Execute(file, d)
		if err != nil {
			log.Fatal(err)
		}
	}

	// create test.go
	file, err := os.Create("int_test.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(`package incremental

import (
	"testing"
)

type user struct {
	i Int
}

func TestIntPtr(t *testing.T) {
	i := &Int{}
	num := i.Next()
	if num != 1 {
		t.Fatal("expected 1, got %d", num)
	}
	num = i.Next()
	if num != 2 {
		t.Fatal("expected 2, got %d", num)
	}
}

func TestAsField(t *testing.T) {
	u := user{}
	num := u.i.Next()
	if num != 1 {
		t.Fatal("expected 1, got %d", num)
	}
	num = u.i.Next()
	if num != 2 {
		t.Fatal("expected 2, got %d", num)
	}
}`)

	// create doc.go
	file, err = os.Create("doc.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(`
// package incremental provides concurency-safe incremental numbers.
//
// This package was created by a simple piece of code located in the gen subdirectory. Please modify that command if you want to modify this package.
package incremental`)
}
