package importer

import (
	"context"
	"sync"

	"github.com/nmertix/osm-importer/dto"
	"github.com/nmertix/osm-importer/filter"
	"github.com/nmertix/osm-importer/osm"
)

type Importer struct {
	reader     osm.Reader
	nodeFilter filter.Node
	batchSize int
}

func NewImporter(
	reader     osm.Reader,
	nodeFilter filter.Node,
	batchSize int,
) *Importer {
	return &Importer{
		reader: reader,
		nodeFilter: nodeFilter,
		batchSize: batchSize,
	}
}

func (i *Importer) Import(ctx context.Context, callback func([]dto.Node)) {
	nodesChannel := make(chan dto.Node)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		i.reader.Read(ctx, func(node dto.Node) {
			nodesChannel <- node
		})
		close(nodesChannel)
	} ()
	go func() {
		defer wg.Done()
		batch := make([]dto.Node, 0, i.batchSize)
		defer func() {
			if len(batch) > 0 {
				callback(batch)
			}
		} ()
		for node := range nodesChannel {
			if i.nodeFilter.IsSuitable(node) {
				batch = append(batch, node)
				if len(batch) == i.batchSize {
					callback(batch)
					batch = make([]dto.Node, 0, i.batchSize)
				}
			}
		}
	
	}()
	wg.Wait()
}
