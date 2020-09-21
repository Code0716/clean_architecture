export class LinksPathnames {}

export namespace LinksPathnames {
  export const HOME: { name: string; path: string } = {
    name: 'HOME',
    path: '/',
  };

  export const IMAGE: { name: string; path: string } = {
    name: 'IMAGE LIST',
    path: '/images',
  };

  export const IMAGE_LIST_UPLOAD: { path: string } = {
    path: '/images/:directory',
  };
  export const IMAGE_DETAILE: { path: string } = {
    path: '/images/list/:id',
  };

  export const TEST: { name: string; path: string } = {
    name: 'テスト用画面',
    path: '/test',
  };

  export const NOT_FOUND: { name: string; path: string } = {
    name: 'NOT FOUND',
    path: 'not_found',
  };
}
