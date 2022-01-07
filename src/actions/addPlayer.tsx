import {Action} from 'redux';
import {ActionType} from './actionType';

export interface AddPlayerAction extends Action {
  type: ActionType.ADD_PLAYER;
  name: string;
}

export function addPlayerAction(name: string): AddPlayerAction {
  return {
    type: ActionType.ADD_PLAYER,
    name,
  };
}
