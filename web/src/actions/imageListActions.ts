import { useCallback } from 'react';
import { Action } from 'redux';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import { push } from 'connected-react-router';
import { cloneDeep } from 'lodash';
import { RootState } from '../store/store';

import { LoadImage, ImageInfo } from '../reducers/imageListReducer';
import { loading, unloading } from '../actions/loadingActions';
import { ActionTypes as shotMessageAT } from './shotMessageActions';

import HttpRequest from '../service/api/HttpRequest';
// Actions
export const ActionTypes = {
  INITIALIZE: 'IMAGE_LIST_INITIALIZE',
  GET_LIST: 'IMAGE_LIST_GET_LIST',
  GET_UPLOAD: 'IMAGE_LIST_GET_UPLOAD_IMAGES',
  UPDATE_PRE_UPLOAD: 'IMAGE_LIST_UPDATE_PRE_UPLOAD',
  UPDATE_UPLOAD_IMAGE: 'IMAGE_LIST_UPDATE_UPLOAD_IMAGE',
  DELETE_PRE_UPLOAD: 'IMAGE_LIST_DELETE_PRE_UPLOAD',
  CHANGE_STATE: 'IMAGE_LIST_CHANGE_STATE',
} as const;

// Action Creators

// Actionの型 Actionを継承 TODO
interface GetList extends Action {
  type: typeof ActionTypes.GET_LIST;
}

export type HttpRequestActionTypes = GetList;

export type RouteParams = {
  directory?: string;
  id?: string;
};

export const imageListActions = () => {
  const dispatch = useDispatch();

  // location params
  const params: RouteParams = useParams();

  //store
  const imageList = useSelector(({ imageList }: RootState) => imageList);

  const { images, preUploadImages, currentBase64 } = imageList;

  const initialize = useCallback(() => {
    dispatch({
      type: ActionTypes.INITIALIZE,
    });
  }, [dispatch]);

  //  画像一覧
  const getImages = useCallback(async () => {
    dispatch({
      type: ActionTypes.GET_LIST,
    });
    try {
      const data = await HttpRequest.get(`/image/upload`);
      dispatch({
        type: ActionTypes.CHANGE_STATE,
        payload: { images: data },
      });
    } finally {
      // TODO
    }
  }, [dispatch]);

  // load preupload image
  const getPreImagesInfo = useCallback(async () => {
    dispatch(loading());
    try {
      const response = await HttpRequest.get('/image/pre_upload');
      dispatch({
        type: ActionTypes.UPDATE_PRE_UPLOAD,
        payload: response,
      });
    } finally {
      dispatch(unloading());
    }
  }, [dispatch, preUploadImages]);

  // 画像を一枚づつ取得する。
  const getImage = useCallback(async () => {
    try {
      const response: { img: string } = await HttpRequest.get(
        `/image/upload/${params.id}`,
      );
      dispatch({
        type: ActionTypes.UPDATE_UPLOAD_IMAGE,
        payload: { ...response },
      });
    } finally {
    }
  }, [dispatch, images, params.id]);

  // add preupload image
  const addPreUploadImages = useCallback(
    async (files: Array<File | null>) => {
      dispatch(loading());
      const postData = new FormData();
      const dateArr: Date[] = [];

      files.forEach(element => {
        // lastModifiedが空のときの処理を入れる
        if (element !== null) {
          dateArr.push(new Date(element.lastModified));
          postData.append('images', element);
        }
      });
      // 撮影日の配列
      postData.append('date', JSON.stringify(dateArr));
      try {
        const response = await HttpRequest.postImg(
          '/image/pre_upload',
          postData,
        );
        dispatch({
          type: ActionTypes.UPDATE_PRE_UPLOAD,
          payload: response.data,
        });
      } finally {
        dispatch(unloading());
      }
    },
    [dispatch, preUploadImages],
  );

  //  PreUpload delete image
  const deletePreImage = useCallback(
    async (id: number) => {
      try {
        const _copyImages = cloneDeep(preUploadImages);
        const newState: LoadImage[] = [];
        let deleteImage!: ImageInfo;

        _copyImages.forEach((elm: LoadImage) => {
          if (elm.info.ID === id) {
            deleteImage = elm.info;
          } else {
            newState.push(elm);
          }
        });

        await HttpRequest.deleteImage(`/image/pre_upload/delete`, deleteImage);
        dispatch({
          type: ActionTypes.DELETE_PRE_UPLOAD,
          payload: newState,
        });
      } finally {
      }
    },
    [dispatch, preUploadImages],
  );

  //　delete all preupload images
  const deletePreupload = useCallback(
    async (tableName: string) => {
      dispatch(loading());
      try {
        const response = await HttpRequest.delete(`/image/delete/${tableName}`);
        dispatch({ type: shotMessageAT.SHOW, message: response.message });

        forwordToImages();
      } finally {
        dispatch(unloading());
      }
    },
    [dispatch, preUploadImages],
  );

  //　preupload to upload commit
  const commitUpload = useCallback(async () => {
    dispatch(loading());
    try {
      const response = await HttpRequest.put('/image/upload');

      dispatch({ type: shotMessageAT.SHOW, message: response.message });

      forwordToImages();
    } finally {
      dispatch(unloading());
    }
  }, [dispatch, preUploadImages]);

  // forword to /list or /upload or /
  const forwordToImages = useCallback(
    (path: string = '') => dispatch(push(`/images/${path}`)),
    [params.directory, params.id],
  );

  // forword to detaile
  const forwordToDetaile = useCallback(
    (id: number) => dispatch(push(`/images/list/${id}`)),
    [params.directory, params.id],
  );

  return {
    //state
    images,
    preUploadImages,
    currentBase64,
    //action
    initialize,
    getPreImagesInfo,
    getImages,
    getImage,
    addPreUploadImages,
    deletePreImage,
    commitUpload,
    deletePreupload,
    forwordToImages,
    forwordToDetaile,
  };
};
