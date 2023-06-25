from threading import Lock
from typing import Any


class SingletonMeta(type):
    """
    スレッドセーフシングルトン
    """

    _instance = None
    _lock: Lock = Lock()

    def __call__(cls, *args: Any, **kwds: Any) -> Any:
        with cls._lock:
            if not cls._instance:
                print("create instance")
                cls._instance = super().__call__(*args, **kwds)

        return cls._instance

    def __new__(cls, name, bases, dict):
        return super().__new__(cls, name, bases, dict)


class Singleton(metaclass=SingletonMeta):
    def __init__(self, value: str) -> None:
        self._value = value

    def __str__(self) -> str:
        return self._value
