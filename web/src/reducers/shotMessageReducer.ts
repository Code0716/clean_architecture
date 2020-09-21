import { ActionTypes } from '../actions/shotMessageActions';

// 型の定義
export type shotMessProps = {
  show: boolean;
  message: string | undefined;
};

const initialState: shotMessProps = {
  show: false,
  message: undefined,
};

export type Action = {
  type: string;
  message: string;
};

// reducer
export function reducer(state = initialState, { type, message }: Action) {
  switch (type) {
    case ActionTypes.INIT:
      return initialState;
    case ActionTypes.SHOW:
      return { ...state, show: true, message: message };
    case ActionTypes.DISAPPEAR:
      return { ...state, show: false };
    default:
      return state;
  }
}
