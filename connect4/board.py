import numpy as np


class Board:

    def __init__(self):
        self.columns = 7
        self.rows = 6
        self.grid = np.array([[0 for y in range(self.rows)] for x in range(self.columns)])

    def clone(self):
        board = Board()
        board.grid = np.copy(self.grid)
        return board

    def add_token(self, column, player):
        for y in range(0, self.rows):
            if self.grid[column][y] == 0:
                self.grid[column][y] = player
                return True
        return False

    def full_board(self):
        for x in range(0, self.columns):
            if self.grid[x][self.rows-1] == 0:
                return False
        return True

    def find4(self):
        for x in range(0, self.columns):
            for y in range(0, self.rows - 3):
                section_sum = sum(map(lambda i: self.grid[x][y + i], list(range(0, 4))))
                if abs(section_sum) == 4:
                    return section_sum / 4

        for x in range(0, self.columns - 3):
            for y in range(0, self.rows):
                section_sum = sum(map(lambda i: self.grid[x + i][y], list(range(0, 4))))
                if abs(section_sum) == 4:
                    return section_sum / 4

        for x in range(0, self.columns - 3):
            for y in range(0, self.rows - 3):
                section_sum = sum(map(lambda i: self.grid[x + i][y + i], list(range(0, 4))))
                if abs(section_sum) == 4:
                    return section_sum / 4

                section_sum = sum(map(lambda i: self.grid[x + i][y + 3 - i], list(range(0, 4))))
                if abs(section_sum) == 4:
                    return section_sum / 4

        return 0
