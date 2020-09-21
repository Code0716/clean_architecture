import { useCallback } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../store/store';
//import { push } from 'connected-react-router';
import HttpRequest from '../service/api/HttpRequest';
import { OtherHttpRequest } from '../service/otherRequest';
import { ActionTypes as shotMessageAT } from './shotMessageActions';

// Actions
export const ActionTypes = {
  GET: 'HOME_GET',
  POST: 'HOME_POST',
  CHANGE_STATE: 'HOME_CHANGE_STATE',
} as const;

// Action Creators

export const testGetPostActions = () => {
  const dispatch = useDispatch();
  //store
  const testGetPost = useSelector(({ testGetPost }: RootState) => testGetPost);

  const { getData, postData, entity } = testGetPost;

  const get = useCallback(
    async (url: string) => {
      dispatch({ type: ActionTypes.GET, url });
      const data = await HttpRequest.get(url);
      // Asiaのみ取得
      dispatch({ type: ActionTypes.CHANGE_STATE, payload: { getData: data } });
    },
    [dispatch],
  );

  const post = useCallback(
    // TODO
    (url, json) => dispatch({ type: ActionTypes.POST, payload: { url, json } }),
    [dispatch],
  );

  const migrate = useCallback(async () => {
    try {
      const response = await OtherHttpRequest.get('/migrate');
      dispatch({ type: shotMessageAT.SHOW, message: response.message });
    } finally {
    }
  }, [dispatch]);

  return {
    //state
    getData,
    postData,
    entity,
    //action
    get,
    post,
    migrate,
  };
};
