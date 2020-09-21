import { testGetPostReducer as testGetPost } from './testGetPostReducer';
import { imageListReducer as imageList } from './imageListReducer';
import { reducer as loading } from './loadingReducer';
import { reducer as shotMess } from './shotMessageReducer';

export const reducers = { testGetPost, imageList, loading, shotMess };
