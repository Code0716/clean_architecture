import * as React from 'react';
import { FC } from 'react';
import * as classnames from 'classnames';
import './style.scss';

type LoadingElmProps = {
  className?: string;
};

export const LoadingElm: FC<LoadingElmProps> = ({ className }) => {
  return (
    <div className={classnames(className, 'outLoading')}>
      <p className="sentence">NOW LODING </p>
      <span className="loading" />
      <div className="loader" />
    </div>
  );
};
