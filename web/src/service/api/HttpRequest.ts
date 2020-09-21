import axios from 'axios';
import config from '../config';
import { ImageInfo } from '../../reducers/imageListReducer';

class HttpRequest {
  get(url: string) {
    return axios
      .get(config.API_SERVER_URL + url)
      .then(response => {
        const payload = response.data;
        return payload;
      })
      .catch((error: Error) => {
        return { error };
      });
  }

  post(data: any) {
    const { url, json } = data.payload;
    const switchObject = JSON.parse(json);
    return axios
      .post(config.API_SERVER_URL + url, { data: switchObject })
      .then(response => {
        const payload = response.data;
        return { payload };
      })
      .catch((error: Error) => {
        return { error };
      });
  }

  put(url: string) {
    return axios
      .put(config.API_SERVER_URL + url)
      .then(response => {
        const payload = response.data;
        return payload;
      })
      .catch((error: Error) => {
        return { error };
      });
  }

  delete(url: string) {
    return axios
      .delete(config.API_SERVER_URL + url)
      .then(response => {
        const payload = response.data;
        return payload;
      })
      .catch((error: Error) => {
        return { error };
      });
  }
  // se
  // send images
  postImg(url: string, postData: FormData) {
    return axios.post(config.API_SERVER_URL + url, postData, {
      headers: {
        'content-type': 'multipart/form-data',
      },
    });
  }

  deleteImage(url: string, deleteData: ImageInfo) {
    return axios
      .delete(config.API_SERVER_URL + url, { data: deleteData })
      .then(response => {
        return response;
      })
      .catch((error: Error) => {
        return { error };
      });
  }
}

export default new HttpRequest();
