import {Action} from 'redux';
import {GetPlayerStateResponse} from '../../api';
import {ActionType} from './actionType';

export interface GetPlayerStateAction extends Action {
  type: ActionType.GET_STATE_REQUEST;
  playerId: string;
}

export function getPlayerState(playerId: string): GetPlayerStateAction {
  return {
    type: ActionType.GET_STATE_REQUEST,
    playerId,
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
