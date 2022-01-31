import {Player} from '../../models';
import {Action} from 'redux';
import {ActionType} from '../actions/actionType';
import {addPlayer, AddPlayerSuccessAction} from '../actions/addPlayer';

const defaultState: Player = {
  money: 0,
  name: '',
  id: '',
};

export default function player(
  state: Player = defaultState,
  action: Action,
): Player {
  switch (action.type) {
    case ActionType.ADD_PLAYER_SUCCESS:
      return addPlayer(state, action as AddPlayerSuccessAction);
    default:
      return state;
  }
}
