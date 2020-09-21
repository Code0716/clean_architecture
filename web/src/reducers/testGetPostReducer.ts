import { ActionTypes } from '../actions/testGetPostActions';

export type State = {
  getData: { code: string; name: string }[];
  postData: { id: number; code: string; name: string }[];
};

export const initialState: State = {
  getData: [],
  postData: [],
};

export type Action = {
  type: String;
  payload?: any;
};

export function testGetPostReducer(
  state = initialState,
  { type, payload }: Action,
) {
  switch (type) {
    case ActionTypes.CHANGE_STATE:
      return {
        ...state,
        ...payload,
      };

    default:
      return state;
  }
}
