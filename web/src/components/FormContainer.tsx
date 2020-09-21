import * as React from 'react';
type Props = {
  children: React.ReactNode;
};
export const FormContainer: React.FC<Props> = (props: Props) => {
  return (
    <div className="contents fade-in">
      <div className="form_container">{props.children}</div>
    </div>
  );
};
