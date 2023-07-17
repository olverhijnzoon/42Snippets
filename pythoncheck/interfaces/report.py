# File: interfaces/report.py
from abc import ABC, abstractmethod

class BaseReport(ABC):

    @abstractmethod
    def log_check(self, check_name: str, result: bool, details: str = ''):
        pass

    @abstractmethod
    def generate_report(self):
        pass
