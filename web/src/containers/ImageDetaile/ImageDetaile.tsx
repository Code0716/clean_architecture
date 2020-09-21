import * as React from 'react';
import { useEffect, useMemo } from 'react';
import { useParams } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import {
  ActionTypes,
  imageListActions,
  RouteParams,
} from '../../actions/imageListActions';
import { LoadImage } from '../../reducers/imageListReducer';
import { LoadingElm } from '../../components/Loading/LoadingElm';
import './style.scss';

export const ImageDetaile: React.FC = () => {
  const dispatch = useDispatch();
  const params: RouteParams = useParams();
  const paramsID = Number(params.id);
  const {
    // store
    getImage,
    images,
    currentBase64,
    // actions
    forwordToDetaile,
  } = imageListActions();

  useEffect(() => {
    getImage();
    return () => {
      dispatch({
        type: ActionTypes.CHANGE_STATE,
        payload: { currentBase64: undefined },
      });
    };
  }, [params.id]);

  // check exist prev or next
  const searchPrevNext = (num: number) =>
    images.some((item: LoadImage, index: number) => {
      // スマートじゃないよねー。
      if (paramsID === item.info.ID) {
        if (images[index + num]) return true;
      }
    });

  const isPrevExist = useMemo<boolean>(() => {
    const result = searchPrevNext(-1);
    return result;
  }, [images, paramsID]);

  const isNextExist = useMemo<boolean>(() => {
    const result = searchPrevNext(1);
    return result;
  }, [images, paramsID]);

  return (
    <div className="flex-box-center">
      {isPrevExist && (
        <button
          className="img_left_arrow"
          onClick={() => forwordToDetaile(paramsID - 1)}
        >
          {'←'}
        </button>
      )}
      <div className="image_box">
        {currentBase64 ? (
          <div className="inner_box">
            <img src={`data:image/png;base64,${currentBase64}`} />
          </div>
        ) : (
          <LoadingElm className="loading_height" />
        )}
      </div>
      {isNextExist && (
        <button
          className="img_right_arrow"
          onClick={() => forwordToDetaile(paramsID + 1)}
        >
          {'→'}
        </button>
      )}
    </div>
  );
};
