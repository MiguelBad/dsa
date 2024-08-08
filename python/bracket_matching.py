def main(input):
    pairs = {'(': ')', '[': ']', '{': '}'}

    '''
        note to self:
        using sets is more performant than lists on membership test
    '''

    left = set(pairs.keys())
    right = set(pairs.values())
    left_so_far = []

    for i in input:
        if i in left:
            left_so_far.append(i)

        elif i in right:
            if not left_so_far:
                return False

            if pairs[left_so_far[-1]] == i:
                left_so_far.pop()
            else:
                return False

        else:
            continue

    return len(left_so_far) == 0


test = ['( )(( )){([( )])}',
        '((( )(( )){([( )])}',
        ')(( )){([( )])}',
        '({[ ])}',
        '(',]

for i in test:
    print(main(i))
