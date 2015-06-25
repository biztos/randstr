// randstr - make a random quotable string.
// -------
// Copyright (c) 2015 Kevin Frost; MIT license; see end of this file.
//
// THIS IS A WORK IN PROGRESS!
// TODO: package-ize, test, etc.

package main

import (
	"crypto/rand"
	"fmt"
	"github.com/docopt/docopt-go"
	"math/big"
	"os"
	"strconv"
)

type printFormat struct {
	Continuing string
	Terminal   string
}

func main() {

	opts := getOpts()

	length, err := strconv.ParseInt(opts["<length>"].(string), 0, 32)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad length value; use -h for help.\n")
		os.Exit(1)

	}

	split64, err := strconv.ParseInt(opts["--split"].(string), 0, 8)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad split value; use -h for help.\n")
		os.Exit(1)
	}
	split := int(split64)

	lang := opts["--lang"].(string)
	if lang != "go" && lang != "perl" && lang != "js" {
		fmt.Fprintf(os.Stderr, "Unsupported langauge %s; use -h for help.\n",
			lang)
		os.Exit(2)
	}

	omit := map[int]bool{}
	if opts["--useall"] == false {
		switch lang {
		case "go":
			set_go_omits(omit)
		case "perl":
			set_perl_omits(omit)
		default:
			panic("unexpected lang: " + lang)
		}
		if opts["--white"] == false {
			omit[32] = true // only worried about space for now.
		}
	}

	pf := &printFormat{}
	switch lang {
	case "go":
		set_go_format(pf)
	case "perl":
		set_perl_format(pf)
	default:
		panic("unexpected lang: " + lang)
	}

	lines := []string{}
	curStr := ""
	for i := int64(1); i <= length; i++ {
		curStr = curStr + randChar(omit)
		if len(curStr) == split {
			lines = append(lines, curStr)
			curStr = ""
		}
	}
	if len(curStr) > 0 {
		lines = append(lines, curStr)
	}

	for idx, line := range lines {
		if opts["--noquote"] == true || opts["--useall"] == true {
			fmt.Println(line)
		} else {
			if idx == len(lines)-1 {
				fmt.Printf(pf.Terminal, line)
			} else {
				fmt.Printf(pf.Continuing, line)
			}
		}

	}

}

func set_go_format(pf *printFormat) {

	pf.Continuing = "`%s` +\n"
	pf.Terminal = "`%s`\n"

}

func set_perl_format(pf *printFormat) {

	pf.Continuing = "'%s' .\n"
	pf.Terminal = "'%s';\n"

}

func set_js_format(pf *printFormat) {

	pf.Continuing = "'%s' +\n"
	pf.Terminal = "'%s';\n"

}

func set_go_omits(omits map[int]bool) {

	omits[34] = true // double-quote
	omits[92] = true // backslash
	omits[96] = true // backtick)

}

func set_perl_omits(omits map[int]bool) {

	omits[34] = true // double-quote
	omits[36] = true // dollar sign
	omits[37] = true // percent sign
	omits[39] = true // single-quote
	omits[64] = true // at sign
	omits[92] = true // backslash

}

func set_js_omits(omits map[int]bool) {

	omits[34] = true // double-quote
	omits[39] = true // single-quote
	omits[92] = true // backslash

}

func randChar(omit map[int]bool) string {

	for {
		r, _ := rand.Int(rand.Reader, big.NewInt(93))
		n := int(r.Int64()) + 32
		if !omit[n] {
			return string(n)
		}
	}

}

func getOpts() map[string]interface{} {

	usage := `randstr - make a random quotable string.

Usage:
  randstr [options] <length>
  randstr -h | --help
  randstr -v | --version

Options:
  -l --lang=<lang>      Target language [default: go].
  -s --split=<width>    Split at this width [default: 60].
  -n --noquote          Do not quote the output.
  -w --white            Allow whitespace.
  -a --useall           Use all printable characters (implies -n).
  -h --help             Show this screen.
  -v --version          Show version.

This program generates a string of "cryptographically secure" random bytes
within the ASCII displayable range of 32-126, excluding characters that could
make quoting difficult in the target language.

Supported target languages include:

* go - Google's Go language, aka Golang (in which this utility is written).
* perl - The Perl language, in all its glory!
* js - Javascript, used all over the Interwebs.

When in doubt, use perl as its set is more conservative in order to allow
interpolation.

Note that the "cryptographically secure" claim is simply echoing Go's
crypto/rand package, which you may or may not choose to trust.

Copyright (c) 2015 Kevin Alan Frost.  License: MIT (see source code).`

	opts, _ := docopt.Parse(usage, nil, true, "randstr 1.0", false)

	return opts
}

/*

LICENSE
-------

Please note that any binary distribution is likely subject to the license
terms of Go itself and of the package github.com/docopt/docopt-go.

The MIT License (MIT)

Copyright (c) 2015 Kevin Alan Frost

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
