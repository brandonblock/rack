package provider

import io "io"
import mock "github.com/stretchr/testify/mock"
import structs "github.com/convox/rack/api/structs"

// MockProvider is an autogenerated mock type for the Provider type
type MockProvider struct {
	mock.Mock
}

// AppDelete provides a mock function with given fields: name
func (_m *MockProvider) AppDelete(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppGet provides a mock function with given fields: name
func (_m *MockProvider) AppGet(name string) (*structs.App, error) {
	ret := _m.Called(name)

	var r0 *structs.App
	if rf, ok := ret.Get(0).(func(string) *structs.App); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.App)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildCreate provides a mock function with given fields: app, method, source, opts
func (_m *MockProvider) BuildCreate(app string, method string, source string, opts structs.BuildOptions) (*structs.Build, error) {
	ret := _m.Called(app, method, source, opts)

	var r0 *structs.Build
	if rf, ok := ret.Get(0).(func(string, string, string, structs.BuildOptions) *structs.Build); ok {
		r0 = rf(app, method, source, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, structs.BuildOptions) error); ok {
		r1 = rf(app, method, source, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildDelete provides a mock function with given fields: app, id
func (_m *MockProvider) BuildDelete(app string, id string) (*structs.Build, error) {
	ret := _m.Called(app, id)

	var r0 *structs.Build
	if rf, ok := ret.Get(0).(func(string, string) *structs.Build); ok {
		r0 = rf(app, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildExport provides a mock function with given fields: app, id, w
func (_m *MockProvider) BuildExport(app string, id string, w io.Writer) error {
	ret := _m.Called(app, id, w)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, io.Writer) error); ok {
		r0 = rf(app, id, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BuildGet provides a mock function with given fields: app, id
func (_m *MockProvider) BuildGet(app string, id string) (*structs.Build, error) {
	ret := _m.Called(app, id)

	var r0 *structs.Build
	if rf, ok := ret.Get(0).(func(string, string) *structs.Build); ok {
		r0 = rf(app, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildImport provides a mock function with given fields: app, r
func (_m *MockProvider) BuildImport(app string, r io.Reader) (*structs.Build, error) {
	ret := _m.Called(app, r)

	var r0 *structs.Build
	if rf, ok := ret.Get(0).(func(string, io.Reader) *structs.Build); ok {
		r0 = rf(app, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, io.Reader) error); ok {
		r1 = rf(app, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildList provides a mock function with given fields: app, limit
func (_m *MockProvider) BuildList(app string, limit int64) (structs.Builds, error) {
	ret := _m.Called(app, limit)

	var r0 structs.Builds
	if rf, ok := ret.Get(0).(func(string, int64) structs.Builds); ok {
		r0 = rf(app, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(structs.Builds)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(app, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildLogs provides a mock function with given fields: app, id, w
func (_m *MockProvider) BuildLogs(app string, id string, w io.Writer) error {
	ret := _m.Called(app, id, w)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, io.Writer) error); ok {
		r0 = rf(app, id, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BuildRelease provides a mock function with given fields: _a0
func (_m *MockProvider) BuildRelease(_a0 *structs.Build) (*structs.Release, error) {
	ret := _m.Called(_a0)

	var r0 *structs.Release
	if rf, ok := ret.Get(0).(func(*structs.Build) *structs.Release); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Release)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*structs.Build) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildSave provides a mock function with given fields: _a0
func (_m *MockProvider) BuildSave(_a0 *structs.Build) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*structs.Build) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CapacityGet provides a mock function with given fields:
func (_m *MockProvider) CapacityGet() (*structs.Capacity, error) {
	ret := _m.Called()

	var r0 *structs.Capacity
	if rf, ok := ret.Get(0).(func() *structs.Capacity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Capacity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CertificateCreate provides a mock function with given fields: pub, key, chain
func (_m *MockProvider) CertificateCreate(pub string, key string, chain string) (*structs.Certificate, error) {
	ret := _m.Called(pub, key, chain)

	var r0 *structs.Certificate
	if rf, ok := ret.Get(0).(func(string, string, string) *structs.Certificate); ok {
		r0 = rf(pub, key, chain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Certificate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(pub, key, chain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CertificateDelete provides a mock function with given fields: id
func (_m *MockProvider) CertificateDelete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CertificateGenerate provides a mock function with given fields: domains
func (_m *MockProvider) CertificateGenerate(domains []string) (*structs.Certificate, error) {
	ret := _m.Called(domains)

	var r0 *structs.Certificate
	if rf, ok := ret.Get(0).(func([]string) *structs.Certificate); ok {
		r0 = rf(domains)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Certificate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(domains)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CertificateList provides a mock function with given fields:
func (_m *MockProvider) CertificateList() (structs.Certificates, error) {
	ret := _m.Called()

	var r0 structs.Certificates
	if rf, ok := ret.Get(0).(func() structs.Certificates); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(structs.Certificates)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnvironmentGet provides a mock function with given fields: app
func (_m *MockProvider) EnvironmentGet(app string) (structs.Environment, error) {
	ret := _m.Called(app)

	var r0 structs.Environment
	if rf, ok := ret.Get(0).(func(string) structs.Environment); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(structs.Environment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventSend provides a mock function with given fields: _a0, _a1
func (_m *MockProvider) EventSend(_a0 *structs.Event, _a1 error) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*structs.Event, error) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FormationGet provides a mock function with given fields: app, process
func (_m *MockProvider) FormationGet(app string, process string) (*structs.ProcessFormation, error) {
	ret := _m.Called(app, process)

	var r0 *structs.ProcessFormation
	if rf, ok := ret.Get(0).(func(string, string) *structs.ProcessFormation); ok {
		r0 = rf(app, process)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.ProcessFormation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, process)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FormationList provides a mock function with given fields: app
func (_m *MockProvider) FormationList(app string) (structs.Formation, error) {
	ret := _m.Called(app)

	var r0 structs.Formation
	if rf, ok := ret.Get(0).(func(string) structs.Formation); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(structs.Formation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FormationSave provides a mock function with given fields: app, pf
func (_m *MockProvider) FormationSave(app string, pf *structs.ProcessFormation) error {
	ret := _m.Called(app, pf)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *structs.ProcessFormation) error); ok {
		r0 = rf(app, pf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IndexDiff provides a mock function with given fields: _a0
func (_m *MockProvider) IndexDiff(_a0 *structs.Index) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*structs.Index) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*structs.Index) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexDownload provides a mock function with given fields: _a0, _a1
func (_m *MockProvider) IndexDownload(_a0 *structs.Index, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*structs.Index, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IndexUpload provides a mock function with given fields: _a0, _a1
func (_m *MockProvider) IndexUpload(_a0 string, _a1 []byte) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Initialize provides a mock function with given fields: opts
func (_m *MockProvider) Initialize(opts structs.ProviderOptions) error {
	ret := _m.Called(opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(structs.ProviderOptions) error); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InstanceList provides a mock function with given fields:
func (_m *MockProvider) InstanceList() (structs.Instances, error) {
	ret := _m.Called()

	var r0 structs.Instances
	if rf, ok := ret.Get(0).(func() structs.Instances); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(structs.Instances)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstanceTerminate provides a mock function with given fields: id
func (_m *MockProvider) InstanceTerminate(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LogStream provides a mock function with given fields: app, w, opts
func (_m *MockProvider) LogStream(app string, w io.Writer, opts structs.LogStreamOptions) error {
	ret := _m.Called(app, w, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, io.Writer, structs.LogStreamOptions) error); ok {
		r0 = rf(app, w, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ObjectDelete provides a mock function with given fields: key
func (_m *MockProvider) ObjectDelete(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ObjectExists provides a mock function with given fields: key
func (_m *MockProvider) ObjectExists(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ObjectFetch provides a mock function with given fields: key
func (_m *MockProvider) ObjectFetch(key string) (io.ReadCloser, error) {
	ret := _m.Called(key)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string) io.ReadCloser); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ObjectList provides a mock function with given fields: prefix
func (_m *MockProvider) ObjectList(prefix string) ([]string, error) {
	ret := _m.Called(prefix)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ObjectStore provides a mock function with given fields: key, r, opts
func (_m *MockProvider) ObjectStore(key string, r io.Reader, opts structs.ObjectOptions) (string, error) {
	ret := _m.Called(key, r, opts)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, io.Reader, structs.ObjectOptions) string); ok {
		r0 = rf(key, r, opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, io.Reader, structs.ObjectOptions) error); ok {
		r1 = rf(key, r, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessExec provides a mock function with given fields: app, pid, command, stream, opts
func (_m *MockProvider) ProcessExec(app string, pid string, command string, stream io.ReadWriter, opts structs.ProcessExecOptions) error {
	ret := _m.Called(app, pid, command, stream, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, io.ReadWriter, structs.ProcessExecOptions) error); ok {
		r0 = rf(app, pid, command, stream, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessList provides a mock function with given fields: app
func (_m *MockProvider) ProcessList(app string) (structs.Processes, error) {
	ret := _m.Called(app)

	var r0 structs.Processes
	if rf, ok := ret.Get(0).(func(string) structs.Processes); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(structs.Processes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessRun provides a mock function with given fields: app, process, opts
func (_m *MockProvider) ProcessRun(app string, process string, opts structs.ProcessRunOptions) (string, error) {
	ret := _m.Called(app, process, opts)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, structs.ProcessRunOptions) string); ok {
		r0 = rf(app, process, opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, structs.ProcessRunOptions) error); ok {
		r1 = rf(app, process, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessStop provides a mock function with given fields: app, pid
func (_m *MockProvider) ProcessStop(app string, pid string) error {
	ret := _m.Called(app, pid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(app, pid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegistryAdd provides a mock function with given fields: server, username, password
func (_m *MockProvider) RegistryAdd(server string, username string, password string) (*structs.Registry, error) {
	ret := _m.Called(server, username, password)

	var r0 *structs.Registry
	if rf, ok := ret.Get(0).(func(string, string, string) *structs.Registry); ok {
		r0 = rf(server, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Registry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(server, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistryDelete provides a mock function with given fields: server
func (_m *MockProvider) RegistryDelete(server string) error {
	ret := _m.Called(server)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(server)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegistryList provides a mock function with given fields:
func (_m *MockProvider) RegistryList() (structs.Registries, error) {
	ret := _m.Called()

	var r0 structs.Registries
	if rf, ok := ret.Get(0).(func() structs.Registries); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(structs.Registries)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseDelete provides a mock function with given fields: app, buildID
func (_m *MockProvider) ReleaseDelete(app string, buildID string) error {
	ret := _m.Called(app, buildID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(app, buildID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReleaseGet provides a mock function with given fields: app, id
func (_m *MockProvider) ReleaseGet(app string, id string) (*structs.Release, error) {
	ret := _m.Called(app, id)

	var r0 *structs.Release
	if rf, ok := ret.Get(0).(func(string, string) *structs.Release); ok {
		r0 = rf(app, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Release)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseList provides a mock function with given fields: app, limit
func (_m *MockProvider) ReleaseList(app string, limit int64) (structs.Releases, error) {
	ret := _m.Called(app, limit)

	var r0 structs.Releases
	if rf, ok := ret.Get(0).(func(string, int64) structs.Releases); ok {
		r0 = rf(app, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(structs.Releases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(app, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleasePromote provides a mock function with given fields: app, id
func (_m *MockProvider) ReleasePromote(app string, id string) (*structs.Release, error) {
	ret := _m.Called(app, id)

	var r0 *structs.Release
	if rf, ok := ret.Get(0).(func(string, string) *structs.Release); ok {
		r0 = rf(app, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Release)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseSave provides a mock function with given fields: _a0
func (_m *MockProvider) ReleaseSave(_a0 *structs.Release) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*structs.Release) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServiceCreate provides a mock function with given fields: name, kind, params
func (_m *MockProvider) ServiceCreate(name string, kind string, params map[string]string) (*structs.Service, error) {
	ret := _m.Called(name, kind, params)

	var r0 *structs.Service
	if rf, ok := ret.Get(0).(func(string, string, map[string]string) *structs.Service); ok {
		r0 = rf(name, kind, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, map[string]string) error); ok {
		r1 = rf(name, kind, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceDelete provides a mock function with given fields: name
func (_m *MockProvider) ServiceDelete(name string) (*structs.Service, error) {
	ret := _m.Called(name)

	var r0 *structs.Service
	if rf, ok := ret.Get(0).(func(string) *structs.Service); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceGet provides a mock function with given fields: name
func (_m *MockProvider) ServiceGet(name string) (*structs.Service, error) {
	ret := _m.Called(name)

	var r0 *structs.Service
	if rf, ok := ret.Get(0).(func(string) *structs.Service); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceLink provides a mock function with given fields: name, app, process
func (_m *MockProvider) ServiceLink(name string, app string, process string) (*structs.Service, error) {
	ret := _m.Called(name, app, process)

	var r0 *structs.Service
	if rf, ok := ret.Get(0).(func(string, string, string) *structs.Service); ok {
		r0 = rf(name, app, process)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(name, app, process)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceList provides a mock function with given fields:
func (_m *MockProvider) ServiceList() (structs.Services, error) {
	ret := _m.Called()

	var r0 structs.Services
	if rf, ok := ret.Get(0).(func() structs.Services); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(structs.Services)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceUnlink provides a mock function with given fields: name, app, process
func (_m *MockProvider) ServiceUnlink(name string, app string, process string) (*structs.Service, error) {
	ret := _m.Called(name, app, process)

	var r0 *structs.Service
	if rf, ok := ret.Get(0).(func(string, string, string) *structs.Service); ok {
		r0 = rf(name, app, process)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(name, app, process)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceUpdate provides a mock function with given fields: name, params
func (_m *MockProvider) ServiceUpdate(name string, params map[string]string) (*structs.Service, error) {
	ret := _m.Called(name, params)

	var r0 *structs.Service
	if rf, ok := ret.Get(0).(func(string, map[string]string) *structs.Service); ok {
		r0 = rf(name, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]string) error); ok {
		r1 = rf(name, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemGet provides a mock function with given fields:
func (_m *MockProvider) SystemGet() (*structs.System, error) {
	ret := _m.Called()

	var r0 *structs.System
	if rf, ok := ret.Get(0).(func() *structs.System); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.System)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemLogs provides a mock function with given fields: w, opts
func (_m *MockProvider) SystemLogs(w io.Writer, opts structs.LogStreamOptions) error {
	ret := _m.Called(w, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Writer, structs.LogStreamOptions) error); ok {
		r0 = rf(w, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SystemProcesses provides a mock function with given fields:
func (_m *MockProvider) SystemProcesses() (structs.Processes, error) {
	ret := _m.Called()

	var r0 structs.Processes
	if rf, ok := ret.Get(0).(func() structs.Processes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(structs.Processes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemReleases provides a mock function with given fields:
func (_m *MockProvider) SystemReleases() (structs.Releases, error) {
	ret := _m.Called()

	var r0 structs.Releases
	if rf, ok := ret.Get(0).(func() structs.Releases); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(structs.Releases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemSave provides a mock function with given fields: system
func (_m *MockProvider) SystemSave(system structs.System) error {
	ret := _m.Called(system)

	var r0 error
	if rf, ok := ret.Get(0).(func(structs.System) error); ok {
		r0 = rf(system)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ Provider = (*MockProvider)(nil)
