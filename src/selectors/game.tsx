import {createSelector} from 'reselect';
import {State} from '../state';

export const selectGame = (state: State) => state.game;

export const getRoundNumber = createSelector(
  selectGame,
  (game) => game.roundNumber,
);
