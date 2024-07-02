package importer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nmertix/osm-importer/dto"
	"github.com/nmertix/osm-importer/filter"
	"github.com/nmertix/osm-importer/osm"
)

type mockReader struct {
}

func (r *mockReader) ReadAsync(ctx context.Context, nodeChannel chan dto.Node) error {
	nodeChannel <- dto.NewNode(1, 1, 11.0, 12.0, map[string]string{
		"shop": "supermarket",
	})
	nodeChannel <- dto.NewNode(2, 2, 21.0, 22.0, map[string]string{
		"tourism": "hotel",
	})
	nodeChannel <- dto.NewNode(3, 3, 31.0, 32.0, map[string]string{
		"amenity": "restaurant",
	})
	close(nodeChannel)
	return nil
}

func (r *mockReader) Read(ctx context.Context, nodeCallback func(dto.Node)) error {
	return nil
}

type hotelFilter struct {
}

func (f *hotelFilter) IsSuitable(node dto.Node) bool {
	return node.Tags()["tourism"] == "hotel"
}

func TestImporter_Import(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name   string
		filter filter.Node
		reader osm.Reader
		expectedNodes []dto.Node
	}{
		{
			"success",
			&hotelFilter{},
			&mockReader{},
			[]dto.Node{
				dto.NewNode(2, 2, 21.0, 22.0, map[string]string{
					"tourism": "hotel",
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Importer{
				nodeFilter: tt.filter,
				reader:     tt.reader,
			}
			i.Import(ctx, func(nodes []dto.Node) {
				assert.Equal(t, tt.expectedNodes, nodes)
			})
		})
	}
}
