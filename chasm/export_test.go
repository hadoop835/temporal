// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package chasm

func (r *Registry) Component(fqn string) (*RegistrableComponent, bool) {
	return r.component(fqn)
}

func (r *Registry) Task(fqn string) (*RegistrableTask, bool) {
	return r.task(fqn)
}

func (r *Registry) ComponentFor(componentInstance any) (*RegistrableComponent, bool) {
	return r.componentFor(componentInstance)
}

func (r *Registry) TaskFor(taskInstance any) (*RegistrableTask, bool) {
	return r.taskFor(taskInstance)
}

func (rc RegistrableComponent) FqType() string {
	return rc.fqType()
}

func (rt RegistrableTask) FqType() string {
	return rt.fqType()
}
