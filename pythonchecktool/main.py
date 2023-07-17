# main.py

from checks import AccessibleFromJumpHost
from utils.console_report import ConsoleReport

def main():
    # Define the host you want to check
    host = "http://localhost:25000"

    # Init report
    report = ConsoleReport()

    # Perform the check
    if AccessibleFromJumpHost.check(host):
        #print(f"{host} is accessible from the jumphost.")
        report.log_check('AccessibleFromJumpHost', True, 'Accessible from Jumphost')
    else:
        #print(f"{host} is NOT accessible from the jumphost.")
        report.log_check('AccessibleFromJumpHost', False, 'Not accessible from Jumphost')

    # Generate report
    report.generate_report()

if __name__ == "__main__":
    main()
