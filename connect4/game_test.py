from connect4.game import Game


def test_clone():
    game = Game()
    game2 = game.clone()
    game.perform_action(0)
    game.done = True
    assert 1 == game.board.grid[0][0]
    assert 1 != game2.board.grid[0][0]
    assert -1 == game.player_index
    assert 1 == game2.player_index
    assert not game2.done


def test_perform_action_on_empty():
    game = Game()
    assert 0 == game.board.grid[0][0]
    assert 1 == game.player_index
    game.perform_action(0)
    assert not game.result
    assert not game.done
    assert 1 == game.state()[0][0]
    assert -1 == game.player_index


def test_perform_action_invalid_move():
    game = Game()
    for i in range(0, 6):
        game.perform_action(0)
    assert not game.result
    assert not game.done
    game.perform_action(0)
    assert game.done
    assert -1 == game.result


def test_perform_action_winner_with_4():
    game = Game()
    game.perform_action(0)
    game.perform_action(1)
    game.perform_action(0)
    game.perform_action(1)
    game.perform_action(0)
    game.perform_action(1)
    assert not game.result
    assert not game.done
    game.perform_action(0)
    assert game.done
    assert 1 == game.result


def test_perform_action_full_board():
    game = Game()
    for x in range(0, 7, 2):
        for y in range(0, 3):
            game.board.grid[x][y] = 1
            game.board.grid[x][y+3] = -1
    for x in range(1, 6, 2):
        for y in range(0, 3):
            game.board.grid[x][y] = -1
            game.board.grid[x][y+3] = 1
    game.board.grid[6][5] = 0
    game.player_index = -1
    assert not game.result
    assert not game.done
    assert not game.board.full_board()

    game.perform_action(6)
    assert game.done
    assert 0 == game.result
    assert game.board.full_board()
