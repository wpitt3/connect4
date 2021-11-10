from az_agent.buffer import ReplayBuffer


def test_store_and_sample():
    buffer = ReplayBuffer(20, [1], 10)
    for i in range(0, 20):
        buffer.store([i], 1)

    states_1, _ = buffer.sample()
    states_2, _ = buffer.sample()
    # ensure randomly selected
    assert any(list(map(lambda j: states_1[j] != states_2[j], list(range(0, 10)))))
    # ensure random order (1/10! chance of failure)
    assert any(list(map(lambda j: states_1[j] > states_1[j+1], list(range(0, 9)))))


def test_store_past_max_replaces_old():
    buffer = ReplayBuffer(20, [1], 1)
    for i in range(0, 40):
        buffer.store([i], 1)
    states, _ = buffer.sample()
    assert states[0] > 19


def test_2d_state():
    buffer = ReplayBuffer(10, [2,2], 1)
    buffer.store([[1., 0.], [-1., 0.]], 1.)
    states, rewards = buffer.sample()
    assert states[0][0][0] == 1.
    assert states[0][0][1] == 0.
    assert states[0][1][0] == -1.
    assert states[0][1][1] == 0.
    assert rewards == [1.]
