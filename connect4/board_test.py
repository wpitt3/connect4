from connect4.board import Board


def test_clone():
    board = Board()
    board2 = board.clone()
    board2.add_token(0, 1)
    assert 1 == board2.grid[0][0]
    assert 1 != board.grid[0][0]


def test_token_can_be_added():
    board = Board()
    assert 0 == board.grid[0][0]
    assert 0 == board.grid[0][1]
    board.add_token(0, 1)
    assert 1 == board.grid[0][0]
    board.add_token(0, -1)
    assert -1 == board.grid[0][1]


def test_full_board():
    board = Board()
    assert not board.full_board()
    board.add_token(0, 1)
    assert not board.full_board()
    for x in range(0, 7):
        for y in range(0, 6):
            board.grid[x][y] = 1
    assert board.full_board()


def test_find4_vertical():
    board = Board()
    assert 0 == board.find4()
    for y in range(0, 4):
        board.grid[0][y] = 1
    assert 1 == board.find4()
    for y in range(0, 4):
        board.grid[0][y] = -1
    assert -1 == board.find4()


def test_find4_horizontal():
    board = Board()
    assert 0 == board.find4()
    for x in range(0, 4):
        board.grid[x][0] = 1
    assert 1 == board.find4()
    for x in range(0, 4):
        board.grid[x][0] = -1
    assert -1 == board.find4()


def test_find4_diagonal_sw_ne():
    board = Board()
    assert 0 == board.find4()
    for i in range(0, 4):
        board.grid[i][i] = 1
    assert 1 == board.find4()


def test_find4_diagonal_nw_se():
    board = Board()
    assert 0 == board.find4()
    for i in range(0, 4):
        board.grid[i][3-i] = 1
    assert 1 == board.find4()
