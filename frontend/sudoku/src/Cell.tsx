import React, { memo } from 'react'

type CellProps = {
  value: number,
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void
}

const Cell = (props: CellProps) => {
  const { value, onChange } = props;

  return (<input
    className="sudokuCell"
    type="number"
    value={Number(value).toString()}
    onChange={onChange}
    maxLength={1}
  />)
}

const isSame = (prevProps: CellProps, nextProps: CellProps): boolean => {
  return prevProps.value === nextProps.value;
}

export default memo(Cell, isSame);