import {put, takeEvery, takeLatest} from 'redux-saga/effects';
import {ActionType} from '../actions/actionType';
import {
  AddBetAction,
  addBetFailureAction,
  addBetSuccessAction,
} from '../actions/addBet';
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
  GetCombatantsResponse,
  GetPlayerStateResponse,
  PostBetResponse,
  PostPlayerResponse,
} from '../../api';
import {
  GetPlayerStateAction,
  getPlayerStateSuccessAction,
} from '../actions/getPlayerState';
import {getPlayersFailureAction} from '../../host/actions/getPlayers';

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

function* postBet(action: AddBetAction) {
  try {
    const json: PostBetResponse = yield fetch(`${baseUrl}/bet`, {
      method: 'POST',
      body: JSON.stringify(action.bet),
    }).then((response) => response.json());
    yield put(addBetSuccessAction(json));
  } catch {
    yield put(addBetFailureAction());
  }
}

function* getPlayerState(action: GetPlayerStateAction) {
  try {
    const json: GetPlayerStateResponse = yield fetch(
      `${baseUrl}/state/${action.playerId}`,
    ).then((response) => response.json());
    yield put(getPlayerStateSuccessAction(json));
  } catch {
    yield put(getPlayersFailureAction());
  }
}

function* saga() {
  yield takeLatest(ActionType.GET_COMBATANTS_REQUEST, getCombatants);
  yield takeEvery(ActionType.ADD_PLAYER_REQUEST, postPlayer);
  yield takeEvery(ActionType.ADD_BET_REQUEST, postBet);
  yield takeLatest(ActionType.GET_STATE_REQUEST, getPlayerState);
}

export default saga;
