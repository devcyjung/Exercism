use std::borrow::Borrow;
use std::io::{ Read, Write, Result };

/// A munger which XORs a key with some data
#[derive(Clone)]
pub struct Xorcism<'a> {
    key: &'a [u8],
    pos: usize,
}

impl<'a> Xorcism<'a> {
    /// Creates a new Xorcism munger from a key
    pub fn new<Key>(key: &'a Key) -> Self
    where
        Key: AsRef<[u8]> + ?Sized
    {
        Self {
            key: key.as_ref(),
            pos: 0,
        }
    }

    /// XOR each byte of the input buffer with a byte from the key.
    ///
    /// Note that this is stateful: repeated calls are likely to produce different results,
    /// even with identical inputs.
    pub fn munge_in_place(&mut self, data: &mut [u8]) {
        let key_len = self.key.len();
        (0..data.len()).for_each(|i| {
            data[i] ^= self.key[self.pos];
            self.pos += 1;
            self.pos %= key_len;
        });
    }

    /// XOR each byte of the data with a byte from the key.
    ///
    /// Note that this is stateful: repeated calls are likely to produce different results,
    /// even with identical inputs.
    ///
    /// Should accept anything which has a cheap conversion to a byte iterator.
    /// Shouldn't matter whether the byte iterator's values are owned or borrowed.
    pub fn munge<Data, Elem>(&mut self, data: Data) -> impl Iterator<Item = u8>
    where
        Data: IntoIterator<Item = Elem>,
        Elem: Borrow<u8>,
    {
        let key_len = self.key.len();
        data.into_iter().map(move |src| {
            let dst = src.borrow() ^ self.key[self.pos];
            self.pos += 1;
            self.pos %= key_len;
            dst
        })
    }

    /// Returns reader stream adaptor
    pub fn reader(self, reader: impl Read) -> impl Read {
        XorcismReader{ xor: self, reader }
    }

    /// Returns writer stream adaptor
    pub fn writer(self, writer: impl Write) -> impl Write {
        XorcismWriter{ xor: self, writer }
    }
}

pub struct XorcismReader<'a, R>
where
    R: Read + ?Sized,
{
    xor: Xorcism<'a>,
    reader: R,
}

pub struct XorcismWriter<'a, W>
where
    W: Write + ?Sized,
{
    xor: Xorcism<'a>,
    writer: W,
}

impl<'a, R> Read for XorcismReader<'a, R>
where
    R: Read + ?Sized,
{
    fn read(&mut self, buf: &mut [u8]) -> Result<usize> {
        let n = self.reader.read(buf)?;
        self.xor.munge_in_place(&mut buf[..n]);
        Ok(n)
    }
}

impl<'a, W> Write for XorcismWriter<'a, W>
where
    W: Write + ?Sized,
{
    fn write(&mut self, buf: &[u8]) -> Result<usize> {
        self.writer.write(self.xor.munge(buf.iter()).collect::<Vec<_>>().as_slice())
    }

    fn flush(&mut self) -> Result<()> {
        self.writer.flush()
    }
}