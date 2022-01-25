import {createSelector} from 'reselect';
import {State} from '../state';

export const selectCombatants = (state: State) => state.combatants;

export const selectCombatantNames = createSelector(
  selectCombatants,
  (combatants) => combatants.map((combatant) => combatant.name),
);
