import axios from 'axios';
import config from './config';

class HttpRequest {
  get(url: string) {
    return axios
      .get(config.OTHER_SERVER_URL + url)
      .then(response => {
        const payload = response.data;
        return payload;
      })
      .catch((error: Error) => {
        return { error };
      });
  }
}

export const OtherHttpRequest = new HttpRequest();
