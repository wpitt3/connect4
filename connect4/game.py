from connect4.board import Board
import numpy as np


class Game:
    def __init__(self):
        self.board = Board()
        self.done = False
        self.result = None
        self.player_index = 1

    def state(self):
        return np.copy(self.board.grid) * self.player_index * -1

    def clone(self):
        game = Game()
        game.board = self.board.clone()
        game.done = self.done
        game.player_index = self.player_index
        return game

    def perform_action(self, action):
        if self.board.add_token(action, self.player_index):
            board_result = self.board.find4()
            if not board_result == 0:
                self.done = True
                self.result = self.player_index
            if self.board.full_board():
                self.done = True
                self.result = 0
        else:
            self.done = True
            self.result = self.player_index * -1

        self.player_index = self.player_index * -1
        return self.state(), self.done, self.result
