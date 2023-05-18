from abc import ABC, abstractmethod
from collections.abc import Iterator


# Product Interface
class ISearcher(ABC, Iterator):
    @abstractmethod
    def __next__(self) -> int:
        pass
