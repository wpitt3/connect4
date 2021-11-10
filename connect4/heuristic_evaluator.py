

class Evaluator:
    def __init__(self):
        self.columns = 7
        self.rows = 6

    # not the smartest heuristic, does not always account for what is playable
    def score_state(self, state):
        score = 0
        for x in range(0, self.columns):
            for y in range(0, self.rows - 3):
                section_sum = sum(map(lambda i: state[x][y + i], list(range(0, 4))))
                if abs(section_sum) == 4:
                    return section_sum * 25
                if section_sum == 3:
                    score += 1
                if section_sum == -3:
                    score -= 10
        for x in range(0, self.columns - 3):
            for y in range(0, self.rows):
                section_sum = sum(map(lambda i: state[x + i][y], list(range(0, 4))))
                if abs(section_sum) == 4:
                    return section_sum * 25
                if section_sum == 3:
                    score += 1
                if section_sum == -3:
                    if y == 0:
                        score -= 7
                    score -= 3

        for x in range(0, self.columns - 3):
            for y in range(0, self.rows - 3):
                section_sum = sum(map(lambda i: state[x + i][y + i], list(range(0, 4))))
                if abs(section_sum) == 4:
                    return section_sum * 25

                if section_sum == 3:
                    score += 1
                if section_sum == -3:
                    score -= 3

                section_sum = sum(map(lambda i: state[x + i][y + 3 - i], list(range(0, 4))))
                if abs(section_sum) == 4:
                    return section_sum * 25

                if section_sum == 3:
                    score += 1
                if section_sum == -3:
                    score -= 3
