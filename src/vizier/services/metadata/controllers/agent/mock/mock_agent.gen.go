// Code generated by MockGen. DO NOT EDIT.
// Source: agent.go

// Package mock_agent is a generated GoMock package.
package mock_agent

import (
	reflect "reflect"

	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
	pl_shared_k8s_metadatapb "pixielabs.ai/pixielabs/src/shared/k8s/metadatapb"
	types "pixielabs.ai/pixielabs/src/shared/types/go"
	messages "pixielabs.ai/pixielabs/src/vizier/messages/messagespb"
	agent "pixielabs.ai/pixielabs/src/vizier/services/metadata/controllers/agent"
	metadatapb "pixielabs.ai/pixielabs/src/vizier/services/metadata/metadatapb"
	storepb "pixielabs.ai/pixielabs/src/vizier/services/metadata/storepb"
	pl_vizier_services_shared_agent "pixielabs.ai/pixielabs/src/vizier/services/shared/agentpb"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAgent mocks base method.
func (m *MockStore) CreateAgent(agentID uuid.UUID, a *pl_vizier_services_shared_agent.Agent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAgent", agentID, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAgent indicates an expected call of CreateAgent.
func (mr *MockStoreMockRecorder) CreateAgent(agentID, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAgent", reflect.TypeOf((*MockStore)(nil).CreateAgent), agentID, a)
}

// DeleteAgent mocks base method.
func (m *MockStore) DeleteAgent(agentID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAgent", agentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAgent indicates an expected call of DeleteAgent.
func (mr *MockStoreMockRecorder) DeleteAgent(agentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgent", reflect.TypeOf((*MockStore)(nil).DeleteAgent), agentID)
}

// GetASID mocks base method.
func (m *MockStore) GetASID() (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetASID")
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetASID indicates an expected call of GetASID.
func (mr *MockStoreMockRecorder) GetASID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetASID", reflect.TypeOf((*MockStore)(nil).GetASID))
}

// GetAgent mocks base method.
func (m *MockStore) GetAgent(agentID uuid.UUID) (*pl_vizier_services_shared_agent.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgent", agentID)
	ret0, _ := ret[0].(*pl_vizier_services_shared_agent.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgent indicates an expected call of GetAgent.
func (mr *MockStoreMockRecorder) GetAgent(agentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgent", reflect.TypeOf((*MockStore)(nil).GetAgent), agentID)
}

// GetAgentIDForHostnamePair mocks base method.
func (m *MockStore) GetAgentIDForHostnamePair(hnPair *agent.HostnameIPPair) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentIDForHostnamePair", hnPair)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentIDForHostnamePair indicates an expected call of GetAgentIDForHostnamePair.
func (mr *MockStoreMockRecorder) GetAgentIDForHostnamePair(hnPair interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentIDForHostnamePair", reflect.TypeOf((*MockStore)(nil).GetAgentIDForHostnamePair), hnPair)
}

// GetAgentIDFromPodName mocks base method.
func (m *MockStore) GetAgentIDFromPodName(podName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentIDFromPodName", podName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentIDFromPodName indicates an expected call of GetAgentIDFromPodName.
func (mr *MockStoreMockRecorder) GetAgentIDFromPodName(podName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentIDFromPodName", reflect.TypeOf((*MockStore)(nil).GetAgentIDFromPodName), podName)
}

// GetAgents mocks base method.
func (m *MockStore) GetAgents() ([]*pl_vizier_services_shared_agent.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgents")
	ret0, _ := ret[0].([]*pl_vizier_services_shared_agent.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgents indicates an expected call of GetAgents.
func (mr *MockStoreMockRecorder) GetAgents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgents", reflect.TypeOf((*MockStore)(nil).GetAgents))
}

// GetAgentsDataInfo mocks base method.
func (m *MockStore) GetAgentsDataInfo() (map[uuid.UUID]*messages.AgentDataInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentsDataInfo")
	ret0, _ := ret[0].(map[uuid.UUID]*messages.AgentDataInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentsDataInfo indicates an expected call of GetAgentsDataInfo.
func (mr *MockStoreMockRecorder) GetAgentsDataInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentsDataInfo", reflect.TypeOf((*MockStore)(nil).GetAgentsDataInfo))
}

// GetComputedSchema mocks base method.
func (m *MockStore) GetComputedSchema() (*storepb.ComputedSchema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComputedSchema")
	ret0, _ := ret[0].(*storepb.ComputedSchema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComputedSchema indicates an expected call of GetComputedSchema.
func (mr *MockStoreMockRecorder) GetComputedSchema() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComputedSchema", reflect.TypeOf((*MockStore)(nil).GetComputedSchema))
}

// GetProcesses mocks base method.
func (m *MockStore) GetProcesses(upids []*types.UInt128) ([]*pl_shared_k8s_metadatapb.ProcessInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProcesses", upids)
	ret0, _ := ret[0].([]*pl_shared_k8s_metadatapb.ProcessInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProcesses indicates an expected call of GetProcesses.
func (mr *MockStoreMockRecorder) GetProcesses(upids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcesses", reflect.TypeOf((*MockStore)(nil).GetProcesses), upids)
}

// PruneComputedSchema mocks base method.
func (m *MockStore) PruneComputedSchema() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PruneComputedSchema")
	ret0, _ := ret[0].(error)
	return ret0
}

// PruneComputedSchema indicates an expected call of PruneComputedSchema.
func (mr *MockStoreMockRecorder) PruneComputedSchema() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PruneComputedSchema", reflect.TypeOf((*MockStore)(nil).PruneComputedSchema))
}

// UpdateAgent mocks base method.
func (m *MockStore) UpdateAgent(agentID uuid.UUID, a *pl_vizier_services_shared_agent.Agent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAgent", agentID, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAgent indicates an expected call of UpdateAgent.
func (mr *MockStoreMockRecorder) UpdateAgent(agentID, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAgent", reflect.TypeOf((*MockStore)(nil).UpdateAgent), agentID, a)
}

// UpdateAgentDataInfo mocks base method.
func (m *MockStore) UpdateAgentDataInfo(agentID uuid.UUID, dataInfo *messages.AgentDataInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAgentDataInfo", agentID, dataInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAgentDataInfo indicates an expected call of UpdateAgentDataInfo.
func (mr *MockStoreMockRecorder) UpdateAgentDataInfo(agentID, dataInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAgentDataInfo", reflect.TypeOf((*MockStore)(nil).UpdateAgentDataInfo), agentID, dataInfo)
}

// UpdateProcesses mocks base method.
func (m *MockStore) UpdateProcesses(processes []*pl_shared_k8s_metadatapb.ProcessInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProcesses", processes)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProcesses indicates an expected call of UpdateProcesses.
func (mr *MockStoreMockRecorder) UpdateProcesses(processes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProcesses", reflect.TypeOf((*MockStore)(nil).UpdateProcesses), processes)
}

// UpdateSchemas mocks base method.
func (m *MockStore) UpdateSchemas(agentID uuid.UUID, schemas []*storepb.TableInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSchemas", agentID, schemas)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSchemas indicates an expected call of UpdateSchemas.
func (mr *MockStoreMockRecorder) UpdateSchemas(agentID, schemas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSchemas", reflect.TypeOf((*MockStore)(nil).UpdateSchemas), agentID, schemas)
}

// MockCIDRInfoProvider is a mock of CIDRInfoProvider interface.
type MockCIDRInfoProvider struct {
	ctrl     *gomock.Controller
	recorder *MockCIDRInfoProviderMockRecorder
}

// MockCIDRInfoProviderMockRecorder is the mock recorder for MockCIDRInfoProvider.
type MockCIDRInfoProviderMockRecorder struct {
	mock *MockCIDRInfoProvider
}

// NewMockCIDRInfoProvider creates a new mock instance.
func NewMockCIDRInfoProvider(ctrl *gomock.Controller) *MockCIDRInfoProvider {
	mock := &MockCIDRInfoProvider{ctrl: ctrl}
	mock.recorder = &MockCIDRInfoProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCIDRInfoProvider) EXPECT() *MockCIDRInfoProviderMockRecorder {
	return m.recorder
}

// GetPodCIDRs mocks base method.
func (m *MockCIDRInfoProvider) GetPodCIDRs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodCIDRs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetPodCIDRs indicates an expected call of GetPodCIDRs.
func (mr *MockCIDRInfoProviderMockRecorder) GetPodCIDRs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodCIDRs", reflect.TypeOf((*MockCIDRInfoProvider)(nil).GetPodCIDRs))
}

// GetServiceCIDR mocks base method.
func (m *MockCIDRInfoProvider) GetServiceCIDR() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceCIDR")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServiceCIDR indicates an expected call of GetServiceCIDR.
func (mr *MockCIDRInfoProviderMockRecorder) GetServiceCIDR() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceCIDR", reflect.TypeOf((*MockCIDRInfoProvider)(nil).GetServiceCIDR))
}

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// ApplyAgentUpdate mocks base method.
func (m *MockManager) ApplyAgentUpdate(update *agent.Update) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyAgentUpdate", update)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyAgentUpdate indicates an expected call of ApplyAgentUpdate.
func (mr *MockManagerMockRecorder) ApplyAgentUpdate(update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyAgentUpdate", reflect.TypeOf((*MockManager)(nil).ApplyAgentUpdate), update)
}

// DeleteAgent mocks base method.
func (m *MockManager) DeleteAgent(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAgent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAgent indicates an expected call of DeleteAgent.
func (mr *MockManagerMockRecorder) DeleteAgent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgent", reflect.TypeOf((*MockManager)(nil).DeleteAgent), arg0)
}

// DeleteAgentUpdateCursor mocks base method.
func (m *MockManager) DeleteAgentUpdateCursor(cursorID uuid.UUID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteAgentUpdateCursor", cursorID)
}

// DeleteAgentUpdateCursor indicates an expected call of DeleteAgentUpdateCursor.
func (mr *MockManagerMockRecorder) DeleteAgentUpdateCursor(cursorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgentUpdateCursor", reflect.TypeOf((*MockManager)(nil).DeleteAgentUpdateCursor), cursorID)
}

// GetActiveAgents mocks base method.
func (m *MockManager) GetActiveAgents() ([]*pl_vizier_services_shared_agent.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveAgents")
	ret0, _ := ret[0].([]*pl_vizier_services_shared_agent.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveAgents indicates an expected call of GetActiveAgents.
func (mr *MockManagerMockRecorder) GetActiveAgents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveAgents", reflect.TypeOf((*MockManager)(nil).GetActiveAgents))
}

// GetAgentIDForHostnamePair mocks base method.
func (m *MockManager) GetAgentIDForHostnamePair(hnPair *agent.HostnameIPPair) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentIDForHostnamePair", hnPair)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentIDForHostnamePair indicates an expected call of GetAgentIDForHostnamePair.
func (mr *MockManagerMockRecorder) GetAgentIDForHostnamePair(hnPair interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentIDForHostnamePair", reflect.TypeOf((*MockManager)(nil).GetAgentIDForHostnamePair), hnPair)
}

// GetAgentUpdates mocks base method.
func (m *MockManager) GetAgentUpdates(cursorID uuid.UUID) ([]*metadatapb.AgentUpdate, *storepb.ComputedSchema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentUpdates", cursorID)
	ret0, _ := ret[0].([]*metadatapb.AgentUpdate)
	ret1, _ := ret[1].(*storepb.ComputedSchema)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAgentUpdates indicates an expected call of GetAgentUpdates.
func (mr *MockManagerMockRecorder) GetAgentUpdates(cursorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentUpdates", reflect.TypeOf((*MockManager)(nil).GetAgentUpdates), cursorID)
}

// GetComputedSchema mocks base method.
func (m *MockManager) GetComputedSchema() (*storepb.ComputedSchema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComputedSchema")
	ret0, _ := ret[0].(*storepb.ComputedSchema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComputedSchema indicates an expected call of GetComputedSchema.
func (mr *MockManagerMockRecorder) GetComputedSchema() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComputedSchema", reflect.TypeOf((*MockManager)(nil).GetComputedSchema))
}

// GetPodCIDRs mocks base method.
func (m *MockManager) GetPodCIDRs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodCIDRs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetPodCIDRs indicates an expected call of GetPodCIDRs.
func (mr *MockManagerMockRecorder) GetPodCIDRs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodCIDRs", reflect.TypeOf((*MockManager)(nil).GetPodCIDRs))
}

// GetServiceCIDR mocks base method.
func (m *MockManager) GetServiceCIDR() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceCIDR")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServiceCIDR indicates an expected call of GetServiceCIDR.
func (mr *MockManagerMockRecorder) GetServiceCIDR() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceCIDR", reflect.TypeOf((*MockManager)(nil).GetServiceCIDR))
}

// MessageActiveAgents mocks base method.
func (m *MockManager) MessageActiveAgents(msg []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageActiveAgents", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// MessageActiveAgents indicates an expected call of MessageActiveAgents.
func (mr *MockManagerMockRecorder) MessageActiveAgents(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageActiveAgents", reflect.TypeOf((*MockManager)(nil).MessageActiveAgents), msg)
}

// MessageAgents mocks base method.
func (m *MockManager) MessageAgents(agentIDs []uuid.UUID, msg []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageAgents", agentIDs, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// MessageAgents indicates an expected call of MessageAgents.
func (mr *MockManagerMockRecorder) MessageAgents(agentIDs, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageAgents", reflect.TypeOf((*MockManager)(nil).MessageAgents), agentIDs, msg)
}

// NewAgentUpdateCursor mocks base method.
func (m *MockManager) NewAgentUpdateCursor() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAgentUpdateCursor")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// NewAgentUpdateCursor indicates an expected call of NewAgentUpdateCursor.
func (mr *MockManagerMockRecorder) NewAgentUpdateCursor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAgentUpdateCursor", reflect.TypeOf((*MockManager)(nil).NewAgentUpdateCursor))
}

// RegisterAgent mocks base method.
func (m *MockManager) RegisterAgent(info *pl_vizier_services_shared_agent.Agent) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAgent", info)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterAgent indicates an expected call of RegisterAgent.
func (mr *MockManagerMockRecorder) RegisterAgent(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAgent", reflect.TypeOf((*MockManager)(nil).RegisterAgent), info)
}

// UpdateConfig mocks base method.
func (m *MockManager) UpdateConfig(arg0, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfig indicates an expected call of UpdateConfig.
func (mr *MockManagerMockRecorder) UpdateConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfig", reflect.TypeOf((*MockManager)(nil).UpdateConfig), arg0, arg1, arg2, arg3)
}

// UpdateHeartbeat mocks base method.
func (m *MockManager) UpdateHeartbeat(agentID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHeartbeat", agentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHeartbeat indicates an expected call of UpdateHeartbeat.
func (mr *MockManagerMockRecorder) UpdateHeartbeat(agentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHeartbeat", reflect.TypeOf((*MockManager)(nil).UpdateHeartbeat), agentID)
}
