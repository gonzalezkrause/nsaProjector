/*
	Copyright 2016 HCN -josef[at]hackercat.ninja-

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	    http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var makeFolder, toLower bool
var prefix, suffix string

func init() {
	flag.BoolVar(&makeFolder, "f", makeFolder, "Creates a folder with the project name")
	flag.BoolVar(&toLower, "l", toLower, "Outputs the name in lowercase")
	flag.StringVar(&prefix, "p", "", "Prefix for the name 'banana' => 'BANANATURRET'")
	flag.StringVar(&suffix, "s", "", "Suffix for the name 'mantle' => 'SEALMANTLE'")
}

func main() {
	var n string

	flag.Parse()

	if prefix != "" {
		n = composeName(prefix, genName())
	} else if suffix != "" {
		n = composeName(genName(), suffix)
	} else {
		n = composeName(genName(), genName())
	}

	if makeFolder {
		fmt.Printf("Creating folder: '%s'\n", n)
		os.Mkdir(n, 0700)
	} else {
		fmt.Println(n)
	}
}

func genName() (a string) {
	s := len(nounlist)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a = strings.Replace(strings.Replace(nounlist[r.Intn(s)], " ", "", -1), "-", "", -1)

	return
}

func composeName(p, s string) (n string) {
	n = fmt.Sprintf("%s%s", p, s)

	if !toLower {
		n = strings.ToUpper(n)
	}

	return
}
