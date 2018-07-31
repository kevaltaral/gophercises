package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime/debug"
	"strconv"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func handler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/", sourceCodeHandler)
	mux.HandleFunc("/panic/", panicDemo)
	//mux.HandleFunc("/panic-after/", panicAfterDemo)
	//mux.HandleFunc("/", hello)
	return mux
}

func main() {

	log.Fatal(http.ListenAndServe(":3000", devMw(handler())))
}

func sourceCodeHandler(w http.ResponseWriter, r *http.Request) {

	path := r.FormValue("path")
	lineStr := r.FormValue("line")
	line, err := strconv.Atoi(lineStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b := bytes.NewBuffer(nil)
	io.Copy(b, file)
	var lines [][2]int
	if line > 0 {
		lines = append(lines, [2]int{line, line})
	}
	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, b.String())
	style := styles.Get("github")
	formatter := html.New(html.TabWidth(2), html.WithLineNumbers(), html.LineNumbersInTable(), html.HighlightLines(lines))
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<style>pre { font-size: 1.2em; }</style>")
	formatter.Format(w, style, iterator)

}

func devMw(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))

				fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, makeLinks(string(stack)))
			}
		}()

		app.ServeHTTP(w, r)
	}
}
func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func makeLinks(stack string) string {
	re := regexp.MustCompile("\t.*:[0-9]*")
	lines := re.FindAllString(stack, -1)

	re = regexp.MustCompile(":")
	for _, line := range lines {
		splits := re.Split(line, -1)
		//fmt.Printf("\npath=%s, line=%s", splits[0], splits[1])
		link := "<a href='/debug?path=" + splits[0] + "&line=" + splits[1] + "'>" + line + "</a>"
		reg := regexp.MustCompile(line)
		stack = reg.ReplaceAllString(stack, link)
	}

	return stack
}
