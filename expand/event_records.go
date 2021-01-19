package expand

import (
	"fmt"
	"github.com/JFJun/stafi-substrate-go/expand/acala"
	"github.com/JFJun/stafi-substrate-go/expand/base"
	"github.com/JFJun/stafi-substrate-go/expand/chainX"
	"github.com/JFJun/stafi-substrate-go/expand/crust"
	"github.com/JFJun/stafi-substrate-go/expand/darwinia"
	"github.com/JFJun/stafi-substrate-go/expand/dot"
	"github.com/JFJun/stafi-substrate-go/expand/ori"
	"github.com/JFJun/stafi-substrate-go/expand/stafi"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"reflect"
	"strings"
)

type IEventRecords interface {
	GetBalancesTransfer() []types.EventBalancesTransfer
	GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess
	GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed
}

/*
扩展： 解析event
*/
func DecodeEventRecords(meta *types.Metadata, rawData string, chainName string) (IEventRecords, error) {
	e := types.EventRecordsRaw(types.MustHexDecodeString(rawData))
	var ier IEventRecords
	switch strings.ToLower(chainName) {
	case "chainx":
		var events chainX.ChainXEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "crab", "darwinia":
		var events darwinia.DarwiniaEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "crust":
		var events crust.CRustEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "mandala": // acala mandala 网络
		var events acala.AcalaEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "node": //stafi
		var events stafi.StafiEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "orion":
		var events ori.OrionEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "polkadot", "ksm":
		var events dot.DotEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	default:
		var events base.BaseEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	}
	return ier, nil
}

func CheckIsImplementedAllEvent(meta *types.Metadata, eventRecordType reflect.Type) (noImplementedEvent []string, isAllImplemented bool) {
	eventList := GetAllImplementedEventList(eventRecordType)
	existFunc := func(eventName string) bool {
		for _, item := range eventList {
			if item == eventName {
				return true
			}
		}

		return false
	}

	// 获取所有没有实现的事件
	for _, moduleItem := range meta.AsMetadataV12.Modules {
		for _, eventItem := range moduleItem.Events {
			eventName := fmt.Sprintf("%v_%v", moduleItem.Name, eventItem.Name)
			if existFunc(eventName) == false {
				noImplementedEvent = append(noImplementedEvent, eventName)
			}
		}
	}

	isAllImplemented = len(noImplementedEvent) > 0
	return
}

func GetAllImplementedEventList(tp reflect.Type) []string {
	eventList := []string{}

	tpObj := reflect.TypeOf(dot.DotEventRecords{})
	_GetEventDetail(tpObj, &eventList)

	return eventList
}

func _GetEventDetail(tp reflect.Type, eventList *[]string) {
	for index := 0; index < tp.NumField(); index++ {
		fieldItem := tp.Field(index)
		if fieldItem.Type.Kind() == reflect.Array || fieldItem.Type.Kind() == reflect.Slice {
			*eventList = append(*eventList, fieldItem.Name)
			continue
		}

		if fieldItem.Type.Kind() == reflect.Struct || fieldItem.Type.Kind() == reflect.Ptr {
			_GetEventDetail(fieldItem.Type, eventList)
		}
	}
}
