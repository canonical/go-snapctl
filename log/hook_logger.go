/*
 * Copyright (C) 2021 Canonical Ltd
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

package log

import (
	"fmt"
	"log/syslog"
	"os"
)

type hookLogger struct {
	writer *syslog.Writer
}

func newHookLogger(label string, debug bool) (*hookLogger, error) {
	snapInstanceName := os.Getenv("SNAP_INSTANCE_NAME")
	if snapInstanceName == "" {
		return nil, fmt.Errorf("SNAP_INSTANCE_NAME environment variable not set")
	}

	// Set syslog tag as "snap.<snap-instance-name>[.<label>]"
	tag := "snap." + snapInstanceName
	if label != "" {
		tag += "." + label
	}

	// The logging priority set here gets overridden by logging functions
	writer, err := syslog.New(syslog.LOG_INFO, tag)
	if err != nil {
		return nil, err
	}

	return &hookLogger{writer}, nil
}

func (l *hookLogger) Print(a ...any) {
	l.writer.Info(fmt.Sprint(a...))
}

func (l *hookLogger) Printf(format string, a ...any) {
	l.Print(fmt.Sprintf(format, a...))
}

func (l *hookLogger) Debug(a ...any) {
	l.writer.Debug(fmt.Sprint(a...))
}

func (l *hookLogger) Debugf(format string, a ...any) {
	l.Debug(fmt.Sprintf(format, a...))
}

func (l *hookLogger) Error(a ...any) {
	msg := fmt.Sprint(a...)
	l.writer.Err(msg)
	// print to stderr as well so that snap command prints them on non-zero exit
	stderr(a...)
}

func (l *hookLogger) Errorf(format string, a ...any) {
	l.Error(fmt.Sprintf(format, a...))
}

func (l *hookLogger) Fatal(a ...any) {
	l.Error(a...)
	os.Exit(1)
}

func (l *hookLogger) Fatalf(format string, a ...any) {
	l.Errorf(format, a...)
	os.Exit(1)
}
