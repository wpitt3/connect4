from az_agent.agent import Agent
from az_agent.buffer import ReplayBuffer
from az_agent.model import Model
from connect4.game import Game
from keras.models import Sequential
from keras.layers import BatchNormalization
import tensorflow as tf
from tensorflow import keras
import os
import numpy as np
from connect4ModelChecker import play_versus_mcst
from monte_carlo_tree_search import Model as McstModel


def warm_buffer(agent):
    print('Filling agent buffer')
    mcst_model = McstModel(5000)
    while not agent.buffer.is_full():
        print(agent.buffer.mem_index)
        connect4 = Game()
        states = []
        rewards = []
        while not connect4.done:
            action = mcst_model.call(connect4.state())
            connect4.perform_action(action)
            states.append(connect4.state())

        if connect4.result == 1:
            for i in range(0, len(states)):
                rewards.append(1.0 - (i % 2) * 2)
        else:
            for i in range(0, len(states)):
                rewards.append(1.0 - ((i + 1) % 2) * 2)
        agent.remember_game(states, rewards)
    agent.buffer.save_buffer()

if __name__ == '__main__':
    state_space = [7, 6]
    action_space = 7

    model = Sequential([
        keras.layers.Conv2D(150, 4, activation="relu", input_shape=[7, 6, 1]),
        keras.layers.Conv2D(200, 2, activation="relu", padding="same"),
        # keras.layers.Conv2D(150, 2, activation="relu", padding="same"),
        # keras.layers.Flatten(input_shape=[7, 6]),
        keras.layers.Flatten(),
        BatchNormalization(),
        # keras.layers.Dense(200, activation="relu"),
        keras.layers.Dense(200, activation="relu"),
        BatchNormalization(),
        keras.layers.Dense(100, activation="relu"),
        BatchNormalization(),
        keras.layers.Dense(1, activation="tanh")
    ])
    model.compile(loss="mse", optimizer=tf.keras.optimizers.Adam(learning_rate=0.00001))
    print(model.summary())

    model_wrapper = Model(model,  'conv')
    model_wrapper.load_model()
    agent = Agent(ReplayBuffer(100000, state_space, 64), model_wrapper, action_space, 0.95, 1.0, 0.1, 0.9995)

    # for i in range(5):
    #     print("START")
    #     print(play_versus_mcst(model_wrapper, 20, 50))

    agent.buffer.load_buffer()
    # warm_buffer(agent)
    # print("buffer warmed")

    # for i in range(100000):
    #     if i % 10000 == 9999:
    #         print(play_versus_mcst(model_wrapper, 20, 10))
    #         print(play_versus_mcst(model_wrapper, 20, 50))
    #         print(play_versus_mcst(model_wrapper, 20, 100))
    #         # print(play_versus_mcst(model_wrapper, 20, 1000))
    #         # print(play_versus_mcst(model_wrapper, 20, 10000))
    #         model_wrapper.save_model()
    #     if i%500 == 0:
    #         print(i)
    #     agent.learn()
    #
    #
    for i in range(50000):
        if i%10 == 0:
            print(i)
            if i % 100 == 0:
                print(agent.epsilon)
                print(play_versus_mcst(model_wrapper, 20, 10))
                print(play_versus_mcst(model_wrapper, 20, 50))
                print(play_versus_mcst(model_wrapper, 20, 100))
        if i % 1000 == 999:
            model_wrapper.save_model()
        connect4 = Game()
        states = []
        rewards = []
        while not connect4.done:
            action = agent.choose_action(connect4)
            connect4.perform_action(action)
            states.append(connect4.state())

        if connect4.result == 1:
            for i in range(0, len(states)):
                rewards.append(1.0 - (i % 2) * 2)
        else:
            for i in range(0, len(states)):
                rewards.append(1.0 - ((i + 1) % 2) * 2)

        agent.remember_game(states, rewards)
        agent.learn()

    model_wrapper.save_model()


