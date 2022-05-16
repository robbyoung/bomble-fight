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

  ferocity: number;
  endurance: number;
  skill: number;
  agility: number;
  speed: number;
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

export enum FightAction {
  Nothing,
  Attack,
  Critical,
  Dodge,
  Block,
}

export interface CombatantStatus {
  id: string;
  health: number;
  loss: number;
  action: FightAction;
}

export interface Fight {
  left: CombatantStatus;
  right: CombatantStatus;
  status: FightStatus;
}
