from az_agent.agent import Agent
from az_agent.buffer import ReplayBuffer
from az_agent.model import Model
from connect4.game import Game
from keras.models import Sequential
from keras.layers import BatchNormalization
import tensorflow as tf
from tensorflow import keras
import numpy as np
from connect4ModelChecker import recognise_patterns, play_versus_random

def warm_buffer(agent):
    for i in range(100):
        connect4 = Game()
        states = []
        rewards = []
        while not connect4.done:
            action = agent.choose_action(connect4)
            state, _, _ = connect4.perform_action(action)
            states.append(state)

        if connect4.result == 1:
            for i in range(0, len(states)):
                rewards.append(1.0 - (i % 2) * 2)
        else:
            for i in range(0, len(states)):
                rewards.append(1.0 - ((i + 1) % 2) * 2)
        agent.remember_game(states[:10], rewards[:10])

if __name__ == '__main__':
    state_space = [7, 6]
    action_space = 7

    model = Sequential([
        keras.layers.Conv2D(150, 4, activation="relu", input_shape=[7, 6, 1]),
        # keras.layers.Conv2D(200, 2, activation="relu", padding="same"),
        # keras.layers.Conv2D(150, 2, activation="relu", padding="same"),
        # keras.layers.Flatten(input_shape=[7, 6]),
        keras.layers.Flatten(),
        BatchNormalization(),
        # keras.layers.Dense(200, activation="relu"),
        BatchNormalization(),
        keras.layers.Dense(200, activation="relu"),
        BatchNormalization(),
        keras.layers.Dense(100, activation="relu"),
        BatchNormalization(),
        keras.layers.Dense(1, activation="tanh")
    ])
    model.compile(loss="mse", optimizer=tf.keras.optimizers.Adam(learning_rate=0.00001))
    print(model.summary())

    model_wrapper = Model(model,  'conv')
    # model_wrapper.load_model()
    agent = Agent(ReplayBuffer(50000, state_space, 64), model_wrapper, action_space, 0.95, 1.0, 0.1, 0.9995)

    warm_buffer(agent)

    for i in range(5000):
        if i%10 == 0:
            print(i)
            if i % 100 == 0:
                print(agent.epsilon)
                # print(play_versus_random(model_wrapper))
                recognise_patterns(model_wrapper, 0.95)
        connect4 = Game()
        states = []
        rewards = []
        while not connect4.done:
            action = agent.choose_action(connect4)
            state, _, _ = connect4.perform_action(action)
            states.append(state)

        if connect4.result == 1:
            for i in range(0, len(states)):
                rewards.append(1.0 - (i % 2) * 2)
        else:
            for i in range(0, len(states)):
                rewards.append(1.0 - ((i + 1) % 2) * 2)

        agent.remember_game(states, rewards)
        agent.learn()

    model_wrapper.save_model()


