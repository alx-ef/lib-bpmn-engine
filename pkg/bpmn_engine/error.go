package bpmn_engine

import "fmt"

type BpmnEngineError struct {
	Msg string
}

func (e *BpmnEngineError) Error() string {
	return e.Msg
}

func NewBpmnEngineError(msg string, args ...any) error {
	return &BpmnEngineError{Msg: fmt.Sprintf(msg, args...)}
}
