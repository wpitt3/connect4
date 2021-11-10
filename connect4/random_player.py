from player import Player
import random


class RandomPlayer(Player):

    def get_move(self, state):
        return random.randint(0, 6)
