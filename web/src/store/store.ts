import { createStore, combineReducers, applyMiddleware, compose } from 'redux';
import { reducers } from '../reducers';

import { connectRouter, routerMiddleware } from 'connected-react-router';
import { createBrowserHistory } from 'history';

const composeEnhancers =
  (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

export const history = createBrowserHistory();

export const rootReducer = combineReducers({
  ...reducers,
  router: connectRouter(history),
});

export function configureStore() {
  const store = createStore(
    rootReducer,
    composeEnhancers(applyMiddleware(routerMiddleware(history))),
  );

  return store;
}

export type RootState = ReturnType<typeof rootReducer>;
