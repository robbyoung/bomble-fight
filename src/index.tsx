import {createStore, applyMiddleware} from 'redux';
import ReactDOM from 'react-dom';
import {Provider} from 'react-redux';
import App from './App';
import rootReducer from './reducers';
import createSagaMiddleware from 'redux-saga';
import saga from './sagas';

const sagaMiddleware = createSagaMiddleware();
const store = createStore(rootReducer, applyMiddleware(sagaMiddleware));
sagaMiddleware.run(saga);

const rootElement = document.getElementById('root');
ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  rootElement,
);
