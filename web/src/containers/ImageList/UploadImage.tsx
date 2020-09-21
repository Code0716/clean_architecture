import * as React from 'react';
import { useEffect, useMemo, useCallback } from 'react';
import { useDispatch } from 'react-redux';
import { Portal } from '../../components/Portal';
import { FormContainer } from '../../components/FormContainer';
import { imageListActions, ActionTypes } from '../../actions/imageListActions';
import { LoadImage } from '../../reducers/imageListReducer';
import { LoadingElm } from '../../components/Loading/LoadingElm';
import { FormatUtil } from '../../util/FormatUtil';
import './style.scss';

export const UploadImage: React.FC = () => {
  const dispatch = useDispatch();

  const {
    // state
    preUploadImages,
    // action
    addPreUploadImages,
    deletePreImage,
    getPreImagesInfo,
    commitUpload,
    deletePreupload,
  } = imageListActions();

  useEffect(() => {
    getPreImagesInfo();
    return () => {
      dispatch({
        type: ActionTypes.CHANGE_STATE,
        payload: { preUploadImagesRender: [] },
      });
    };
  }, []);

  const handleChangeFile = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const target: HTMLInputElement = e.target as HTMLInputElement;
      let files: Array<File | null> = [];

      if (target.files !== null) {
        for (let index = 0; index < target.files.length; index++) {
          files.push(target.files.item(index));
        }
      }
      addPreUploadImages(files);
    },
    [preUploadImages],
  );

  const preUploadImagesRender = useMemo(() => {
    const imgs = preUploadImages.map((elemnt: LoadImage, index: number) => (
      <div className="mb-10 flex-box" key={`pre-up-img${index}`}>
        <div className="preupload-img">
          <button
            className="button"
            onClick={() => {
              deletePreImage(elemnt.info.ID);
            }}
          >
            削除
          </button>
          {elemnt.info.Path ? (
            <img className="fade-in" src={elemnt.info.Path} />
          ) : (
            <LoadingElm className={'image_size_loading'} />
          )}
        </div>
        <div className="preupload-description">
          <span className="d-block">Title : {elemnt.info.Title}</span>
          <span className="d-block">ID : {elemnt.info.ID}</span>
          <span className="d-block">
            Date : {FormatUtil.formatedDate(elemnt.info.ShotDate)}
          </span>
          <span className="d-block">Path : {elemnt.info.Path}</span>
        </div>
      </div>
    ));
    return <div className="preupload-imgbox mb-10">{imgs}</div>;
  }, [preUploadImages]);

  return (
    <React.Fragment>
      <label className="image_up_label mb-5">
        ＋写真を選択
        <input
          className="d-none"
          type="file"
          accept=".png,.jpg,.gif"
          multiple
          onChange={handleChangeFile}
        />
      </label>
      {preUploadImages.length !== 0 && (
        <Portal domId={'preupload'} className="modal">
          <FormContainer>
            {preUploadImagesRender}
            <div className="image_commit_box">
              <div>Upload confirm</div>
              <label className="image_up_label mb-5">
                ＋写真を選択
                <input
                  className="d-none"
                  type="file"
                  accept=".png,.jpg,.gif"
                  multiple
                  onChange={handleChangeFile}
                />
              </label>
              <button className="action_button" onClick={commitUpload}>
                <span>Commit Upload</span>
              </button>
              <button
                className="action_button"
                onClick={() => deletePreupload('preupload')}
              >
                <span>Delete all images</span>
              </button>
            </div>
          </FormContainer>
        </Portal>
      )}
    </React.Fragment>
  );
};
