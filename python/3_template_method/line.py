from otp import OneTimePassword


class Line(OneTimePassword):
    def __init__(self) -> None:
        super().__init__()

    # テンプレートメソッド以外ならオーバーライド可能
    def _generatePassword(self):
        self._password = "line-otp-123456-9384-111"
        print(f"Line: Generate otp. [{self._password}]")

    # 必要ならhookを実装できる
    def _saveToCache(self) -> None:
        print("Line: Save otp to cache.")

    def _getMessage(self) -> str:
        print("Line: Get message.")
        return f"Line: one time password: {self._password}"

    def _sendNotification(self, message) -> None:
        print(f"Line: Send otp by Line. message [{message}]")
