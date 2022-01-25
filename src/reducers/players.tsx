import {Player} from '../state';
import {Action} from 'redux';
import {ActionType} from '../actions/actionType';
import {addPlayer, AddPlayerSuccessAction} from '../actions/addPlayer';
import {getPlayers, GetPlayersSuccessAction} from '../actions/getPlayers';

const defaultState: Player[] = [];

export default function players(
  state: Player[] = defaultState,
  action: Action,
): Player[] {
  switch (action.type) {
    case ActionType.ADD_PLAYER_SUCCESS:
      return addPlayer(state, action as AddPlayerSuccessAction);
    case ActionType.GET_PLAYERS_SUCCESS:
      return getPlayers(state, action as GetPlayersSuccessAction);
    default:
      return state;
  }
}
