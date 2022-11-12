import {Action} from 'redux';
import {GetPlayersResponse} from '../../api';
import {Player} from '../../models';
import {ActionType} from './actionType';

export function startPlayersPollAction(): Action {
  return {
    type: ActionType.START_PLAYERS_POLL,
  };
}

export function stopPlayersPollAction(): Action {
  return {
    type: ActionType.STOP_PLAYERS_POLL,
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
