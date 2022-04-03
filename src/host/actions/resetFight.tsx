import {Action} from 'redux';
import {ActionType} from './actionType';

export interface ResetFightAction extends Action {
  type: ActionType.RESET_FIGHT_REQUEST;
}

export function resetFightAction(): ResetFightAction {
  return {
    type: ActionType.RESET_FIGHT_REQUEST,
  };
}

export interface ResetFightSuccessAction extends Action {
  type: ActionType.RESET_FIGHT_SUCCESS;
}

export function resetFightSuccessAction(): ResetFightSuccessAction {
  return {
    type: ActionType.RESET_FIGHT_SUCCESS,
  };
}

export interface ResetFightFailureAction extends Action {
  type: ActionType.RESET_FIGHT_FAILURE;
}

export function resetFightFailureAction(): ResetFightFailureAction {
  return {
    type: ActionType.RESET_FIGHT_FAILURE,
  };
}
