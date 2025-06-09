/// IncompleteNumber when the last byte in the sequence has MSB=1
/// Overflow when more than 5 contiguous bytes in the sequence has MSB=1
#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    IncompleteNumber,
    Overflow,
}

const ENCODING: u8 = 1u8 << 7;
const MASK: [u32; 5] = [0xF0000000, 0xFE00000, 0x1FC000, 0x3F80, 0x7F];
const OFFSET: [u32; 5] = [28, 21, 14, 7, 0];
const DECODING: u8 = !ENCODING;

/// Convert a list of numbers to a stream of bytes encoded with variable length encoding.
pub fn to_bytes(values: &[u32]) -> Vec<u8> {
    let mut result = Vec::with_capacity(5 * values.len());
    for value in values {
        // Scenario
        // leading zeros = 32 -> 0
        // leading zeros = [32-7, 32-1] -> Use MASK/OFFSET[4]
        // leading zeros = [32-14, 32-8] -> Use MASK/OFFSET[3], [4]
        // leading zeros = [0, 32-29] -> Use MASK/OFFSET[0], [1], [2], [3], [4], [5]
        for i in (5 - (32 - value.leading_zeros()).div_ceil(7).max(1)) as usize..5 {
            if i < 4 {
                result.push(((value & MASK[i]) >> OFFSET[i]) as u8 | ENCODING);
            } else {
                result.push(((value & MASK[i]) >> OFFSET[i]) as u8);
            }
        }
    }
    result.shrink_to_fit();
    result
}

/// Given a stream of bytes, extract all numbers which are encoded in there.
pub fn from_bytes(bytes: &[u8]) -> Result<Vec<u32>, Error> {
    let mut index = 0;
    let mut result = Vec::with_capacity(bytes.len());
    while index < bytes.len() {
        let mut offset = 0;
        loop {
            // Scenario (MSB)
            // 1, 1, 1 (bytes end) (illegal)
            // 1, 1, 1, 1, 1 (illegal)
            // 1, 1, 1, 1, 0 (legal)
            if index + offset == bytes.len() {
                return Err(Error::IncompleteNumber);
            }
            if bytes[index + offset] & ENCODING == 0 {
                break;
            }
            if offset == 4 {
                return Err(Error::Overflow);
            }
            offset += 1;
        }
        // Scenario
        // index = 26, offset = 3
        // bytes[26], [27], [28], [29]
        // OFFSET[1], [2], [3], [4]
        // index = 26, offset = 2
        // bytes[26], [27], [28]
        // OFFSET[2], [3], [4]
        let mut data = 0u32;
        for (bytes_index, offset_index) in (index..=(index + offset)).zip((4 - offset)..5) {
            data |= ((bytes[bytes_index] & DECODING) as u32) << OFFSET[offset_index];
        }
        result.push(data);
        index += offset + 1;
    }
    result.shrink_to_fit();
    Ok(result)
}