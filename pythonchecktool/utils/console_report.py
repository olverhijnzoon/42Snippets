# File: utils/console_report.py
from interfaces.report import BaseReport

class ConsoleReport(BaseReport):

    def __init__(self):
        self.checks = {}

    def log_check(self, check_name: str, result: bool, details: str = ''):
        self.checks[check_name] = {
            'result': 'Success' if result else 'Failure',
            'details': details
        }

    def generate_report(self):
        print(f"---Report------------------------------------")
        for check_name, check_info in self.checks.items():
            print(f'{check_name}: {check_info["result"]}')
            if check_info['details']:
                print(f'  Details: {check_info["details"]}')
