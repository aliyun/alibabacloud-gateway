package client

import (
	"testing"

	"github.com/gogo/protobuf/proto"
)

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
