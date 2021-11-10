from connect4.game import Game
import numpy as np


def recognise_patterns(model, gamma):
    v_score = 0.
    h_score = 0.
    d_score = 0.

    for i in range(0, 7):
        game = Game()
        game.board.grid[i][0] = -1
        game.board.grid[i][1] = -1
        game.board.grid[i][2] = -1
        v_score += abs(gamma - model.score_state(game.state()))
        game.board.grid[i][3] = -1
        v_score += abs(1 - model.score_state(game.state()))
        game.board.grid[i][0] = 1
        game.board.grid[i][1] = 1
        game.board.grid[i][2] = 1
        game.board.grid[i][3] = 1
        v_score += abs(-1 - model.score_state(game.state()))

    for i in range(0, 4):
        game = Game()
        game.board.grid[0 + i][0] = -1
        game.board.grid[1 + i][0] = -1
        game.board.grid[2 + i][0] = -1
        game.board.grid[3 + i][0] = -1
        h_score += abs(1 - model.score_state(game.state()))

    for i in range(0, 4):
        game = Game()
        game.board.grid[0 + i][0] = 1
        game.board.grid[1 + i][0] = 1
        game.board.grid[2 + i][0] = 1
        game.board.grid[3 + i][0] = 1
        h_score += abs(-1 - model.score_state(game.state()))

    for i in range(0, 4):
        game = Game()
        game.board.grid[0 + i][0] = -1
        game.board.grid[1 + i][1] = -1
        game.board.grid[2 + i][2] = -1
        game.board.grid[3 + i][3] = -1
        d_score += abs(1 - model.score_state(game.state()))

    for i in range(0, 4):
        game = Game()
        game.board.grid[0 + i][3] = -1
        game.board.grid[1 + i][2] = -1
        game.board.grid[2 + i][1] = -1
        game.board.grid[3 + i][0] = -1
        d_score += abs(1 - model.score_state(game.state()))

    print(v_score / 21)
    print(h_score / 8)
    print(d_score / 8)


def play_versus_random(model):
    action_space = [i for i in range(7)]
    score = 0
    for i in range(0, 10):
        connect4 = Game()
        while not connect4.done:
            action = choose_action(model, connect4)
            state, done, result = connect4.perform_action(action)
            if not done:
                action = np.random.choice(list(filter(lambda i: connect4.board.grid[i][5] == 0, action_space)))
                state, done, result = connect4.perform_action(action)
                if done:
                    score -= result
            else:
                score += result
    for i in range(0, 10):
        connect4 = Game()
        while not connect4.done:
            action = np.random.choice(list(filter(lambda i: connect4.board.grid[i][5] == 0, action_space)))
            _, done, result = connect4.perform_action(action)
            if not done:
                action = choose_action(model, connect4)
                _, done, result = connect4.perform_action(action)
                if done:
                    score += result
            else:
                score -= result
    return score


def choose_action(model, game):
    max_score = -100.0
    max_action = 0
    for i in range(0, 7):
        new_game = game.clone()
        state, done, result = new_game.perform_action(i)
        score = model.score_state(state)
        if score > max_score:
            max_action = i
    return max_action
