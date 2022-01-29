import {Bet, Combatant, Player} from '../models';

export interface State {
  player: Player;
  combatants: Combatant[];
  bet: Bet;
}
