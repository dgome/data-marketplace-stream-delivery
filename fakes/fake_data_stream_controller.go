//
// Copyright (c) 2019 LG Electronics Inc.
// SPDX-License-Identifier: Apache-2.0
//

// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	context "context"
	sync "sync"

	resources "github.com/lgsvl/data-marketplace-stream-delivery/resources"
)

type FakeDataStreamController struct {
	CreateDataStreamStub        func(context.Context, resources.CreateDataStreamRequest) resources.CreateDataStreamResponse
	createDataStreamMutex       sync.RWMutex
	createDataStreamArgsForCall []struct {
		arg1 context.Context
		arg2 resources.CreateDataStreamRequest
	}
	createDataStreamReturns struct {
		result1 resources.CreateDataStreamResponse
	}
	createDataStreamReturnsOnCall map[int]struct {
		result1 resources.CreateDataStreamResponse
	}
	DeleteDataStreamStub        func(context.Context, resources.DeleteDataStreamRequest) resources.DeleteDataStreamResponse
	deleteDataStreamMutex       sync.RWMutex
	deleteDataStreamArgsForCall []struct {
		arg1 context.Context
		arg2 resources.DeleteDataStreamRequest
	}
	deleteDataStreamReturns struct {
		result1 resources.DeleteDataStreamResponse
	}
	deleteDataStreamReturnsOnCall map[int]struct {
		result1 resources.DeleteDataStreamResponse
	}
	GetDataStreamStub        func(context.Context, resources.GetDataStreamRequest) resources.GetDataStreamResponse
	getDataStreamMutex       sync.RWMutex
	getDataStreamArgsForCall []struct {
		arg1 context.Context
		arg2 resources.GetDataStreamRequest
	}
	getDataStreamReturns struct {
		result1 resources.GetDataStreamResponse
	}
	getDataStreamReturnsOnCall map[int]struct {
		result1 resources.GetDataStreamResponse
	}
	ListDataStreamsStub        func(context.Context, resources.ListDataStreamsRequest) resources.ListDataStreamsResponse
	listDataStreamsMutex       sync.RWMutex
	listDataStreamsArgsForCall []struct {
		arg1 context.Context
		arg2 resources.ListDataStreamsRequest
	}
	listDataStreamsReturns struct {
		result1 resources.ListDataStreamsResponse
	}
	listDataStreamsReturnsOnCall map[int]struct {
		result1 resources.ListDataStreamsResponse
	}
	UpdateDataStreamStub        func(context.Context, resources.UpdateDataStreamRequest) resources.UpdateDataStreamResponse
	updateDataStreamMutex       sync.RWMutex
	updateDataStreamArgsForCall []struct {
		arg1 context.Context
		arg2 resources.UpdateDataStreamRequest
	}
	updateDataStreamReturns struct {
		result1 resources.UpdateDataStreamResponse
	}
	updateDataStreamReturnsOnCall map[int]struct {
		result1 resources.UpdateDataStreamResponse
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDataStreamController) CreateDataStream(arg1 context.Context, arg2 resources.CreateDataStreamRequest) resources.CreateDataStreamResponse {
	fake.createDataStreamMutex.Lock()
	ret, specificReturn := fake.createDataStreamReturnsOnCall[len(fake.createDataStreamArgsForCall)]
	fake.createDataStreamArgsForCall = append(fake.createDataStreamArgsForCall, struct {
		arg1 context.Context
		arg2 resources.CreateDataStreamRequest
	}{arg1, arg2})
	fake.recordInvocation("CreateDataStream", []interface{}{arg1, arg2})
	fake.createDataStreamMutex.Unlock()
	if fake.CreateDataStreamStub != nil {
		return fake.CreateDataStreamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createDataStreamReturns
	return fakeReturns.result1
}

func (fake *FakeDataStreamController) CreateDataStreamCallCount() int {
	fake.createDataStreamMutex.RLock()
	defer fake.createDataStreamMutex.RUnlock()
	return len(fake.createDataStreamArgsForCall)
}

func (fake *FakeDataStreamController) CreateDataStreamCalls(stub func(context.Context, resources.CreateDataStreamRequest) resources.CreateDataStreamResponse) {
	fake.createDataStreamMutex.Lock()
	defer fake.createDataStreamMutex.Unlock()
	fake.CreateDataStreamStub = stub
}

func (fake *FakeDataStreamController) CreateDataStreamArgsForCall(i int) (context.Context, resources.CreateDataStreamRequest) {
	fake.createDataStreamMutex.RLock()
	defer fake.createDataStreamMutex.RUnlock()
	argsForCall := fake.createDataStreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDataStreamController) CreateDataStreamReturns(result1 resources.CreateDataStreamResponse) {
	fake.createDataStreamMutex.Lock()
	defer fake.createDataStreamMutex.Unlock()
	fake.CreateDataStreamStub = nil
	fake.createDataStreamReturns = struct {
		result1 resources.CreateDataStreamResponse
	}{result1}
}

func (fake *FakeDataStreamController) CreateDataStreamReturnsOnCall(i int, result1 resources.CreateDataStreamResponse) {
	fake.createDataStreamMutex.Lock()
	defer fake.createDataStreamMutex.Unlock()
	fake.CreateDataStreamStub = nil
	if fake.createDataStreamReturnsOnCall == nil {
		fake.createDataStreamReturnsOnCall = make(map[int]struct {
			result1 resources.CreateDataStreamResponse
		})
	}
	fake.createDataStreamReturnsOnCall[i] = struct {
		result1 resources.CreateDataStreamResponse
	}{result1}
}

func (fake *FakeDataStreamController) DeleteDataStream(arg1 context.Context, arg2 resources.DeleteDataStreamRequest) resources.DeleteDataStreamResponse {
	fake.deleteDataStreamMutex.Lock()
	ret, specificReturn := fake.deleteDataStreamReturnsOnCall[len(fake.deleteDataStreamArgsForCall)]
	fake.deleteDataStreamArgsForCall = append(fake.deleteDataStreamArgsForCall, struct {
		arg1 context.Context
		arg2 resources.DeleteDataStreamRequest
	}{arg1, arg2})
	fake.recordInvocation("DeleteDataStream", []interface{}{arg1, arg2})
	fake.deleteDataStreamMutex.Unlock()
	if fake.DeleteDataStreamStub != nil {
		return fake.DeleteDataStreamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteDataStreamReturns
	return fakeReturns.result1
}

func (fake *FakeDataStreamController) DeleteDataStreamCallCount() int {
	fake.deleteDataStreamMutex.RLock()
	defer fake.deleteDataStreamMutex.RUnlock()
	return len(fake.deleteDataStreamArgsForCall)
}

func (fake *FakeDataStreamController) DeleteDataStreamCalls(stub func(context.Context, resources.DeleteDataStreamRequest) resources.DeleteDataStreamResponse) {
	fake.deleteDataStreamMutex.Lock()
	defer fake.deleteDataStreamMutex.Unlock()
	fake.DeleteDataStreamStub = stub
}

func (fake *FakeDataStreamController) DeleteDataStreamArgsForCall(i int) (context.Context, resources.DeleteDataStreamRequest) {
	fake.deleteDataStreamMutex.RLock()
	defer fake.deleteDataStreamMutex.RUnlock()
	argsForCall := fake.deleteDataStreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDataStreamController) DeleteDataStreamReturns(result1 resources.DeleteDataStreamResponse) {
	fake.deleteDataStreamMutex.Lock()
	defer fake.deleteDataStreamMutex.Unlock()
	fake.DeleteDataStreamStub = nil
	fake.deleteDataStreamReturns = struct {
		result1 resources.DeleteDataStreamResponse
	}{result1}
}

func (fake *FakeDataStreamController) DeleteDataStreamReturnsOnCall(i int, result1 resources.DeleteDataStreamResponse) {
	fake.deleteDataStreamMutex.Lock()
	defer fake.deleteDataStreamMutex.Unlock()
	fake.DeleteDataStreamStub = nil
	if fake.deleteDataStreamReturnsOnCall == nil {
		fake.deleteDataStreamReturnsOnCall = make(map[int]struct {
			result1 resources.DeleteDataStreamResponse
		})
	}
	fake.deleteDataStreamReturnsOnCall[i] = struct {
		result1 resources.DeleteDataStreamResponse
	}{result1}
}

func (fake *FakeDataStreamController) GetDataStream(arg1 context.Context, arg2 resources.GetDataStreamRequest) resources.GetDataStreamResponse {
	fake.getDataStreamMutex.Lock()
	ret, specificReturn := fake.getDataStreamReturnsOnCall[len(fake.getDataStreamArgsForCall)]
	fake.getDataStreamArgsForCall = append(fake.getDataStreamArgsForCall, struct {
		arg1 context.Context
		arg2 resources.GetDataStreamRequest
	}{arg1, arg2})
	fake.recordInvocation("GetDataStream", []interface{}{arg1, arg2})
	fake.getDataStreamMutex.Unlock()
	if fake.GetDataStreamStub != nil {
		return fake.GetDataStreamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getDataStreamReturns
	return fakeReturns.result1
}

func (fake *FakeDataStreamController) GetDataStreamCallCount() int {
	fake.getDataStreamMutex.RLock()
	defer fake.getDataStreamMutex.RUnlock()
	return len(fake.getDataStreamArgsForCall)
}

func (fake *FakeDataStreamController) GetDataStreamCalls(stub func(context.Context, resources.GetDataStreamRequest) resources.GetDataStreamResponse) {
	fake.getDataStreamMutex.Lock()
	defer fake.getDataStreamMutex.Unlock()
	fake.GetDataStreamStub = stub
}

func (fake *FakeDataStreamController) GetDataStreamArgsForCall(i int) (context.Context, resources.GetDataStreamRequest) {
	fake.getDataStreamMutex.RLock()
	defer fake.getDataStreamMutex.RUnlock()
	argsForCall := fake.getDataStreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDataStreamController) GetDataStreamReturns(result1 resources.GetDataStreamResponse) {
	fake.getDataStreamMutex.Lock()
	defer fake.getDataStreamMutex.Unlock()
	fake.GetDataStreamStub = nil
	fake.getDataStreamReturns = struct {
		result1 resources.GetDataStreamResponse
	}{result1}
}

func (fake *FakeDataStreamController) GetDataStreamReturnsOnCall(i int, result1 resources.GetDataStreamResponse) {
	fake.getDataStreamMutex.Lock()
	defer fake.getDataStreamMutex.Unlock()
	fake.GetDataStreamStub = nil
	if fake.getDataStreamReturnsOnCall == nil {
		fake.getDataStreamReturnsOnCall = make(map[int]struct {
			result1 resources.GetDataStreamResponse
		})
	}
	fake.getDataStreamReturnsOnCall[i] = struct {
		result1 resources.GetDataStreamResponse
	}{result1}
}

func (fake *FakeDataStreamController) ListDataStreams(arg1 context.Context, arg2 resources.ListDataStreamsRequest) resources.ListDataStreamsResponse {
	fake.listDataStreamsMutex.Lock()
	ret, specificReturn := fake.listDataStreamsReturnsOnCall[len(fake.listDataStreamsArgsForCall)]
	fake.listDataStreamsArgsForCall = append(fake.listDataStreamsArgsForCall, struct {
		arg1 context.Context
		arg2 resources.ListDataStreamsRequest
	}{arg1, arg2})
	fake.recordInvocation("ListDataStreams", []interface{}{arg1, arg2})
	fake.listDataStreamsMutex.Unlock()
	if fake.ListDataStreamsStub != nil {
		return fake.ListDataStreamsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.listDataStreamsReturns
	return fakeReturns.result1
}

func (fake *FakeDataStreamController) ListDataStreamsCallCount() int {
	fake.listDataStreamsMutex.RLock()
	defer fake.listDataStreamsMutex.RUnlock()
	return len(fake.listDataStreamsArgsForCall)
}

func (fake *FakeDataStreamController) ListDataStreamsCalls(stub func(context.Context, resources.ListDataStreamsRequest) resources.ListDataStreamsResponse) {
	fake.listDataStreamsMutex.Lock()
	defer fake.listDataStreamsMutex.Unlock()
	fake.ListDataStreamsStub = stub
}

func (fake *FakeDataStreamController) ListDataStreamsArgsForCall(i int) (context.Context, resources.ListDataStreamsRequest) {
	fake.listDataStreamsMutex.RLock()
	defer fake.listDataStreamsMutex.RUnlock()
	argsForCall := fake.listDataStreamsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDataStreamController) ListDataStreamsReturns(result1 resources.ListDataStreamsResponse) {
	fake.listDataStreamsMutex.Lock()
	defer fake.listDataStreamsMutex.Unlock()
	fake.ListDataStreamsStub = nil
	fake.listDataStreamsReturns = struct {
		result1 resources.ListDataStreamsResponse
	}{result1}
}

func (fake *FakeDataStreamController) ListDataStreamsReturnsOnCall(i int, result1 resources.ListDataStreamsResponse) {
	fake.listDataStreamsMutex.Lock()
	defer fake.listDataStreamsMutex.Unlock()
	fake.ListDataStreamsStub = nil
	if fake.listDataStreamsReturnsOnCall == nil {
		fake.listDataStreamsReturnsOnCall = make(map[int]struct {
			result1 resources.ListDataStreamsResponse
		})
	}
	fake.listDataStreamsReturnsOnCall[i] = struct {
		result1 resources.ListDataStreamsResponse
	}{result1}
}

func (fake *FakeDataStreamController) UpdateDataStream(arg1 context.Context, arg2 resources.UpdateDataStreamRequest) resources.UpdateDataStreamResponse {
	fake.updateDataStreamMutex.Lock()
	ret, specificReturn := fake.updateDataStreamReturnsOnCall[len(fake.updateDataStreamArgsForCall)]
	fake.updateDataStreamArgsForCall = append(fake.updateDataStreamArgsForCall, struct {
		arg1 context.Context
		arg2 resources.UpdateDataStreamRequest
	}{arg1, arg2})
	fake.recordInvocation("UpdateDataStream", []interface{}{arg1, arg2})
	fake.updateDataStreamMutex.Unlock()
	if fake.UpdateDataStreamStub != nil {
		return fake.UpdateDataStreamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateDataStreamReturns
	return fakeReturns.result1
}

func (fake *FakeDataStreamController) UpdateDataStreamCallCount() int {
	fake.updateDataStreamMutex.RLock()
	defer fake.updateDataStreamMutex.RUnlock()
	return len(fake.updateDataStreamArgsForCall)
}

func (fake *FakeDataStreamController) UpdateDataStreamCalls(stub func(context.Context, resources.UpdateDataStreamRequest) resources.UpdateDataStreamResponse) {
	fake.updateDataStreamMutex.Lock()
	defer fake.updateDataStreamMutex.Unlock()
	fake.UpdateDataStreamStub = stub
}

func (fake *FakeDataStreamController) UpdateDataStreamArgsForCall(i int) (context.Context, resources.UpdateDataStreamRequest) {
	fake.updateDataStreamMutex.RLock()
	defer fake.updateDataStreamMutex.RUnlock()
	argsForCall := fake.updateDataStreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDataStreamController) UpdateDataStreamReturns(result1 resources.UpdateDataStreamResponse) {
	fake.updateDataStreamMutex.Lock()
	defer fake.updateDataStreamMutex.Unlock()
	fake.UpdateDataStreamStub = nil
	fake.updateDataStreamReturns = struct {
		result1 resources.UpdateDataStreamResponse
	}{result1}
}

func (fake *FakeDataStreamController) UpdateDataStreamReturnsOnCall(i int, result1 resources.UpdateDataStreamResponse) {
	fake.updateDataStreamMutex.Lock()
	defer fake.updateDataStreamMutex.Unlock()
	fake.UpdateDataStreamStub = nil
	if fake.updateDataStreamReturnsOnCall == nil {
		fake.updateDataStreamReturnsOnCall = make(map[int]struct {
			result1 resources.UpdateDataStreamResponse
		})
	}
	fake.updateDataStreamReturnsOnCall[i] = struct {
		result1 resources.UpdateDataStreamResponse
	}{result1}
}

func (fake *FakeDataStreamController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createDataStreamMutex.RLock()
	defer fake.createDataStreamMutex.RUnlock()
	fake.deleteDataStreamMutex.RLock()
	defer fake.deleteDataStreamMutex.RUnlock()
	fake.getDataStreamMutex.RLock()
	defer fake.getDataStreamMutex.RUnlock()
	fake.listDataStreamsMutex.RLock()
	defer fake.listDataStreamsMutex.RUnlock()
	fake.updateDataStreamMutex.RLock()
	defer fake.updateDataStreamMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDataStreamController) recordInvocation(key string, args []interface{}) {
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

var _ resources.DataStreamController = new(FakeDataStreamController)
