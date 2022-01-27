import {Action} from 'redux';
import {GetCombatantsResponse} from '../sagas/api';
import {ActionType} from './actionType';
import {Combatant} from '../state';

export interface GetCombatantsRequestAction extends Action {
  type: ActionType.GET_COMBATANTS_REQUEST;
}

export function getCombatantsAction(): GetCombatantsRequestAction {
  return {
    type: ActionType.GET_COMBATANTS_REQUEST,
  };
}

export interface GetCombatantsSuccessAction extends Action {
  type: ActionType.GET_COMBATANTS_SUCCESS;
  response: GetCombatantsResponse;
}

export function getCombatantsSuccessAction(
  response: GetCombatantsResponse,
): GetCombatantsSuccessAction {
  return {
    type: ActionType.GET_COMBATANTS_SUCCESS,
    response,
  };
}

export interface GetCombatantsFailureAction extends Action {
  type: ActionType.GET_COMBATANTS_FAILURE;
}

export function getCombatantsFailureAction(): GetCombatantsFailureAction {
  return {
    type: ActionType.GET_COMBATANTS_FAILURE,
  };
}

export function getCombatants(action: GetCombatantsSuccessAction): Combatant[] {
  return action.response.combatants.map((combatant): Combatant => {
    return {
      name: combatant.Name,
      health: combatant.Health,
      id: combatant.Id,
    };
  });
}
