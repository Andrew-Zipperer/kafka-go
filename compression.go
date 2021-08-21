package kafka

import (
	"errors"

	"github.com/Andrew-Zipperer/kafka-go/compress"
)

type Compression = compress.Compression

const (
	Gzip   Compression = compress.Gzip
	Snappy Compression = compress.Snappy
	Lz4    Compression = compress.Lz4
	Zstd   Compression = compress.Zstd
)

type CompressionCodec = compress.Codec

// TODO: this file should probably go away once the internals of the package
// have moved to use the protocol package.
const (
	compressionCodecMask = 0x07
)

var (
	errUnknownCodec = errors.New("the compression code is invalid or its codec has not been imported")
)

// resolveCodec looks up a codec by Code()
func resolveCodec(code int8) (CompressionCodec, error) {
	codec := compress.Compression(code).Codec()
	if codec == nil {
		return nil, errUnknownCodec
	}
	return codec, nil
}
