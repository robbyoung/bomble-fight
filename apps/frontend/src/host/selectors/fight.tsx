import {createSelector} from 'reselect';
import {FightStatus} from '../../models';
import {selectCombatants} from '../../player/selectors/combatants';
import {State} from '../state';

export const selectFight = (state: State) => state.fight;

export const selectFightStatus = createSelector(
  selectFight,
  (fight) => fight.status,
);

export const selectFightStatusDescription = createSelector(
  selectFight,
  (fight) => {
    switch (fight.status) {
      case FightStatus.Pending:
        return 'Pending';
      case FightStatus.Starting:
        return 'Starting';
      case FightStatus.Active:
        return 'Active';
      case FightStatus.Finished:
        return 'Finished';
    }
  },
);

export const selectFightStepDetails = createSelector(
  selectCombatants,
  selectFight,
  (combatants, fight) => {
    return combatants.map((c) => {
      if (c.id === fight.left.id) {
        return {
          name: c.name,
          maxHealth: c.health,
          currentHealth: fight.left.health,
          loss: fight.left.loss,
        };
      } else {
        return {
          name: c.name,
          maxHealth: c.health,
          currentHealth: fight.right.health,
          loss: fight.right.loss,
        };
      }
    });
  },
);
