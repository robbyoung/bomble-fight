import {put, takeEvery, takeLatest} from 'redux-saga/effects';
import {ActionType} from '../actions/actionType';
import {
  AddPlayerAction,
  addPlayerFailureAction,
  addPlayerSuccessAction,
} from '../actions/addPlayer';
import {
  getCombatantsFailureAction,
  getCombatantsSuccessAction,
} from '../actions/getCombatants';
import {
  getPlayersFailureAction,
  getPlayersSuccessAction,
} from '../actions/getPlayers';
import {
  GetCombatantsResponse,
  GetPlayersResponse,
  PostPlayerResponse,
} from './api';

const baseUrl = 'http://localhost:3001';

function* getCombatants() {
  try {
    const json: GetCombatantsResponse = yield fetch(
      `${baseUrl}/combatants`,
    ).then((response) => response.json());
    yield put(getCombatantsSuccessAction(json));
  } catch {
    yield put(getCombatantsFailureAction());
  }
}

function* getPlayers() {
  try {
    const json: GetPlayersResponse = yield fetch(`${baseUrl}/players`).then(
      (response) => response.json(),
    );
    yield put(getPlayersSuccessAction(json));
  } catch {
    yield put(getPlayersFailureAction());
  }
}

function* postPlayer(action: AddPlayerAction) {
  try {
    const json: PostPlayerResponse = yield fetch(`${baseUrl}/player`, {
      method: 'POST',
      body: JSON.stringify(action.player),
    }).then((response) => response.json());
    yield put(addPlayerSuccessAction(json));
  } catch {
    yield put(addPlayerFailureAction());
  }
}

function* saga() {
  yield takeLatest(ActionType.GET_COMBATANTS_REQUEST, getCombatants);
  yield takeLatest(ActionType.GET_PLAYERS_REQUEST, getPlayers);
  yield takeEvery(ActionType.ADD_PLAYER_REQUEST, postPlayer);
}

export default saga;
