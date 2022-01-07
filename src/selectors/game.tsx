import {createSelector} from 'reselect';
import {State} from '../state';

const selectGame = (state: State) => state.game;

export const selectRoundNumber = createSelector(
  selectGame,
  (game) => game.roundNumber,
);
