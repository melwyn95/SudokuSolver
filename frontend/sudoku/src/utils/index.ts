// covert array -> puzzle string
export const convertBoardToString = (board: number[][]) =>
  board.reduce((rowAcc: string, row: number[]) =>
    row.reduce((colAcc: string, cell: number) =>
      colAcc + (cell ? cell : '.'), rowAcc), '');


// covert puzzle string -> array 
export const convertStringToBoard = (puzzle: string) => {
  const emptyBoard = Array(9).fill(Array(9).fill(0));
  return emptyBoard.map((row: number[][], rowIndex: number) => {
    return row.map((cell: number[], colIndex: number) => {
      const cellChar = puzzle[9 * rowIndex + colIndex];
      return cellChar === '.' ? 0 : Number(cellChar);
    })
  })
}

const getRow = (board: number[][], rowIndex: number, colIndex: number) =>
  board[rowIndex].filter((_, cellIndex: number) =>
    cellIndex !== colIndex);

const getCol = (board: number[][], rowIndex: number, colIndex: number) =>
  Array.from(Array(9).keys()).reduce((acc: number[], index: number) =>
    [...acc, ...(index !== rowIndex ? [board[index][colIndex]] : [])], [])

const getCenterCoOrds = (rowIndex: number, colIndex: number) =>
  [3 * +Math.floor(rowIndex / 3) + 1, 3 * +Math.floor(colIndex / 3) + 1]

const getBox = (board: number[][], rowIndex: number, colIndex: number) => {
  const [cX, cY] = getCenterCoOrds(rowIndex, colIndex);
  const box = [];
  for (let rowOffset = -1; rowOffset < 2; rowOffset++) {
    for (let colOffset = -1; colOffset < 2; colOffset++) {
      const r = cX + rowOffset, c = cY + colOffset;
      if (r !== rowIndex && c !== colIndex) {
        box.push(board[r][c]);
      }
    }
  }
  return box;
}

const check = (board: number[][], rowIndex: number, colIndex: number) => {
  const row = getRow(board, rowIndex, colIndex);
  const col = getCol(board, rowIndex, colIndex);
  const box = getBox(board, rowIndex, colIndex);
  const value = board[rowIndex][colIndex];
  return !(row.includes(value) || col.includes(value) || box.includes(value));
}

// validate puzzle while checking
export const validateCheck = (board: number[][]) =>
  board.reduce((rowAcc: boolean, row: number[], rowIndex: number) =>
    row.reduce((colAcc: boolean, _: number, colIndex: number) =>
      colAcc && check(board, rowIndex, colIndex), rowAcc), true);

// validate puzzle to have atlease 17 values
export const validateSolve = (board: number[][]) =>
  board.reduce((rowAcc: number[], row: number[]) =>
    row.reduce((colAcc: number[], cell: number) =>
      cell ? [...colAcc, cell] : colAcc, rowAcc), []).length >= 17;