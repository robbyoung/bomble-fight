import {Bet, Combatant, FightStatus, Player} from '../models';

export interface State {
  player: Player;
  combatants: Combatant[];
  bet: Bet;
  fightStatus: FightStatus;
}
