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

package sp

import (
	"path/filepath"
	"testing"

	"github.com/descope-dev/lite-idp/idp"
	"github.com/spf13/viper"
)

func Test_serviceProvider_GetRedirect(t *testing.T) {
	viper.Set("tls-certificate", filepath.Join("testdata", "certificate.pem"))
	viper.Set("tls-private-key", filepath.Join("testdata", "key.pem"))
	tlsConfigClient, err := idp.ConfigureTLS()
	if err != nil {
		t.Fatal(err)
	}
	serviceProvider, err := New(Configuration{
		EntityID:                    "https://www.jw.dev.gfclab.com/user",
		AssertionConsumerServiceURL: "http://test",
		TLSConfig:                   tlsConfigClient,
	})
	_, err = serviceProvider.GetRedirect([]byte("mystate"))
	if err != nil {
		t.Fatal(err)
	}
}
