package osm

import (
	"context"

	"github.com/nmertix/osm-importer/dto"
)

type Reader interface {
	Read(ctx context.Context, nodeCallback func(dto.Node)) error
}
