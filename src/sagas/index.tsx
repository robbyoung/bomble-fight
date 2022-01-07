import {put, takeLatest} from 'redux-saga/effects';
import {ActionType} from '../actions/actionType';
import {
  getCombatantsFailureAction,
  getCombatantsSuccessAction,
} from '../actions/getCombatants';
import {GetCombatantsResponse} from './api';

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

function* saga() {
  yield takeLatest(ActionType.GET_COMBATANTS_REQUEST, getCombatants);
}

export default saga;
