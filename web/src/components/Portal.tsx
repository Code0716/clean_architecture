import * as React from 'react';
import { FC, useEffect } from 'react';
import * as ReactDOM from 'react-dom';

type PortalProps = {
  domId: string;
  className?: string;
  children: React.ReactNode;
};

export const Portal: FC<PortalProps> = ({ domId, className, children }) => {
  useEffect(() => {
    return () => {
      document.body.removeChild(target);
    };
  }, []);

  const target = document.createElement('div');
  if (className) target.setAttribute('class', className);
  target.setAttribute('id', domId);

  const potalDom = document.getElementById(domId);

  if (!potalDom) {
    // BODY直下に新しいdiv要素を作成する
    document.body.appendChild(target);
    // 作成したdivの子要素としてchildrenを描写する
    return ReactDOM.createPortal(children, target);
  }

  return ReactDOM.createPortal(children, potalDom);
};
