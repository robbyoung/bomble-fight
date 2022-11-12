import {createSelector} from 'reselect';
import {State} from '../state';

export const selectBet = (state: State) => state.bet;

export const selectBetExists = createSelector(
  selectBet,
  (bet) => bet.amount !== 0,
);
