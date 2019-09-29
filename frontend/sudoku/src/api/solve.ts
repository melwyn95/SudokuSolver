import { BASE_API_URL, Response } from "../common";
import { convertStringToBoard } from "../utils";

const solveApi = (puzzle: string) => `${BASE_API_URL}/solve?puzzle=${puzzle}`;

export const solve = (puzzle: string,
  setError: React.Dispatch<React.SetStateAction<string>>,
  setBoard: React.Dispatch<React.SetStateAction<any[]>>) =>
  fetch(solveApi(puzzle))
    .then(response => response.json())
    .then((r: Response) => {
      const { puzzle, solved } = r;
      if (solved) {
        setError('');
        setBoard(convertStringToBoard(puzzle));
      } else {
        setError('Wrong Answer');
      }
    })
    .catch(() => setError('Error: Unable to solved the puzzle'));