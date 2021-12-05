import {Action} from 'redux';
import {ActionType} from './actionType';

export interface StartGameAction extends Action {
  type: ActionType.START_GAME;
}

export function startGameAction(): StartGameAction {
  return {
    type: ActionType.START_GAME,
  };
}
