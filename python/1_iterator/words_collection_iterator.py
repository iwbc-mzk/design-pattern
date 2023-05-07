from __future__ import annotations
from collections.abc import Iterator
from typing import TYPE_CHECKING


if TYPE_CHECKING:
    from words_collection import WordsCollection


class WordsCollectionIterator(Iterator):
    """
    具象Iteratorクラス。
    集合体の探索方法を実装する。
    __next__メソッドで順に集合体から要素を取り出す。
    """
    def __init__(
        self,
        collection: WordsCollection,
        reverse: bool = False,
    ) -> None:
        super().__init__()
        self._collection = collection
        self._reverse = reverse
        self._index = -1 if reverse else 0

    def __next__(self):
        """
        集合体中のアイテムを端から順番に1つずつ取り出す。
        取り出せなくなったら StopIteration エラーを送出する。
        => Pythonの言語仕様。Iterator参照。
        """
        try:
            value = self._collection.get_item_at(self._index)
            self._index += -1 if self._reverse else 1
        except IndexError:
            raise StopIteration()

        return value
