package filter

import "github.com/nmertix/osm-importer/dto"

type Node interface {
	IsSuitable(dto.Node) bool
}
