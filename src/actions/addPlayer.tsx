import {Action} from 'redux';
import {PostPlayerResponse} from '../sagas/api';
import {Player} from '../state';
import {ActionType} from './actionType';
import {v4} from 'uuid';

export interface AddPlayerAction extends Action {
  type: ActionType.ADD_PLAYER_REQUEST;
  player: Player;
}

export function addPlayerAction(name: string): AddPlayerAction {
  return {
    type: ActionType.ADD_PLAYER_REQUEST,
    player: {
      name,
      id: v4(),
      money: 50,
    },
  };
}

export interface AddPlayerSuccessAction extends Action {
  type: ActionType.ADD_PLAYER_SUCCESS;
  response: PostPlayerResponse;
}

export function addPlayerSuccessAction(
  response: PostPlayerResponse,
): AddPlayerSuccessAction {
  return {
    type: ActionType.ADD_PLAYER_SUCCESS,
    response,
  };
}

export interface AddPlayerFailureAction extends Action {
  type: ActionType.ADD_PLAYER_FAILURE;
}

export function addPlayerFailureAction(): AddPlayerFailureAction {
  return {
    type: ActionType.ADD_PLAYER_FAILURE,
  };
}

export function addPlayer(
  state: Player[],
  action: AddPlayerSuccessAction,
): Player[] {
  const player = action.response;
  return [...state, {id: player.Id, name: player.Name, money: player.Money}];
}
