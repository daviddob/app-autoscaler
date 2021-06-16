// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"autoscaler/db"
	"autoscaler/models"
	"database/sql"
	"sync"
)

type FakePolicyDB struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteCredentialStub        func(string) error
	deleteCredentialMutex       sync.RWMutex
	deleteCredentialArgsForCall []struct {
		arg1 string
	}
	deleteCredentialReturns struct {
		result1 error
	}
	deleteCredentialReturnsOnCall map[int]struct {
		result1 error
	}
	DeletePolicyStub        func(string) error
	deletePolicyMutex       sync.RWMutex
	deletePolicyArgsForCall []struct {
		arg1 string
	}
	deletePolicyReturns struct {
		result1 error
	}
	deletePolicyReturnsOnCall map[int]struct {
		result1 error
	}
	GetAppIdsStub        func() (map[string]bool, error)
	getAppIdsMutex       sync.RWMutex
	getAppIdsArgsForCall []struct {
	}
	getAppIdsReturns struct {
		result1 map[string]bool
		result2 error
	}
	getAppIdsReturnsOnCall map[int]struct {
		result1 map[string]bool
		result2 error
	}
	GetAppPolicyStub        func(string) (*models.ScalingPolicy, error)
	getAppPolicyMutex       sync.RWMutex
	getAppPolicyArgsForCall []struct {
		arg1 string
	}
	getAppPolicyReturns struct {
		result1 *models.ScalingPolicy
		result2 error
	}
	getAppPolicyReturnsOnCall map[int]struct {
		result1 *models.ScalingPolicy
		result2 error
	}
	GetCredentialStub        func(string) (*models.Credential, error)
	getCredentialMutex       sync.RWMutex
	getCredentialArgsForCall []struct {
		arg1 string
	}
	getCredentialReturns struct {
		result1 *models.Credential
		result2 error
	}
	getCredentialReturnsOnCall map[int]struct {
		result1 *models.Credential
		result2 error
	}
	GetDBStatusStub        func() sql.DBStats
	getDBStatusMutex       sync.RWMutex
	getDBStatusArgsForCall []struct {
	}
	getDBStatusReturns struct {
		result1 sql.DBStats
	}
	getDBStatusReturnsOnCall map[int]struct {
		result1 sql.DBStats
	}
	RetrievePoliciesStub        func() ([]*models.PolicyJson, error)
	retrievePoliciesMutex       sync.RWMutex
	retrievePoliciesArgsForCall []struct {
	}
	retrievePoliciesReturns struct {
		result1 []*models.PolicyJson
		result2 error
	}
	retrievePoliciesReturnsOnCall map[int]struct {
		result1 []*models.PolicyJson
		result2 error
	}
	SaveAppPolicyStub        func(string, string, string) error
	saveAppPolicyMutex       sync.RWMutex
	saveAppPolicyArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	saveAppPolicyReturns struct {
		result1 error
	}
	saveAppPolicyReturnsOnCall map[int]struct {
		result1 error
	}
	SaveCredentialStub        func(string, models.Credential) error
	saveCredentialMutex       sync.RWMutex
	saveCredentialArgsForCall []struct {
		arg1 string
		arg2 models.Credential
	}
	saveCredentialReturns struct {
		result1 error
	}
	saveCredentialReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePolicyDB) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePolicyDB) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakePolicyDB) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakePolicyDB) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePolicyDB) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePolicyDB) DeleteCredential(arg1 string) error {
	fake.deleteCredentialMutex.Lock()
	ret, specificReturn := fake.deleteCredentialReturnsOnCall[len(fake.deleteCredentialArgsForCall)]
	fake.deleteCredentialArgsForCall = append(fake.deleteCredentialArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteCredentialStub
	fakeReturns := fake.deleteCredentialReturns
	fake.recordInvocation("DeleteCredential", []interface{}{arg1})
	fake.deleteCredentialMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePolicyDB) DeleteCredentialCallCount() int {
	fake.deleteCredentialMutex.RLock()
	defer fake.deleteCredentialMutex.RUnlock()
	return len(fake.deleteCredentialArgsForCall)
}

func (fake *FakePolicyDB) DeleteCredentialCalls(stub func(string) error) {
	fake.deleteCredentialMutex.Lock()
	defer fake.deleteCredentialMutex.Unlock()
	fake.DeleteCredentialStub = stub
}

func (fake *FakePolicyDB) DeleteCredentialArgsForCall(i int) string {
	fake.deleteCredentialMutex.RLock()
	defer fake.deleteCredentialMutex.RUnlock()
	argsForCall := fake.deleteCredentialArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePolicyDB) DeleteCredentialReturns(result1 error) {
	fake.deleteCredentialMutex.Lock()
	defer fake.deleteCredentialMutex.Unlock()
	fake.DeleteCredentialStub = nil
	fake.deleteCredentialReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePolicyDB) DeleteCredentialReturnsOnCall(i int, result1 error) {
	fake.deleteCredentialMutex.Lock()
	defer fake.deleteCredentialMutex.Unlock()
	fake.DeleteCredentialStub = nil
	if fake.deleteCredentialReturnsOnCall == nil {
		fake.deleteCredentialReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteCredentialReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePolicyDB) DeletePolicy(arg1 string) error {
	fake.deletePolicyMutex.Lock()
	ret, specificReturn := fake.deletePolicyReturnsOnCall[len(fake.deletePolicyArgsForCall)]
	fake.deletePolicyArgsForCall = append(fake.deletePolicyArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeletePolicyStub
	fakeReturns := fake.deletePolicyReturns
	fake.recordInvocation("DeletePolicy", []interface{}{arg1})
	fake.deletePolicyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePolicyDB) DeletePolicyCallCount() int {
	fake.deletePolicyMutex.RLock()
	defer fake.deletePolicyMutex.RUnlock()
	return len(fake.deletePolicyArgsForCall)
}

func (fake *FakePolicyDB) DeletePolicyCalls(stub func(string) error) {
	fake.deletePolicyMutex.Lock()
	defer fake.deletePolicyMutex.Unlock()
	fake.DeletePolicyStub = stub
}

func (fake *FakePolicyDB) DeletePolicyArgsForCall(i int) string {
	fake.deletePolicyMutex.RLock()
	defer fake.deletePolicyMutex.RUnlock()
	argsForCall := fake.deletePolicyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePolicyDB) DeletePolicyReturns(result1 error) {
	fake.deletePolicyMutex.Lock()
	defer fake.deletePolicyMutex.Unlock()
	fake.DeletePolicyStub = nil
	fake.deletePolicyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePolicyDB) DeletePolicyReturnsOnCall(i int, result1 error) {
	fake.deletePolicyMutex.Lock()
	defer fake.deletePolicyMutex.Unlock()
	fake.DeletePolicyStub = nil
	if fake.deletePolicyReturnsOnCall == nil {
		fake.deletePolicyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deletePolicyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePolicyDB) GetAppIds() (map[string]bool, error) {
	fake.getAppIdsMutex.Lock()
	ret, specificReturn := fake.getAppIdsReturnsOnCall[len(fake.getAppIdsArgsForCall)]
	fake.getAppIdsArgsForCall = append(fake.getAppIdsArgsForCall, struct {
	}{})
	stub := fake.GetAppIdsStub
	fakeReturns := fake.getAppIdsReturns
	fake.recordInvocation("GetAppIds", []interface{}{})
	fake.getAppIdsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePolicyDB) GetAppIdsCallCount() int {
	fake.getAppIdsMutex.RLock()
	defer fake.getAppIdsMutex.RUnlock()
	return len(fake.getAppIdsArgsForCall)
}

func (fake *FakePolicyDB) GetAppIdsCalls(stub func() (map[string]bool, error)) {
	fake.getAppIdsMutex.Lock()
	defer fake.getAppIdsMutex.Unlock()
	fake.GetAppIdsStub = stub
}

func (fake *FakePolicyDB) GetAppIdsReturns(result1 map[string]bool, result2 error) {
	fake.getAppIdsMutex.Lock()
	defer fake.getAppIdsMutex.Unlock()
	fake.GetAppIdsStub = nil
	fake.getAppIdsReturns = struct {
		result1 map[string]bool
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyDB) GetAppIdsReturnsOnCall(i int, result1 map[string]bool, result2 error) {
	fake.getAppIdsMutex.Lock()
	defer fake.getAppIdsMutex.Unlock()
	fake.GetAppIdsStub = nil
	if fake.getAppIdsReturnsOnCall == nil {
		fake.getAppIdsReturnsOnCall = make(map[int]struct {
			result1 map[string]bool
			result2 error
		})
	}
	fake.getAppIdsReturnsOnCall[i] = struct {
		result1 map[string]bool
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyDB) GetAppPolicy(arg1 string) (*models.ScalingPolicy, error) {
	fake.getAppPolicyMutex.Lock()
	ret, specificReturn := fake.getAppPolicyReturnsOnCall[len(fake.getAppPolicyArgsForCall)]
	fake.getAppPolicyArgsForCall = append(fake.getAppPolicyArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetAppPolicyStub
	fakeReturns := fake.getAppPolicyReturns
	fake.recordInvocation("GetAppPolicy", []interface{}{arg1})
	fake.getAppPolicyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePolicyDB) GetAppPolicyCallCount() int {
	fake.getAppPolicyMutex.RLock()
	defer fake.getAppPolicyMutex.RUnlock()
	return len(fake.getAppPolicyArgsForCall)
}

func (fake *FakePolicyDB) GetAppPolicyCalls(stub func(string) (*models.ScalingPolicy, error)) {
	fake.getAppPolicyMutex.Lock()
	defer fake.getAppPolicyMutex.Unlock()
	fake.GetAppPolicyStub = stub
}

func (fake *FakePolicyDB) GetAppPolicyArgsForCall(i int) string {
	fake.getAppPolicyMutex.RLock()
	defer fake.getAppPolicyMutex.RUnlock()
	argsForCall := fake.getAppPolicyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePolicyDB) GetAppPolicyReturns(result1 *models.ScalingPolicy, result2 error) {
	fake.getAppPolicyMutex.Lock()
	defer fake.getAppPolicyMutex.Unlock()
	fake.GetAppPolicyStub = nil
	fake.getAppPolicyReturns = struct {
		result1 *models.ScalingPolicy
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyDB) GetAppPolicyReturnsOnCall(i int, result1 *models.ScalingPolicy, result2 error) {
	fake.getAppPolicyMutex.Lock()
	defer fake.getAppPolicyMutex.Unlock()
	fake.GetAppPolicyStub = nil
	if fake.getAppPolicyReturnsOnCall == nil {
		fake.getAppPolicyReturnsOnCall = make(map[int]struct {
			result1 *models.ScalingPolicy
			result2 error
		})
	}
	fake.getAppPolicyReturnsOnCall[i] = struct {
		result1 *models.ScalingPolicy
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyDB) GetCredential(arg1 string) (*models.Credential, error) {
	fake.getCredentialMutex.Lock()
	ret, specificReturn := fake.getCredentialReturnsOnCall[len(fake.getCredentialArgsForCall)]
	fake.getCredentialArgsForCall = append(fake.getCredentialArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetCredentialStub
	fakeReturns := fake.getCredentialReturns
	fake.recordInvocation("GetCredential", []interface{}{arg1})
	fake.getCredentialMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePolicyDB) GetCredentialCallCount() int {
	fake.getCredentialMutex.RLock()
	defer fake.getCredentialMutex.RUnlock()
	return len(fake.getCredentialArgsForCall)
}

func (fake *FakePolicyDB) GetCredentialCalls(stub func(string) (*models.Credential, error)) {
	fake.getCredentialMutex.Lock()
	defer fake.getCredentialMutex.Unlock()
	fake.GetCredentialStub = stub
}

func (fake *FakePolicyDB) GetCredentialArgsForCall(i int) string {
	fake.getCredentialMutex.RLock()
	defer fake.getCredentialMutex.RUnlock()
	argsForCall := fake.getCredentialArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePolicyDB) GetCredentialReturns(result1 *models.Credential, result2 error) {
	fake.getCredentialMutex.Lock()
	defer fake.getCredentialMutex.Unlock()
	fake.GetCredentialStub = nil
	fake.getCredentialReturns = struct {
		result1 *models.Credential
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyDB) GetCredentialReturnsOnCall(i int, result1 *models.Credential, result2 error) {
	fake.getCredentialMutex.Lock()
	defer fake.getCredentialMutex.Unlock()
	fake.GetCredentialStub = nil
	if fake.getCredentialReturnsOnCall == nil {
		fake.getCredentialReturnsOnCall = make(map[int]struct {
			result1 *models.Credential
			result2 error
		})
	}
	fake.getCredentialReturnsOnCall[i] = struct {
		result1 *models.Credential
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyDB) GetDBStatus() sql.DBStats {
	fake.getDBStatusMutex.Lock()
	ret, specificReturn := fake.getDBStatusReturnsOnCall[len(fake.getDBStatusArgsForCall)]
	fake.getDBStatusArgsForCall = append(fake.getDBStatusArgsForCall, struct {
	}{})
	stub := fake.GetDBStatusStub
	fakeReturns := fake.getDBStatusReturns
	fake.recordInvocation("GetDBStatus", []interface{}{})
	fake.getDBStatusMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePolicyDB) GetDBStatusCallCount() int {
	fake.getDBStatusMutex.RLock()
	defer fake.getDBStatusMutex.RUnlock()
	return len(fake.getDBStatusArgsForCall)
}

func (fake *FakePolicyDB) GetDBStatusCalls(stub func() sql.DBStats) {
	fake.getDBStatusMutex.Lock()
	defer fake.getDBStatusMutex.Unlock()
	fake.GetDBStatusStub = stub
}

func (fake *FakePolicyDB) GetDBStatusReturns(result1 sql.DBStats) {
	fake.getDBStatusMutex.Lock()
	defer fake.getDBStatusMutex.Unlock()
	fake.GetDBStatusStub = nil
	fake.getDBStatusReturns = struct {
		result1 sql.DBStats
	}{result1}
}

func (fake *FakePolicyDB) GetDBStatusReturnsOnCall(i int, result1 sql.DBStats) {
	fake.getDBStatusMutex.Lock()
	defer fake.getDBStatusMutex.Unlock()
	fake.GetDBStatusStub = nil
	if fake.getDBStatusReturnsOnCall == nil {
		fake.getDBStatusReturnsOnCall = make(map[int]struct {
			result1 sql.DBStats
		})
	}
	fake.getDBStatusReturnsOnCall[i] = struct {
		result1 sql.DBStats
	}{result1}
}

func (fake *FakePolicyDB) RetrievePolicies() ([]*models.PolicyJson, error) {
	fake.retrievePoliciesMutex.Lock()
	ret, specificReturn := fake.retrievePoliciesReturnsOnCall[len(fake.retrievePoliciesArgsForCall)]
	fake.retrievePoliciesArgsForCall = append(fake.retrievePoliciesArgsForCall, struct {
	}{})
	stub := fake.RetrievePoliciesStub
	fakeReturns := fake.retrievePoliciesReturns
	fake.recordInvocation("RetrievePolicies", []interface{}{})
	fake.retrievePoliciesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePolicyDB) RetrievePoliciesCallCount() int {
	fake.retrievePoliciesMutex.RLock()
	defer fake.retrievePoliciesMutex.RUnlock()
	return len(fake.retrievePoliciesArgsForCall)
}

func (fake *FakePolicyDB) RetrievePoliciesCalls(stub func() ([]*models.PolicyJson, error)) {
	fake.retrievePoliciesMutex.Lock()
	defer fake.retrievePoliciesMutex.Unlock()
	fake.RetrievePoliciesStub = stub
}

func (fake *FakePolicyDB) RetrievePoliciesReturns(result1 []*models.PolicyJson, result2 error) {
	fake.retrievePoliciesMutex.Lock()
	defer fake.retrievePoliciesMutex.Unlock()
	fake.RetrievePoliciesStub = nil
	fake.retrievePoliciesReturns = struct {
		result1 []*models.PolicyJson
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyDB) RetrievePoliciesReturnsOnCall(i int, result1 []*models.PolicyJson, result2 error) {
	fake.retrievePoliciesMutex.Lock()
	defer fake.retrievePoliciesMutex.Unlock()
	fake.RetrievePoliciesStub = nil
	if fake.retrievePoliciesReturnsOnCall == nil {
		fake.retrievePoliciesReturnsOnCall = make(map[int]struct {
			result1 []*models.PolicyJson
			result2 error
		})
	}
	fake.retrievePoliciesReturnsOnCall[i] = struct {
		result1 []*models.PolicyJson
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyDB) SaveAppPolicy(arg1 string, arg2 string, arg3 string) error {
	fake.saveAppPolicyMutex.Lock()
	ret, specificReturn := fake.saveAppPolicyReturnsOnCall[len(fake.saveAppPolicyArgsForCall)]
	fake.saveAppPolicyArgsForCall = append(fake.saveAppPolicyArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.SaveAppPolicyStub
	fakeReturns := fake.saveAppPolicyReturns
	fake.recordInvocation("SaveAppPolicy", []interface{}{arg1, arg2, arg3})
	fake.saveAppPolicyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePolicyDB) SaveAppPolicyCallCount() int {
	fake.saveAppPolicyMutex.RLock()
	defer fake.saveAppPolicyMutex.RUnlock()
	return len(fake.saveAppPolicyArgsForCall)
}

func (fake *FakePolicyDB) SaveAppPolicyCalls(stub func(string, string, string) error) {
	fake.saveAppPolicyMutex.Lock()
	defer fake.saveAppPolicyMutex.Unlock()
	fake.SaveAppPolicyStub = stub
}

func (fake *FakePolicyDB) SaveAppPolicyArgsForCall(i int) (string, string, string) {
	fake.saveAppPolicyMutex.RLock()
	defer fake.saveAppPolicyMutex.RUnlock()
	argsForCall := fake.saveAppPolicyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePolicyDB) SaveAppPolicyReturns(result1 error) {
	fake.saveAppPolicyMutex.Lock()
	defer fake.saveAppPolicyMutex.Unlock()
	fake.SaveAppPolicyStub = nil
	fake.saveAppPolicyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePolicyDB) SaveAppPolicyReturnsOnCall(i int, result1 error) {
	fake.saveAppPolicyMutex.Lock()
	defer fake.saveAppPolicyMutex.Unlock()
	fake.SaveAppPolicyStub = nil
	if fake.saveAppPolicyReturnsOnCall == nil {
		fake.saveAppPolicyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveAppPolicyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePolicyDB) SaveCredential(arg1 string, arg2 models.Credential) error {
	fake.saveCredentialMutex.Lock()
	ret, specificReturn := fake.saveCredentialReturnsOnCall[len(fake.saveCredentialArgsForCall)]
	fake.saveCredentialArgsForCall = append(fake.saveCredentialArgsForCall, struct {
		arg1 string
		arg2 models.Credential
	}{arg1, arg2})
	stub := fake.SaveCredentialStub
	fakeReturns := fake.saveCredentialReturns
	fake.recordInvocation("SaveCredential", []interface{}{arg1, arg2})
	fake.saveCredentialMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePolicyDB) SaveCredentialCallCount() int {
	fake.saveCredentialMutex.RLock()
	defer fake.saveCredentialMutex.RUnlock()
	return len(fake.saveCredentialArgsForCall)
}

func (fake *FakePolicyDB) SaveCredentialCalls(stub func(string, models.Credential) error) {
	fake.saveCredentialMutex.Lock()
	defer fake.saveCredentialMutex.Unlock()
	fake.SaveCredentialStub = stub
}

func (fake *FakePolicyDB) SaveCredentialArgsForCall(i int) (string, models.Credential) {
	fake.saveCredentialMutex.RLock()
	defer fake.saveCredentialMutex.RUnlock()
	argsForCall := fake.saveCredentialArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePolicyDB) SaveCredentialReturns(result1 error) {
	fake.saveCredentialMutex.Lock()
	defer fake.saveCredentialMutex.Unlock()
	fake.SaveCredentialStub = nil
	fake.saveCredentialReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePolicyDB) SaveCredentialReturnsOnCall(i int, result1 error) {
	fake.saveCredentialMutex.Lock()
	defer fake.saveCredentialMutex.Unlock()
	fake.SaveCredentialStub = nil
	if fake.saveCredentialReturnsOnCall == nil {
		fake.saveCredentialReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveCredentialReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePolicyDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.deleteCredentialMutex.RLock()
	defer fake.deleteCredentialMutex.RUnlock()
	fake.deletePolicyMutex.RLock()
	defer fake.deletePolicyMutex.RUnlock()
	fake.getAppIdsMutex.RLock()
	defer fake.getAppIdsMutex.RUnlock()
	fake.getAppPolicyMutex.RLock()
	defer fake.getAppPolicyMutex.RUnlock()
	fake.getCredentialMutex.RLock()
	defer fake.getCredentialMutex.RUnlock()
	fake.getDBStatusMutex.RLock()
	defer fake.getDBStatusMutex.RUnlock()
	fake.retrievePoliciesMutex.RLock()
	defer fake.retrievePoliciesMutex.RUnlock()
	fake.saveAppPolicyMutex.RLock()
	defer fake.saveAppPolicyMutex.RUnlock()
	fake.saveCredentialMutex.RLock()
	defer fake.saveCredentialMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePolicyDB) recordInvocation(key string, args []interface{}) {
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

var _ db.PolicyDB = new(FakePolicyDB)
