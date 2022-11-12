import {Action} from 'redux';
import {PostFightResponse} from '../../api';
import {Fight} from '../../models';
import {ActionType} from './actionType';

export interface ProgressFightAction extends Action {
  type: ActionType.PROGRESS_FIGHT_REQUEST;
}

export function progressFightAction(): ProgressFightAction {
  return {
    type: ActionType.PROGRESS_FIGHT_REQUEST,
  };
}

export interface ProgressFightSuccessAction extends Action {
  type: ActionType.PROGRESS_FIGHT_SUCCESS;
  response: PostFightResponse;
}

export function progressFightSuccessAction(
  response: PostFightResponse,
): ProgressFightSuccessAction {
  return {
    type: ActionType.PROGRESS_FIGHT_SUCCESS,
    response,
  };
}

export interface ProgressFightFailureAction extends Action {
  type: ActionType.PROGRESS_FIGHT_FAILURE;
}

export function progressFightFailureAction(): ProgressFightFailureAction {
  return {
    type: ActionType.PROGRESS_FIGHT_FAILURE,
  };
}

export function progressFight(
  _state: Fight,
  action: ProgressFightSuccessAction,
): Fight {
  return {
    status: action.response.FightStatus,
    left: {
      id: action.response.Left.Id,
      health: action.response.Left.Health,
      loss: action.response.Left.Loss,
      action: action.response.Left.Action,
    },
    right: {
      id: action.response.Right.Id,
      health: action.response.Right.Health,
      loss: action.response.Right.Loss,
      action: action.response.Right.Action,
    },
  };
}
