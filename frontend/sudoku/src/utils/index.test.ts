import { convertBoardToString, convertStringToBoard, validateCheck, validateSolve } from './';

describe('Convert Board[Puzzle] to String[Puzzle]', () => {

  it('should convert empty board to string with all dots', () => {
    const board = [
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0]
    ];
    const answer = '.................................................................................';
    expect(convertBoardToString(board)).toBe(answer);
  });

  it('should convert empty board to string with all dots', () => {
    const board = [
      [1, 2, 3, 4, 5, 6, 7, 8, 9],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [9, 8, 7, 6, 5, 4, 3, 2, 1]
    ];
    const answer = '123456789...............................................................987654321';
    expect(convertBoardToString(board)).toBe(answer);
  });

});

describe('Convert Board[Puzzle] to String[Puzzle]', () => {

  it('should convert empty board to string with all dots', () => {
    const board = [
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0]
    ];
    const answer = '.................................................................................';
    expect(convertStringToBoard(answer)).toStrictEqual(board);
  });

  it('should convert board to string', () => {
    const board = [
      [1, 2, 3, 4, 5, 6, 7, 8, 9],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [9, 8, 7, 6, 5, 4, 3, 2, 1]
    ];
    const answer = '123456789...............................................................987654321';
    expect(convertStringToBoard(answer)).toStrictEqual(board);
  });

});

describe('Validate Board[Puzzle] Check', () => {

  it('should return false for invalid puzzle', () => {
    const board = [
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0]
    ];
    expect(validateCheck(board)).toStrictEqual(false);
  });

  it('should return true for valid puzzle', () => {
    const board = [
      [8, 2, 5, 4, 7, 1, 3, 9, 6],
      [1, 9, 4, 3, 2, 6, 5, 7, 8],
      [3, 7, 6, 9, 8, 5, 2, 4, 1],
      [5, 1, 9, 7, 4, 3, 8, 6, 2],
      [6, 3, 2, 5, 9, 8, 4, 1, 7],
      [4, 8, 7, 6, 1, 2, 9, 3, 5],
      [2, 6, 3, 1, 5, 9, 7, 8, 4],
      [9, 4, 8, 2, 6, 7, 1, 5, 3],
      [7, 5, 1, 8, 3, 4, 6, 2, 9]
    ];
    expect(validateCheck(board)).toStrictEqual(true);
  });

});

describe('Validate Board[Puzzle] Solve', () => {

  it('should return false for invalid puzzle', () => {
    const board = [
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0]
    ];
    expect(validateSolve(board)).toStrictEqual(false);
  });

  it('should return true for valid puzzle', () => {
    const board = [
      [8, 2, 5, 4, 7, 1, 3, 9, 6],
      [1, 9, 4, 3, 2, 6, 5, 7, 8],
      [3, 7, 6, 9, 8, 5, 2, 4, 1],
      [5, 1, 9, 7, 4, 3, 8, 6, 2],
      [6, 3, 2, 5, 9, 8, 4, 1, 7],
      [4, 8, 7, 6, 1, 2, 9, 3, 5],
      [2, 6, 3, 1, 5, 9, 7, 8, 4],
      [9, 4, 8, 2, 6, 7, 1, 5, 3],
      [7, 5, 1, 8, 3, 4, 6, 2, 9]
    ];
    expect(validateSolve(board)).toStrictEqual(true);
  });

});