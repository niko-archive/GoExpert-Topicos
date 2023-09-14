package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name string
	Data interface{}
	Time time.Time
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetData() interface{} {
	return e.Data
}

func (e *TestEvent) GetTime() time.Time {
	return e.Time
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {

}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event01         TestEvent
	event02         TestEvent
	handler01       TestEventHandler
	handler02       TestEventHandler
	handler03       TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.event02 = TestEvent{
		Name: "test_event_01",
		Data: "test_data_01",
		Time: time.Now(),
	}
	suite.event01 = TestEvent{
		Name: "test_event_02",
		Data: "test_data_02",
		Time: time.Now(),
	}
	suite.handler02 = TestEventHandler{ID: 1}
	suite.handler01 = TestEventHandler{ID: 2}
	suite.handler03 = TestEventHandler{ID: 3}

	suite.eventDispatcher = NewEventDispatcher()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	// Mock
	eh01 := &MockHandler{}
	eh01.On("Handle", &suite.event01)

	eh02 := &MockHandler{}
	eh02.On("Handle", &suite.event01)

	suite.eventDispatcher.Register(suite.event01.GetName(), eh01)
	suite.eventDispatcher.Register(suite.event01.GetName(), eh02)

	suite.eventDispatcher.Dispatch(&suite.event01)

	eh01.AssertExpectations(suite.T())
	eh02.AssertExpectations(suite.T())

	eh01.AssertNumberOfCalls(suite.T(), "Handle", 1)
	eh02.AssertNumberOfCalls(suite.T(), "Handle", 1)

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event02.GetName(), &suite.handler02)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event02.GetName()]))

	err = suite.eventDispatcher.Register(suite.event02.GetName(), &suite.handler01)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event02.GetName()]))

	assert.Equal(suite.T(), &suite.handler02, suite.eventDispatcher.handlers[suite.event02.GetName()][0])
	assert.Equal(suite.T(), &suite.handler01, suite.eventDispatcher.handlers[suite.event02.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.eventDispatcher.Register(suite.event02.GetName(), &suite.handler02)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event02.GetName()]))

	err = suite.eventDispatcher.Register(suite.event02.GetName(), &suite.handler02)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event02.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	// Event 01
	err := suite.eventDispatcher.Register(suite.event01.GetName(), &suite.handler01)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event01.GetName()]))

	err = suite.eventDispatcher.Register(suite.event01.GetName(), &suite.handler02)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event01.GetName()]))

	// Event 02
	err = suite.eventDispatcher.Register(suite.event02.GetName(), &suite.handler01)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event02.GetName()]))

	err = suite.eventDispatcher.Register(suite.event02.GetName(), &suite.handler02)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event02.GetName()]))

	//
	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.handlers))

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	// Event 01
	err := suite.eventDispatcher.Register(suite.event01.GetName(), &suite.handler01)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event01.GetName()]))

	err = suite.eventDispatcher.Register(suite.event01.GetName(), &suite.handler02)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event01.GetName()]))

	// Event 02
	err = suite.eventDispatcher.Register(suite.event02.GetName(), &suite.handler01)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event02.GetName()]))

	err = suite.eventDispatcher.Register(suite.event02.GetName(), &suite.handler02)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event02.GetName()]))

	// Assert
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event01.GetName(), &suite.handler01))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event01.GetName(), &suite.handler02))
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event01.GetName(), &suite.handler03))

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	// Event 01
	err := suite.eventDispatcher.Register(suite.event01.GetName(), &suite.handler01)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event01.GetName()]))

	err = suite.eventDispatcher.Register(suite.event01.GetName(), &suite.handler02)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event01.GetName()]))

	// Event 02
	err = suite.eventDispatcher.Register(suite.event02.GetName(), &suite.handler03)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event02.GetName()]))

	// Remove
	suite.eventDispatcher.Remove(suite.event01.GetName(), &suite.handler01)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event01.GetName()]))
	assert.Equal(suite.T(), &suite.handler02, suite.eventDispatcher.handlers[suite.event01.GetName()][0])

	suite.eventDispatcher.Remove(suite.event01.GetName(), &suite.handler02)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event01.GetName()]))

	suite.eventDispatcher.Remove(suite.event02.GetName(), &suite.handler03)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event02.GetName()]))

}

func TestSuit(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
