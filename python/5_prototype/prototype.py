from typing import Any
import copy


class Prototype:
    """
    Pythonではプロトタイプのインターフェースは`copy.copy`, `copy.deepcopy`を通して実行される。
    基本的に自分で実装する必要はないが、何らかのカスタマイズしたコピーを実行したい場合は`__copy__`, `__deepcopy__`をオーバーライドする。
    コピー処理をコピーするオブジェクト自身に任せることで、利用側はオブジェクトの詳細に関わらず簡単にコピーを作成できる。
    """

    def __init__(self, name: str, some_list_of_object: list[Any]) -> None:
        self._name = name
        self._some_list_of_object = some_list_of_object

    def __copy__(self) -> "Prototype":
        """
        shallowコピーを作成する。
        `copy.copy(object)`の形で呼び出されたときに実行され、shallowコピーを返す。
        """
        print(f"Create copy of {self._name} object.")

        some_list_of_object = copy.copy(self._some_list_of_object)
        new = self.__class__(self._name, some_list_of_object)
        new.__dict__.update(self.__dict__)

        return new

    def __deepcopy__(self, memo: dict = {}) -> "Prototype":
        """
        deepコピーを作成する。
        `copy.deepcopy(object)`の形で呼び出されたときに実行され、deepコピーを返す。
        memoは循環参照オブジェクトの時にcopy.deepcopy内で無限再帰に陥るのを避けるために使われる。
        """
        print(f"Create deep copy of {self._name} object.")

        some_list_of_object = copy.deepcopy(self._some_list_of_object, memo)

        new = self.__class__(self._name, some_list_of_object)
        new.__dict__ = copy.deepcopy(self.__dict__, memo)

        return new

    def __str__(self) -> str:
        return f"name: {self._name}, some_list_of_objects: {self._some_list_of_object}"
