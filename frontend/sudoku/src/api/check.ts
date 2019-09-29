import { BASE_API_URL, Response } from "../common";

const checkApi = (puzzle: string) => `${BASE_API_URL}/check?puzzle=${puzzle}`;

export const check = (puzzle: string, setError: React.Dispatch<React.SetStateAction<string>>) =>
  fetch(checkApi(puzzle))
    .then(response => response.json())
    .then((r: Response) => {
      const { status } = r;
      if (status) {
        alert('Correct Answer');
      } else {
        setError('Wrong Answer');
      }
    })
    .catch(() => setError('Error: Unable to check the puzzle'));