// Copyright © 2016 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package npm

import (
	"encoding/json"
	"time"
)

type All map[string]*ShortPackageDefinition

type ShortPackageDefinition struct {
	Name string `json:"name,omitempty"`
	//Description string           `json:"description,omitempty"`
	//DistTags    *json.RawMessage `json:"dist-tags,omitempty"`
	//Maintainers *json.RawMessage `json:"maintainers,omitempty"`
	//Author      *json.RawMessage `json:"author,omitempty"`
	//Users       *json.RawMessage `json:"users,omitempty"`
	//Repository  struct {
	//	Type string `json:"type,omitempty"`
	//	URL  string `json:"url,omitempty"`
	//} `json:"repository,omitempty"`
	//Homepage       string           `json:"homepage,omitempty"`
	//Bugs           *json.RawMessage `json:"bugs,omitempty"`
	//ReadmeFilename string           `json:"readmeFilename,omitempty"`
	//Keywords       []string         `json:"keywords,omitempty"`
	//License        string           `json:"license,omitempty"`
	Time struct {
		Modified time.Time `json:"modified,omitempty"`
	} `json:"time,omitempty"`
	//Versions map[string]string `json:"versions,omitempty"`
	FullPackageDefinition FullPackageDefinition `json:"-"`
}

type PackageVersionDefinition struct {
	Name            string           `json:"name,omitempty"`
	Description     string           `json:"description,omitempty"`
	Version         string           `json:"version,omitempty"`
	Homepage        *json.RawMessage `json:"homepage,omitempty"`
	Repository      *json.RawMessage `json:"repository,omitempty"`
	Author          *json.RawMessage `json:"author,omitempty"`
	Tags            *json.RawMessage `json:"tags,omitempty"`
	Files           *json.RawMessage `json:"files,omitempty"`
	Bin             *json.RawMessage `json:"bin,omitempty"`
	Man             *json.RawMessage `json:"man,omitempty"`
	Dependencies    *json.RawMessage `json:"dependencies,omitempty"`
	DevDependencies *json.RawMessage `json:"devDependencies,omitempty"`
	Scripts         *json.RawMessage `json:"scripts,omitempty"`
	Engines         *json.RawMessage `json:"engines,omitempty"`
	License         *json.RawMessage `json:"license,omitempty"`
	GitHead         *json.RawMessage `json:"gitHead,omitempty"`
	Bugs            *json.RawMessage `json:"bugs,omitempty"`
	ID              *json.RawMessage `json:"_id,omitempty"`
	Shasum          *json.RawMessage `json:"_shasum,omitempty"`
	From            *json.RawMessage `json:"_from,omitempty"`
	NpmVersion      *json.RawMessage `json:"_npmVersion,omitempty"`
	NodeVersion     *json.RawMessage `json:"_nodeVersion,omitempty"`
	NpmUser         *json.RawMessage `json:"_npmUser,omitempty"`
	Maintainers     *json.RawMessage `json:"maintainers,omitempty"`
	Dist            struct {
		Shasum  string `json:"shasum,omitempty"`
		Tarball string `json:"tarball,omitempty"`
	} `json:"dist,omitempty"`
	NpmOperationalInternal *json.RawMessage `json:"_npmOperationalInternal,omitempty"`
	Directories            *json.RawMessage `json:"directories,omitempty"`
}

type FullPackageDefinition struct {
	ID             string                               `json:"_id,omitempty"`
	Rev            string                               `json:"_rev,omitempty"`
	Name           string                               `json:"name,omitempty"`
	Description    *json.RawMessage                     `json:"description,omitempty"`
	DistTags       *json.RawMessage                     `json:"dist-tags"`
	Versions       map[string]*PackageVersionDefinition `json:"versions,omitempty"`
	Readme         *json.RawMessage                     `json:"readme,omitempty"`
	Maintainers    *json.RawMessage                     `json:"maintainers,omitempty"`
	Time           *json.RawMessage                     `json:"time,omitempty"`
	Author         *json.RawMessage                     `json:"author,omitempty"`
	Repository     *json.RawMessage                     `json:"repository,omitempty"`
	Users          *json.RawMessage                     `json:"users,omitempty"`
	ReadmeFilename *json.RawMessage                     `json:"readmeFilename,omitempty"`
	Homepage       *json.RawMessage                     `json:"homepage,omitempty"`
	Bugs           *json.RawMessage                     `json:"bugs,omitempty"`
	License        *json.RawMessage                     `json:"license,omitempty"`
	Attachments    *json.RawMessage                     `json:"_attachments,omitempty"`
}
