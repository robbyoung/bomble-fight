import {createSelector} from 'reselect';
import {FightStatus} from '../../models';
import {State} from '../state';

export const selectFightStatus = (state: State) => state.fightStatus;

export const selectFightStatusDescription = createSelector(
  selectFightStatus,
  (status) => {
    switch (status) {
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
