// Copyright © 2017 Aaron Donovan <amdonov@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package idp

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/descope-dev/lite-idp/model"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewAttributeSource(t *testing.T) {
	in, err := os.Open(filepath.Join("testdata", "users.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(in)
	if err != nil {
		t.Fatal(err)
	}
	attSrc, err := NewAttributeSource()
	if err != nil {
		t.Fatal(err)
	}
	user := &model.User{Name: "john"}
	if err = attSrc.AddAttributes(user, nil); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(user.Attributes), "expected 3 attributes")
}
