package xml

import (
	"context"
	"encoding/xml"
	"io"

	"github.com/nmertix/osm-importer/dto"
	"github.com/nmertix/osm-importer/osm"
)

const (
	nodeElement = "node"
)

type xmlReader struct {
	source io.Reader
}

func NewXmlReader(source io.Reader) osm.Reader {
	return &xmlReader{source: source}
}



func (r *xmlReader) Read(ctx context.Context, nodeCallback func(dto.Node)) error {
	decoder := xml.NewDecoder(r.source)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		} else if t == nil {
			break
		}

		switch el := t.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case nodeElement:
				var node Node
				err = decoder.DecodeElement(&node, &el)
				if err != nil {
					return err
				} 
				tags := make(map[string]string, len(node.Tags))
				for _, tag := range node.Tags {
					tags[tag.Key] = tag.Value
				}
				nodeData := dto.NewNode(node.Id, node.Version, node.Latitude, node.Longitude, tags)
				nodeCallback(nodeData)
			}
		}
	}

	return nil
}