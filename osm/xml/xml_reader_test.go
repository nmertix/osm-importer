package xml

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nmertix/osm-importer/dto"
)

//go:embed testdata/test.osm.xml
var testOsm []byte

//go:embed testdata/wrong.osm.xml
var wrongXml []byte

func Test_xmlReader_Read(t *testing.T) {
	type args struct {
		ctx         context.Context
		source      io.Reader
	}
	tests := []struct {
		name          string
		args          args
		want []dto.Node
		wantErr bool
	}{
		{
			"success",
			args{
				context.Background(),
				bytes.NewReader(testOsm),
			},
			[]dto.Node{
				dto.NewNode(45161410, 12, 42.6178377, 1.73608, map[string]string{
					"alt_name": "Pic de Juclar",
					"ele": "2617",
					"name": "Pic de Ruf",
					"natural": "peak",
					"source": "Bing;GPS",
				}),
				dto.NewNode(45161413, 10, 42.6117099, 1.7384087, map[string]string{
					"ele": "2748",
					"name": "Pic de Noè",
					"natural": "peak",
					"wikidata": "Q21330042",
				}),
				dto.NewNode(45161419, 6, 42.6125608, 1.7379897, map[string]string{}),
			},
			false,
		},
		{
			"wrong xml",
			args{
				context.Background(),
				bytes.NewReader(wrongXml),
			},
			[]dto.Node{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewXmlReader(tt.args.source)
			nodes := make([]dto.Node, 0)
			err := r.Read(tt.args.ctx, func(node dto.Node) {
				nodes = append(nodes, node)
			})
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, nodes)
		})
	}
}
