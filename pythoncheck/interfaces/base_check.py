# interfaces/base_check.py

from abc import ABC, abstractmethod

class BaseCheck(ABC):
    @abstractmethod
    def check(self):
        pass
