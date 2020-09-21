import * as React from 'react';

type Props = {
  value: string;
  option: { value: string }[];
  onChange: (event: React.ChangeEvent<HTMLSelectElement>) => void;
};

export const Select = ({ value, option, onChange }: Props) => {
  const options = option.map(opt => (
    <option key={opt.value} value={opt.value}>
      {opt.value}
    </option>
  ));

  return (
    <select value={value} onChange={onChange}>
      {options}
    </select>
  );
};
