from typing import Literal

from concrete.bfs_creator import BFSCreator
from concrete.dfs_creator import DFSCreator
from interface.gragh_search import IGraghSearch


class App:
    def __init__(
        self, mode: Literal["BFS", "DFS"], edges: list[tuple], start: int
    ) -> None:
        self._mode = mode

        if mode == "BFS":
            self._searcher: IGraghSearch = BFSCreator(edges, start)
        elif mode == "DFS":
            self._searcher: IGraghSearch = DFSCreator(edges, start)
        else:
            raise Exception()

    def main(self):
        print(self._mode)
        self._searcher.search()


def main():
    """
    1 - 3 - 2 - 4 - 8
    |   |
    6   5
    |
    7 - 9
    """
    edges = [(1, 3), (3, 2), (2, 4), (3, 5), (1, 6), (6, 7), (4, 8), (7, 9)]
    start = 1

    for mode in ["BFS", "DFS"]:
        app = App(mode, edges, start)
        app.main()


if __name__ == "__main__":
    main()
