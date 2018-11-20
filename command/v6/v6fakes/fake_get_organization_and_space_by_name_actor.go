// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v6"
)

type FakeGetOrganizationAndSpaceByNameActor struct {
	GetOrganizationByNameStub        func(name string) (v3action.Organization, v3action.Warnings, error)
	getOrganizationByNameMutex       sync.RWMutex
	getOrganizationByNameArgsForCall []struct {
		name string
	}
	getOrganizationByNameReturns struct {
		result1 v3action.Organization
		result2 v3action.Warnings
		result3 error
	}
	getOrganizationByNameReturnsOnCall map[int]struct {
		result1 v3action.Organization
		result2 v3action.Warnings
		result3 error
	}
	GetSpaceByNameAndOrganizationStub        func(spaceName string, orgGUID string) (v3action.Space, v3action.Warnings, error)
	getSpaceByNameAndOrganizationMutex       sync.RWMutex
	getSpaceByNameAndOrganizationArgsForCall []struct {
		spaceName string
		orgGUID   string
	}
	getSpaceByNameAndOrganizationReturns struct {
		result1 v3action.Space
		result2 v3action.Warnings
		result3 error
	}
	getSpaceByNameAndOrganizationReturnsOnCall map[int]struct {
		result1 v3action.Space
		result2 v3action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGetOrganizationAndSpaceByNameActor) GetOrganizationByName(name string) (v3action.Organization, v3action.Warnings, error) {
	fake.getOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.getOrganizationByNameReturnsOnCall[len(fake.getOrganizationByNameArgsForCall)]
	fake.getOrganizationByNameArgsForCall = append(fake.getOrganizationByNameArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("GetOrganizationByName", []interface{}{name})
	fake.getOrganizationByNameMutex.Unlock()
	if fake.GetOrganizationByNameStub != nil {
		return fake.GetOrganizationByNameStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getOrganizationByNameReturns.result1, fake.getOrganizationByNameReturns.result2, fake.getOrganizationByNameReturns.result3
}

func (fake *FakeGetOrganizationAndSpaceByNameActor) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakeGetOrganizationAndSpaceByNameActor) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return fake.getOrganizationByNameArgsForCall[i].name
}

func (fake *FakeGetOrganizationAndSpaceByNameActor) GetOrganizationByNameReturns(result1 v3action.Organization, result2 v3action.Warnings, result3 error) {
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v3action.Organization
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGetOrganizationAndSpaceByNameActor) GetOrganizationByNameReturnsOnCall(i int, result1 v3action.Organization, result2 v3action.Warnings, result3 error) {
	fake.GetOrganizationByNameStub = nil
	if fake.getOrganizationByNameReturnsOnCall == nil {
		fake.getOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v3action.Organization
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getOrganizationByNameReturnsOnCall[i] = struct {
		result1 v3action.Organization
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGetOrganizationAndSpaceByNameActor) GetSpaceByNameAndOrganization(spaceName string, orgGUID string) (v3action.Space, v3action.Warnings, error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	ret, specificReturn := fake.getSpaceByNameAndOrganizationReturnsOnCall[len(fake.getSpaceByNameAndOrganizationArgsForCall)]
	fake.getSpaceByNameAndOrganizationArgsForCall = append(fake.getSpaceByNameAndOrganizationArgsForCall, struct {
		spaceName string
		orgGUID   string
	}{spaceName, orgGUID})
	fake.recordInvocation("GetSpaceByNameAndOrganization", []interface{}{spaceName, orgGUID})
	fake.getSpaceByNameAndOrganizationMutex.Unlock()
	if fake.GetSpaceByNameAndOrganizationStub != nil {
		return fake.GetSpaceByNameAndOrganizationStub(spaceName, orgGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getSpaceByNameAndOrganizationReturns.result1, fake.getSpaceByNameAndOrganizationReturns.result2, fake.getSpaceByNameAndOrganizationReturns.result3
}

func (fake *FakeGetOrganizationAndSpaceByNameActor) GetSpaceByNameAndOrganizationCallCount() int {
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	return len(fake.getSpaceByNameAndOrganizationArgsForCall)
}

func (fake *FakeGetOrganizationAndSpaceByNameActor) GetSpaceByNameAndOrganizationArgsForCall(i int) (string, string) {
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	return fake.getSpaceByNameAndOrganizationArgsForCall[i].spaceName, fake.getSpaceByNameAndOrganizationArgsForCall[i].orgGUID
}

func (fake *FakeGetOrganizationAndSpaceByNameActor) GetSpaceByNameAndOrganizationReturns(result1 v3action.Space, result2 v3action.Warnings, result3 error) {
	fake.GetSpaceByNameAndOrganizationStub = nil
	fake.getSpaceByNameAndOrganizationReturns = struct {
		result1 v3action.Space
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGetOrganizationAndSpaceByNameActor) GetSpaceByNameAndOrganizationReturnsOnCall(i int, result1 v3action.Space, result2 v3action.Warnings, result3 error) {
	fake.GetSpaceByNameAndOrganizationStub = nil
	if fake.getSpaceByNameAndOrganizationReturnsOnCall == nil {
		fake.getSpaceByNameAndOrganizationReturnsOnCall = make(map[int]struct {
			result1 v3action.Space
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getSpaceByNameAndOrganizationReturnsOnCall[i] = struct {
		result1 v3action.Space
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGetOrganizationAndSpaceByNameActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeGetOrganizationAndSpaceByNameActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.GetOrganizationAndSpaceByNameActor = new(FakeGetOrganizationAndSpaceByNameActor)
