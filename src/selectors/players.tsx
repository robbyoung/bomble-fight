import {createSelector} from 'reselect';
import {State} from '../state';

export const selectPlayers = (state: State) => state.players;

export const selectPlayerCount = createSelector(
  selectPlayers,
  (players) => players.length,
);

export const selectPlayerNames = createSelector(selectPlayers, (players) =>
  players.map((player) => player.name),
);
