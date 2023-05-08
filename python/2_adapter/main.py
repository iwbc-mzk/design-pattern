from target import Target
from adapter import AdapterByInheritance, AdapterByDelegate


def main():
    url = "https://example.co.jp"
    port = "501"

    # By Inheritance
    print("=== By Inheritance ===")

    target = Target()
    print("Target: ", target.http_request(url, port))

    ada_inhe = AdapterByInheritance()
    print("Adapter: ", ada_inhe.http_request(url, port))

    print()

    # By Delegate
    print("=== By Delegate ===")

    target = Target()
    print("Target: ", target.http_request(url, port))

    ada_dele = AdapterByDelegate()
    print("Adapter: ", ada_dele.http_request(url, port))


if __name__ == "__main__":
    main()
