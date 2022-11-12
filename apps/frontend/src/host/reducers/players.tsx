import {Player} from '../../models';
import {Action} from 'redux';
import {ActionType} from '../actions/actionType';
import {getPlayers, GetPlayersSuccessAction} from '../actions/getPlayers';

const defaultState: Player[] = [];

export default function players(
  state: Player[] = defaultState,
  action: Action,
): Player[] {
  switch (action.type) {
    case ActionType.GET_PLAYERS_SUCCESS:
      return getPlayers(state, action as GetPlayersSuccessAction);
    default:
      return state;
  }
}
