/*
 * Copyright 2021 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package remote

import (
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/bytedance/gopkg/lang/dirtmake"
)

// Mask bits.
const (
	BitReadable = 1 << iota
	BitWritable
)

var bytebufPool sync.Pool

func init() {
	bytebufPool.New = newDefaultByteBuffer
}

// Next reads n bytes sequentially, returns the original address.
func (b *defaultByteBuffer) Next(n int) (buf []byte, err error) {
	if b.status&BitReadable == 0 {
		return nil, errors.New("unreadable buffer, cannot support Next")
	}
	buf, err = b.Peek(n)
	if err == nil {
		b.readIdx += n
	}
	return buf, err
}