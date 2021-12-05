import {Player} from '../state';
import {Action} from 'redux';
import {ActionType} from '../actions/actionType';
import {AddPlayerAction} from '../actions/addPlayer';

const defaultState: Player[] = [];

export default function players(
  state: Player[] = defaultState,
  action: Action,
): Player[] {
  switch (action.type) {
    case ActionType.ADD_PLAYER:
      return addPlayer(state, action as AddPlayerAction);
    default:
      return state;
  }
}

function addPlayer(state: Player[], action: AddPlayerAction): Player[] {
  return [...state, {id: `${state.length}`, name: action.name, money: 0}];
}
