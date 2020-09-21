'use strict';
import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { Route, Switch } from 'react-router';
import { ConnectedRouter } from 'connected-react-router';
import { configureStore, history } from './store/store';

// constant
import { LinksPathnames } from './constant/commonConstant';

import { Home } from '../src/containers/Home';
import { ImageList } from './containers/ImageList/ImageList';
import TestGetPost from '../src/containers/TestGetPost';
import { NotFound } from './containers/NotFound.jsx';
import { SideMenu } from './components/SideMenu/SideMenu.jsx';
import { Loading } from './components/Loading/Loading';
import { ShotMessage } from './util/ShotMessage/ShotMessage';
import './sass/style.scss';

const store = configureStore();

ReactDOM.render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
      <Switch>
        <Route exact path={LinksPathnames.HOME.path} component={Home} />
        <Route exact path={LinksPathnames.IMAGE.path} component={ImageList} />
        <Route
          exact
          path={LinksPathnames.IMAGE_LIST_UPLOAD.path}
          component={ImageList}
        />
        <Route
          exact
          path={LinksPathnames.IMAGE_DETAILE.path}
          component={ImageList}
        />
        <Route exact path={LinksPathnames.TEST.path} component={TestGetPost} />
        <Route component={NotFound} />
      </Switch>
      <SideMenu />
      <Loading />
      <ShotMessage />
    </ConnectedRouter>
  </Provider>,
  document.getElementById('root'),
);
