import {createSelector} from 'reselect';
import {State} from '../state';

export const selectFight = (state: State) => state.fight;

export const selectFightStatus = createSelector(
  selectFight,
  (fight) => fight.status,
);
