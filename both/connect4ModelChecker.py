from connect4.game import Game
from monte_carlo_tree_search import Model


def play_versus_mcst(model, games, mcstmoves):
    score = 0
    start_player = 1
    mcst_model = Model(mcstmoves)

    for i in range(0, games):
        connect4 = Game()
        player = start_player
        while not connect4.done:
            if player == 1:
                action = choose_action(model, connect4)
                connect4.perform_action(action)
            else:
                action = mcst_model.call(connect4.state())
                connect4.perform_action(action)
            player *= -1
        score += connect4.result * start_player

        start_player *= -1
    return score


def choose_action(model, game):
    max_score = -100.0
    max_action = 0
    for i in range(0, 7):
        new_game = game.clone()
        new_game.perform_action(i)
        score = model.score_state(new_game.state())
        if score > max_score:
            max_action = i
    return max_action
