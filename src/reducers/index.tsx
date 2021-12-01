import {Action} from 'redux';
import { State } from '../state';

const defaultState: State = {
    game: {roundNumber: 0},
    combatants: [],
    players: [],
    currentRound: {
        bets: [],
        combatantIds: [],
        seed: ''
    }
}

export default function rootReducer(
    state: State = defaultState,
    _action: Action,
): State {
    return state;
}