import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import axios from "axios"




class Board extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            board: Array(9).fill(null),
            winner: ""

        };
    }

    handleClick(i) {


        if (this.state.board[i] === "X" || this.state.winner !== "") {
            return
        }
        const game = this.state;
        game.board[i] = 'X';

        axios.post(process.env.REACT_APP_API_HOST + ":" + process.env.REACT_APP_API_PORT + process.env.REACT_APP_API_ENDPOINT, JSON.stringify(game)).then((r) => {

            this.setState(r.data);

        })

    }

    renderSquare(i) {
        return (
            <Square
                value={this.state.board[i]}
                onClick={() => this.handleClick(i)}
            />
        );
    }

    render() {
        const winner = this.state.winner

        return (
            <div>
                <div className="board-row">
                    {this.renderSquare(0)}
                    {this.renderSquare(1)}
                    {this.renderSquare(2)}
                </div>
                <div className="board-row">
                    {this.renderSquare(3)}
                    {this.renderSquare(4)}
                    {this.renderSquare(5)}
                </div>
                <div className="board-row">
                    {this.renderSquare(6)}
                    {this.renderSquare(7)}
                    {this.renderSquare(8)}
                </div>
                {winner === "DRAW" && <h1 className="winner">Game can't be won by either player</h1>}
                {winner !== "" && winner !== "DRAW" && <h1 className="winner">{winner} is the winner</h1>}
            </div >
        );
    }
}

function Square(props) {
    return (
        <button className="square" onClick={props.onClick}>
            {props.value}
        </button>
    );
}



class Game extends React.Component {
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

// ========================================

ReactDOM.render(
    <Game />,
    document.getElementById('root')
);
