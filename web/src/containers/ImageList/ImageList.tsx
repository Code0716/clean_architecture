import * as React from 'react';
import { useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { FormContainer } from '../../components/FormContainer';
import { imageListActions, RouteParams } from '../../actions/imageListActions';
import { UploadImage } from './UploadImage';
import { ListImages } from './ListImages';
import { ImageDetaile } from '../ImageDetaile/ImageDetaile';

import './style.scss';

export const ImageList: React.FC = () => {
  const params: RouteParams = useParams();
  const {
    //state
    //action
    initialize,
    getImages,
    forwordToImages,
  } = imageListActions();

  useEffect(() => {
    getImages();
    return initialize;
  }, [params.directory]);

  return (
    <FormContainer>
      <div className="flex-box mb-10 p-10">
        <button className="button" onClick={() => forwordToImages('upload')}>
          Upload images
        </button>
        <button className="button" onClick={() => forwordToImages('list')}>
          Image list
        </button>
      </div>
      {params.directory === 'upload' && <UploadImage />}
      {params.directory === 'list' && <ListImages />}
      {params.id && <ImageDetaile />}
    </FormContainer>
  );
};
