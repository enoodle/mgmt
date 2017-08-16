// Mgmt
// Copyright (C) 2013-2017+ James Shubin and the project contributors
// Written by James Shubin <james@shubin.ca> and the project contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package yamlgraph2 provides the facilities for loading a graph from a yaml file.
package scap

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/purpleidea/mgmt/pgraph"
	"github.com/purpleidea/mgmt/resources"

	errwrap "github.com/pkg/errors"
	//"gopkg.in/yaml.v2"
)

type ScapConfig struct {
	OS         string
	OutputFile string
	FedoraStr  string
	OtherOSStr string
}

func ParseConfigFromFile(filename string) *ScapConfig {
}
