// Copyright 2014 Properties authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Properties is a package for read properties
package Properties

import (
	"log"
	"testing"
)

func TestProperties(t *testing.T) {
	p, err := NewProperties("./settings1.properties")
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(p.GetFilepath())
	p.SetFilepath("./settings.properties")
	log.Println(p.GetFilepath())

	log.Println(p.items)
	p.ReadProperties()
	log.Println(p.items)

	s := p.Get("string")
	log.Println(s)

	log.Println(p.GetProperties())
}
