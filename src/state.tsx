export interface Game {
  roundNumber: number;
}

export interface Player {
  id: string;
  name: string;
  money: number;
}

export interface Combatant {
  id: string;
  name: string;
  health: number;
}

export interface Bet {
  playerId: string;
  combatantId: string;
  amount: number;
}

export interface Round {
  combatantIds: string[];
}

export interface State {
  game: Game;
  players: Player[];
  combatants: Combatant[];
  bets: Bet[];
  currentRound: Round;
}
