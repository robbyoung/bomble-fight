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
    amount: string;
}

export interface Round {
    combatantIds: string[];
    bets: Bet[];
    seed: string;
}

export interface State {
    game: Game;
    players: Player[];
    combatants: Combatant[];
    currentRound: Round;
}