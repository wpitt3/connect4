

class SearchTree:
    def __init__(self, action_space):
        self.action_space = action_space

    def choose_action(self, game, model):
        max_score = -100.0
        max_action = 0
        for i in range(0, self.action_space):
            new_game = game.clone()
            new_game.perform_action(i)
            score = model.score_state(new_game.state())
            if new_game.done:
                score = new_game.result
            if score > max_score:
                max_action = i
        return max_action
