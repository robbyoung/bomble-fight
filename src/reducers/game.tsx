import {ActionType} from '../actions/actionType';
import {Game} from '../state';
import {Action} from 'redux';
import {StartGameAction} from '../actions/startGame';

const defaultState: Game = {
  roundNumber: 0,
};

export default function game(state: Game = defaultState, action: Action): Game {
  switch (action.type) {
    case ActionType.START_GAME:
      return startGame(action as StartGameAction);
    default:
      return state;
  }
}

function startGame(action: StartGameAction): Game {
  return {
    roundNumber: 1,
  };
}
