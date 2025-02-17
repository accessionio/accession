// Copyright 2025 The Accession Authors
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

package models

import "github.com/oklog/ulid/v2"

// Fine-grained Access Control
type Permission struct {
	identifier ulid.ULID
}

func CreatePermission() Permission {
	return Permission{
		identifier: ulid.MustNew(ulid.Now(), ulid.DefaultEntropy()),
	}
}
