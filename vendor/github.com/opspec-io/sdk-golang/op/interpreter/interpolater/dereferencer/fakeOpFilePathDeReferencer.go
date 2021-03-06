// Code generated by counterfeiter. DO NOT EDIT.
package dereferencer

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type fakeOpFilePathDeReferencer struct {
	DeReferenceOpFilePathStub        func(ref string, scope map[string]*model.Value, opHandle model.DataHandle) (string, bool, error)
	deReferenceOpFilePathMutex       sync.RWMutex
	deReferenceOpFilePathArgsForCall []struct {
		ref      string
		scope    map[string]*model.Value
		opHandle model.DataHandle
	}
	deReferenceOpFilePathReturns struct {
		result1 string
		result2 bool
		result3 error
	}
	deReferenceOpFilePathReturnsOnCall map[int]struct {
		result1 string
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeOpFilePathDeReferencer) DeReferenceOpFilePath(ref string, scope map[string]*model.Value, opHandle model.DataHandle) (string, bool, error) {
	fake.deReferenceOpFilePathMutex.Lock()
	ret, specificReturn := fake.deReferenceOpFilePathReturnsOnCall[len(fake.deReferenceOpFilePathArgsForCall)]
	fake.deReferenceOpFilePathArgsForCall = append(fake.deReferenceOpFilePathArgsForCall, struct {
		ref      string
		scope    map[string]*model.Value
		opHandle model.DataHandle
	}{ref, scope, opHandle})
	fake.recordInvocation("DeReferenceOpFilePath", []interface{}{ref, scope, opHandle})
	fake.deReferenceOpFilePathMutex.Unlock()
	if fake.DeReferenceOpFilePathStub != nil {
		return fake.DeReferenceOpFilePathStub(ref, scope, opHandle)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.deReferenceOpFilePathReturns.result1, fake.deReferenceOpFilePathReturns.result2, fake.deReferenceOpFilePathReturns.result3
}

func (fake *fakeOpFilePathDeReferencer) DeReferenceOpFilePathCallCount() int {
	fake.deReferenceOpFilePathMutex.RLock()
	defer fake.deReferenceOpFilePathMutex.RUnlock()
	return len(fake.deReferenceOpFilePathArgsForCall)
}

func (fake *fakeOpFilePathDeReferencer) DeReferenceOpFilePathArgsForCall(i int) (string, map[string]*model.Value, model.DataHandle) {
	fake.deReferenceOpFilePathMutex.RLock()
	defer fake.deReferenceOpFilePathMutex.RUnlock()
	return fake.deReferenceOpFilePathArgsForCall[i].ref, fake.deReferenceOpFilePathArgsForCall[i].scope, fake.deReferenceOpFilePathArgsForCall[i].opHandle
}

func (fake *fakeOpFilePathDeReferencer) DeReferenceOpFilePathReturns(result1 string, result2 bool, result3 error) {
	fake.DeReferenceOpFilePathStub = nil
	fake.deReferenceOpFilePathReturns = struct {
		result1 string
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *fakeOpFilePathDeReferencer) DeReferenceOpFilePathReturnsOnCall(i int, result1 string, result2 bool, result3 error) {
	fake.DeReferenceOpFilePathStub = nil
	if fake.deReferenceOpFilePathReturnsOnCall == nil {
		fake.deReferenceOpFilePathReturnsOnCall = make(map[int]struct {
			result1 string
			result2 bool
			result3 error
		})
	}
	fake.deReferenceOpFilePathReturnsOnCall[i] = struct {
		result1 string
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *fakeOpFilePathDeReferencer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deReferenceOpFilePathMutex.RLock()
	defer fake.deReferenceOpFilePathMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *fakeOpFilePathDeReferencer) recordInvocation(key string, args []interface{}) {
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
