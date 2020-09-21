import * as React from 'react';

type Props = {
  name: string;
};

export const Hello: React.FC<Props> = ({ name }) => {
  return <div>HELLO,{name}</div>;
};
