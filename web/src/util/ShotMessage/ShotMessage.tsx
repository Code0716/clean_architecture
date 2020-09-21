import * as React from 'react';
import { useState, useEffect, useMemo } from 'react';
import { Portal } from '../../components/Portal';
import { shotMessageAction } from '../../actions/shotMessageActions';
import './style.scss';

type Props = {
  message: string;
};

const _sleep = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

export const ShotMessage: React.FC = () => {
  const [portalStyle, setPortalStyle] = useState('shot_message_box');
  const {
    // stre
    show,
    message,
    // action
    initShotMess,
    disappear,
  } = shotMessageAction();

  useMemo(async () => {
    if (message !== undefined) {
      await _sleep(500);
      setPortalStyle('shot_message_box show_mess');
      await _sleep(3200);
      setPortalStyle('shot_message_box hide_mess');
      await _sleep(1000);
      disappear();
      initShotMess();
    }
  }, [message]);

  return (
    <React.Fragment>
      {show && (
        <Portal domId="shot_message" className="shot_message_portal">
          <div className={portalStyle}>
            <div className="massage_box">
              {'Message : '}
              <span>{message}</span>
            </div>
          </div>
        </Portal>
      )}
    </React.Fragment>
  );
};
