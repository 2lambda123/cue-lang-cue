//go:build go1.19

// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package par

import "sync/atomic"

type atomicBool = atomic.Bool