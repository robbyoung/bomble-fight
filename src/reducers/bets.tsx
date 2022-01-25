import {Bet} from '../state';
import {Action} from 'redux';
import {ActionType} from '../actions/actionType';
import {addBet, AddBetSuccessAction} from '../actions/addBet';

const defaultState: Bet[] = [];

export default function bets(
  state: Bet[] = defaultState,
  action: Action,
): Bet[] {
  switch (action.type) {
    case ActionType.ADD_BET_SUCCESS:
      return addBet(state, action as AddBetSuccessAction);
    default:
      return state;
  }
}
