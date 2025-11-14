package client

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gogo/protobuf/proto"
)

func serializeLogGroupToPB(logGroup interface{}) ([]byte, error) {
	if logGroup == nil {
		return nil, errors.New("fail to serialize LogGroup to protobuf, logGroup is nil")
	}

	m, ok := logGroup.(map[string]interface{})
	if !ok {
		return nil, errors.New("fail to serialize LogGroup to protobuf, logGroup is not a map")
	}

	lg := &LogGroup{}

	// topic
	topic, err := getString(m, "Topic")
	if err != nil {
		return nil, err
	}
	if topic != "" {
		lg.Topic = &topic
	}

	// source
	source, err := getString(m, "Source")
	if err != nil {
		return nil, err
	}
	if source != "" {
		lg.Source = &source
	}

	// logTags
	if err := serializeLogTags(lg, m); err != nil {
		return nil, err
	}

	// logItems
	if err := serializeLogItems(lg, m); err != nil {
		return nil, err
	}

	// to bytes
	data, err := lg.Marshal()
	if err != nil {
		return nil, fmt.Errorf("fail to serialize LogGroup to protobuf, %w", err)
	}
	return data, nil
}

func serializeLogItems(lg *LogGroup, logGroup map[string]interface{}) error {
	// check if logItems is nil, regard as empty logItems if nil
	logItemsValue, ok := logGroup["LogItems"]
	if !ok {
		return nil
	}

	logItems, ok := logItemsValue.([]interface{})
	if !ok {
		return errors.New("fail to serialize LogGroup to protobuf, LogItems field is not a slice")
	}

	// empty logItems is valid
	for _, logItem := range logItems {
		if err := serializeLogItem(lg, logItem); err != nil {
			return err
		}
	}
	return nil
}

func serializeLogItem(lg *LogGroup, logItem interface{}) error {
	if logItem == nil {
		return errors.New("fail to serialize LogGroup to protobuf, logItem is nil")
	}

	m, ok := logItem.(map[string]interface{})
	if !ok {
		return errors.New("fail to serialize LogGroup to protobuf, logItem is not a map")
	}

	logPb := &Log{}

	if time, exists, err := getInt(m, "Time"); err != nil {
		return err
	} else if !exists {
		return errors.New("fail to serialize LogGroup to protobuf, log time is missing")
	} else {
		logPb.Time = proto.Uint32(uint32(time))
	}

	if timeNano, exists, err := getInt(m, "TimeNs"); err != nil {
		return err
	} else if exists {
		logPb.TimeNs = proto.Uint32(uint32(timeNano))
	}

	if err := serializeLogContents(logPb, m); err != nil {
		return err
	}
	lg.Logs = append(lg.Logs, logPb)
	return nil
}

func serializeLogContents(logPb *Log, logItem map[string]interface{}) error {
	// check if logContents is nil, regard as empty logContents if nil
	contentsValue, ok := logItem["Contents"]
	if !ok {
		return errors.New("fail to serialize LogGroup to protobuf, log Contents is missing")
	}
	contents, ok := contentsValue.([]interface{})
	if !ok {
		return errors.New("fail to serialize LogGroup to protobuf, logContents is not a slice")
	}

	// empty logContents is valid
	for _, content := range contents {
		if err := serializeLogContent(logPb, content); err != nil {
			return err
		}
	}

	return nil
}

func serializeLogContent(logPb *Log, content interface{}) error {
	if content == nil {
		return errors.New("fail to serialize LogGroup to protobuf, logContent is nil")
	}

	m, ok := content.(map[string]interface{})
	if !ok {
		return errors.New("fail to serialize LogGroup to protobuf, logContent is not a map")
	}

	key, value, err := getKeyValuePair(m)
	if err != nil {
		return err
	}

	logPb.Contents = append(logPb.Contents, &LogContent{Key: &key, Value: &value})
	return nil
}

func serializeLogTags(lg *LogGroup, logGroup map[string]interface{}) error {
	// check if logTags is nil, regard as empty logTags if nil
	logTagsValue, ok := logGroup["LogTags"]
	if !ok {
		return nil
	}

	logTags, ok := logTagsValue.([]interface{})
	if !ok {
		return errors.New("fail to serialize LogGroup to protobuf, LogTags field is not a slice")
	}

	// empty logTags is valid
	for _, logTag := range logTags {
		if err := serializeLogTag(lg, logTag); err != nil {
			return err
		}
	}
	return nil
}

func serializeLogTag(lg *LogGroup, logTag interface{}) error {
	if logTag == nil {
		return errors.New("fail to serialize LogGroup to protobuf, logTag is nil")
	}

	m, ok := logTag.(map[string]interface{})
	if !ok {
		return errors.New("fail to serialize LogGroup to protobuf, logTag is not a map")
	}

	key, value, err := getKeyValuePair(m)
	if err != nil {
		return err
	}

	lg.LogTags = append(lg.LogTags, &LogTag{Key: &key, Value: &value})
	return nil
}

func getKeyValuePair(m map[string]interface{}) (key string, value string, err error) {
	key, err = getString(m, "Key")
	if err != nil {
		return "", "", err
	}

	value, err = getString(m, "Value")
	if err != nil {
		return "", "", err
	}

	return key, value, nil
}

func getString(m map[string]interface{}, key string) (string, error) {
	v, ok := m[key]
	if !ok {
		return "", nil
	}

	s, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("value of %s is not a string (type: %T)", key, v)
	}
	return s, nil
}

func getInt(m map[string]interface{}, key string) (int64, bool, error) {
	v, ok := m[key]
	if !ok {
		return 0, false, nil
	}

	if i, ok := v.(json.Number); ok {
		if vv, err := i.Int64(); err != nil {
			return 0, false, err
		} else {
			return vv, true, nil
		}
	}

	if i, ok := v.(int64); ok {
		return i, true, nil
	}

	if i, ok := v.(int); ok {
		return int64(i), true, nil
	}

	if i, ok := v.(int32); ok {
		return int64(i), true, nil
	}

	return 0, false, fmt.Errorf("value of %s is not a number (type: %T, value: %v)", key, v, v)
}
