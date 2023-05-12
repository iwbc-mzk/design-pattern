from mail import Email
from line import Line


def main():
    email_otp = Email()
    email_otp.genAndSendOTP()

    print()

    line_otp = Line()
    line_otp.genAndSendOTP()


if __name__ == "__main__":
    main()
