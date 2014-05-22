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
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
	"strings"
)

type Property struct {
	file  string
	items map[string]interface{}
}

//NewProperties create new Properties
func NewProperties(file string) (*Property, error) {
	if len(file) == 0 {
		return nil, errors.New("Filepath cam't be empty")
	}
	p := new(Property)
	p.file = file
	return p, nil
}

//SetFilepath set the properties filepath
func (p *Property) SetFilepath(file string) *Property {
	if len(file) != 0 {
		p.file = file
	}
	return p
}

//GetFilepath get the properties filepath
func (p *Property) GetFilepath() string {
	return p.file
}

//GetProperties get all the properties
func (p *Property) GetProperties() map[string]interface{} {
	if p.items != nil {
		return p.items
	} else {
		return nil
	}
}

//ReadProperties read setting form the properties file
func (p *Property) ReadProperties() (*Property, error) {

	f, err := os.Open(p.file)

	defer f.Close()
	if err != nil {
		return nil, err
	}

	reg := regexp.MustCompile(`^#.*`)
	p.items = make(map[string]interface{})

	r := bufio.NewReader(f)
	for {
		s, _, err := r.ReadLine()
		if err != nil {
			if io.EOF == err {
				break
			} else {
				return nil, errors.New("Read the properties error " + err.Error())
			}
		}
		c := string(s)
		if 0 == len(c) || "\r\n" == c || reg.FindAllString(c, -1) != nil {
			continue
		}
		items := strings.Split(strings.TrimSpace(c), "=")
		p.items[items[0]] = items[1]
	}
	return p, nil
}

//Get get the one setting from properties file
func (p *Property) Get(key string) interface{} {
	if len(key) != 0 {
		return p.items[key]
	}
	return nil
}
