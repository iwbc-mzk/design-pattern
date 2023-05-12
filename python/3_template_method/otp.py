from abc import ABC, abstractmethod
from uuid import uuid4


class OneTimePassword(ABC):
    def _generatePassword(self) -> None:
        self._password = uuid4()
        print(f"Generate OTP. [{self._password}]")

    # hook
    def _saveToCache(self) -> None:
        pass

    @abstractmethod
    def _getMessage(self) -> str:
        pass

    @abstractmethod
    def _sendNotification(self, message) -> None:
        pass

    # Template Method
    def genAndSendOTP(self) -> None:
        self._generatePassword()
        self._saveToCache()  # hook
        msg = self._getMessage()
        self._sendNotification(msg)
