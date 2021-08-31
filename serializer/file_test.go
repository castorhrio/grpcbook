package serializer

import (
	"testing"

	"github.com/castorhrio/grpcpcbook/grpcbook"
	"github.com/castorhrio/grpcpcbook/sample"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"
	laptop := sample.NewLaptop()

	err := WriteProtobufToBinaryFile(laptop, binaryFile)
	require.NoError(t, err)

	laptop2 := &grpcbook.Laptop{}
	err = ReadProtobuffFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop, laptop2))

	err = WriteProtobufToJSONFile(laptop, jsonFile)
	require.NoError(t, err)
}
