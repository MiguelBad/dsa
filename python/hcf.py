def iterative_hcf(a, b):
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


a = 9
b = 6

main(a, b)
iterative_hcf(a, b)
