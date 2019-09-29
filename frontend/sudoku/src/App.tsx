import React from 'react';
import './App.css';
import Board from './Board';

const App: React.FC = () => {
  return (
    <>
      <header className="sudokuSolverHeader">
        <h1>Sudoku Solver</h1>
      </header>
      <div className="App">
        <Board />
      </div>
    </>
  );
}

export default App;
