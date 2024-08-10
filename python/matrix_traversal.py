def traverse(matrix, start):
    directions = [(2, 0), (-1, 1), (-1, 1), (2, 0)]
    current_pos = start

    print('Starting Position:', current_pos)
    draw_matrix(matrix, current_pos)

    for direction in directions:
        for _ in range(abs(direction[0])):
            current_pos[0] += int(direction[0] / abs(direction[0]))
            draw_matrix(matrix, current_pos)
        for _ in range(abs(direction[1])):
            current_pos[1] += int(direction[1] / abs(direction[1]))
            draw_matrix(matrix, current_pos)

    print('\nMy final position: ', current_pos)


def draw_matrix(matrix, curr):
    print('Position: ', curr)
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            if [j, i] == curr:
                print('*', end='')
            else:
                print('=', end='')
        print()
    print()


matrix = [[0, 0, 0],
          [0, 0, 0],
          [0, 0, 0]]

starting_pos = [0, 0]

traverse(matrix, starting_pos)
