import numpy as np
import ctypes
import numpy.ctypeslib as npct


class Model:
    def __init__(self):
        self.c4mcts = npct.load_library("c4mcts.so", "../go")
        array_1d_int = npct.ndpointer(dtype=np.int_, ndim=1, flags='CONTIGUOUS')
        self.c4mcts.Connect4MCTS.restype = ctypes.c_int
        self.c4mcts.Connect4MCTS.argtypes = [array_1d_int, ctypes.c_int]

    def call(self, state):
        flat_state = state.flatten() * -1
        return self.c4mcts.Connect4MCTS(flat_state, 5000)
