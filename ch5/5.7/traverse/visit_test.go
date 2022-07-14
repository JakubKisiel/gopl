package traverse_test

import (
	"bytes"
	"gopl/ch5/5.7/traverse"
	"log"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const testHTML = `<html><head><title>test</title></head><body>
	<a href="http://test"></a><!--comment--><div>text</div></body></html>`
const outputHTML = `<html>
  <head>
    <title>
      testxdddd
    </title>
  </head>
  <body>
    <a href="http://test"/>
    <!--comment-->
    <div>
      textxdddd
    </div>
  </body>
</html>
`

func TestVisitPrinting(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(testHTML))
	if err != nil {
		log.Fatalf("error :%v", err)
	}
	formatted := captureOutput(func() {
		depth := 0
		traverse.Visit(doc, &depth, traverse.BeforeElement, traverse.AfterElement)
	})
	if formatted != outputHTML {
		t.Fail()
	}
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetFlags(0)
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
