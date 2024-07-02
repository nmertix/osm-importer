Simple importer for OpenStreetMap data to another format.

Usage example:
```go
type hotelFilter struct {
}

func (f *hotelFilter) IsSuitable(node dto.Node) bool {
	return node.Tags()["tourism"] == "hotel"
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	osmFile := "data.osm"
	file, err := os.Open(osmFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open OSM file: %v\n", err)
		os.Exit(1)
	}

	csvFile, err := os.Create("hotels.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create CSV file: %v\n", err)
		os.Exit(1)
	}

	importer := importer.NewImporter(xml.NewXmlReader(file), &hotelFilter{}, 100)
	writer := csv.NewWriter(csvFile)
	writer.Comma = ';'
	lines := [][]string{{"node_id", "version", "latitude", "longitude", "tags"}}

	importer.Import(ctx, func(nodes []dto.Node) {
		for _, node := range nodes {
			lines = append(lines, []string{
				strconv.FormatInt(node.Id(), 10), 
				strconv.FormatInt(int64(node.Version()), 10),
				strconv.FormatFloat(node.Latitude(), 'g', -1, 64),
				strconv.FormatFloat(node.Longitude(), 'g', -1, 64),
				fmt.Sprintf("%v", node.Tags()),
			})
		}
	})

	writer.WriteAll(lines)
	if err := writer.Error(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot write CSV file: %v\n", err)
		os.Exit(1)
	}
}
```