// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
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

package x25519

import (
	"crypto/rand"
	"io"

	"github.com/agl/ed25519/extra25519"
)

// RandUniform - Generates a secret key whose corresponding public key has a uniform representative.
func RandUniform() (_ [32]byte, _ [32]byte, _ [32]byte, err error) {
	var sk, pk, upk [32]byte
	for {
		_, err = io.ReadFull(rand.Reader, sk[:])
		if err != nil {
			return
		}
		if extra25519.ScalarBaseMult(&pk, &upk, &sk) {
			break
		}
	}
	return sk, pk, upk, nil
}
