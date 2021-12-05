import {Action} from 'redux';
import {ActionType} from './actionType';

export interface StartGameAction extends Action {
  type: ActionType.START_GAME;
  numPlayers: number;
}

export function startGameAction(numPlayers: number): StartGameAction {
  return {
    type: ActionType.START_GAME,
    numPlayers,
  };
}
