import numpy as np
import tensorflow as tf
import os

class ReplayBuffer:
    def __init__(self, max_buffer_size, state_shape, batch_size):
        self.mem_size = max_buffer_size
        self.mem_index = 0
        self.batch_size = batch_size
        self.state_shape = state_shape
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

    def is_full(self):
        return self.mem_index > self.mem_size

    def save_buffer(self):
        dataset = tf.data.Dataset.from_tensor_slices((self.state_memory, self.reward_memory))
        path = os.path.join('.', "buffer_data")
        tf.data.experimental.save(dataset, path)

    def load_buffer(self):
        self.state_memory = np.zeros((self.mem_size, *self.state_shape))
        self.reward_memory = np.zeros(self.mem_size)

        path = os.path.join('.', "buffer_data")
        buffer_data = tf.data.experimental.load(path)
        index = 0
        for elem in buffer_data:
            state = elem[0].numpy()
            reward = elem[1].numpy()
            if not elem[1].numpy() == 0.0:
                self.state_memory[index] = state
                self.reward_memory[index] = reward
                index += 1
            else:
                self.mem_index = index
                return
        self.mem_index = index

