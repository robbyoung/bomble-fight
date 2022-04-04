import {Action} from 'redux';
import {GetPlayerStateResponse} from '../../api';
import {ActionType} from './actionType';

export function startPlayerStatePollAction(): Action {
  return {
    type: ActionType.START_STATE_POLL,
  };
}

export function stopPlayerStatePollAction(): Action {
  return {
    type: ActionType.STOP_STATE_POLL,
  };
}

export interface GetPlayerStateSuccessAction extends Action {
  type: ActionType.GET_STATE_SUCCESS;
  response: GetPlayerStateResponse;
}

export function getPlayerStateSuccessAction(
  response: GetPlayerStateResponse,
): GetPlayerStateSuccessAction {
  return {
    type: ActionType.GET_STATE_SUCCESS,
    response,
  };
}

export interface GetPlayerStateFailureAction extends Action {
  type: ActionType.GET_STATE_FAILURE;
}

export function getPlayerStateFailureAction(): GetPlayerStateFailureAction {
  return {
    type: ActionType.GET_STATE_FAILURE,
  };
}
