import {Action} from 'redux';
import {Bet} from '../../models';
import {ActionType} from './actionType';
import {PostBetResponse} from '../../api';

export interface AddBetAction extends Action {
  type: ActionType.ADD_BET_REQUEST;
  bet: Bet;
}

export function addBetAction(
  playerId: string,
  combatantId: string,
  amount: number,
): AddBetAction {
  return {
    type: ActionType.ADD_BET_REQUEST,
    bet: {
      playerId,
      combatantId,
      amount,
    },
  };
}

export interface AddBetSuccessAction extends Action {
  type: ActionType.ADD_BET_SUCCESS;
  response: PostBetResponse;
}

export function addBetSuccessAction(
  response: PostBetResponse,
): AddBetSuccessAction {
  return {
    type: ActionType.ADD_BET_SUCCESS,
    response,
  };
}

export interface AddBetFailureAction extends Action {
  type: ActionType.ADD_BET_FAILURE;
}

export function addBetFailureAction(): AddBetFailureAction {
  return {
    type: ActionType.ADD_BET_FAILURE,
  };
}

export function addBet(_state: Bet, action: AddBetSuccessAction): Bet {
  const bet = action.response;
  return {
    playerId: bet.PlayerId,
    combatantId: bet.CombatantId,
    amount: bet.Amount,
  };
}
