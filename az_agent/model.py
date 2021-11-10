from keras.models import load_model
import numpy as np


class Model:
    def __init__(self, model, name):
        self.model = model
        self.name = name

    def learn(self, states, rewards):
        self.model.fit(np.expand_dims(states, axis=3), rewards, verbose=0)

    def score_state(self, state):
        return self.model.predict(np.expand_dims(np.expand_dims(state, axis=2), axis=0))

    def save_model(self):
        self.model.save(self.name)

    def load_model(self):
        self.model = load_model(self.name)
