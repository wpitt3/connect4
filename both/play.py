from az_agent.agent import Agent
from az_agent.buffer import ReplayBuffer
from az_agent.model import Model
from connect4.game import Game
from keras.models import Sequential
from keras.layers import BatchNormalization
import tensorflow as tf
from tensorflow import keras
import numpy as np
from monte_carlo_tree_search import Model
import os

if __name__ == '__main__':

    # for i in range(20, 0, -1):
    #     print(i)
    game = Game()

    game.perform_action(5)
    game.perform_action(6)
    Model(50).call(game.state())
    Model(50).call(game.state())
