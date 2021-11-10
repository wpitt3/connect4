import React from 'react';
import './App.css';

class Board extends React.Component {
  // computer is always 1 (and is red?)

  constructor(props) {
    super(props);

    let firstPlayer = -1
    let player = firstPlayer
    let grid = []
    let done = false
    let result = 0
    for(let i = 0; i < 7; i++) {
      grid.push([]);
      for(let j = 0; j < 6; j++) {
          grid[i][j] = 0;
      }
    }

    this.state = {
      grid,
      player,
      firstPlayer,
      done,
      result
    };
    this.handleColumnClick = this.handleColumnClick.bind(this);
    this.handleFPClick = this.handleFPClick.bind(this);
    this.handleRestart = this.handleRestart.bind(this);
    this.performMove = this.performMove.bind(this);
  }

  async handleColumnClick(index) {
    let {grid, player, done} = this.state;

    let indexOfSpace = grid[index].indexOf(0);
    if (indexOfSpace !== -1 && player === -1 && !done) {
      grid[index][indexOfSpace] = player
      player = player * -1
      this.setState({
        grid,
        player
      })
      this.performMove(grid, player)
    }
  }

  handleFPClick() {
    let {firstPlayer, grid} = this.state;

    if (!grid.some(x => x[0] !== 0)) {
      firstPlayer = firstPlayer * -1
      this.setState({
        firstPlayer,
        player: firstPlayer,
      })
    }
    this.waitBeforeGettingMove()
  }

  handleRestart() {
    let {firstPlayer} = this.state;
    let grid = []
    for(let i = 0; i < 7; i++) {
      grid.push([]);
      for(let j = 0; j < 6; j++) {
          grid[i][j] = 0;
      }
    }
    this.setState({
      grid,
      player: firstPlayer,
      done: false
    })
    this.waitBeforeGettingMove()
  }

  waitBeforeGettingMove() {
    setTimeout(function() {
      let {grid, player} = this.state;
      if (player === 1) {
        this.performMove(grid, player)
      }
    }.bind(this), 700)
  }

  async performMove(grid, player) {
    let {done} = this.state;
    if (!done) {
      const response = await fetch("http://localhost:8080/", {
        method: "POST",
        body: JSON.stringify(grid),
        mode: 'cors',
        headers: {
          'Content-Type': 'application/json',
          'Access-Control-Allow-Origin': '*',
        },
      });
      let {BestMove, Full, Result} = await response.json();
      player = this.state.player;
      if (BestMove !== -1) {
        let indexOfSpace = grid[BestMove].indexOf(0);
        if (indexOfSpace !== -1 && player === 1 && !done) {
          grid[BestMove][indexOfSpace] = player
          player = player * -1
        }
      }

      this.setState({
        grid,
        player,
        done: Full || done,
        result: Result
      })
    }
  }

  render() {
    return (
      <div>
        <div className="next-player">Next player: { this.state.player === 1 ? <div className="red-token"/> : <div className="yellow-token"/> }</div>
        <div className="first-player" onClick={() => this.handleFPClick()}>First player: { this.state.firstPlayer === 1 ? <div className="red-token"/> : <div className="yellow-token"/> }</div>
        <div className="restart" onClick={() => this.handleRestart()}>Restart</div>
        <div className="grid">
          {this.state.grid.map( (column, index) => {
            return <div key={index} index={index} onClick={() => this.handleColumnClick(index)} className='column'>
              {column.reverse().map( (cell, cindex) => {
                let inputStyle = {}
                if (cell === 1) {
                  inputStyle = {
                    background: '#d21404'
                  }
                }
                if (cell === -1) {
                  inputStyle = {
                    background: '#fce245'
                  }
                }
                return <div key={cindex} className='cell' style={inputStyle}></div>
              })}
            </div>
          })}
        </div>
        {!this.state.done || <div className="winner">{this.state.result === -1 ? 'You Won' : (this.state.result === 0 ? 'DRAW' : 'You Lost')  }</div> }
      </div>
    );
  }
}

class App extends React.Component {
  render() {
    return (
      <div className="game">
        <div className="game-board">
          <Board />
        </div>
        <div className="game-info">
          <div>{/* status */}</div>
          <ol>{/* TODO */}</ol>
        </div>
      </div>
    );
  }
}


export default App;
