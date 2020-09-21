import { ActionTypes } from '../actions/imageListActions';

// 型の定義 imageと
export type ImageInfo = {
  ID: number;
  Title: string;
  ShotDate: Date;
  Path: string;
};

export interface LoadImage {
  info: ImageInfo;
}

export interface State {
  images: LoadImage[];
  preUploadImages: LoadImage[];
  currentBase64: string | undefined;
}

// 初期値
const initialState: State = {
  images: [],
  preUploadImages: [],
  currentBase64: undefined,
};

export interface Action {
  type: String;
  payload?: any;
}

export function imageListReducer(
  state = initialState,
  { type, payload }: Action,
) {
  switch (type) {
    case ActionTypes.CHANGE_STATE:
      return {
        ...state,
        ...payload,
      };
    case ActionTypes.INITIALIZE:
      return initialState;
    case ActionTypes.UPDATE_PRE_UPLOAD:
      return {
        ...state,
        preUploadImages: [...state.preUploadImages, ...payload],
      };
    case ActionTypes.UPDATE_UPLOAD_IMAGE:
      // 重くなるので都度取得するようにした。

      return {
        ...state,
        currentBase64: payload.img,
      };
    case ActionTypes.DELETE_PRE_UPLOAD:
      return {
        ...state,
        preUploadImages: [...payload],
      };

    default:
      return state;
  }
}
