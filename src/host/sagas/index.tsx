import {
  call,
  delay,
  put,
  race,
  take,
  takeEvery,
  takeLatest,
} from 'redux-saga/effects';
import {ActionType} from '../actions/actionType';
import {
  getCombatantsFailureAction,
  getCombatantsSuccessAction,
} from '../actions/getCombatants';
import {
  getPlayersFailureAction,
  getPlayersSuccessAction,
} from '../actions/getPlayers';
import {
  ProgressFightAction,
  progressFightFailureAction,
  progressFightSuccessAction,
} from '../actions/progressFight';
import {
  GetCombatantsResponse,
  GetPlayersResponse,
  PostFightResponse,
} from '../../api';
import {
  resetFightFailureAction,
  resetFightSuccessAction,
} from '../actions/resetFight';

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

function* pollPlayers() {
  while (true) {
    try {
      const json: GetPlayersResponse = yield fetch(`${baseUrl}/players`).then(
        (response) => response.json(),
      );
      yield put(getPlayersSuccessAction(json));
      yield delay(2000);
    } catch {
      yield put(getPlayersFailureAction());
    }
  }
}

function* postFight(_action: ProgressFightAction) {
  try {
    const json: PostFightResponse = yield fetch(`${baseUrl}/fight`, {
      method: 'POST',
    }).then((response) => response.json());
    yield put(progressFightSuccessAction(json));
  } catch {
    yield put(progressFightFailureAction());
  }
}

function* resetFight(_action: ProgressFightAction) {
  try {
    yield fetch(`${baseUrl}/reset`, {
      method: 'POST',
    }).then((response) => response.json());
    yield put(resetFightSuccessAction());
  } catch {
    yield put(resetFightFailureAction());
  }
}

function* playersPollWatcher() {
  while (true) {
    yield take(ActionType.START_PLAYERS_POLL);
    yield race([call(pollPlayers), take(ActionType.STOP_PLAYERS_POLL)]);
  }
}

function* saga() {
  yield playersPollWatcher();
  yield takeLatest(ActionType.GET_COMBATANTS_REQUEST, getCombatants);
  yield takeEvery(ActionType.PROGRESS_FIGHT_REQUEST, postFight);
  yield takeEvery(ActionType.RESET_FIGHT_REQUEST, resetFight);
}

export default saga;
