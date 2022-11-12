import {Combatant, Fight, Player} from '../models';

export interface State {
  players: Player[];
  combatants: Combatant[];
  fight: Fight;
}
