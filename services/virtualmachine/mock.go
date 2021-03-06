package virtualmachine

import (
	"github.com/maraino/go-mock"
	"github.com/orbs-network/go-scaffold/services/statestorage"
	"github.com/orbs-network/go-scaffold/types/services/virtualmachine"
)

type MockService struct {
	mock.Mock
}

func (s *MockService) Start(stateStorage statestorage.Service, stop *chan error) {
	s.Called(stateStorage, stop)
}

func (s *MockService) Stop() {
	s.Called()
}

func (s *MockService) IsStarted() bool {
	return s.Called().Bool(0)
}

func (s *MockService) ProcessTransaction(input *virtualmachine.ProcessTransactionInput) (*virtualmachine.ProcessTransactionOutput, error) {
	ret := s.Called(input)
	return ret.Get(0).(*virtualmachine.ProcessTransactionOutput), ret.Error(1)
}

func (s *MockService) CallContract(input *virtualmachine.CallContractInput) (*virtualmachine.CallContractOutput, error) {
	ret := s.Called(input)
	return ret.Get(0).(*virtualmachine.CallContractOutput), ret.Error(1)
}