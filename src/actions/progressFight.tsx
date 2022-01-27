import {Action} from 'redux';
import {PostFightResponse} from '../sagas/api';
import {Fight} from '../state';
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
    attackerId: action.response.AttackerId,
    defenderId: action.response.DefenderId,
    attackerDamage: action.response.AttackerDamage,
    defenderDamage: action.response.DefenderDamage,
    attackerHealth: action.response.AttackerHealth,
    defenderHealth: action.response.DefenderHealth,
  };
}
