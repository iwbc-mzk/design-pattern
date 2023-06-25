from threading import Thread

from singleton import Singleton


def test_singleton(value: str) -> None:
    s = Singleton(value)
    print(s)


def main():
    p1 = Thread(target=test_singleton, args=("singleton1",))
    p2 = Thread(target=test_singleton, args=("singleton2",))

    p1.start()
    p2.start()

    p3 = Thread(target=test_singleton, args=("singleton3",))
    p3.start()


if __name__ == "__main__":
    main()
