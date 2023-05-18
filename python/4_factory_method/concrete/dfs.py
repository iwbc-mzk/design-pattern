from queue import LifoQueue
from collections import defaultdict

from interface.searcher import ISearcher


class DFS(ISearcher):
    def __init__(self, edges: list[tuple], start: int) -> None:
        super().__init__()
        self._queue: LifoQueue = LifoQueue()
        self._queue.put(start)

        self._visited: set = set()
        self._visited.add(start)

        self._create_gragh(edges)

    def _create_gragh(self, edges: list[tuple]):
        self._gragh = defaultdict(set)
        for u, v in edges:
            self._gragh[u].add(v)
            self._gragh[v].add(u)

    def __next__(self) -> int:
        if self._queue.empty():
            raise StopIteration()

        next = self._queue.get()
        for v in self._gragh[next]:
            if v not in self._visited:
                self._queue.put(v)
                self._visited.add(v)

        return next
