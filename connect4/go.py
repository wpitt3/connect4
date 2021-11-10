from connect4.game import Game
from random_player import RandomPlayer


player = RandomPlayer()

# print(connect4.board.grid)

results = [0, 0, 0]
for i in range(0, 10000):
    game = Game()
    while not game.done:
        game.perform_action(player.get_move(game.board.grid))
    results[game.result+1] = results[game.result+1] + 1

print(results)