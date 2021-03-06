// Copyright 2019 SumUp Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ostest

import (
	"io"
	stdOs "os"
	"os/user"
	"testing"

	"github.com/sumup-oss/go-pkgs/os"

	"github.com/stretchr/testify/mock"
)

var _ os.OsExecutor = (*FakeOsExecutor)(nil)

type FakeOsExecutor struct {
	mock.Mock
}

func NewFakeOsExecutor(t *testing.T) *FakeOsExecutor {
	fake := &FakeOsExecutor{}
	fake.Test(t)

	return fake
}

func (f *FakeOsExecutor) Getwd() (string, error) {
	args := f.Called()
	return args.String(0), args.Error(1)
}

func (f *FakeOsExecutor) Chdir(dir string) error {
	args := f.Called(dir)
	return args.Error(0)
}

func (f *FakeOsExecutor) Mkdir(dirname string, perm stdOs.FileMode) error {
	args := f.Called(dirname, perm)
	return args.Error(0)
}

func (f *FakeOsExecutor) Execute(
	cmd string,
	arg []string,
	env []string,
	dir string,
) ([]byte, []byte, error) {
	args := f.Called(cmd, arg, env, dir)
	rawStdout := args.Get(0)
	rawStderr := args.Get(1)
	returnErr := args.Error(2)

	var returnStdout, returnStderr []byte
	if rawStdout != nil {
		returnStdout = rawStdout.([]byte)
	}
	if rawStderr != nil {
		returnStderr = rawStderr.([]byte)
	}

	return returnStdout, returnStderr, returnErr
}

func (f *FakeOsExecutor) MkdirAll(dirname string, perm stdOs.FileMode) error {
	args := f.Called(dirname, perm)
	return args.Error(0)
}

func (f *FakeOsExecutor) Exit(statusCode int) {
	f.Called(statusCode)
}

func (f *FakeOsExecutor) Stderr() io.Writer {
	args := f.Called()
	returnValue := args.Get(0)
	if returnValue == nil {
		return nil
	}

	return returnValue.(io.Writer)
}

func (f *FakeOsExecutor) Stdin() io.Reader {
	args := f.Called()
	returnValue := args.Get(0)
	if returnValue == nil {
		return nil
	}

	return returnValue.(io.Reader)
}

func (f *FakeOsExecutor) Stdout() io.Writer {
	args := f.Called()
	returnValue := args.Get(0)
	if returnValue == nil {
		return nil
	}

	return returnValue.(io.Writer)
}

func (f *FakeOsExecutor) Args() []string {
	args := f.Called()
	returnValue := args.Get(0)
	if returnValue == nil {
		return nil
	}

	return returnValue.([]string)
}

func (f *FakeOsExecutor) Stat(filepath string) (stdOs.FileInfo, error) {
	args := f.Called(filepath)
	returnValue := args.Get(0)
	err := args.Error(1)

	if returnValue == nil {
		return nil, err
	}

	return returnValue.(stdOs.FileInfo), err
}

func (f *FakeOsExecutor) IsNotExist(err error) bool {
	args := f.Called(err)
	return args.Bool(0)
}

func (f *FakeOsExecutor) OpenFile(path string, flag int, perm stdOs.FileMode) (*stdOs.File, error) {
	args := f.Called(path, flag, perm)

	returnValue := args.Get(0)
	err := args.Error(1)

	if returnValue == nil {
		return nil, err
	}

	return returnValue.(*stdOs.File), err
}

func (f *FakeOsExecutor) WriteFile(path string, data []byte, perm stdOs.FileMode) error {
	args := f.Called(path, data, perm)
	return args.Error(0)
}

func (f *FakeOsExecutor) ExpandTilde(path string) (string, error) {
	args := f.Called(path)
	return args.String(0), args.Error(1)
}

func (f *FakeOsExecutor) Getenv(key string) string {
	args := f.Called(key)
	return args.String(0)
}

func (f *FakeOsExecutor) GetOS() string {
	args := f.Called()
	return args.String(0)
}

func (f *FakeOsExecutor) ExecuteWithStreams(
	cmd string,
	arg []string,
	env []string,
	dir string,
	stdout io.Writer,
	stderr io.Writer,
) error {
	args := f.Called(cmd, arg, env, dir, stdout, stderr)
	return args.Error(0)
}

func (f *FakeOsExecutor) ResolvePath(path string) (string, error) {
	args := f.Called(path)
	return args.String(0), args.Error(1)
}

func (f *FakeOsExecutor) Remove(path string) error {
	args := f.Called(path)
	return args.Error(0)
}

func (f *FakeOsExecutor) CurrentUser() (*user.User, error) {
	args := f.Called()
	returnValue := args.Get(0)
	err := args.Error(1)
	if returnValue == nil {
		return nil, err
	}

	return returnValue.(*user.User), err
}

func (f *FakeOsExecutor) Create(name string) (*stdOs.File, error) {
	args := f.Called(name)
	returnValue := args.Get(0)
	err := args.Error(1)
	if returnValue == nil {
		return nil, err
	}

	return returnValue.(*stdOs.File), err
}

func (f *FakeOsExecutor) ReadFile(filename string) (bytes []byte, e error) {
	args := f.Called(filename)
	returnValue := args.Get(0)
	err := args.Error(1)
	if returnValue == nil {
		return nil, err
	}

	return returnValue.([]byte), err
}

func (f *FakeOsExecutor) IsDir(path string) error {
	args := f.Called(path)
	return args.Error(0)
}

func (f *FakeOsExecutor) IsFile(path string) error {
	args := f.Called(path)
	return args.Error(0)
}

func (f *FakeOsExecutor) RemoveAll(path string) error {
	args := f.Called(path)
	return args.Error(0)
}
