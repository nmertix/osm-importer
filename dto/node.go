package dto

type Node struct {
	id int64
	version int16
	latitude float64
	longitude float64
	tags map[string]string
}

func NewNode(
	id int64,
	version int16,
	latitude float64,
	longitude float64,
	tags map[string]string,
) Node {
	return Node{
		id: id,
		version: version,
		latitude: latitude,
		longitude: longitude,
		tags: tags,
	}
}

func (n Node) Id() int64 {
	return n.id
}

func (n Node) Version() int16 {
	return n.version
}

func (n Node) Latitude() float64 {
	return n.latitude
}

func (n Node) Longitude() float64 {
	return n.longitude
}

func (n Node) Tags() map[string]string {
	return n.tags
}