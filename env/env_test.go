/*
 * Copyright (C) 2025 Canonical Ltd
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 * SPDX-License-Identifier: Apache-2.0'
 */

package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvVars(t *testing.T) {
	// SNAP
	assert.Regexp(t, "/snap/go-snapctl-tester/x\\d+", Snap())
	// SNAP_COMMON
	assert.Regexp(t, "/var/snap/go-snapctl-tester/common", SnapCommon())
	// SNAP_DATA
	assert.Regexp(t, "/var/snap/go-snapctl-tester/x\\d+", SnapData())
	// SNAP_INSTANCE_NAME
	assert.Equal(t, "go-snapctl-tester", SnapInstanceName())
	// SNAP_NAME
	assert.Equal(t, "go-snapctl-tester", SnapName())
	// SNAP_REVISION
	assert.Regexp(t, "x\\d+", SnapRevision())
}
