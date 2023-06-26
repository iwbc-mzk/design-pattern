import copy

from prototype import Prototype


if __name__ == "__main__":
    """
    Prototypeパターンのおかげでオブジェクトの詳細について知らなくとも、copy.copy, copy.deepcopyを呼び出すだけで
    オブジェクトのコピーを作成できる。
    """
    list_of_objects = [1, [2, 34], [111, 394, 2]]

    prototype = Prototype("prototype1", list_of_objects)
    print("Original: ", prototype)

    print("\n==== Shallow Copy ====")
    shallow_copy = copy.copy(prototype)
    print("Shallow copy: ", shallow_copy)


    print("\n==== Add object to list ====")
    shallow_copy._some_list_of_object.append([1, 2, 3, 4])
    print("Original: ", prototype)
    print("Shallow copy: ", shallow_copy)

    print("\n==== Deep Copy ====")
    deep_copy = copy.deepcopy(prototype)
    print("Deep copy: ", deep_copy)

    print("\n==== Add object to list ====")
    deep_copy._some_list_of_object.append([9, 99, 999])
    print("Original: ", prototype)
    print("Deep copy: ", deep_copy)

