#include "intergalactic_transmission.h"

static uint8_t read_bit(const uint8_t byte, const uint8_t bit_pos) {
    return (byte & 1 << (7 - bit_pos)) == 0 ? 0 : 1;
}

static uint8_t write_bit(uint8_t* byte, const uint8_t bit_pos) {
    return *byte |= 1 << (7 - bit_pos);
}

int transmit_sequence(uint8_t *buffer, const uint8_t *message,
                      const int message_length) {
    if (!buffer || !message || !message_length)
        return 0;
    int out = 0;
    uint8_t out_bit = 0;
    buffer[0] = 0;
    uint8_t parity = 0;
    for (int in = 0; in < message_length; ++in) {
        for (uint8_t in_bit = 0; in_bit < 8; ++in_bit) {
            const uint8_t read = read_bit(message[in], in_bit);
            if (read) {
                write_bit(&buffer[out], out_bit);
                parity ^= 1;
            }
            ++out_bit;
            if (out_bit == 7) {
                if (parity) {
                    write_bit(&buffer[out], 7);
                }
                out_bit = 0;
                parity = 0;
                buffer[++out] = 0;
            }
        }
    }
    if (out_bit != 0) {
        if (parity) {
            write_bit(&buffer[out], 7);
        }
        ++out;
    }
    return out;
}

int decode_message(uint8_t *buffer, const uint8_t *message,
                   int message_length) {
    if (!buffer || !message || !message_length)
        return 0;
    int out = 0;
    uint8_t out_bit = 0;
    buffer[0] = 0;
    int out_length = 1;
    for (int in = 0; in < message_length; ++in) {
        uint8_t parity = 0;
        for (uint8_t in_bit = 0; in_bit < 7; ++in_bit) {
            if (read_bit(message[in], in_bit)) {
                write_bit(&buffer[out], out_bit);
                out_length = out + 1;
                parity ^= 1;
            }
            if (++out_bit == 8) {
                out_bit = 0;
                buffer[++out] = 0;
            }
        }
        if (parity != read_bit(message[in], 7)) {
            return WRONG_PARITY;
        }
    }
    return out_length;
}
