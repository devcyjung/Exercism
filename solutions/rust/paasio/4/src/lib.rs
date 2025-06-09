use std::io::{Read, Result, Write};

pub struct ReadStats<R>
where
    R: Read,
{
    reader: R,
    call_count: usize,
    bytes_read: usize,
}

impl<R> ReadStats<R>
where
    R: Read,
{
    pub fn new(wrapped: R) -> Self {
        Self {
            reader: wrapped,
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

impl<R> Read for ReadStats<R>
where
    R: Read,
{
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
where
    W: Write,
{
    writer: W,
    call_count: usize,
    bytes_written: usize,
}

impl<W> WriteStats<W>
where
    W: Write,
{
    pub fn new(wrapped: W) -> Self {
        Self {
            writer: wrapped,
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

impl<W> Write for WriteStats<W>
where
    W: Write,
{
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