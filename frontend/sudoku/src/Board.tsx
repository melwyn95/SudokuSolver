import React, { useState, useCallback } from 'react';
import Cell from './Cell';

import { solve } from './api/solve';
import { check } from './api/check';
import { convertBoardToString, validateCheck, validateSolve } from './utils';

const getInitialBoard = () => Array(9).fill(Array(9).fill(0))

//() => convertStringToBoard('53..7....6..195....98....6.8...6...34..8.3..17...2...6.6....28....419..5....8..79')

const Board = () => {
  const [board, setBoard] = useState(getInitialBoard);
  const [error, setError] = useState('');
  const onChangeFactory = useCallback((rowIndex: number, colIndex: number) => {
    return (e: React.ChangeEvent<HTMLInputElement>) => {
      e.preventDefault();
      setError('');
      const { target: { value } } = e;
      let nValue = Number(value) % 10
      nValue = nValue === 0 ? Number(value) / 10 : nValue;
      setBoard((_board: number[][]) =>
        _board.map((row: number[], index: number) => {
          if (index !== rowIndex) {
            return row;
          }
          return row.map((cell: number, cellIndex: number) => {
            if (cellIndex !== colIndex) {
              return cell
            }
            return nValue;
          })
        }));
    }
  }, []);

  const solveClicked = () => validateSolve(board) ?
    solve(convertBoardToString(board), setError, setBoard) :
    setError('Invalid Sudoku');

  const checkClicked = () => validateCheck(board) ?
    check(convertBoardToString(board), setError) :
    setError('Invalid Sudoku');

  return <>
    <div className="sudokuContainer">
      {board.map((row: number[], rowIndex: number) => {
        return row.map((cell: number, colIndex: number) => {
          return (
            <Cell
              key={`${rowIndex}_${colIndex}`}
              value={cell}
              onChange={onChangeFactory(rowIndex, colIndex)} />)
        })
      })}
    </div>
    <div className="actionPanel">
      <button className="actionButton" onClick={solveClicked}>Solve</button>
      <button className="actionButton" onClick={checkClicked}>Check</button>
    </div>
    <div className="errorMessage">
      {error}
    </div>
  </>
}

export default Board;