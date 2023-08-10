package bpmn_engine

import (
	"github.com/nitram509/lib-bpmn-engine/pkg/bpmn_engine/exporter"
	"github.com/nitram509/lib-bpmn-engine/pkg/spec/BPMN20"
)

// AddEventExporter registers an EventExporter instance
func (state *BpmnEngineState) AddEventExporter(exporter exporter.EventExporter) {
	state.exporters = append(state.exporters, exporter)
}

func (state *BpmnEngineState) exportNewProcessEvent(processInfo ProcessInfo, xmlData []byte, resourceName string, checksum string) {
	event := exporter.ProcessEvent{
		ProcessId:    processInfo.BpmnProcessId,
		ProcessKey:   processInfo.ProcessKey,
		Version:      processInfo.Version,
		XmlData:      xmlData,
		ResourceName: resourceName,
		Checksum:     checksum,
	}
	for _, exp := range state.exporters {
		exp.NewProcessEvent(&event)
	}
}

func (state *BpmnEngineState) exportIncident(process ProcessInfo, processInstance ProcessInstanceInfo, err error) {
	event := exporter.ProcessInstanceEvent{
		ProcessId:          process.BpmnProcessId,
		ProcessKey:         process.ProcessKey,
		Version:            process.Version,
		ProcessInstanceKey: processInstance.instanceKey,
	}
	for _, exp := range state.exporters {
		exp.ProcessIncident(&event, err)
	}
}

func (state *BpmnEngineState) exportEndProcessEvent(process ProcessInfo, processInstance ProcessInstanceInfo) {
	event := exporter.ProcessInstanceEvent{
		ProcessId:          process.BpmnProcessId,
		ProcessKey:         process.ProcessKey,
		Version:            process.Version,
		ProcessInstanceKey: processInstance.instanceKey,
	}
	for _, exp := range state.exporters {
		exp.EndProcessEvent(&event)
	}
}

func (state *BpmnEngineState) exportProcessInstanceEvent(process ProcessInfo, processInstance ProcessInstanceInfo) {
	event := exporter.ProcessInstanceEvent{
		ProcessId:          process.BpmnProcessId,
		ProcessKey:         process.ProcessKey,
		Version:            process.Version,
		ProcessInstanceKey: processInstance.instanceKey,
	}
	for _, exp := range state.exporters {
		exp.NewProcessInstanceEvent(&event)
	}
}

func (state *BpmnEngineState) exportElementEvent(process ProcessInfo, processInstance ProcessInstanceInfo, element BPMN20.BaseElement, intent exporter.Intent) {
	event := exporter.ProcessInstanceEvent{
		ProcessId:          process.BpmnProcessId,
		ProcessKey:         process.ProcessKey,
		Version:            process.Version,
		ProcessInstanceKey: processInstance.instanceKey,
	}
	info := exporter.ElementInfo{
		BpmnElementType: string(element.GetType()),
		ElementId:       element.GetId(),
		Intent:          string(intent),
	}
	for _, exp := range state.exporters {
		exp.NewElementEvent(&event, &info)
	}
}

func (state *BpmnEngineState) exportSequenceFlowEvent(process ProcessInfo, processInstance ProcessInstanceInfo, flow BPMN20.TSequenceFlow) {
	event := exporter.ProcessInstanceEvent{
		ProcessId:          process.BpmnProcessId,
		ProcessKey:         process.ProcessKey,
		Version:            process.Version,
		ProcessInstanceKey: processInstance.instanceKey,
	}
	info := exporter.ElementInfo{
		BpmnElementType: string(BPMN20.SequenceFlow),
		ElementId:       flow.Id,
		Intent:          string(exporter.SequenceFlowTaken),
	}
	for _, exp := range state.exporters {
		exp.NewElementEvent(&event, &info)
	}
}
