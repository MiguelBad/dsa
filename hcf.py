"""
def main(a, b):
    while a != b:
        if a > b:
            a = a - b

        else:
            b = b - a

    print(a)

def main(a, b):
    while b != 0:
        a, b = b, a % b

    print(a)
"""


def main(a, b):
    while b != 0:
        a, b = b, a % b

    print(a)


main(9, 6)
