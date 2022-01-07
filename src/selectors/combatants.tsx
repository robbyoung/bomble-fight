import {createSelector} from 'reselect';
import {State} from '../state';

const selectCombatants = (state: State) => state.combatants;

export const selectCombatantNames = createSelector(
  selectCombatants,
  (players) => players.map((combatant) => combatant.name),
);
