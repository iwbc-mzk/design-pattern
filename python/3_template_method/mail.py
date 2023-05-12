from otp import OneTimePassword


class Email(OneTimePassword):
    def __init__(self) -> None:
        super().__init__()

    def _getMessage(self) -> str:
        print("Email: Get message.")
        return f"Email: one time password: {self._password}"

    def _sendNotification(self, message) -> None:
        print(f"Email: Send otp by Email. message [{message}]")
