import React, { useState, useCallback } from 'react';
import Cell from './Cell';

const getInitialBoard = () => Array(9).fill(Array(9).fill(0))

const Board = () => {
    const [board, setBoard] = useState(getInitialBoard);
    const onChangeFactory = useCallback((rowIndex: number, colIndex: number) =>
        (e: React.ChangeEvent<HTMLInputElement>) => {
            e.preventDefault();
            const { target: { value } } = e;
            let nValue = Number(value) % 10
            nValue = nValue === 0 ? Number(value) / 10 : nValue;
            setBoard(board.map((row: number[], index: number) => {
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
        }, [board]);
    return <div className="sudokuContainer">
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
}

export default Board;