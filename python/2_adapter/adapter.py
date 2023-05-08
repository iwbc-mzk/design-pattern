from target import Target, TagetInterface
from adaptee import Adaptee


class AdapterByInheritance(TagetInterface, Adaptee):
    def http_request(self, url, port) -> str:
        return super().request(f"{url}:{port}")


class AdapterByDelegate(Target):
    def __init__(self) -> None:
        super().__init__()
        self._adaptee = Adaptee()

    def http_request(self, url, port) -> str:
        return self._adaptee.request(f"{url}:{port}")
