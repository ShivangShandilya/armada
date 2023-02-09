// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apache/pulsar-client-go/pulsar (interfaces: Client,Producer,Message)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	pulsar "github.com/apache/pulsar-client-go/pulsar"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// CreateProducer mocks base method.
func (m *MockClient) CreateProducer(arg0 pulsar.ProducerOptions) (pulsar.Producer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProducer", arg0)
	ret0, _ := ret[0].(pulsar.Producer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProducer indicates an expected call of CreateProducer.
func (mr *MockClientMockRecorder) CreateProducer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProducer", reflect.TypeOf((*MockClient)(nil).CreateProducer), arg0)
}

// CreateReader mocks base method.
func (m *MockClient) CreateReader(arg0 pulsar.ReaderOptions) (pulsar.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReader", arg0)
	ret0, _ := ret[0].(pulsar.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReader indicates an expected call of CreateReader.
func (mr *MockClientMockRecorder) CreateReader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReader", reflect.TypeOf((*MockClient)(nil).CreateReader), arg0)
}

// CreateTableView mocks base method.
func (m *MockClient) CreateTableView(arg0 pulsar.TableViewOptions) (pulsar.TableView, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTableView", arg0)
	ret0, _ := ret[0].(pulsar.TableView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTableView indicates an expected call of CreateTableView.
func (mr *MockClientMockRecorder) CreateTableView(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTableView", reflect.TypeOf((*MockClient)(nil).CreateTableView), arg0)
}

// Subscribe mocks base method.
func (m *MockClient) Subscribe(arg0 pulsar.ConsumerOptions) (pulsar.Consumer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(pulsar.Consumer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockClientMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockClient)(nil).Subscribe), arg0)
}

// TopicPartitions mocks base method.
func (m *MockClient) TopicPartitions(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopicPartitions", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopicPartitions indicates an expected call of TopicPartitions.
func (mr *MockClientMockRecorder) TopicPartitions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopicPartitions", reflect.TypeOf((*MockClient)(nil).TopicPartitions), arg0)
}

// MockProducer is a mock of Producer interface.
type MockProducer struct {
	ctrl     *gomock.Controller
	recorder *MockProducerMockRecorder
}

// MockProducerMockRecorder is the mock recorder for MockProducer.
type MockProducerMockRecorder struct {
	mock *MockProducer
}

// NewMockProducer creates a new mock instance.
func NewMockProducer(ctrl *gomock.Controller) *MockProducer {
	mock := &MockProducer{ctrl: ctrl}
	mock.recorder = &MockProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducer) EXPECT() *MockProducerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockProducer) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockProducerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockProducer)(nil).Close))
}

// Flush mocks base method.
func (m *MockProducer) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockProducerMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockProducer)(nil).Flush))
}

// LastSequenceID mocks base method.
func (m *MockProducer) LastSequenceID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastSequenceID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// LastSequenceID indicates an expected call of LastSequenceID.
func (mr *MockProducerMockRecorder) LastSequenceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastSequenceID", reflect.TypeOf((*MockProducer)(nil).LastSequenceID))
}

// Name mocks base method.
func (m *MockProducer) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockProducerMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockProducer)(nil).Name))
}

// Send mocks base method.
func (m *MockProducer) Send(arg0 context.Context, arg1 *pulsar.ProducerMessage) (pulsar.MessageID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(pulsar.MessageID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockProducerMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockProducer)(nil).Send), arg0, arg1)
}

// SendAsync mocks base method.
func (m *MockProducer) SendAsync(arg0 context.Context, arg1 *pulsar.ProducerMessage, arg2 func(pulsar.MessageID, *pulsar.ProducerMessage, error)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendAsync", arg0, arg1, arg2)
}

// SendAsync indicates an expected call of SendAsync.
func (mr *MockProducerMockRecorder) SendAsync(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAsync", reflect.TypeOf((*MockProducer)(nil).SendAsync), arg0, arg1, arg2)
}

// Topic mocks base method.
func (m *MockProducer) Topic() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Topic")
	ret0, _ := ret[0].(string)
	return ret0
}

// Topic indicates an expected call of Topic.
func (mr *MockProducerMockRecorder) Topic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Topic", reflect.TypeOf((*MockProducer)(nil).Topic))
}

// MockMessage is a mock of Message interface.
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMockRecorder
}

// MockMessageMockRecorder is the mock recorder for MockMessage.
type MockMessageMockRecorder struct {
	mock *MockMessage
}

// NewMockMessage creates a new mock instance.
func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &MockMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessage) EXPECT() *MockMessageMockRecorder {
	return m.recorder
}

// BrokerPublishTime mocks base method.
func (m *MockMessage) BrokerPublishTime() *time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BrokerPublishTime")
	ret0, _ := ret[0].(*time.Time)
	return ret0
}

// BrokerPublishTime indicates an expected call of BrokerPublishTime.
func (mr *MockMessageMockRecorder) BrokerPublishTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BrokerPublishTime", reflect.TypeOf((*MockMessage)(nil).BrokerPublishTime))
}

// EventTime mocks base method.
func (m *MockMessage) EventTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// EventTime indicates an expected call of EventTime.
func (mr *MockMessageMockRecorder) EventTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventTime", reflect.TypeOf((*MockMessage)(nil).EventTime))
}

// GetEncryptionContext mocks base method.
func (m *MockMessage) GetEncryptionContext() *pulsar.EncryptionContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEncryptionContext")
	ret0, _ := ret[0].(*pulsar.EncryptionContext)
	return ret0
}

// GetEncryptionContext indicates an expected call of GetEncryptionContext.
func (mr *MockMessageMockRecorder) GetEncryptionContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEncryptionContext", reflect.TypeOf((*MockMessage)(nil).GetEncryptionContext))
}

// GetReplicatedFrom mocks base method.
func (m *MockMessage) GetReplicatedFrom() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicatedFrom")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetReplicatedFrom indicates an expected call of GetReplicatedFrom.
func (mr *MockMessageMockRecorder) GetReplicatedFrom() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicatedFrom", reflect.TypeOf((*MockMessage)(nil).GetReplicatedFrom))
}

// GetSchemaValue mocks base method.
func (m *MockMessage) GetSchemaValue(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchemaValue", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetSchemaValue indicates an expected call of GetSchemaValue.
func (mr *MockMessageMockRecorder) GetSchemaValue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchemaValue", reflect.TypeOf((*MockMessage)(nil).GetSchemaValue), arg0)
}

// ID mocks base method.
func (m *MockMessage) ID() pulsar.MessageID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(pulsar.MessageID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockMessageMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockMessage)(nil).ID))
}

// Index mocks base method.
func (m *MockMessage) Index() *uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index")
	ret0, _ := ret[0].(*uint64)
	return ret0
}

// Index indicates an expected call of Index.
func (mr *MockMessageMockRecorder) Index() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockMessage)(nil).Index))
}

// IsReplicated mocks base method.
func (m *MockMessage) IsReplicated() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsReplicated")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsReplicated indicates an expected call of IsReplicated.
func (mr *MockMessageMockRecorder) IsReplicated() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsReplicated", reflect.TypeOf((*MockMessage)(nil).IsReplicated))
}

// Key mocks base method.
func (m *MockMessage) Key() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Key")
	ret0, _ := ret[0].(string)
	return ret0
}

// Key indicates an expected call of Key.
func (mr *MockMessageMockRecorder) Key() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockMessage)(nil).Key))
}

// OrderingKey mocks base method.
func (m *MockMessage) OrderingKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderingKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// OrderingKey indicates an expected call of OrderingKey.
func (mr *MockMessageMockRecorder) OrderingKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderingKey", reflect.TypeOf((*MockMessage)(nil).OrderingKey))
}

// Payload mocks base method.
func (m *MockMessage) Payload() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Payload")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Payload indicates an expected call of Payload.
func (mr *MockMessageMockRecorder) Payload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Payload", reflect.TypeOf((*MockMessage)(nil).Payload))
}

// ProducerName mocks base method.
func (m *MockMessage) ProducerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProducerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ProducerName indicates an expected call of ProducerName.
func (mr *MockMessageMockRecorder) ProducerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProducerName", reflect.TypeOf((*MockMessage)(nil).ProducerName))
}

// Properties mocks base method.
func (m *MockMessage) Properties() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Properties")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Properties indicates an expected call of Properties.
func (mr *MockMessageMockRecorder) Properties() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Properties", reflect.TypeOf((*MockMessage)(nil).Properties))
}

// PublishTime mocks base method.
func (m *MockMessage) PublishTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// PublishTime indicates an expected call of PublishTime.
func (mr *MockMessageMockRecorder) PublishTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishTime", reflect.TypeOf((*MockMessage)(nil).PublishTime))
}

// RedeliveryCount mocks base method.
func (m *MockMessage) RedeliveryCount() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedeliveryCount")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// RedeliveryCount indicates an expected call of RedeliveryCount.
func (mr *MockMessageMockRecorder) RedeliveryCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedeliveryCount", reflect.TypeOf((*MockMessage)(nil).RedeliveryCount))
}

// Topic mocks base method.
func (m *MockMessage) Topic() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Topic")
	ret0, _ := ret[0].(string)
	return ret0
}

// Topic indicates an expected call of Topic.
func (mr *MockMessageMockRecorder) Topic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Topic", reflect.TypeOf((*MockMessage)(nil).Topic))
}
