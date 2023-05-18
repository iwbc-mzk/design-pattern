from abc import ABC, abstractmethod

from interface.searcher import ISearcher


# Creator Interface
class IGraghSearch(ABC):
    def __init__(self, edges: list[tuple], start: int) -> None:
        super().__init__()
        self._edges = edges
        self._start = start

    # Factory Method
    @abstractmethod
    def _factory_method(self, *args, **kwargs) -> ISearcher:
        pass

    # ファクトリーメソッドから返されるオブジェクトに依存したビジネスロジック
    # クライアント側が必要なのはこのビジネスロジックの部分
    # ファクトリーメソッドから異なる種類のプロダクトが返されることで、種類に応じたビジネスロジックが実行される
    # このビジネスロジックではプロダクトインターフェースを介して処理を行う
    def search(self) -> None:
        s = self._factory_method()
        print(*[i for i in s])
