package gene

// GeneFactory is a point to creates genes
type GeneFactory struct {
}

// Create new gene by its type
func (factory *GeneFactory) Create(geneType GeneType) (gene gene) {
	switch geneType {
	case geneTypePhotosynthesize:
		gene = &photosynthesizeGene{}
	case geneTypeMoveNorth:
		gene = &moveNorthGene{}
	case geneTypeMoveEast:
		gene = &moveEastGene{}
	case geneTypeMoveSouth:
		gene = &moveSouthGene{}
	case geneTypeMoveWest:
		gene = &moveWestGene{}
	}

	return
}
