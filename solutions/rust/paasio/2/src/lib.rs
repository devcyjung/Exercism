use std::io::{Read, Result, Write};

pub struct ReadStats<R>
where R: Read
{
    reader: R,
    call_count: usize,
    bytes_read: usize,
}

impl<R: Read> ReadStats<R> {
    pub fn new(_wrapped: R) -> Self {
        Self {
            reader: _wrapped,
            call_count: 0,
            bytes_read: 0,
        }
    }

    pub fn get_ref(&self) -> &R {
        &self.reader
    }

    pub fn bytes_through(&self) -> usize {
        self.bytes_read
    }

    pub fn reads(&self) -> usize {
        self.call_count
    }
}

impl<R: Read> Read for ReadStats<R> {
    fn read(&mut self, buf: &mut [u8]) -> Result<usize> {
        let res = self.reader.read(buf);
        if let Ok(n) = res {
            self.bytes_read += n;
        }
        self.call_count += 1;
        res
    }
}

pub struct WriteStats<W>
where W: Write
{
    writer: W,
    call_count: usize,
    bytes_written: usize,
}

impl<W: Write> WriteStats<W> {
    pub fn new(_wrapped: W) -> Self {
        Self {
            writer: _wrapped,
            call_count: 0,
            bytes_written: 0,
        }
    }

    pub fn get_ref(&self) -> &W {
        &self.writer
    }

    pub fn bytes_through(&self) -> usize {
        self.bytes_written
    }

    pub fn writes(&self) -> usize {
        self.call_count
    }
}

impl<W: Write> Write for WriteStats<W> {
    fn write(&mut self, buf: &[u8]) -> Result<usize> {
        let res = self.writer.write(buf);
        if let Ok(n) = res {
            self.bytes_written += n;
        }
        self.call_count += 1;
        res
    }

    fn flush(&mut self) -> Result<()> {
        self.writer.flush()
    }
}
