import createSagaMiddleware from 'redux-saga';
import App from './App';
import {createStore, applyMiddleware} from 'redux';
import {Provider} from 'react-redux';
import rootReducer from './reducers';
import saga from './sagas';

const sagaMiddleware = createSagaMiddleware();
const store = createStore(rootReducer, applyMiddleware(sagaMiddleware));
sagaMiddleware.run(saga);

function Player() {
  return (
    <Provider store={store}>
      <App />
    </Provider>
  );
}

export default Player;
