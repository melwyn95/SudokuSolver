# Solve API
**Endpoint**: https://melwyn-sudoku-solver.herokuapp.com/solve  
**Method**: GET  
**Query Params**: puzzle: string  
**Response Type**: { puzzle: string, solved: boolean, status: 'success' | 'failure' }  

**Example**: https://melwyn-sudoku-solver.herokuapp.com/solve?puzzle=53..7....6..195....98....6.8...6...34..8.3..17...2...6.6....28....419..5....8..79  
**Response**: {"puzzle":"534678912672195348198342567859761423426853791713924856961537284287419635345286179","solved":true,"status":"success"}  
**The puzzle query param is a string representaion of the Sudoku Grid row-by-row**  

# Check API
**Endpoint**: https://melwyn-sudoku-solver.herokuapp.com/check  
**Method**: GET  
**Query Params**: puzzle: string  
**Response Type**: { puzzle: string, solved: boolean, status: 'success' | 'failure' }  

**Example**: https://melwyn-sudoku-solver.herokuapp.com/check?puzzle=534678912672195348198342567859761423426853791713924856961537284287419635345286179  
**Response**: {"puzzle":"534678912672195348198342567859761423426853791713924856961537284287419635345286179","solved":true,"status":"success"}  
**The puzzle query param is a string representaion of the Sudoku Grid row-by-row**  
