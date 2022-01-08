export interface GetCombatantsResponse {
  combatants: {
    Id: string;
    Name: string;
    Health: number;
    Streak: number;
  }[];
  count: number;
}

export interface PostPlayerResponse {
  Id: string;
  Name: string;
  Money: number;
}
