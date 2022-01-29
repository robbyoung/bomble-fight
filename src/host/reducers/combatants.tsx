import {Combatant} from '../../models';
import {Action} from 'redux';
import {ActionType} from '../actions/actionType';
import {
  getCombatants,
  GetCombatantsSuccessAction,
} from '../actions/getCombatants';

const defaultState: Combatant[] = [];

export default function combatants(
  state: Combatant[] = defaultState,
  action: Action,
): Combatant[] {
  switch (action.type) {
    case ActionType.GET_COMBATANTS_SUCCESS:
      return getCombatants(action as GetCombatantsSuccessAction);
    default:
      return state;
  }
}
