# 現在利用しているTargetからこちらのクラスの機能に切り替えたい
class Adaptee:
    def request(self, url) -> str:
        return f"Adaptee: request [url:{url}]"
