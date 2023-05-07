from words_collection import WordsCollection


def main() -> None:
    collection = WordsCollection()
    collection.add_item("Around the world")
    collection.add_item("Bible")
    collection.add_item("Cinderella")

    # クライアント側はコレクションが要素をどのように管理しているか(リストや辞書など)に関わらず
    # 同じ方法でイテレーションできる。
    print("=== Straight order ===")
    for c in collection:
        print(c)

    print()

    print("=== Reverse order ===")
    for c in collection.get_reverse_iterator():
        print(c)


if __name__ == "__main__":
    main()
