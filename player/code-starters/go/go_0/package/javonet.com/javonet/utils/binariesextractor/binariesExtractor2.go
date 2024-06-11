//go:build excludeBinaries
// +build excludeBinaries

package binariesextractor

type BinariesExtractor struct {
	binariesZipPath string
}

func NewBinariesExtractor() *BinariesExtractor {
	return &BinariesExtractor{}
}

func (bu *BinariesExtractor) ExtractFromBinariesZip(runtimeName string) error {
	return nil
}
