import numpy as np


class ReplayBuffer:
    def __init__(self, max_buffer_size, state_shape, batch_size):
        self.mem_size = max_buffer_size
        self.mem_index = 0
        self.batch_size = batch_size
        self.state_memory = np.zeros((self.mem_size, *state_shape))
        self.reward_memory = np.zeros(self.mem_size)

    def store(self, state, reward):
        index = self.mem_index % self.mem_size
        self.state_memory[index] = state
        self.reward_memory[index] = reward
        self.mem_index = self.mem_index + 1

    def sample(self):
        max_mem = min(self.mem_index, self.mem_size)
        batch = np.random.choice(max_mem, self.batch_size)
        states = self.state_memory[batch]
        rewards = self.reward_memory[batch]
        return states, rewards

    def can_sample(self):
        return self.mem_index > self.batch_size * 20
