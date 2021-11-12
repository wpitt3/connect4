from az_agent.model import Model
from az_agent.search_tree import SearchTree
import numpy as np


# The Agent is based off of Alpha Zero and contains, a model, a search tree and a buffer
# The Agent plays a game using the search tree and model to decide the action or a random action based upon epsilon
# The action is stored in the buffer and the model can then be trained upon that buffer every x steps
class Agent:
    def __init__(self, buffer, model, action_space, gamma, epsilon_init=1.0, epsilon_end=0.01, epsilon_dec=0.999):
        self.buffer = buffer
        self.model = model
        self.search_tree = SearchTree(action_space)
        self.epsilon = epsilon_init
        self.epsilon_end = epsilon_end
        self.epsilon_dec = epsilon_dec
        self.action_space = [i for i in range(action_space)]
        self.gamma = gamma

    def remember_game(self, states, rewards):
        multiplier = 1.0
        for i in range(len(states)-1, -1, -1):
            self.buffer.store(states[i], rewards[i] * multiplier)
            self.buffer.store(np.flip(states[i], axis=0), rewards[i] * multiplier)
            self.buffer.store(states[i]*-1, rewards[i] * -1 * multiplier)
            self.buffer.store(np.flip(states[i]*-1, axis=0), rewards[i] * -1 * multiplier)
            multiplier = multiplier * self.gamma

    def learn(self):
        if self.buffer.can_sample():
            states, results = self.buffer.sample()
            self.model.learn(states, results)
        self.epsilon = self.epsilon * self.epsilon_dec if self.epsilon > self.epsilon_end else self.epsilon_end

    def choose_action(self, game):
        rand = np.random.random()
        if rand < self.epsilon:
            return np.random.choice(list(filter(lambda i: game.board.grid[i][5] == 0, self.action_space)))
        else:
            return self.search_tree.choose_action(game, self.model)
