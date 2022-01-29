import {Bet} from '../../models';
import {Action} from 'redux';
import {ActionType} from '../actions/actionType';
import {addBet, AddBetSuccessAction} from '../actions/addBet';

const defaultState: Bet = {
  combatantId: '',
  amount: 0,
  playerId: '',
};

export default function bet(state: Bet = defaultState, action: Action): Bet {
  switch (action.type) {
    case ActionType.ADD_BET_SUCCESS:
      return addBet(state, action as AddBetSuccessAction);
    default:
      return state;
  }
}
