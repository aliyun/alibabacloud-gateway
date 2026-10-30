package client

import (
	"bytes"
	"encoding/xml"
	"io"
	"strings"

	"github.com/alibabacloud-go/tea/dara"
)

func parseXml(raw *string) map[string]interface{} {
	dec := xml.NewDecoder(bytes.NewReader([]byte(dara.StringValue(raw))))
	dec.CharsetReader = func(_ string, input io.Reader) (io.Reader, error) {
		return input, nil
	}
	var start xml.StartElement
	for {
		tok, err := dec.Token()
		if err != nil {
			return nil
		}
		se, ok := tok.(xml.StartElement)
		if ok {
			start = se
			break
		}
	}
	node, err := readXML(dec, start)
	if err != nil {
		return nil
	}
	return map[string]interface{}{start.Name.Local: node}
}

func readXML(dec *xml.Decoder, start xml.StartElement) (interface{}, error) {
	var text strings.Builder
	children := map[string][]interface{}{}
	hasChild := false
	for {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			hasChild = true
			child, err := readXML(dec, t)
			if err != nil {
				return nil, err
			}
			name := t.Name.Local
			children[name] = append(children[name], child)
		case xml.EndElement:
			if !hasChild {
				return strings.TrimSpace(text.String()), nil
			}
			out := map[string]interface{}{}
			for key, values := range children {
				if len(values) == 1 {
					out[key] = values[0]
				} else {
					out[key] = values
				}
			}
			return out, nil
		case xml.CharData:
			text.Write(t)
		}
	}
}
