from interface.gragh_search import IGraghSearch
from interface.searcher import ISearcher
from concrete.bfs import BFS


class BFSCreator(IGraghSearch):
    def _factory_method(self) -> ISearcher:
        return BFS(self._edges, self._start)
