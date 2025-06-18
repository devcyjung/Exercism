from io import BufferedRandom
from types import TracebackType
from typing import Protocol, Type, runtime_checkable

class MeteredFile(BufferedRandom):

    def __init__(self, *args, **kwargs) -> None:
        super().__init__(*args, **kwargs)
        self.__read_ops: int = 0
        self.__read_bytes: int = 0
        self.__write_ops: int = 0
        self.__write_bytes: int = 0

    def __enter__(self) -> 'MeteredFile':
        return self

    def __exit__(
        self,
        exc_type: Type[BaseException] | None,
        exc_val: BaseException | None,
        exc_tb: TracebackType | None
    ) -> bool | None:
        return super().__exit__(exc_type, exc_val, exc_tb)

    def __iter__(self) -> 'MeteredFile':
        return self

    def __next__(self) -> bytes:
        result: bytes = super().readline()
        if not result:
            raise StopIteration
        self.__read_ops += 1
        self.__read_bytes += len(result)
        return result

    def read(self, size: int = -1) -> bytes:
        result: bytes = super().read(size)
        self.__read_ops += 1
        self.__read_bytes += len(result)
        return result

    @property
    def read_bytes(self) -> int:
        return self.__read_bytes

    @property
    def read_ops(self) -> int:
        return self.__read_ops

    def write(self, b: bytes) -> int:
        written: int = super().write(b)
        self.__write_ops += 1
        self.__write_bytes += written
        return written

    @property
    def write_bytes(self) -> int:
        return self.__write_bytes

    @property
    def write_ops(self) -> int:
        return self.__write_ops

@runtime_checkable
class ContextManager(Protocol):
    def __enter__(self) -> 'ContextManager': ...
    def __exit__(
        self,
        exc_type: Type[BaseException] | None,
        exc_val: BaseException | None,
        exc_tb: TracebackType | None
    ) -> bool | None: ...

@runtime_checkable
class SendRecv(Protocol):
    def send(self, data: bytes, flags: int | None) -> int: ...
    def recv(self, bufsize: int, flags: int | None) -> bytes: ...

@runtime_checkable
class Socket(ContextManager, SendRecv, Protocol): ...

class MeteredSocket:

    def __init__(self, socket: Socket) -> None:
        self.__socket: Socket = socket
        self.__recv_bytes: int = 0
        self.__recv_ops: int = 0
        self.__send_bytes: int = 0
        self.__send_ops: int = 0

    def __enter__(self) -> 'MeteredSocket':
        return self

    def __exit__(
        self,
        exc_type: Type[BaseException] | None,
        exc_val: BaseException | None,
        exc_tb: TracebackType | None
    ):
        return self.__socket.__exit__(exc_type, exc_val, exc_tb)

    def recv(self, bufsize: int, flags: int = 0) -> bytes:
        result: bytes = self.__socket.recv(bufsize, flags)
        self.__recv_bytes += len(result)
        self.__recv_ops += 1
        return result

    @property
    def recv_bytes(self) -> int:
        return self.__recv_bytes

    @property
    def recv_ops(self) -> int:
        return self.__recv_ops

    def send(self, data: bytes, flags: int = 0) -> int:
        sent_bytes: int = self.__socket.send(data, flags)
        self.__send_bytes += sent_bytes
        self.__send_ops += 1
        return sent_bytes

    @property
    def send_bytes(self) -> int:
        return self.__send_bytes

    @property
    def send_ops(self) -> int:
        return self.__send_ops
