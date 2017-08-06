// This file was generated by counterfeiter
package core

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type fakeOpKiller struct {
	KillStub        func(req model.KillOpReq)
	killMutex       sync.RWMutex
	killArgsForCall []struct {
		req model.KillOpReq
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeOpKiller) Kill(req model.KillOpReq) {
	fake.killMutex.Lock()
	fake.killArgsForCall = append(fake.killArgsForCall, struct {
		req model.KillOpReq
	}{req})
	fake.recordInvocation("Kill", []interface{}{req})
	fake.killMutex.Unlock()
	if fake.KillStub != nil {
		fake.KillStub(req)
	}
}

func (fake *fakeOpKiller) KillCallCount() int {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	return len(fake.killArgsForCall)
}

func (fake *fakeOpKiller) KillArgsForCall(i int) model.KillOpReq {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	return fake.killArgsForCall[i].req
}

func (fake *fakeOpKiller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	return fake.invocations
}

func (fake *fakeOpKiller) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}