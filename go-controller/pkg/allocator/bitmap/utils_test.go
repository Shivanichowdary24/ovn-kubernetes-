// SPDX-FileCopyrightText: Copyright The OVN-Kubernetes Contributors
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package bitmap

import (
	"math/big"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		name     string
		n        *big.Int
		expected int
	}{
		{
			name:     "zero",
			n:        big.NewInt(0),
			expected: 0,
		},
		{
			name:     "single byte",
			n:        big.NewInt(0xff),
			expected: 8,
		},
		{
			name:     "multiple bytes",
			n:        big.NewInt(0xffffffffff),
			expected: 40,
		},
		{
			name:     "single bit",
			n:        big.NewInt(1),
			expected: 1,
		},
		{
			name:     "alternating bits",
			n:        big.NewInt(0xaaaaaaaa),
			expected: 16,
		},
		{
			name:     "byte boundary bits",
			n:        big.NewInt(0x0101010101),
			expected: 5,
		},
		{
			name:     "high bit in each byte",
			n:        big.NewInt(0x8080808080),
			expected: 5,
		},
		{
			name:     "all bits across multiple bytes",
			n:        new(big.Int).SetUint64(0xffffffffffffffff),
			expected: 64,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := countBits(test.n)
			if test.expected != actual {
				t.Errorf("%s should have %d bits but recorded as %d",
					test.n, test.expected, actual)
			}
		})
	}
}
