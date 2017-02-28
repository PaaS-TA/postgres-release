// This file was generated by counterfeiter
package directorfakes

import (
	"io"
	"sync"

	"github.com/cloudfoundry/bosh-cli/director"
	"github.com/cloudfoundry/bosh-cli/ui"
)

type FakeFileReporter struct {
	TrackUploadStub        func(int64, io.ReadCloser) ui.ReadSeekCloser
	trackUploadMutex       sync.RWMutex
	trackUploadArgsForCall []struct {
		arg1 int64
		arg2 io.ReadCloser
	}
	trackUploadReturns struct {
		result1 ui.ReadSeekCloser
	}
	TrackDownloadStub        func(int64, io.Writer) io.Writer
	trackDownloadMutex       sync.RWMutex
	trackDownloadArgsForCall []struct {
		arg1 int64
		arg2 io.Writer
	}
	trackDownloadReturns struct {
		result1 io.Writer
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFileReporter) TrackUpload(arg1 int64, arg2 io.ReadCloser) ui.ReadSeekCloser {
	fake.trackUploadMutex.Lock()
	fake.trackUploadArgsForCall = append(fake.trackUploadArgsForCall, struct {
		arg1 int64
		arg2 io.ReadCloser
	}{arg1, arg2})
	fake.recordInvocation("TrackUpload", []interface{}{arg1, arg2})
	fake.trackUploadMutex.Unlock()
	if fake.TrackUploadStub != nil {
		return fake.TrackUploadStub(arg1, arg2)
	}
	return fake.trackUploadReturns.result1
}

func (fake *FakeFileReporter) TrackUploadCallCount() int {
	fake.trackUploadMutex.RLock()
	defer fake.trackUploadMutex.RUnlock()
	return len(fake.trackUploadArgsForCall)
}

func (fake *FakeFileReporter) TrackUploadArgsForCall(i int) (int64, io.ReadCloser) {
	fake.trackUploadMutex.RLock()
	defer fake.trackUploadMutex.RUnlock()
	return fake.trackUploadArgsForCall[i].arg1, fake.trackUploadArgsForCall[i].arg2
}

func (fake *FakeFileReporter) TrackUploadReturns(result1 ui.ReadSeekCloser) {
	fake.TrackUploadStub = nil
	fake.trackUploadReturns = struct {
		result1 ui.ReadSeekCloser
	}{result1}
}

func (fake *FakeFileReporter) TrackDownload(arg1 int64, arg2 io.Writer) io.Writer {
	fake.trackDownloadMutex.Lock()
	fake.trackDownloadArgsForCall = append(fake.trackDownloadArgsForCall, struct {
		arg1 int64
		arg2 io.Writer
	}{arg1, arg2})
	fake.recordInvocation("TrackDownload", []interface{}{arg1, arg2})
	fake.trackDownloadMutex.Unlock()
	if fake.TrackDownloadStub != nil {
		return fake.TrackDownloadStub(arg1, arg2)
	}
	return fake.trackDownloadReturns.result1
}

func (fake *FakeFileReporter) TrackDownloadCallCount() int {
	fake.trackDownloadMutex.RLock()
	defer fake.trackDownloadMutex.RUnlock()
	return len(fake.trackDownloadArgsForCall)
}

func (fake *FakeFileReporter) TrackDownloadArgsForCall(i int) (int64, io.Writer) {
	fake.trackDownloadMutex.RLock()
	defer fake.trackDownloadMutex.RUnlock()
	return fake.trackDownloadArgsForCall[i].arg1, fake.trackDownloadArgsForCall[i].arg2
}

func (fake *FakeFileReporter) TrackDownloadReturns(result1 io.Writer) {
	fake.TrackDownloadStub = nil
	fake.trackDownloadReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeFileReporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.trackUploadMutex.RLock()
	defer fake.trackUploadMutex.RUnlock()
	fake.trackDownloadMutex.RLock()
	defer fake.trackDownloadMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFileReporter) recordInvocation(key string, args []interface{}) {
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

var _ director.FileReporter = new(FakeFileReporter)