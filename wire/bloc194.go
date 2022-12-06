package wire

import "io"

// BlOC-194 - expose internal methods
func ReadBlockHeader(r io.Reader, pver uint32, bh *BlockHeader) error {
	return readBlockHeader(r, pver, bh)
}

// BlOC-194 - expose internal methods
func GetMaxTxPerBlock() uint64 {
	return maxTxPerBlock
}
