// Copyright 2022 The prometheus-operator Authors
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

package validation

import (
	"encoding/json"
	"fmt"

	"github.com/prometheus/alertmanager/config"
)

// ValidateURL against the config.URL
// This could potentially become a regex and be validated via OpenAPI
// but right now, since we know we need to unmarshal into an upstream type
// after conversion, we validate we don't error when doing so.
func ValidateURL(url string) (*config.URL, error) {
	var u config.URL
	err := json.Unmarshal([]byte(fmt.Sprintf(`"%s"`, url)), &u)
	if err != nil {
		return nil, fmt.Errorf("validate url from string failed for %s: %w", url, err)
	}
	return &u, nil
}

// ValidateSecretURL against config.URL
// This is for URLs which are retrieved from secrets and should not
// logged as part of the err.
func ValidateSecretURL(url string) error {
	var u config.SecretURL

	err := u.UnmarshalJSON([]byte(fmt.Sprintf(`"%s"`, url)))
	if err != nil {
		return fmt.Errorf("validate url from string failed with error: %w", err)
	}

	return nil
}
