import {createSelector} from 'reselect';
import {State} from '../state';
import {selectCombatants} from './combatants';
import {selectPlayers} from './players';

export const selectBets = (state: State) => state.bets;

export const selectBetsByPlayer = createSelector(
  selectBets,
  selectPlayers,
  selectCombatants,
  (bets, players, combatants) =>
    players.map((player) => {
      const bet = bets.find((b) => b.playerId === player.id);
      const combatantName = bet
        ? combatants.find((c) => c.id === bet.combatantId)?.name
        : 'None';

      return {
        amount: bet ? bet.amount : 0,
        playerName: player.name,
        combatantName,
      };
    }),
);
