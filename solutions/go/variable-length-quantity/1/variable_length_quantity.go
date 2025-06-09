package variablelengthquantity

import (
    "errors"
    "slices"
)

const (
    encoding byte = 0x80
    decoding byte = ^encoding
)

var (
    masks = [5]uint32{
        0xFE000000, 0xFE00000, 0x1FC000, 0x3F80, 0x7F,
    }
    thresholds = [5]uint32{
        masks[1] | masks[2] | masks[3] | masks[4],
        masks[2] | masks[3] | masks[4],
        masks[3] | masks[4],
        masks[4],
        0,
    }
    offsets = [5]int{
        28, 21, 14, 7, 0,
    }
    incompleteSequenceError = errors.New("encoded sequence is incomplete")
    overflowError = errors.New("encoded sequence has a value exceeding uint32 upper bound")
)

func EncodeVarint(input []uint32) []byte {
	u8s := make([]byte, 0, len(input) * 5)
    for _, u32 := range input {
        if u32 == 0 {
            u8s = append(u8s, byte(0))
            continue
        }
        for i := 0; i < 5; i++ {
            if thresholds[i] < u32 {
                if i == 4 {
                    u8s = append(u8s, byte((u32 & masks[i]) >> offsets[i]))
                } else {
                    u8s = append(u8s, byte((u32 & masks[i]) >> offsets[i]) | encoding)
                }
            }
        }
    }
    slices.Clip(u8s)
    return u8s
}

func DecodeVarint(input []byte) ([]uint32, error) {
	u32s := make([]uint32, 0, len(input))
    index, offset, total := 0, 0, len(input)
    var data uint32
    for index < total {
        for {
            if index + offset == total {
                return nil, incompleteSequenceError
            }
            if input[index + offset] & encoding == 0 {
                break
            }
            if offset == 4 {
                return nil, overflowError
            }
            offset++
        }
        for i := 0; i <= offset; i++ {
            data |= uint32(input[index + i] & decoding) << offsets[4 - offset + i]
        }
        u32s = append(u32s, data)
        index += offset + 1
        offset = 0
        data = 0
    }
    slices.Clip(u32s)
    return u32s, nil
}