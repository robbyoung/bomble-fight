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

export enum FightStatus {
  Pending,
  Starting,
  Active,
  Finished,
}

export interface Fight {
  attackerId: string;
  defenderId: string;
  attackerHealth: number;
  defenderHealth: number;
  attackerDamage: number;
  defenderDamage: number;
  status: FightStatus;
}
