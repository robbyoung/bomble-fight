import {Action} from 'redux';
import {GetPlayersResponse} from '../sagas/api';
import {Player} from '../state';
import {ActionType} from './actionType';

export interface GetPlayersAction extends Action {
  type: ActionType.GET_PLAYERS_REQUEST;
}

export function getPlayersAction(): GetPlayersAction {
  return {
    type: ActionType.GET_PLAYERS_REQUEST,
  };
}

export interface GetPlayersSuccessAction extends Action {
  type: ActionType.GET_PLAYERS_SUCCESS;
  response: GetPlayersResponse;
}

export function getPlayersSuccessAction(
  response: GetPlayersResponse,
): GetPlayersSuccessAction {
  return {
    type: ActionType.GET_PLAYERS_SUCCESS,
    response,
  };
}

export interface GetPlayersFailureAction extends Action {
  type: ActionType.GET_PLAYERS_FAILURE;
}

export function getPlayersFailureAction(): GetPlayersFailureAction {
  return {
    type: ActionType.GET_PLAYERS_FAILURE,
  };
}

export function getPlayers(
  _state: Player[],
  action: GetPlayersSuccessAction,
): Player[] {
  return action.response.players.map((player) => {
    return {id: player.Id, name: player.Name, money: player.Money};
  });
}
