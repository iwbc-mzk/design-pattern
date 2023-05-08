from abc import ABC, abstractmethod
# 既存サービスでこのクラスのインターフェースを利用している
class Target:
    def http_request(self, url, port) -> str:
        return f"Target http_request: [url:{url}, port:{port}]"
    

class TagetInterface(ABC):
    @abstractmethod
    def http_request(self, url, port) -> str:
        pass
