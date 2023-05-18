from interface.gragh_search import IGraghSearch
from interface.searcher import ISearcher
from concrete.dfs import DFS


class DFSCreator(IGraghSearch):
    def _factory_method(self) -> ISearcher:
        return DFS(self._edges, self._start)
