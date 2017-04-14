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

	if len(prefix) > 0 {
		n = fmt.Sprintf("%s%s",
			strings.ToUpper(prefix),
			strings.ToUpper(genName()),
		)
	} else if len(suffix) > 0 {
		n = fmt.Sprintf("%s%s",
			strings.ToUpper(genName()),
			strings.ToUpper(suffix),
		)

	} else {
		n = fmt.Sprintf("%s%s",
			strings.ToUpper(genName()),
			strings.ToUpper(genName()),
		)
	}

	if toLower {
		n = strings.ToLower(n)
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
