// Copyright © 2016 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"testing"

	"net/http"

	"github.com/stretchr/testify/assert"
	"goji.io/pattern"
	"golang.org/x/net/context"
)

func mustReq(method, path string) (context.Context, *http.Request) {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		panic(err)
	}
	ctx := pattern.SetPath(context.Background(), req.URL.EscapedPath())

	return ctx, req
}

func Test_Composer_Pat_Archive(t *testing.T) {
	p := &GitPat{}

	c, r := mustReq("GET", "/git/github.com/kevinlebrun/colors.php/cb9b6666a2dfd9b6074b4a5caec7902fe3033578.zip")

	result := p.Match(c, r)

	assert.NotNil(t, result)
	assert.Equal(t, "github.com/kevinlebrun/colors.php", result.Value(pattern.Variable("path")))
	assert.Equal(t, "cb9b6666a2dfd9b6074b4a5caec7902fe3033578", result.Value(pattern.Variable("ref")))
	assert.Equal(t, "zip", result.Value(pattern.Variable("format")))
}
