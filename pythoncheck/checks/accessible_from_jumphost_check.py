# checks/accessible_from_jumphost_check.py

from interfaces.base_check import BaseCheck

import requests

class AccessibleFromJumpHost(BaseCheck):
    @staticmethod
    def check(hostname):
        try:
            response = requests.get(hostname)
            # If the response was successful, no Exception will be raised
            response.raise_for_status()
        except requests.exceptions.HTTPError as http_err:
            #print(f'HTTP error occurred: {http_err}')
            return False
        except Exception as err:
            #print(f'Other error occurred: {err}')
            return False
        else:
            #print('Success!')
            return True
