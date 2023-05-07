from collections.abc import Iterable
from typing import Any, List

from words_collection_iterator import WordsCollectionIterator


class WordsCollection(Iterable):
    """
    具象Iterableクラス。
    __iter__メソッドでイテレータを返す。
    """
    def __init__(self, collection: List[Any] = []) -> None:
        super().__init__()
        self._collection = collection

    def __iter__(self) -> WordsCollectionIterator:
        return WordsCollectionIterator(self)

    def get_reverse_iterator(self) -> WordsCollectionIterator:
        return WordsCollectionIterator(self, True)

    def add_item(self, item: Any):
        self._collection.append(item)

    def get_item_at(self, index: int) -> Any:
        return self._collection[index]
