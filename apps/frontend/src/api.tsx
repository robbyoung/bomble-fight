import {FightStatus} from './models';

export interface GetCombatantsResponse {
  combatants: {
    Id: string;
    Name: string;
    Health: number;
    Streak: number;

    Ferocity: number;
    Endurance: number;
    Skill: number;
    Agility: number;
    Speed: number;
  }[];
  count: number;
}

export interface GetPlayersResponse {
  players: {
    Id: string;
    Name: string;
    Money: number;
  }[];
  count: number;
}

export interface PostPlayerResponse {
  Id: string;
  Name: string;
  Money: number;
}

export interface PostBetResponse {
  PlayerId: string;
  CombatantId: string;
  Amount: number;
}

export interface PostFightResponse {
  Left: {
    Id: string;
    Health: number;
    Loss: number;
    Action: number;
  };
  Right: {
    Id: string;
    Health: number;
    Loss: number;
    Action: number;
  };
  FightStatus: FightStatus;
}

export interface GetPlayerStateResponse {
  Player: {
    Id: string;
    Name: string;
    Money: number;
  };
  Bet: {
    PlayerId: string;
    CombatantId: string;
    Amount: number;
  };
  FightStatus: FightStatus;
}

export interface ErrorResponse {
  status: number;
  message: string;
}
