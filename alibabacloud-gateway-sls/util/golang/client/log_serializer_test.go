package client

import (
	"reflect"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestDeserializeLogGroupListFromPB_Empty(t *testing.T) {
	result, err := deserializeLogGroupListFromPB([]byte{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	m := result.(map[string]interface{})
	logGroups := m["logGroupList"].([]interface{})
	if len(logGroups) != 0 {
		t.Errorf("expected empty logGroupList, got %d items", len(logGroups))
	}
}

func TestDeserializeLogGroupListFromPB_InvalidData(t *testing.T) {
	_, err := deserializeLogGroupListFromPB([]byte{0xff, 0xfe, 0xfd})
	if err == nil {
		t.Fatal("expected error for invalid protobuf data")
	}
}

func TestDeserializeLogGroupListFromPB_SingleLogGroup(t *testing.T) {
	topic := "test-topic"
	source := "test-source"
	time := uint32(1678888888)
	timeNs := uint32(123456)
	key := "content-key"
	value := "content-value"
	tagKey := "tag-key"
	tagValue := "tag-value"

	lgl := &LogGroupList{
		LogGroups: []*LogGroup{
			{
				Topic:  &topic,
				Source: &source,
				LogTags: []*LogTag{
					{Key: &tagKey, Value: &tagValue},
				},
				Logs: []*Log{
					{
						Time:   &time,
						TimeNs: &timeNs,
						Contents: []*LogContent{
							{Key: &key, Value: &value},
						},
					},
				},
			},
		},
	}
	data, err := lgl.Marshal()
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	result, err := deserializeLogGroupListFromPB(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	m := result.(map[string]interface{})
	logGroups := m["logGroupList"].([]interface{})
	if len(logGroups) != 1 {
		t.Fatalf("expected 1 logGroup, got %d", len(logGroups))
	}

	lg := logGroups[0].(map[string]interface{})
	if lg["Topic"] != "test-topic" {
		t.Errorf("expected Topic 'test-topic', got %v", lg["Topic"])
	}
	if lg["Source"] != "test-source" {
		t.Errorf("expected Source 'test-source', got %v", lg["Source"])
	}

	tags := lg["LogTags"].([]interface{})
	if len(tags) != 1 {
		t.Fatalf("expected 1 tag, got %d", len(tags))
	}
	tag := tags[0].(map[string]interface{})
	if tag["Key"] != "tag-key" || tag["Value"] != "tag-value" {
		t.Errorf("unexpected tag: %v", tag)
	}

	items := lg["LogItems"].([]interface{})
	if len(items) != 1 {
		t.Fatalf("expected 1 logItem, got %d", len(items))
	}
	item := items[0].(map[string]interface{})
	if item["Time"] != int(1678888888) {
		t.Errorf("expected Time 1678888888, got %v", item["Time"])
	}
	if item["TimeNs"] != int(123456) {
		t.Errorf("expected TimeNs 123456, got %v", item["TimeNs"])
	}
	contents := item["Contents"].([]interface{})
	if len(contents) != 1 {
		t.Fatalf("expected 1 content, got %d", len(contents))
	}
	c := contents[0].(map[string]interface{})
	if c["Key"] != "content-key" || c["Value"] != "content-value" {
		t.Errorf("unexpected content: %v", c)
	}
}

func TestDeserializeLogGroupListFromPB_MultipleLogGroups(t *testing.T) {
	topic1 := "topic-1"
	topic2 := "topic-2"
	time1 := uint32(1000)
	time2 := uint32(2000)
	k1, v1 := "k1", "v1"
	k2, v2 := "k2", "v2"

	lgl := &LogGroupList{
		LogGroups: []*LogGroup{
			{
				Topic: &topic1,
				Logs: []*Log{
					{Time: &time1, Contents: []*LogContent{{Key: &k1, Value: &v1}}},
				},
			},
			{
				Topic: &topic2,
				Logs: []*Log{
					{Time: &time2, Contents: []*LogContent{{Key: &k2, Value: &v2}}},
				},
			},
		},
	}
	data, _ := lgl.Marshal()

	result, err := deserializeLogGroupListFromPB(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	m := result.(map[string]interface{})
	logGroups := m["logGroupList"].([]interface{})
	if len(logGroups) != 2 {
		t.Fatalf("expected 2 logGroups, got %d", len(logGroups))
	}

	lg1 := logGroups[0].(map[string]interface{})
	lg2 := logGroups[1].(map[string]interface{})
	if lg1["Topic"] != "topic-1" || lg2["Topic"] != "topic-2" {
		t.Errorf("unexpected topics: %v, %v", lg1["Topic"], lg2["Topic"])
	}
}

func TestDeserializeLogGroupListFromPB_NoOptionalFields(t *testing.T) {
	time := uint32(1000)
	k, v := "key", "value"

	lgl := &LogGroupList{
		LogGroups: []*LogGroup{
			{
				Logs: []*Log{
					{Time: &time, Contents: []*LogContent{{Key: &k, Value: &v}}},
				},
			},
		},
	}
	data, _ := lgl.Marshal()

	result, err := deserializeLogGroupListFromPB(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	m := result.(map[string]interface{})
	lg := m["logGroupList"].([]interface{})[0].(map[string]interface{})

	if _, ok := lg["Topic"]; ok {
		t.Error("Topic should not be present")
	}
	if _, ok := lg["Source"]; ok {
		t.Error("Source should not be present")
	}
	if _, ok := lg["LogTags"]; ok {
		t.Error("LogTags should not be present")
	}

	item := lg["LogItems"].([]interface{})[0].(map[string]interface{})
	if _, ok := item["TimeNs"]; ok {
		t.Error("TimeNs should not be present")
	}
}

func TestDeserializeLogGroupListFromPB_MultipleContents(t *testing.T) {
	time := uint32(1000)
	k1, v1 := "level", "INFO"
	k2, v2 := "message", "hello"
	k3, v3 := "host", "10.0.0.1"

	lgl := &LogGroupList{
		LogGroups: []*LogGroup{
			{
				Logs: []*Log{
					{
						Time: &time,
						Contents: []*LogContent{
							{Key: &k1, Value: &v1},
							{Key: &k2, Value: &v2},
							{Key: &k3, Value: &v3},
						},
					},
				},
			},
		},
	}
	data, _ := lgl.Marshal()

	result, err := deserializeLogGroupListFromPB(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	m := result.(map[string]interface{})
	item := m["logGroupList"].([]interface{})[0].(map[string]interface{})["LogItems"].([]interface{})[0].(map[string]interface{})
	contents := item["Contents"].([]interface{})
	if len(contents) != 3 {
		t.Fatalf("expected 3 contents, got %d", len(contents))
	}

	expected := []map[string]interface{}{
		{"Key": "level", "Value": "INFO"},
		{"Key": "message", "Value": "hello"},
		{"Key": "host", "Value": "10.0.0.1"},
	}
	for i, c := range contents {
		cm := c.(map[string]interface{})
		if !reflect.DeepEqual(cm, expected[i]) {
			t.Errorf("content[%d]: expected %v, got %v", i, expected[i], cm)
		}
	}
}

func TestDeserializeLogGroupListFromPB_RoundTrip(t *testing.T) {
	logGroupMap := map[string]interface{}{
		"Topic":  "round-trip-topic",
		"Source": "round-trip-source",
		"LogTags": []interface{}{
			map[string]interface{}{"Key": "env", "Value": "prod"},
		},
		"LogItems": []interface{}{
			map[string]interface{}{
				"Time":   int(1678888888),
				"TimeNs": int(999),
				"Contents": []interface{}{
					map[string]interface{}{"Key": "msg", "Value": "hello world"},
				},
			},
		},
	}

	data, err := serializeLogGroupToPB(logGroupMap)
	if err != nil {
		t.Fatalf("serialize failed: %v", err)
	}

	lgl := &LogGroupList{LogGroups: []*LogGroup{{}}}
	if err := lgl.LogGroups[0].Unmarshal(data); err != nil {
		t.Fatalf("unmarshal to LogGroup failed: %v", err)
	}
	listData, err := lgl.Marshal()
	if err != nil {
		t.Fatalf("marshal LogGroupList failed: %v", err)
	}

	result, err := deserializeLogGroupListFromPB(listData)
	if err != nil {
		t.Fatalf("deserialize failed: %v", err)
	}

	m := result.(map[string]interface{})
	lg := m["logGroupList"].([]interface{})[0].(map[string]interface{})
	if lg["Topic"] != "round-trip-topic" {
		t.Errorf("Topic mismatch: %v", lg["Topic"])
	}
	if lg["Source"] != "round-trip-source" {
		t.Errorf("Source mismatch: %v", lg["Source"])
	}
	item := lg["LogItems"].([]interface{})[0].(map[string]interface{})
	if item["Time"] != int(1678888888) {
		t.Errorf("Time mismatch: %v", item["Time"])
	}
	if item["TimeNs"] != int(999) {
		t.Errorf("TimeNs mismatch: %v", item["TimeNs"])
	}
	c := item["Contents"].([]interface{})[0].(map[string]interface{})
	if c["Key"] != "msg" || c["Value"] != "hello world" {
		t.Errorf("Content mismatch: %v", c)
	}
}

func TestSerializeLogGroupToPB(t *testing.T) {
	// Prepare test data
	logGroupMap := map[string]interface{}{
		"Topic":  "test-topic",
		"Source": "test-source",
		"LogTags": []interface{}{
			map[string]interface{}{
				"Key":   "tag-key",
				"Value": "tag-value",
			},
		},
		"LogItems": []interface{}{
			map[string]interface{}{
				"Time":   int(1678888888),
				"TimeNs": int(123456),
				"Contents": []interface{}{
					map[string]interface{}{
						"Key":   "content-key",
						"Value": "content-value",
					},
				},
			},
		},
	}

	// Execute
	data, err := serializeLogGroupToPB(logGroupMap)
	if err != nil {
		t.Fatalf("serializeLogGroupToPB failed: %v", err)
	}

	// Validate
	if len(data) == 0 {
		t.Fatal("serialized data is empty")
	}

	var lg LogGroup
	if err := proto.Unmarshal(data, &lg); err != nil {
		t.Fatalf("failed to unmarshal serialized data: %v", err)
	}

	if lg.GetTopic() != "test-topic" {
		t.Errorf("expected topic 'test-topic', got %s", lg.GetTopic())
	}

	if lg.GetSource() != "test-source" {
		t.Errorf("expected source 'test-source', got %s", lg.GetSource())
	}

	if len(lg.LogTags) != 1 {
		t.Fatalf("expected 1 tag, got %d", len(lg.LogTags))
	}
	if lg.LogTags[0].GetKey() != "tag-key" || lg.LogTags[0].GetValue() != "tag-value" {
		t.Errorf("unexpected tag content: %v", lg.LogTags[0])
	}

	if len(lg.Logs) != 1 {
		t.Fatalf("expected 1 log item, got %d", len(lg.Logs))
	}
	if lg.Logs[0].GetTime() != 1678888888 {
		t.Errorf("expected time 1678888888, got %d", lg.Logs[0].GetTime())
	}
}
